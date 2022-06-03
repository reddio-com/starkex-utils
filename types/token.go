package types

type TokenType = string

const (
	ETH     = TokenType("ETH")
	ERC20   = TokenType("ERC20")
	ERC721  = TokenType("ERC721")
	ERC20M  = TokenType("ERC20M")
	ERC721M = TokenType("ERC721M")
)
