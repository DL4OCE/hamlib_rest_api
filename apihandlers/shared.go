package apihandlers

import (
	"hamlib_rest_api/apihandlers/rigctld"
	"hamlib_rest_api/apihandlers/rotctld"
	"net/http"
)

func RegisterRoutesRigctld(mux *http.ServeMux) {
	// core routes
	mux.HandleFunc("GET /trx/{trx_id}/frequency", rigctld.HandleGetFrequency)
	mux.HandleFunc("POST /trx/{trx_id}/frequency", rigctld.HandleSetFrequency)
	mux.HandleFunc("GET /trx/{trx_id}/mode", rigctld.HandleGetMode)
	mux.HandleFunc("POST /trx/{trx_id}/mode", rigctld.HandleSetMode)
	mux.HandleFunc("GET /trx/{trx_id}/split_frequency_mode", rigctld.HandleGetSplitFrequencyMode)
	mux.HandleFunc("POST /trx/{trx_id}/split_frequency_mode", rigctld.HandleSetSplitFrequencyMode)
	mux.HandleFunc("GET /trx/{trx_id}/split_frequency", rigctld.HandleGetSplitFrequency)
	mux.HandleFunc("POST /trx/{trx_id}/split_frequency", rigctld.HandleSetSplitFrequency)
	mux.HandleFunc("GET /trx/{trx_id}/split_mode", rigctld.HandleGetSplitMode)
	mux.HandleFunc("POST /trx/{trx_id}/split_mode", rigctld.HandleSetSplitMode)
	mux.HandleFunc("GET /trx/{trx_id}/level/{level_param}", rigctld.HandleGetLevel)
	mux.HandleFunc("POST /trx/{trx_id}/level/{level_param}", rigctld.HandleSetLevel)
	mux.HandleFunc("GET /trx/{trx_id}/level/list", rigctld.HandleGetLevelList)
	mux.HandleFunc("GET /trx/{trx_id}/tuningstep", rigctld.HandleGetTuningStep)
	mux.HandleFunc("POST /trx/{trx_id}/tuningstep", rigctld.HandleSetTuningStep)
	mux.HandleFunc("GET /trx/{trx_id}/split_vfo", rigctld.HandleGetSplitVFO)
	mux.HandleFunc("POST /trx/{trx_id}/split_vfo", rigctld.HandleSetSplitVFO)

	// Functions, Parameters, Scans & Transceive
	mux.HandleFunc("GET /trx/{trx_id}/function/list", rigctld.HandleGetFunctionList)
	mux.HandleFunc("GET /trx/{trx_id}/function/{param}", rigctld.HandleGetFunction)
	mux.HandleFunc("POST /trx/{trx_id}/function/{param}", rigctld.HandleSetFunction)
	mux.HandleFunc("GET /trx/{trx_id}/parameter/list", rigctld.HandleGetParameterList)
	mux.HandleFunc("GET /trx/{trx_id}/parameter/{param}", rigctld.HandleGetParameter)
	mux.HandleFunc("POST /trx/{trx_id}/parameter/{param}", rigctld.HandleSetParameter)
	mux.HandleFunc("GET /trx/{trx_id}/scan/list", rigctld.HandleGetScanList)
	mux.HandleFunc("GET /trx/{trx_id}/scan/{param}", rigctld.HandleGetScan)
	mux.HandleFunc("POST /trx/{trx_id}/scan/{param}", rigctld.HandleSetScan)
	mux.HandleFunc("GET /trx/{trx_id}/transceive/list", rigctld.HandleGetTransceiveList)
	mux.HandleFunc("GET /trx/{trx_id}/transceive", rigctld.HandleGetTransceive)
	mux.HandleFunc("POST /trx/{trx_id}/transceive", rigctld.HandleSetTransceive)

	// Repeater Shift, Offset, Tones & VFO
	mux.HandleFunc("GET /trx/{trx_id}/repeater/shift", rigctld.HandleGetRepeaterShift)
	mux.HandleFunc("POST /trx/{trx_id}/repeater/shift", rigctld.HandleSetRepeaterShift)
	mux.HandleFunc("GET /trx/{trx_id}/repeater/offset", rigctld.HandleGetRepeaterOffset)
	mux.HandleFunc("POST /trx/{trx_id}/repeater/offset", rigctld.HandleSetRepeaterOffset)
	mux.HandleFunc("GET /trx/{trx_id}/tone/ctcss", rigctld.HandleGetCtcssTone)
	mux.HandleFunc("POST /trx/{trx_id}/tone/ctcss", rigctld.HandleSetCtcssTone)
	mux.HandleFunc("GET /trx/{trx_id}/tone/dcs", rigctld.HandleGetDcsTone)
	mux.HandleFunc("POST /trx/{trx_id}/tone/dcs", rigctld.HandleSetDcsTone)
	mux.HandleFunc("GET /trx/{trx_id}/vfo", rigctld.HandleGetVFO)
	mux.HandleFunc("POST /trx/{trx_id}/vfo", rigctld.HandleSetVFO)

	// PTT, Memory, Channels, RIT/XIT & Antenna
	mux.HandleFunc("GET /trx/{trx_id}/ptt", rigctld.HandleGetPTT)
	mux.HandleFunc("POST /trx/{trx_id}/ptt", rigctld.HandleSetPTT)
	mux.HandleFunc("GET /trx/{trx_id}/memory", rigctld.HandleGetMemory)
	mux.HandleFunc("POST /trx/{trx_id}/memory", rigctld.HandleSetMemory)
	mux.HandleFunc("GET /trx/{trx_id}/channel", rigctld.HandleGetChannel)
	mux.HandleFunc("POST /trx/{trx_id}/channel", rigctld.HandleSetChannel)
	mux.HandleFunc("GET /trx/{trx_id}/info", rigctld.HandleGetInfo)
	mux.HandleFunc("GET /trx/{trx_id}/rit", rigctld.HandleGetRit)
	mux.HandleFunc("POST /trx/{trx_id}/rit", rigctld.HandleSetRit)
	mux.HandleFunc("GET /trx/{trx_id}/xit", rigctld.HandleGetXit)
	mux.HandleFunc("POST /trx/{trx_id}/xit", rigctld.HandleSetXit)
	mux.HandleFunc("GET /trx/{trx_id}/antenna", rigctld.HandleGetAntenna)
	mux.HandleFunc("POST /trx/{trx_id}/antenna", rigctld.HandleSetAntenna)

	// Raw Commands, Morse & Power Conversions
	mux.HandleFunc("POST /trx/{trx_id}/raw", rigctld.HandleSetRawCommand)
	mux.HandleFunc("POST /trx/{trx_id}/raw_rx", rigctld.HandleSetRawCommandRx)
	mux.HandleFunc("POST /trx/{trx_id}/power/to_factor", rigctld.HandleGetMwPower)
	mux.HandleFunc("POST /trx/{trx_id}/power/to_mw", rigctld.HandleGetPowerMw)
	mux.HandleFunc("GET /trx/{trx_id}/capabilities", rigctld.HandleGetCapabilities)
	mux.HandleFunc("GET /trx/{trx_id}/configuration", rigctld.HandleGetConfiguration)
	mux.HandleFunc("POST /trx/{trx_id}/morse", rigctld.HandleSetMorse)
	mux.HandleFunc("POST /trx/{trx_id}/morse/stop", rigctld.HandleSetMorseStop)

	// SQL Extensions, Rig State & Misc
	mux.HandleFunc("GET /trx/{trx_id}/sql/ctcss", rigctld.HandleGetCtcssSql)
	mux.HandleFunc("POST /trx/{trx_id}/sql/ctcss", rigctld.HandleSetCtcssSql)
	mux.HandleFunc("GET /trx/{trx_id}/sql/dcs", rigctld.HandleGetDcsSql)
	mux.HandleFunc("POST /trx/{trx_id}/sql/dcs", rigctld.HandleSetDcsSql)
	mux.HandleFunc("GET /trx/{trx_id}/dtmf", rigctld.HandleGetDtmf)
	mux.HandleFunc("POST /trx/{trx_id}/dtmf", rigctld.HandleSetDtmf)
	mux.HandleFunc("GET /trx/{trx_id}/morse/wait", rigctld.HandleGetMorseWait)
	mux.HandleFunc("GET /trx/{trx_id}/dcd", rigctld.HandleGetDcd)
	mux.HandleFunc("GET /trx/{trx_id}/twiddle", rigctld.HandleGetTwiddle)
	mux.HandleFunc("POST /trx/{trx_id}/twiddle", rigctld.HandleSetTwiddle)
	mux.HandleFunc("GET /trx/{trx_id}/cache", rigctld.HandleGetCache)
	mux.HandleFunc("POST /trx/{trx_id}/cache", rigctld.HandleSetCache)
	mux.HandleFunc("POST /trx/{trx_id}/state/dump", rigctld.HandleSetStateDump)
	mux.HandleFunc("GET /trx/{trx_id}/rig_info", rigctld.HandleGetRigInfo)
	mux.HandleFunc("GET /trx/{trx_id}/modes", rigctld.HandleGetModes)
	mux.HandleFunc("GET /trx/{trx_id}/power_state", rigctld.HandleGetPowerState)
	mux.HandleFunc("POST /trx/{trx_id}/power_state", rigctld.HandleSetPowerState)
	mux.HandleFunc("POST /trx/{trx_id}/voice_mem", rigctld.HandleSetVoiceMem)

	// Systemctl Routen für rigctld Dienste
	mux.HandleFunc("POST /trx/{trx_id}/service/start", rigctld.HandleStartRigctld)
	mux.HandleFunc("POST /trx/{trx_id}/service/stop", rigctld.HandleStopRigctld)
	mux.HandleFunc("GET /trxs", rigctld.HandleListRigs)
}

