package powerstatus

type PowerStatus string

const (
	POWER_OFF PowerStatus = "0000000000000000"
	POWER_ON  PowerStatus = "0000000000000001"
	ERROR     PowerStatus = "FFFFFFFFFFFFFFFF"
)
