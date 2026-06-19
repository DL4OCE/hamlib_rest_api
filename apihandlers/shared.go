package apihandlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Core Routen
	mux.HandleFunc("GET /trx/{trx_id}/frequency", handleGetFrequency)
	mux.HandleFunc("POST /trx/{trx_id}/frequency", handleSetFrequency)
	mux.HandleFunc("GET /trx/{trx_id}/mode", handleGetMode)
	mux.HandleFunc("POST /trx/{trx_id}/mode", handleSetMode)
	mux.HandleFunc("GET /trx/{trx_id}/split_frequency_mode", handleGetSplitFrequencyMode)
	mux.HandleFunc("POST /trx/{trx_id}/split_frequency_mode", handleSetSplitFrequencyMode)
	mux.HandleFunc("GET /trx/{trx_id}/split_frequency", handleGetSplitFrequency)
	mux.HandleFunc("POST /trx/{trx_id}/split_frequency", handleSetSplitFrequency)
	mux.HandleFunc("GET /trx/{trx_id}/split_mode", handleGetSplitMode)
	mux.HandleFunc("POST /trx/{trx_id}/split_mode", handleSetSplitMode)
	mux.HandleFunc("GET /trx/{trx_id}/level/{level_param}", handleGetLevel)
	mux.HandleFunc("POST /trx/{trx_id}/level/{level_param}", handleSetLevel)
	mux.HandleFunc("GET /trx/{trx_id}/level/list", handleGetLevelList)
	mux.HandleFunc("GET /trx/{trx_id}/tuningstep", handleGetTuningStep)
	mux.HandleFunc("POST /trx/{trx_id}/tuningstep", handleSetTuningStep)
	mux.HandleFunc("GET /trx/{trx_id}/split_vfo", handleGetSplitVFO)
	mux.HandleFunc("POST /trx/{trx_id}/split_vfo", handleSetSplitVFO)

	// Functions, Parameters, Scans & Transceive
	mux.HandleFunc("GET /trx/{trx_id}/function/list", handleGetFunctionList)
	mux.HandleFunc("GET /trx/{trx_id}/function/{param}", handleGetFunction)
	mux.HandleFunc("POST /trx/{trx_id}/function/{param}", handleSetFunction)
	mux.HandleFunc("GET /trx/{trx_id}/parameter/list", handleGetParameterList)
	mux.HandleFunc("GET /trx/{trx_id}/parameter/{param}", handleGetParameter)
	mux.HandleFunc("POST /trx/{trx_id}/parameter/{param}", handleSetParameter)
	mux.HandleFunc("GET /trx/{trx_id}/scan/list", handleGetScanList)
	mux.HandleFunc("GET /trx/{trx_id}/scan/{param}", handleGetScan)
	mux.HandleFunc("POST /trx/{trx_id}/scan/{param}", handleSetScan)
	mux.HandleFunc("GET /trx/{trx_id}/transceive/list", handleGetTransceiveList)
	mux.HandleFunc("GET /trx/{trx_id}/transceive", handleGetTransceive)
	mux.HandleFunc("POST /trx/{trx_id}/transceive", handleSetTransceive)

	// Repeater Shift, Offset, Tones & VFO
	mux.HandleFunc("GET /trx/{trx_id}/repeater/shift", handleGetRepeaterShift)
	mux.HandleFunc("POST /trx/{trx_id}/repeater/shift", handleSetRepeaterShift)
	mux.HandleFunc("GET /trx/{trx_id}/repeater/offset", handleGetRepeaterOffset)
	mux.HandleFunc("POST /trx/{trx_id}/repeater/offset", handleSetRepeaterOffset)
	mux.HandleFunc("GET /trx/{trx_id}/tone/ctcss", handleGetCtcssTone)
	mux.HandleFunc("POST /trx/{trx_id}/tone/ctcss", handleSetCtcssTone)
	mux.HandleFunc("GET /trx/{trx_id}/tone/dcs", handleGetDcsTone)
	mux.HandleFunc("POST /trx/{trx_id}/tone/dcs", handleSetDcsTone)
	mux.HandleFunc("GET /trx/{trx_id}/vfo", handleGetVFO)
	mux.HandleFunc("POST /trx/{trx_id}/vfo", handleSetVFO)

	// PTT, Memory, Channels, RIT/XIT & Antenna
	mux.HandleFunc("GET /trx/{trx_id}/ptt", handleGetPTT)
	mux.HandleFunc("POST /trx/{trx_id}/ptt", handleSetPTT)
	mux.HandleFunc("GET /trx/{trx_id}/memory", handleGetMemory)
	mux.HandleFunc("POST /trx/{trx_id}/memory", handleSetMemory)
	mux.HandleFunc("GET /trx/{trx_id}/channel", handleGetChannel)
	mux.HandleFunc("POST /trx/{trx_id}/channel", handleSetChannel)
	mux.HandleFunc("GET /trx/{trx_id}/info", handleGetInfo)
	mux.HandleFunc("GET /trx/{trx_id}/rit", handleGetRit)
	mux.HandleFunc("POST /trx/{trx_id}/rit", handleSetRit)
	mux.HandleFunc("GET /trx/{trx_id}/xit", handleGetXit)
	mux.HandleFunc("POST /trx/{trx_id}/xit", handleSetXit)
	mux.HandleFunc("GET /trx/{trx_id}/antenna", handleGetAntenna)
	mux.HandleFunc("POST /trx/{trx_id}/antenna", handleSetAntenna)

	// Raw Commands, Morse & Power Conversions
	mux.HandleFunc("POST /trx/{trx_id}/raw", handleSetRawCommand)
	mux.HandleFunc("POST /trx/{trx_id}/raw_rx", handleSetRawCommandRx)
	mux.HandleFunc("POST /trx/{trx_id}/power/to_factor", handleGetMwPower)
	mux.HandleFunc("POST /trx/{trx_id}/power/to_mw", handleGetPowerMw)
	mux.HandleFunc("GET /trx/{trx_id}/capabilities", handleGetCapabilities)
	mux.HandleFunc("GET /trx/{trx_id}/configuration", handleGetConfiguration)
	mux.HandleFunc("POST /trx/{trx_id}/morse", handleSetMorse)
	mux.HandleFunc("POST /trx/{trx_id}/morse/stop", handleSetMorseStop)

	// SQL Extensions, Rig State & Misc
	mux.HandleFunc("GET /trx/{trx_id}/sql/ctcss", handleGetCtcssSql)
	mux.HandleFunc("POST /trx/{trx_id}/sql/ctcss", handleSetCtcssSql)
	mux.HandleFunc("GET /trx/{trx_id}/sql/dcs", handleGetDcsSql)
	mux.HandleFunc("POST /trx/{trx_id}/sql/dcs", handleSetDcsSql)
	mux.HandleFunc("GET /trx/{trx_id}/dtmf", handleGetDtmf)
	mux.HandleFunc("POST /trx/{trx_id}/dtmf", handleSetDtmf)
	mux.HandleFunc("GET /trx/{trx_id}/morse/wait", handleGetMorseWait)
	mux.HandleFunc("GET /trx/{trx_id}/dcd", handleGetDcd)
	mux.HandleFunc("GET /trx/{trx_id}/twiddle", handleGetTwiddle)
	mux.HandleFunc("POST /trx/{trx_id}/twiddle", handleSetTwiddle)
	mux.HandleFunc("GET /trx/{trx_id}/cache", handleGetCache)
	mux.HandleFunc("POST /trx/{trx_id}/cache", handleSetCache)
	mux.HandleFunc("POST /trx/{trx_id}/state/dump", handleSetStateDump)
	mux.HandleFunc("GET /trx/{trx_id}/rig_info", handleGetRigInfo)
	mux.HandleFunc("GET /trx/{trx_id}/modes", handleGetModes)
	mux.HandleFunc("GET /trx/{trx_id}/power_state", handleGetPowerState)
	mux.HandleFunc("POST /trx/{trx_id}/power_state", handleSetPowerState)
	mux.HandleFunc("POST /trx/{trx_id}/voice_mem", handleSetVoiceMem)

	// Systemctl Routen für rigctld Dienste
	mux.HandleFunc("POST /trx/{trx_id}/start", handleStartRigctld)
	mux.HandleFunc("POST /trx/{trx_id}/stop", handleStopRigctld)
	mux.HandleFunc("GET /trxs", HandleListRigs)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func pollTrx(trxID int, command string) (string, error) {
	targetPort := 4532 + trxID
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", targetPort))
	if err != nil {
		return "", fmt.Errorf("rigctld auf Port %d nicht erreichbar", targetPort)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", strings.TrimSpace(command))

	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(resp), nil
}
