package inbounds

type Inbound struct {
	Tag            string         `json:"tag"`
	Listen         string         `json:"listen,omitempty"`
	Port           int            `json:"port"`
	Protocol       string         `json:"protocol"`
	Settings       Settings       `json:"settings"`
	StreamSettings StreamSettings `json:"streamSettings,omitempty"`
}

type Settings struct {
	Clients []Client `json:"clients,omitempty"`
}

type Client struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Method   string `json:"method,omitempty"`
}

type StreamSettings struct {
	Network  string `json:"network,omitempty"`
	Security string `json:"security,omitempty"`

	RealitySettings *RealitySettings `json:"realitySettings,omitempty"`
}

func BaseInbound(
	tag string,
	protocol string,
	port int,
	settings Settings,
	stream StreamSettings,
) Inbound {

	return Inbound{
		Tag:            tag,
		Port:           port,
		Protocol:       protocol,
		Settings:       settings,
		StreamSettings: stream,
	}
}
