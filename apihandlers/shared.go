package apihandlers

import (
	"hamlib_rest_api/apihandlers/rigctld"
	"hamlib_rest_api/apihandlers/rotctld"
	"hamlib_rest_api/config"
	"net/http"
)

func RegisterRoutesRigctld(mux *http.ServeMux) {
	// services routines
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/service/start", rigctld.HandleStartRigctld)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/service/stop", rigctld.HandleStopRigctld)
	mux.HandleFunc("GET "+config.ApiBasePath+"/devices", HandleListDevices)
	mux.HandleFunc("GET "+config.ApiBasePath+"/devices/rigs", rigctld.HandleListRigs)

	// Basic Control Endpoints
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/frequency", rigctld.HandleGetFrequency)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/frequency", rigctld.HandleSetFrequency)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/mode", rigctld.HandleGetMode)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/mode", rigctld.HandleSetMode)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/split-frequency-mode", rigctld.HandleGetSplitFrequencyMode)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/split-frequency-mode", rigctld.HandleSetSplitFrequencyMode)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/split-frequency", rigctld.HandleGetSplitFrequency)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/split-frequency", rigctld.HandleSetSplitFrequency)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/split-modes", rigctld.HandleGetSplitModes)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/split-modes/{split_mode_name}", rigctld.HandleGetSplitMode)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/split-modes/{split_mode_name}", rigctld.HandleSetSplitMode)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/level/{level_param}", rigctld.HandleGetLevel)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/level/{level_param}", rigctld.HandleSetLevel)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/levels", rigctld.HandleGetLevels)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/tuning-step", rigctld.HandleGetTuningStep)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/tuning-step", rigctld.HandleSetTuningStep)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/split-vfo", rigctld.HandleGetSplitVFO)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/split-vfo", rigctld.HandleSetSplitVFO)

	// Functions, Parameters, Scans & Transceive
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/functions", rigctld.HandleGetFunctions)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/functions/{function_name}", rigctld.HandleGetFunction)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/functions/{function_name}", rigctld.HandleSetFunction)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/parameters", rigctld.HandleGetParameters)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/parameters/{parameter_name}", rigctld.HandleGetParameter)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/parameters/{parameter_name}", rigctld.HandleSetParameter)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/scan/functions", rigctld.HandleGetScanFunctions)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/scan/{scan_param}", rigctld.HandleGetScan)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/scan/{scan_param}", rigctld.HandleSetScan)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/transceive-modes", rigctld.HandleGetTransceive)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/transceive-modes/{transceive_mode}", rigctld.HandleGetTransceive)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/transceive-modes/{transceive_mode}", rigctld.HandleSetTransceive)

	// Repeater Shift, Offset, Tones & VFO
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/repeater-shift", rigctld.HandleGetRepeaterShift)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/repeater-shift", rigctld.HandleSetRepeaterShift)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/repeater-offset", rigctld.HandleGetRepeaterOffset)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/repeater-offset", rigctld.HandleSetRepeaterOffset)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/tone-ctcss", rigctld.HandleGetCtcssTone)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/tone-ctcss", rigctld.HandleSetCtcssTone)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/tone-dcs", rigctld.HandleGetDcsTone)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/tone-dcs", rigctld.HandleSetDcsTone)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/vfo", rigctld.HandleGetVFO)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/vfo", rigctld.HandleSetVFO)

	// PTT, Memory, Channels, RIT/XIT & Antenna
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/ptt", rigctld.HandleGetPTT)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/ptt", rigctld.HandleSetPTT)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/memory-channel-number", rigctld.HandleGetMemoryChannelNumber)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/memory-channel-number", rigctld.HandleSetMemoryChannelNumber)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/channel", rigctld.HandleGetChannel)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/channel", rigctld.HandleSetChannel)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/info", rigctld.HandleGetInfo)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/rit", rigctld.HandleGetRit)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/rit", rigctld.HandleSetRit)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/xit", rigctld.HandleGetXit)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/xit", rigctld.HandleSetXit)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/antenna-number", rigctld.HandleGetAntennaNumber)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/antenna-number", rigctld.HandleSetAntenna)

	// Raw Commands, Morse & Power Conversions
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/raw", rigctld.HandleSetRawCommand)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/raw-rx", rigctld.HandleSetRawCommandRx)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/power-to-factor", rigctld.HandleGetMw2Power)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/power-to-mw", rigctld.HandleGetPower2Mw)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/capabilities", rigctld.HandleGetCapabilities)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/configuration", rigctld.HandleGetConfiguration)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/morse", rigctld.HandleSetMorse)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/morse/stop", rigctld.HandleSetMorseStop)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/morse/wait", rigctld.HandleGetMorseWait)

	// SQL Extensions, Rig State & Misc
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/sql-ctcss", rigctld.HandleGetCtcssSql)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/sql-ctcss", rigctld.HandleSetCtcssSql)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/sql-dcs", rigctld.HandleGetDcsSql)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/sql-dcs", rigctld.HandleSetDcsSql)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/dtmf", rigctld.HandleGetDtmf)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/dtmf", rigctld.HandleSetDtmf)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/dcd", rigctld.HandleGetDcd)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/twiddle", rigctld.HandleGetTwiddle)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/twiddle", rigctld.HandleSetTwiddle)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/cache", rigctld.HandleGetCache)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/cache", rigctld.HandleSetCache)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/state/dump", rigctld.HandleGetStateDump)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/rig-info", rigctld.HandleGetRigInfo)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/modes", rigctld.HandleGetModes)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rigs/{trx_id}/power-state", rigctld.HandleGetPowerState)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/power-state", rigctld.HandleSetPowerState)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rigs/{trx_id}/voice-mem", rigctld.HandleSetVoiceMem)


}

func RegisterRoutesRotctld(mux *http.ServeMux) {
	mux.HandleFunc("GET "+config.ApiBasePath+"/devices/rotators", rotctld.HandleListRotators)
	// Central raw command gateway (analogous to rigctld)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/command", rotctld.HandleRawCommand)

	// Position and movement control
	mux.HandleFunc("GET "+config.ApiBasePath+"/rotators/{rotator_id}/position", rotctld.HandleGetPosition)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/position", rotctld.HandleSetPosition)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/park", rotctld.HandlePark)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/stop", rotctld.HandleStop)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/move", rotctld.HandleMove)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/reset", rotctld.HandleReset)

	// Metadata and hardware state profiles
	mux.HandleFunc("GET "+config.ApiBasePath+"/rotators/{rotator_id}/info", rotctld.HandleGetInfo)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rotators/{rotator_id}/status", rotctld.HandleGetStatus)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rotators/{rotator_id}/state", rotctld.HandleGetState)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rotators/{rotator_id}/capabilities", rotctld.HandleGetCapabilities)

	// Advanced hardware level and auxiliary function mappings
	mux.HandleFunc("GET "+config.ApiBasePath+"/rotators/{rotator_id}/levels/{level}", rotctld.HandleGetLevel)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/levels/{level}", rotctld.HandleSetLevel)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rotators/{rotator_id}/functions/{function_name}", rotctld.HandleGetFunction)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/functions/{function_name}", rotctld.HandleSetFunction)
	mux.HandleFunc("GET "+config.ApiBasePath+"/rotators/{rotator_id}/parameters/{parameter_name}", rotctld.HandleGetParameter)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/parameters/{parameter_name}", rotctld.HandleSetParameter)

	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/start-service", rotctld.HandleStartService)
	mux.HandleFunc("POST "+config.ApiBasePath+"/rotators/{rotator_id}/stop-service", rotctld.HandleStopService)
}
