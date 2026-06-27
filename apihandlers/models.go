package apihandlers

import (
	"hamlib_rest_api/apihandlers/rigctld"
	"hamlib_rest_api/apihandlers/rotctld"
)

type DeviceList struct {
	Rigs     []rigctld.TransceiverConfig `json:"rigs"`
	Rotators []rotctld.RotatorConfig     `json:"rotators"`
}

// type RigInfo struct {
// 	ID    int `json:"id"`
// 	Model int `json:"model"`
// 	Port  int `json:"port"`
// }

// type RotatorInfo struct {
// 	ID    int `json:"id"`
// 	Model int `json:"model"`
// 	Port  int `json:"port"`
// }

// type DeviceConfig struct {
// 	ID       int    `json:"id"`
// 	Model    int    `json:"model"`
// 	Device   string `json:"device"`
// 	Baudrate int    `json:"baudrate"`
// 	Port     int    `json:"port"`
// 	Status   string `json:"status"` // Wird beim Laden des Zustands gesetzt
// }