func RegisterRoutesRotctld(mux *http.ServeMux) {
	mux.HandleFunc("GET /rotators", rotctld.HandleListRotators)
	// Central raw command gateway (analogous to rigctld)
	mux.HandleFunc("POST /rotator/{rotator_id}/command", rotctld.HandleRawCommand)

	// Position and movement control
	mux.HandleFunc("GET /rotator/{rotator_id}/position", rotctld.HandleGetPosition)
	mux.HandleFunc("POST /rotator/{rotator_id}/position", rotctld.HandleSetPosition)
	mux.HandleFunc("POST /rotator/{rotator_id}/park", rotctld.HandlePark)
	mux.HandleFunc("POST /rotator/{rotator_id}/stop", rotctld.HandleStop)
	mux.HandleFunc("POST /rotator/{rotator_id}/move", rotctld.HandleMove)
	mux.HandleFunc("POST /rotator/{rotator_id}/reset", rotctld.HandleReset)

	// Metadata and hardware state profiles
	mux.HandleFunc("GET /rotator/{rotator_id}/info", rotctld.HandleGetInfo)
	mux.HandleFunc("GET /rotator/{rotator_id}/status", rotctld.HandleGetStatus)
	mux.HandleFunc("GET /rotator/{rotator_id}/state", rotctld.HandleGetState)
	mux.HandleFunc("GET /rotator/{rotator_id}/capabilities", rotctld.HandleGetCapabilities)

	// Advanced hardware level and auxiliary function mappings
	mux.HandleFunc("GET /rotator/{rotator_id}/level/{level}", rotctld.HandleGetLevel)
	mux.HandleFunc("POST /rotator/{rotator_id}/level/{level}", rotctld.HandleSetLevel)
	mux.HandleFunc("GET /rotator/{rotator_id}/function/{function}", rotctld.HandleGetFunction)
	mux.HandleFunc("POST /rotator/{rotator_id}/function/{function}", rotctld.HandleSetFunction)
	mux.HandleFunc("GET /rotator/{rotator_id}/parameter/{parameter}", rotctld.HandleGetParameter)
	mux.HandleFunc("POST /rotator/{rotator_id}/parameter/{parameter}", rotctld.HandleSetParameter)

	mux.HandleFunc("POST /rotator/{rotator_id}/service/start", rotctld.HandleStartService)
	mux.HandleFunc("POST /rotator/{rotator_id}/service/stop", rotctld.HandleStopService)
}
