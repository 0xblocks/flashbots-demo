module github.com/0xblocks/flashbots-demo

go 1.16

require (
	contracts/checker v0.0.0-00010101000000-000000000000
	contracts/weth v0.0.0-00010101000000-000000000000
	github.com/0xblocks/flashbots-bundle v0.1.2
	github.com/ethereum/go-ethereum v1.10.3
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
)

replace (
	contracts/checker => ./contracts/checker
	contracts/weth => ./contracts/weth
)
