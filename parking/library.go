package parking

type CredentialLibrary struct {
	credentials map[string]Credential
}

func NewCredentialLibrary(credentials ...Credential) *CredentialLibrary {
	library := &CredentialLibrary{credentials: make(map[string]Credential, len(credentials))}
	for _, credential := range credentials {
		library.Add(credential)
	}
	return library
}

func (l *CredentialLibrary) Add(credential Credential) {
	if l.credentials == nil {
		l.credentials = make(map[string]Credential)
	}
	l.credentials[credential.Digest()] = credential
}

func (l *CredentialLibrary) Lookup(digest string) (Credential, bool) {
	credential, ok := l.credentials[digest]
	return credential, ok
}

func (l *CredentialLibrary) Find(licensePlate string) (Credential, bool) {
	for _, credential := range l.credentials {
		if credential.LicensePlate == licensePlate {
			return credential, true
		}
	}
	return Credential{}, false
}
