# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
?   	parking-offline/cmd/parking-capi	[no test files]
?   	parking-offline/cmd/parking-offline	[no test files]
--- FAIL: TestExitRejectsCredentialAfterItWasReleased (0.00s)
    parking_test.go:55: second verification mismatch (-want +got):
          parking.Verification{
        - 	State:       "credential_used",
        + 	State:       "valid",
          	Credential:  {LicensePlate: "沪A12345", EntryTime: s"2025-06-01 08:15:00 +0000 UTC", ZoneCode: "B2"},
          	Digest:      "d57e4190721de5a10a967e99ea2335a54cb5df3cc14f1c5ae68f2f8ea576768e",
        - 	FeeCents:    0,
        + 	FeeCents:    300,
        - 	ChargeReady: false,
        + 	ChargeReady: true,
          }
FAIL
FAIL	parking-offline/parking	0.001s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/parking-capi): exit `0`
- Go run smoke (cmd/parking-offline): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/parking-capi): exit `0`
- Go run smoke (cmd/parking-offline): exit `0`
