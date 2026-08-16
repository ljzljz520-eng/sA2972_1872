package parking

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Credential struct {
	LicensePlate string
	EntryTime    time.Time
	ZoneCode     string
}

var (
	ErrLicensePlate = errors.New("license plate is required")
	ErrEntryTime    = errors.New("entry time must be set")
	ErrZoneCode     = errors.New("zone code is required")
	ErrCredential   = errors.New("credential is malformed")
)

func NewCredential(licensePlate string, entryTime time.Time, zoneCode string) (Credential, error) {
	credential := Credential{
		LicensePlate: strings.TrimSpace(licensePlate),
		EntryTime:    entryTime.UTC(),
		ZoneCode:     strings.TrimSpace(zoneCode),
	}
	if credential.LicensePlate == "" {
		return Credential{}, ErrLicensePlate
	}
	if credential.EntryTime.IsZero() {
		return Credential{}, ErrEntryTime
	}
	if credential.ZoneCode == "" {
		return Credential{}, ErrZoneCode
	}
	return credential, nil
}

func (c Credential) canonical() string {
	return fmt.Sprintf("%s|%s|%s", c.LicensePlate, c.EntryTime.UTC().Format(time.RFC3339Nano), c.ZoneCode)
}

func (c Credential) Digest() string {
	sum := sha256.Sum256([]byte(c.canonical()))
	return hex.EncodeToString(sum[:])
}

func (c Credential) Encode() string {
	return c.canonical()
}

func ParseCredential(encoded string) (Credential, error) {
	parts := strings.Split(encoded, "|")
	if len(parts) != 3 {
		return Credential{}, ErrCredential
	}
	entryTime, err := time.Parse(time.RFC3339Nano, parts[1])
	if err != nil {
		return Credential{}, ErrCredential
	}
	credential, err := NewCredential(parts[0], entryTime, parts[2])
	if err != nil {
		return Credential{}, ErrCredential
	}
	return credential, nil
}
