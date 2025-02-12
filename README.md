# Chain library
This library provides common utils for all components in the Lombard ecosystem.

The library is meant to not introduce new dependencies, so it only uses the Go Standard Library.

## LChainId
The `LChainId` type models the constraints and definitions of Lombard chain identifiers. Handy constructors are provided for chains we support.

### Supported Chains

- Ethereum `0x0000000000000000000000000000000000000000000000000000000000000001`
- Ethereum Sepolia `0x0000000000000000000000000000000000000000000000000000000000aa36a7`
- Ethereum Holesky `0x0000000000000000000000000000000000000000000000000000000000004268`
- Base `0x0000000000000000000000000000000000000000000000000000000000002105`
- Base Sepolia `0x0000000000000000000000000000000000000000000000000000000000014a34`
- BSC `0x0000000000000000000000000000000000000000000000000000000000000038`
- Sui `0x0100000000000000000000000000000000000000000000000000000035834a8a`
- Sui Testnet `0x010000000000000000000000000000000000000000000000000000004c78adac`
- Bitcoin `0xff0000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f`
- Bitcoin Signet `0xff000008819873e925422c1ff0f99f7cc9bbb232af63a077a480a3633bee1ef6`

## Address

The `Address` interface provides all the functionalities required by some data that carries information about a blockchain address. The address types of each supported chain implement this interface.