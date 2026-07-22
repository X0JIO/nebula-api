package devices

import db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

func ToResponse(device db.Device) DeviceResponse {

	ip := ""

	if device.LastIp != nil {
		ip = device.LastIp.String()
	}

	return DeviceResponse{
		ID:       device.ID.String(),
		Name:     device.Name,
		Platform: device.Platform,
		LastIP:   ip,
		LastSeen: device.LastSeen.Time,
	}
}

func ToResponses(devices []db.Device) []DeviceResponse {

	result := make([]DeviceResponse, 0, len(devices))

	for _, d := range devices {
		result = append(result, ToResponse(d))
	}

	return result
}
