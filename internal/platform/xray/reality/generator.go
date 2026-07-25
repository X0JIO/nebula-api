package reality

func Generate() (*Credentials, error) {

	keys, err := GenerateKeyPair()
	if err != nil {
		return nil, err
	}

	shortID, err := GenerateShortID()
	if err != nil {
		return nil, err
	}

	return &Credentials{
		UUID:       GenerateUUID(),
		PrivateKey: keys.PrivateKey,
		PublicKey:  keys.PublicKey,
		ShortID:    shortID,
	}, nil
}
