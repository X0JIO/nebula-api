package generator

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateAll(
	identity *Identity,
	server ServerEndpoint,
) map[string]string {

	return map[string]string{
		"vless":       GenerateVLESS(identity, server),
		"vmess":       GenerateVMess(identity, server),
		"trojan":      GenerateTrojan(identity, server),
		"shadowsocks": GenerateShadowsocks(identity, server),
		"reality":     GenerateReality(identity, server),
	}
}
