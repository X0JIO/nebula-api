package reality

type KeyPair struct {
	PrivateKey string
	PublicKey  string
}

type Credentials struct {
	UUID       string
	PrivateKey string
	PublicKey  string
	ShortID    string
}
