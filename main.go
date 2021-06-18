package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strings"
	"time"

	flashbots "github.com/0xblocks/flashbots-bundle"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	checker "contracts/checker"
	weth "contracts/weth"
)

const (
	WETH           = "0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6"
	CheckerAddress = "0x04E86113484B66E638D3f497e7695274C7041BD4"

	FlashbotsAttempts = 9 // Target 9 blocks since there are a limited number of miners on Goerli
)

type Bot struct {
	Client     *ethclient.Client
	SigningKey *ecdsa.PrivateKey
	PrivateKey *ecdsa.PrivateKey
	Address    *common.Address
}

func NewBot(signingKey string, privateKey string, providerURL string) (*Bot, error) {
	var err error
	bot := Bot{}

	bot.Client, err = ethclient.Dial(providerURL)
	if err != nil {
		return nil, err
	}

	bot.SigningKey, err = crypto.HexToECDSA(signingKey)
	if err != nil {
		return nil, err
	}

	bot.PrivateKey, bot.Address, err = wallet(privateKey)
	if err != nil {
		return nil, err
	}

	return &bot, nil
}

func main() {

	fmt.Println("Starting up")

	// Load config from env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	privateKey := os.Getenv("BOT_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("BOT_PRIVATE_KEY not set")
	}

	signingKey := os.Getenv("FLASHBOTS_SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("FLASHBOTS_SIGNING_KEY not set")
	}

	providerURL := os.Getenv("PROVIDER_URL")
	if providerURL == "" {
		log.Fatal("PROVIDER_URL not set")
	}

	bot, err := NewBot(signingKey, privateKey, providerURL)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %s\n", err)
	}

	tokenAddress := common.HexToAddress(WETH)
	wethABI, err := abi.JSON(strings.NewReader(weth.WETHABI))
	if err != nil {
		log.Fatal(err)
	}

	myBalance, err := bot.getBalance(bot.Address)
	log.Printf("my balance: %f\n", toFloat(myBalance))

	// Build meat transactions
	depositAmount, _ := new(big.Int).SetString("1000000000000000000", 10)
	tx1, err := bot.deposit(depositAmount, 0, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bundle tx1 (deposit): %s\n", tx1.Hash().Hex())

	withdrawAmount, _ := new(big.Int).SetString("500000000000000000", 10)
	tx2, err := bot.withdraw(withdrawAmount, tx1.Nonce()+1, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bundle tx2 (withdraw): %s\n", tx2.Hash().Hex())

	// Build check and send transaction
	// This makes sure that our previous txs did what we expected them to and
	// only pays the miner if they did. In this case, we're just checking that
	// the WETH balance is what we expect it to be
	dataBalance, err := wethABI.Pack("balanceOf", bot.Address)
	if err != nil {
		log.Fatal(err)
	}

	bribe, _ := new(big.Int).SetString("100000000000000000", 10)
	balance, _ := bot.getBalance(bot.Address)
	remanderAmount := new(big.Int).Add(balance, depositAmount)
	remanderAmount.Sub(remanderAmount, withdrawAmount)

	// CheckAndSend takes the data and addresses required to make the contract method calls itself.
	// Build the data structures needed to pass that in
	paddedRemander := common.LeftPadBytes(remanderAmount.Bytes(), 32)
	buffer := [32]byte{}
	for i := range paddedRemander {
		buffer[i] = paddedRemander[i]
	}

	addresses := []common.Address{tokenAddress}
	payloads := [][]byte{dataBalance}
	matches := [][32]byte{buffer}

	tx3, err := bot.checkAndBribe(bribe, addresses, payloads, matches, tx1.Nonce()+2, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bundle tx3 (checkAndBribe): %s\n", tx2.Hash().Hex())

	// Now we have a list of signed transactions. Send these to the flashbots relay
	txs := []*types.Transaction{tx1, tx2, tx3}
	err = bot.sendFlashbotsBundle(txs, FlashbotsAttempts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bundle summitted")

	// Pause and wait until the transaction to be mined
	err = bot.waitForTx(tx3, 300)
	if err != nil {
		log.Fatal(err)
	}
}

func (bot *Bot) getBalance(who *common.Address) (*big.Int, error) {

	tokenAddress := common.HexToAddress(WETH)
	instance, err := weth.NewWETH(tokenAddress, bot.Client)
	if err != nil {
		return nil, err
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, *who)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func (bot *Bot) deposit(amount *big.Int, nonce uint64, zeroGas bool) (*types.Transaction, error) {

	tokenAddress := common.HexToAddress(WETH)
	contractABI, err := abi.JSON(strings.NewReader(weth.WETHABI))
	if err != nil {
		return nil, err
	}

	data, err := contractABI.Pack("deposit")
	if err != nil {
		return nil, err
	}

	signedTx, err := bot.buildTransaction(&tokenAddress, amount, data, nonce, zeroGas)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (bot *Bot) withdraw(amount *big.Int, nonce uint64, zeroGas bool) (*types.Transaction, error) {

	tokenAddress := common.HexToAddress(WETH)
	contractABI, err := abi.JSON(strings.NewReader(weth.WETHABI))
	if err != nil {
		return nil, err
	}

	data, err := contractABI.Pack("withdraw", amount)
	if err != nil {
		return nil, err
	}

	signedTx, err := bot.buildTransaction(&tokenAddress, big.NewInt(0), data, nonce, zeroGas)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (bot *Bot) checkAndBribe(bribe *big.Int, addresses []common.Address, payloads [][]byte, matches [][32]byte, nonce uint64, zeroGas bool) (*types.Transaction, error) {

	contractAddress := common.HexToAddress(CheckerAddress)
	contractABI, err := abi.JSON(strings.NewReader(checker.FlashbotsCheckAndSendABI))
	if err != nil {
		return nil, err
	}

	data, err := contractABI.Pack("check32BytesAndSendMulti", addresses, payloads, matches)
	if err != nil {
		return nil, err
	}

	signedTx, err := bot.buildTransaction(&contractAddress, bribe, data, nonce, zeroGas)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (bot *Bot) buildTransaction(to *common.Address, value *big.Int, data []byte, nonce uint64, zeroGas bool) (*types.Transaction, error) {

	var err error
	if nonce == 0 {
		nonce, err = bot.Client.PendingNonceAt(context.Background(), *bot.Address)
		if err != nil {
			return nil, err
		}
	}

	gasPrice := big.NewInt(0)
	if !zeroGas {
		gasPrice, err = bot.Client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
	}

	gasLimit, err := bot.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: *bot.Address,
		To:   to,
		Data: data,
	})
	if err != nil {
		log.Printf("Caught error estimating gas: %s\n", err)
		gasLimit = 100000
	}

	chainID, err := bot.Client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, *to, value, gasLimit*2, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), bot.PrivateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (bot *Bot) sendFlashbotsBundle(signedTxs []*types.Transaction, attempts int64) error {

	start := time.Now()

	fb := flashbots.NewProvider(bot.SigningKey, bot.PrivateKey, flashbots.TestRelayURL)

	txs := []string{}
	for _, tx := range signedTxs {
		data, err := tx.MarshalBinary()
		if err != nil {
			return err
		}
		txs = append(txs, hexutil.Encode(data))
	}

	block, err := bot.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	// First simulate the transaction to make sure it doesn't revert
	resp, err := fb.Simulate(txs, block.Number, "latest", 0)
	if err != nil {
		return err
	}

	err = resp.HasError()
	if err != nil {
		return err
	}

	cb, _ := new(big.Float).SetString(resp.Result.CoinbaseDiff)
	eth := new(big.Float).Quo(cb, big.NewFloat(math.Pow10(18)))
	wei, _ := resp.EffectiveGasPrice()
	gwei := toGwei(wei)

	fmt.Printf("Simulation completed in %fs. Cost: %f eth, effective price: %d gwei\n", time.Since(start).Seconds(), eth, gwei)

	// Send the bundle to the FB relay for inclusion in a block
	// Unless your tx is only valid for a single block, you should target several blocks since flashbots
	// isn't available on 100% of the hashing power. On Goerli, it's only available on a single validator,
	// so target at least 9 blocks to ensure it gets picked up
	for i := int64(0); i < attempts; i++ {
		targetBlockNumber := new(big.Int).Add(block.Number, big.NewInt(int64(i)))
		_, err := fb.SendBundle(txs, targetBlockNumber, &flashbots.Options{})
		if err != nil {
			return err
		}
		fmt.Printf("submitted for block: %d\n", targetBlockNumber)
	}
	return nil
}

func (bot *Bot) waitForTx(tx *types.Transaction, maxWaitSeconds uint) error {

	log.Println("Waiting for tx to complete...")
	loops := uint(0)
	for {
		_, isPending, err := bot.Client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil && err != ethereum.NotFound {
			return err
		}

		if !isPending {
			// It's not pending, so check if it's been mined
			receipt, err := bot.Client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil && err != ethereum.NotFound {
				return err
			}
			if receipt != nil {
				log.Println("tx complete. it may take a few minutes to appear in etherscan.")
				log.Printf("https://goerli.etherscan.io/tx/%s\n", tx.Hash().Hex())
				break
			}
		}

		time.Sleep(1 * time.Second)

		loops = loops + 1
		if loops > maxWaitSeconds {
			log.Printf("timed out after %d seconds. check manually here:\n", maxWaitSeconds)
			log.Printf("https://goerli.etherscan.io/tx/%s\n", tx.Hash().Hex())
			break
		}
	}

	return nil
}

func wallet(private string) (*ecdsa.PrivateKey, *common.Address, error) {

	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return nil, nil, err
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	return privateKey, &fromAddress, nil
}

func toFloat(val *big.Int) *big.Float {

	fval := new(big.Float)
	fval.SetString(val.String())
	return new(big.Float).Quo(fval, big.NewFloat(math.Pow10(int(18))))
}

func toGwei(wei *big.Int) *big.Int {

	return new(big.Int).Div(wei, big.NewInt(1000000000))
}
