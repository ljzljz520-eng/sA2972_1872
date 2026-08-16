package parking

import "time"

type Fixture struct {
	Credential Credential
	ExitTime   time.Time
	ExitZone   string
	RateCents  int
}

func NewFixture() Fixture {
	entry := time.Date(2025, time.June, 1, 8, 15, 0, 0, time.UTC)
	credential, _ := NewCredential("沪A12345", entry, "B2")
	return Fixture{
		Credential: credential,
		ExitTime:   time.Date(2025, time.June, 1, 10, 45, 0, 0, time.UTC),
		ExitZone:   "B2",
		RateCents:  100,
	}
}

func (f Fixture) Library() *CredentialLibrary {
	return NewCredentialLibrary(f.Credential)
}

func (f Fixture) Verifier() *OfflineVerifier {
	verifier, _ := NewOfflineVerifier(f.Library(), f.RateCents)
	return verifier
}
