# flashbots-demo

This is a demo bot written in go that uses flashbots v0.2 on the Goerli testnet, that works out of the box. Feel free to take this and use it as the starting point for your own MEV bot.

## Demo flow

This bot creates a simple bundle that does the following:

1. Deposit 1 eth into the WETH token. It's using the same WETH contract used by uniswap on goerli.
2. Withdraw 0.5 WETH back into eth, leaving 0.5 as WETH.
3. Call `FlashbotsCheckAndSend` to check that the remaining WETH balance is the previous balance + 0.5.
4. If the check pass, transfer 0.1 eth to the miner.

I've deployed [FlashbotsCheckAndSend](https://github.com/flashbots/searcher-sponsored-tx/blob/main/contracts/FlashbotsCheckAndSend.sol) to goerli and included the address in this demo.

Feel free to deploy your own version, or modify the code to use your own.

Since this demo is leaving half of each deposit in WETH, you will eventually run out. I'm leaving it as an exercise to the reader to figure out how to withdraw the rest of your eth.

## Setup

First, you'll need a wallet for your bot. You can generate a new wallet using the following:

```
go run init/init.go > .env
```

Update the created `.env` file with the correct URL for your provider.

Next, you need some eth. You can request some from the [Goerli social faucet](https://faucet.goerli.mudit.blog/). You'll need to tweet your address and provide the link to the faucet.

## Running the demo

Once you have a wallet with some eth, and you've setup your `.env` file, you're ready to run the bot

```
make run
```