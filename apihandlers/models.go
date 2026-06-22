package apihandlers

type DeviceList struct {
	Rigs     []RigInfo     `json:"rigs"`
	Rotators []RotatorInfo `json:"rotators"`
}

type RigInfo struct {
	ID    int    `json:"id"`
	Model string `json:"model"`
	Port  string `json:"port"`
}

type RotatorInfo struct {
	ID    int `json:"id"`
	Model int `json:"model"`
	Port  int `json:"port"`
}
