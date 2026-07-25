package inbounds

type Client struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`

	Flow string `json:"flow,omitempty"`

	Password string `json:"password,omitempty"`

	Method string `json:"method,omitempty"`
}

type Settings struct {
	Clients []Client `json:"clients"`
}

type Sniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride"`
}
