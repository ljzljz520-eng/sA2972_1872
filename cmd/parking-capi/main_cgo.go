//go:build cgo

package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"sync"
	"time"
	"unsafe"

	"parking-offline/parking"
)

var (
	serviceMu       sync.Mutex
	serviceLibrary  = parking.NewCredentialLibrary()
	serviceVerifier = mustVerifier(serviceLibrary)
)

func main() {}

func mustVerifier(library *parking.CredentialLibrary) *parking.OfflineVerifier {
	verifier, err := parking.NewOfflineVerifier(library, 100)
	if err != nil {
		panic(err)
	}
	return verifier
}

//export ParkingCredentialCreate
func ParkingCredentialCreate(licensePlate *C.char, entryUnix int64, zoneCode *C.char) *C.char {
	credential, err := parking.NewCredential(C.GoString(licensePlate), time.Unix(entryUnix, 0).UTC(), C.GoString(zoneCode))
	if err != nil {
		return C.CString("")
	}
	serviceMu.Lock()
	serviceLibrary.Add(credential)
	serviceMu.Unlock()
	return C.CString(credential.Encode())
}

//export ParkingCredentialValidate
func ParkingCredentialValidate(encoded *C.char, exitUnix int64, exitZone *C.char) *C.char {
	credential, err := parking.ParseCredential(C.GoString(encoded))
	if err != nil {
		return C.CString(string(parking.StateInvalidCredential))
	}
	serviceMu.Lock()
	result := serviceVerifier.Validate(credential, time.Unix(exitUnix, 0).UTC(), C.GoString(exitZone))
	serviceMu.Unlock()
	return C.CString(string(result.State))
}

//export ParkingStringFree
func ParkingStringFree(value *C.char) {
	C.free(unsafe.Pointer(value))
}
