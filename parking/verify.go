package parking

import (
	"errors"
	"math"
	"time"
)

type ValidationState string

const (
	StateValid             ValidationState = "valid"
	StateCredentialUsed    ValidationState = "credential_used"
	StateUnknownCredential ValidationState = "unknown_credential"
	StateInvalidCredential ValidationState = "invalid_credential"
	StateWrongZone         ValidationState = "wrong_zone"
	StateInvalidExitTime   ValidationState = "invalid_exit_time"
)

type Verification struct {
	State       ValidationState
	Credential  Credential
	Digest      string
	FeeCents    int
	ChargeReady bool
}

type OfflineVerifier struct {
	library    *CredentialLibrary
	usedDigest map[string]struct{}
	rateCents  int
}

func NewOfflineVerifier(library *CredentialLibrary, hourlyRateCents int) (*OfflineVerifier, error) {
	if library == nil {
		return nil, errors.New("credential library is required")
	}
	if hourlyRateCents <= 0 {
		return nil, errors.New("hourly rate must be positive")
	}
	return &OfflineVerifier{
		library:    library,
		usedDigest: make(map[string]struct{}),
		rateCents:  hourlyRateCents,
	}, nil
}

func (v *OfflineVerifier) Validate(credential Credential, exitTime time.Time, exitZone string) Verification {
	digest := credential.Digest()
	if credential.LicensePlate == "" || credential.EntryTime.IsZero() || credential.ZoneCode == "" {
		return Verification{State: StateInvalidCredential, Credential: credential, Digest: digest}
	}
	registered, ok := v.library.Lookup(digest)
	if !ok || registered != credential {
		return Verification{State: StateUnknownCredential, Credential: credential, Digest: digest}
	}
	if registered.ZoneCode != exitZone {
		return Verification{State: StateWrongZone, Credential: registered, Digest: digest}
	}
	if exitTime.Before(registered.EntryTime) {
		return Verification{State: StateInvalidExitTime, Credential: registered, Digest: digest}
	}
	duration := exitTime.Sub(registered.EntryTime)
	hours := int(math.Ceil(duration.Hours()))
	if hours < 1 {
		hours = 1
	}
	return Verification{
		State:       StateValid,
		Credential:  registered,
		Digest:      digest,
		FeeCents:    hours * v.rateCents,
		ChargeReady: true,
	}
}
