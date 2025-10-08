# [v0.5.1](https://github.com/lombard-finance/chain/releases/tag/v0.5.1)
- Add LChainId for Avalanche C-Chain and Fuji testnet
- Add LChainId for Lombard Ledger testnets
# [v0.5.0](https://github.com/lombard-finance/chain/releases/tag/v0.5.0)
- Support `LChainId` and `Address` for Starknet chains
- Add `Address` constructors for Zero address on supported chains
- Add `LChainId` constructors for Ink, Sonic, BSC and their testnets
- Add constructor for Babylon `LChainId`
- Truncate by default EVM addresses to 20 bytes and check truncated bytes are zeroes.
# [v0.4.1](https://github.com/lombard-finance/chain/releases/tag/v0.4.1)
- Accept both chain Id and chain name in Cosmos `LChainId` constructor
# [v0.4.0](https://github.com/lombard-finance/chain/releases/tag/v0.4.0)
- Support `LChainId` and `Address` for Cosmos chains
# [v0.3.0](https://github.com/lombard-finance/chain/releases/tag/v0.3.0)
- Implementation for Solana ecosystem in `LChainId` and `Address`
- `base58` library
# [v0.2.0](https://github.com/lombard-finance/chain/releases/tag/v0.2.0)
- Remove support enforcement on `ledger-utils` side by introducing `GenericLChainId` and `GenericAddress` types
# [v0.1.2](https://github.com/lombard-finance/chain/releases/tag/v0.1.2)
- Add `NewEvmAddressTruncating` to create address from longer byte slice
# [v0.1.1](https://github.com/lombard-finance/chain/releases/tag/v0.1.1)
- Fix to Sui address length
# [v0.1.0](https://github.com/lombard-finance/chain/releases/tag/v0.1.0)
- Introduce the `chainid` package, the `LChainId` interface and the implementations for all the supported chains
- Introduce the `address` package, the `Address` interface and the implementations for all the supported ecosystems