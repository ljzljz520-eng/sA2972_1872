package parking

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestEntranceGeneratesOfflineCredential(t *testing.T) {
	entry := time.Date(2025, 6, 1, 8, 15, 0, 0, time.UTC)
	credential, err := NewCredential("沪A12345", entry, "B2")
	if err != nil {
		t.Fatalf("NewCredential returned error: %v", err)
	}
	want := Credential{LicensePlate: "沪A12345", EntryTime: entry, ZoneCode: "B2"}
	if diff := cmp.Diff(want, credential); diff != "" {
		t.Fatalf("credential mismatch (-want +got):\n%s", diff)
	}
	if credential.Digest() == "" {
		t.Fatal("credential digest is empty")
	}
}

func TestExitValidatesKnownCredentialBeforeCharging(t *testing.T) {
	fixture := NewFixture()
	result := fixture.Verifier().Validate(fixture.Credential, fixture.ExitTime, fixture.ExitZone)
	want := Verification{State: StateValid, Credential: fixture.Credential, Digest: fixture.Credential.Digest(), FeeCents: 300, ChargeReady: true}
	if diff := cmp.Diff(want, result); diff != "" {
		t.Fatalf("verification mismatch (-want +got):\n%s", diff)
	}
}

func TestExitRejectsCredentialFromAnotherZone(t *testing.T) {
	fixture := NewFixture()
	result := fixture.Verifier().Validate(fixture.Credential, fixture.ExitTime, "C1")
	if result.State != StateWrongZone {
		t.Fatalf("state = %q, want %q", result.State, StateWrongZone)
	}
	if result.ChargeReady {
		t.Fatal("charge should not be ready")
	}
}

func TestExitRejectsCredentialAfterItWasReleased(t *testing.T) {
	fixture := NewFixture()
	verifier := fixture.Verifier()
	first := verifier.Validate(fixture.Credential, fixture.ExitTime, fixture.ExitZone)
	second := verifier.Validate(fixture.Credential, fixture.ExitTime, fixture.ExitZone)
	if first.State != StateValid || !first.ChargeReady {
		t.Fatalf("first verification = %#v, want a charge-ready valid result", first)
	}
	want := Verification{State: StateCredentialUsed, Credential: fixture.Credential, Digest: fixture.Credential.Digest()}
	if diff := cmp.Diff(want, second); diff != "" {
		t.Fatalf("second verification mismatch (-want +got):\n%s", diff)
	}
}
