# ledger-utils

Go library of Lombard chain identifiers (`LChainId`) and blockchain addresses (`Address`), imported by Lombard components such as notaryd and the ledger.
Stack: Go (version pinned in `go.mod`), standard library only.

## Commands

The `Makefile` is the command list. CI runs `make build` and `make test`.

## Map

- `chainid/chain_id.go` — `LChainId` interface, `Ecosystem` values and `NewLChainId`. Read when adding an ecosystem.
- `chainid/` — one file per ecosystem with its predefined `LChainId` constructors. Read the ecosystem's file when adding a chain.
- `address/address.go` — `Address` interface and the `NewAddress` dispatch by ecosystem. Read when adding an address type.
- `common/base58/` — base58 copied from mr-tron/base58. Read when changing Solana or Bitcoin encoding.

## Conventions and gotchas

- Add a case to `TestLChainId` in `chainid/chain_id_test.go` for every new `LChainId` constructor — for test coverage.
- Use `LChainId` for Lombard identifiers and keep "chain id" for an ecosystem's own value — auditors flagged `ChainId` as confusable with ecosystem chain ids.
- Keep external dependencies to a minimum.

## Boundaries

Ask first:

- `make create-release` or any push of a tag.
- Changing or removing an exported identifier in `chainid` or `address`.

## Done means

- `make build` passes.
- `make test` passes.
