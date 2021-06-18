files := main.go

build-checker:
	solc --bin contracts/checker/sol/FlashbotsCheckAndSend.sol -o contracts/checker/sol/bin/ --overwrite
	solc --abi contracts/checker/sol/FlashbotsCheckAndSend.sol -o contracts/checker/sol/abi/ --overwrite
	abigen --bin=contracts/checker/sol/bin/FlashbotsCheckAndSend.bin --abi=contracts/checker/sol/abi/FlashbotsCheckAndSend.abi --pkg=FlashbotsCheckAndSend > contracts/checker/FlashbotsCheckAndSend.go

build-weth:
	solc --bin contracts/weth/sol/WETH.sol -o contracts/weth/sol/bin/ --overwrite
	solc --abi contracts/weth/sol/WETH.sol -o contracts/weth/sol/abi/ --overwrite
	abigen --bin=contracts/weth/sol/bin/WETH.bin --abi=contracts/weth/sol/abi/WETH.abi --pkg=WETH > contracts/weth/WETH.go

run:
	go run $(files)