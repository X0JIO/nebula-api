package inbounds

type RealitySettings struct {
	Show        bool     `json:"show"`
	Dest        string   `json:"dest"`
	Xver        int      `json:"xver"`
	ServerNames []string `json:"serverNames"`

	PrivateKey string   `json:"privateKey"`
	ShortIds   []string `json:"shortIds"`
}

type StreamSettings struct {
	Network  string `json:"network"`
	Security string `json:"security,omitempty"`

	RealitySettings *RealitySettings `json:"realitySettings,omitempty"`
}
