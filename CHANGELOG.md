# Changelog

## [0.6.0] - 2026-11-17

### Changed

- **Breaking:** store the chain Id as a fixed `[32]byte` array so `LChainId` values are comparable and usable as map keys ([#18](https://github.com/lombard-finance/ledger-utils/pull/18)) (hashxtree)
- **Breaking:** return values instead of pointers from the Cosmos, EVM, Solana, Starknet and Sui `LChainId` constructors ([#19](https://github.com/lombard-finance/ledger-utils/pull/19)) (hashxtree)
- Bump minimum Go version to 1.26.11 ([#18](https://github.com/lombard-finance/ledger-utils/pull/18)) (hashxtree)

### Added

- Add `LChainId` constructors for Monad, MegaETH, Stable and their testnets ([#25](https://github.com/lombard-finance/ledger-utils/pull/25)) (Diver229)
- Add `LChainId` constructors for Lombard Ledger Gastald testnet and staging devnet ([#17](https://github.com/lombard-finance/ledger-utils/pull/17)) (Antonio Russo)
- Add `LChainId` constructors for Avalanche and Avalanche Fuji testnet ([#16](https://github.com/lombard-finance/ledger-utils/pull/16)) (Petr Murzin, Antonio Russo)

### Fixed

- Reject Starknet addresses that are not lower than the Starknet field prime, so an address has a single byte representation ([#26](https://github.com/lombard-finance/ledger-utils/pull/26)) (andrey)

## [0.5.0] - 2025-06-18

### Changed

- **Breaking:** return an error from `NewGenericAddress` and reject empty addresses ([#11](https://github.com/lombard-finance/ledger-utils/pull/11)) (Antonio Russo)
- Accept byte slices longer than 20 bytes in `NewEvmAddress` and `NewAddress`, truncating leading bytes that must be zero ([#11](https://github.com/lombard-finance/ledger-utils/pull/11), [#13](https://github.com/lombard-finance/ledger-utils/pull/13)) (Antonio Russo)

### Added

- Add Starknet support to `LChainId` and `Address` ([#8](https://github.com/lombard-finance/ledger-utils/pull/8)) (Antonio Russo)
- Add `LChainId` constructors for Ink, Sonic, BSC and their testnets ([#12](https://github.com/lombard-finance/ledger-utils/pull/12)) (Petr Murzin, Antonio Russo)
- Add `NewZeroAddress` constructor for supported ecosystems ([#10](https://github.com/lombard-finance/ledger-utils/pull/10)) (Antonio Russo)
- Add Babylon `LChainId` constructor ([#9](https://github.com/lombard-finance/ledger-utils/pull/9)) (Antonio Russo)

### Removed

- **Breaking:** remove `NewEvmAddressTruncating`, use `NewEvmAddress` instead ([#13](https://github.com/lombard-finance/ledger-utils/pull/13)) (Antonio Russo)

## [0.4.1] - 2025-03-19

### Changed

- Accept both chain Id and chain name in `NewCosmosLChainId` ([#7](https://github.com/lombard-finance/ledger-utils/pull/7)) (Antonio Russo)

## [0.4.0] - 2025-03-14

### Added

- Add Cosmos support to `LChainId` and `Address` ([#3](https://github.com/lombard-finance/ledger-utils/pull/3)) (Antonio Russo)

### Fixed

- Fix `Ecosystem.String()` for Solana ([#4](https://github.com/lombard-finance/ledger-utils/pull/4)) (Antonio Russo)

## [0.3.0] - 2025-02-26

### Added

- Add Solana support to `LChainId` and `Address` ([#2](https://github.com/lombard-finance/ledger-utils/pull/2)) (Antonio Russo)
- Add `base58` package ([#2](https://github.com/lombard-finance/ledger-utils/pull/2)) (Antonio Russo)

## [0.2.0] - 2025-02-14

### Changed

- **Breaking:** rename module from `github.com/lombard-finance/chain` to `github.com/lombard-finance/ledger-utils` ([`830ba0a`](https://github.com/lombard-finance/ledger-utils/commit/830ba0a)) (Antonio Russo)
- **Breaking:** replace `UnsafeLChainId` with `GenericLChainId` ([`95a3365`](https://github.com/lombard-finance/ledger-utils/commit/95a3365)) (Antonio Russo)
- Return `GenericLChainId` and `GenericAddress` for unsupported ecosystems instead of an error ([`95a3365`](https://github.com/lombard-finance/ledger-utils/commit/95a3365), [`4f8f4af`](https://github.com/lombard-finance/ledger-utils/commit/4f8f4af)) (Antonio Russo)

## [0.1.2] - 2025-02-14

### Added

- Add `NewEvmAddressTruncating` to create an EVM address from a longer byte slice ([`43a6679`](https://github.com/lombard-finance/ledger-utils/commit/43a6679)) (Antonio Russo)

## [0.1.1] - 2025-02-14

### Fixed

- Fix `Length()` of Sui addresses, which returned the EVM address length ([`ed1b874`](https://github.com/lombard-finance/ledger-utils/commit/ed1b874)) (Antonio Russo)

## [0.1.0] - 2025-02-12

_Initial release._

### Added

- Add `chainid` package with the `LChainId` interface and implementations for all supported chains ([`1aa8036`](https://github.com/lombard-finance/ledger-utils/commit/1aa8036)) (Antonio Russo)
- Add `address` package with the `Address` interface and implementations for all supported ecosystems ([`6da981b`](https://github.com/lombard-finance/ledger-utils/commit/6da981b), [`7e00b55`](https://github.com/lombard-finance/ledger-utils/commit/7e00b55)) (Antonio Russo)

[0.6.0]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.6.0
[0.5.0]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.5.0
[0.4.1]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.4.1
[0.4.0]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.4.0
[0.3.0]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.3.0
[0.2.0]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.2.0
[0.1.2]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.1.2
[0.1.1]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.1.1
[0.1.0]: https://github.com/lombard-finance/ledger-utils/releases/tag/v0.1.0
