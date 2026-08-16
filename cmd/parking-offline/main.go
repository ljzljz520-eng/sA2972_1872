package main

import (
	"fmt"

	"parking-offline/parking"
)

func main() {
	fixture := parking.NewFixture()
	result := fixture.Verifier().Validate(fixture.Credential, fixture.ExitTime, fixture.ExitZone)
	fmt.Printf("credential_status=%s\n", result.State)
	fmt.Printf("credential_digest=%s\n", result.Digest)
	fmt.Printf("fee_cents=%d\n", result.FeeCents)
	fmt.Printf("charge_ready=%t\n", result.ChargeReady)
}
