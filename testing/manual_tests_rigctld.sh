#!/bin/bash

test_get_api() {
    echo -e "curl -s -XGET $1 $2"
    curl -s -X GET "$1"
    echo -e "------------------------------------------------------------------------------------------------------------------------------------------------------"
}

test_post_api() {
    echo -e "curl -s -XPOST $1 -H \"Content-Type: application/json\" -d \"$3\" $2"
    if [ -n "$3" ]; then
        curl -s -X POST "$1" -H "Content-Type: application/json" -d "$3"
    else
        curl -s -X POST "$1"
    fi
    echo -e "------------------------------------------------------------------------------------------------------------------------------------------------------"
}

# ==============================================================================
# CONFIGURATION
# ==============================================================================
API_URL="http://localhost:8080"
API_BASE="/api/v1"
RIG_BASE_URL="$API_URL$API_BASE/rigs/1"
ROTATOR_BASE_URL="$API_URL$API_BASE/rotators/1"
HEADER="Content-Type: application/json"

echo "=== STARTING HAMLIB REST API MANUAL CURL TESTS ==="
echo "Target API URL: $API_URL"
echo "Target TRX URL: $BASE_URL"
echo "--------------------------------------------------"

# ==============================================================================
# === NEW === SYSTEM & SERVICE MANAGEMENT (GLOBAL ENDPOINTS)
# ==============================================================================
echo -e "[0. System & Service Management]"
test_get_api $API_URL$API_BASE/devices/rigs "(list rigs (transceivers))"
test_get_api $API_URL$API_BASE/devices/rotators "(list rotators)"
test_post_api $RIG_BASE_URL/service/stop "(Stop TRX 1)"
test_post_api $RIG_BASE_URL/service/start "(Start TRX 1)"
test_post_api $ROTATOR_BASE_URL/service/stop "(Stop Rotator 1)"
test_post_api $ROTATOR_BASE_URL/service/start "(Start Rotator 1)"
sleep 2

# ==============================================================================
# 1. CORE FUNCTIONS (FREQUENCY & MODE)
# ==============================================================================

echo -e "\n[1. Frequency & Mode]"
test_get_api $RIG_BASE_URL/frequency "(Get current frequency)"
test_post_api $RIG_BASE_URL/frequency "(Set frequency to 145.500 MHz)" '{"newValue":"145500000"}'
test_get_api $RIG_BASE_URL/mode "(Get current mode)"
test_post_api $RIG_BASE_URL/mode "(Set mode to USB with 2.4kHz passband)" '{"mode":"USB","passband":"2400"}'


# ==============================================================================
# 2. SPLIT REQS
# ==============================================================================
echo -e "\n[2. Split Settings]"
test_get_api $RIG_BASE_URL/split_frequency "(Get current split frequency)"
test_post_api $RIG_BASE_URL/split_frequency "(Set split frequency to 145.525 MHz)" '{"newValue":"145525000"}'

test_get_api $RIG_BASE_URL/split_mode "(Get current split mode)"
test_post_api $RIG_BASE_URL/split_mode "(Set split mode to LSB with 2.7kHz passband)" '{"mode":"LSB","passband":"2700"}'

test_get_api $RIG_BASE_URL/split_frequency_mode "(Get current split mode)"
test_post_api $RIG_BASE_URL/split_frequency_mode "(Set split mode to CW with 500 Hz passband)" '{"frequency":"145425000","mode":"CW","passband":"500"}'

test_get_api $RIG_BASE_URL/split_vfo "(Get current VFO mode)"
test_post_api $RIG_BASE_URL/split_vfo "(Set VFO mode to split enabled and active VFO = VFO B)" '{"split_mode":"1","tx_vfo":"VFOB"}'

# ==============================================================================
# 3. TUNING & LEVELS
# ==============================================================================
echo -e "\n\n[3. Tuning & Levels]"
test_get_api $RIG_BASE_URL/tuningstep "(Get tuning step size)"
test_post_api $RIG_BASE_URL/tuningstep "(Set tuning step size)" '{"newValue":"500"}'

test_get_api $RIG_BASE_URL/level/list "(Get list of stored levels)"
test_get_api $RIG_BASE_URL/level/RFPOWER "(Get current TX power)"
test_post_api $RIG_BASE_URL/level/RFPOWER "(Set TX power)" '{"newValue":"0.500000"}'

# ==============================================================================
# 4. RIG CAPABILITIES (FUNCTIONS, PARAMS, SCANS, TRANSCEIVE)
# ==============================================================================
echo -e "\n\n[4. Functions, Parameters & Scans]"

test_get_api $RIG_BASE_URL/function/list "(Get list of functions)"
test_get_api $RIG_BASE_URL/function/RIT "(Get current RIT)"
test_post_api $RIG_BASE_URL/function/RIT "(Set current RIT)" '{"newValue":"1"}'

test_get_api $RIG_BASE_URL/parameter/list "(Get list of parameters)"
test_get_api $RIG_BASE_URL/parameter/APO "(Get current APO duration value)"
test_post_api $RIG_BASE_URL/parameter/APO "(Get current APO duration value)" '{"newValue":"60"}'

test_get_api $RIG_BASE_URL/scan/list "(Get list of scan channels)"
# test_get_api $RIG_BASE_URL/scan/MEM "(Is memory scan mode enabled)"
# test_post_api $RIG_BASE_URL/scan/MEM "(Set memory scan mode)" '{"newValue":"1"}'

test_get_api $RIG_BASE_URL/transceive/list "(Get list of scan channels)"
test_get_api $RIG_BASE_URL/transceive "(Is transceive mode enabled)"
test_post_api $RIG_BASE_URL/transceive "(Enable / disable transceive mode)" '{"newValue":"1"}'

# ==============================================================================
# 5. REPEATER, TONES & VFO
# ==============================================================================
echo -e "\n\n[5. Repeater & Tones]"

test_get_api $RIG_BASE_URL/repeater/shift "(Get repeater shift direction)"
test_post_api $RIG_BASE_URL/repeater/shift "(Set repeater shift direction)" '{"newValue":"+"}'

test_get_api $RIG_BASE_URL/repeater/offset "(Get repeater offset)"
test_post_api $RIG_BASE_URL/repeater/offset "(Set repeater offset)" '{"newValue":"600000"}'

test_get_api $RIG_BASE_URL/tone/ctcss "(Get repeater CTCSS tone)"
test_post_api $RIG_BASE_URL/tone/ctcss "(Set repeater CTCSS tone)" '{"newValue":"88.5"}'

test_get_api $RIG_BASE_URL/tone/dcs "(Get repeater DCS tone)"
test_post_api $RIG_BASE_URL/tone/dcs "(Set repeater DCS tone)" '{"newValue":"023"}'

test_get_api $RIG_BASE_URL/vfo "(Get current VFO)"
test_post_api $RIG_BASE_URL/vfo "(Set current VFO to VFOA)" '{"newValue":"VFOA"}'

# ==============================================================================
# 6. HARDWARE STATES (PTT, MEMORY, CHANNEL, ANTENNA)
# ==============================================================================
echo -e "\n\n[6. Hardware States]"

test_get_api $RIG_BASE_URL/ptt "(Get current PTT state)"
test_post_api $RIG_BASE_URL/ptt "(Set PTT state to 0)" '{"newValue":"0"}'

test_get_api $RIG_BASE_URL/memory "(Get current memory state)"
test_post_api $RIG_BASE_URL/memory "(Set memory state to 1)" '{"newValue":"1"}'

# test_get_api $RIG_BASE_URL/channel "(Get current channel state)"
# test_post_api $RIG_BASE_URL/channel "(Set channel state to 12)" '{"newValue":"12"}'

test_get_api $RIG_BASE_URL/info "(Get current rig info)"

test_get_api $RIG_BASE_URL/rit "(Get current RIT state)"
test_post_api $RIG_BASE_URL/rit "(Set RIT state to 0)" '{"newValue":"0"}'

test_get_api $RIG_BASE_URL/xit "(Get current XIT state)"
test_post_api $RIG_BASE_URL/xit "(Set XIT state to 0)" '{"newValue":"0"}'

# test_get_api $RIG_BASE_URL/antenna "(Get current antenna state)"
# test_post_api $RIG_BASE_URL/antenna "(Set antenna state to A)" '{"newValue":"A"}'

# ==============================================================================
# 7. MORSE, RIG INFO & CACHE
# ==============================================================================
echo -e "\n\n[7. Morse, Rig State & Conversions]"

test_post_api $RIG_BASE_URL/morse "(Send Morse code: CQ CQ DE DL4OCE)" '{"text":"CQ CQ DE DL4OCE"}'
test_post_api $RIG_BASE_URL/morse/stop "(Stop Morse code)"

# test_get_api $RIG_BASE_URL/morse/wait "(Get Morse wait time)"
# test_get_api $RIG_BASE_URL/dcd "(Get DCD state)"

# test_get_api $RIG_BASE_URL/twiddle "(Get twiddle state)"
test_post_api $RIG_BASE_URL/twiddle "(Set twiddle state to 1)" '{"newValue":"1"}'

# test_get_api $RIG_BASE_URL/cache "(Get cache state)"
test_post_api $RIG_BASE_URL/cache "(Set cache state to 10)" '{"newValue":"10"}'
test_get_api $RIG_BASE_URL/capabilities "(Get rig capabilities)"
test_get_api $RIG_BASE_URL/configuration "(Get rig configuration)"
test_post_api $RIG_BASE_URL/state/dump "(Dump rig state)" '{}'
# test_get_api $RIG_BASE_URL/rig_info "(Get rig info)"

test_get_api $RIG_BASE_URL/modes "(Get list of supported modes)"
# test_get_api $RIG_BASE_URL/power_state "(Get current power state)"
test_post_api $RIG_BASE_URL/power_state "(Set power state to 1)" '{"newValue":"1"}'

# ==============================================================================
# 8. SQL EXTENSIONS & DTMF
# ==============================================================================
echo -e "\n\n[8. SQL Extensions & DTMF]"
# test_get_api $RIG_BASE_URL/sql/ctcss "(Get current CTCSS tone)"
test_post_api $RIG_BASE_URL/sql/ctcss "(Set CTCSS tone to 1)" '{"newValue":"1"}'

# test_get_api $RIG_BASE_URL/sql/dcs "(Get current DCS tone)"
test_post_api $RIG_BASE_URL/sql/dcs "(Set DCS tone to 1)" '{"newValue":"1"}'

test_get_api $RIG_BASE_URL/dtmf "(Get current DTMF state)"
test_post_api $RIG_BASE_URL/dtmf "(Set DTMF state to 123)" '{"newValue":"123"}'
test_post_api $RIG_BASE_URL/voice_mem "(Set voice memory to 1)" '{"newValue":"1"}'

# ==============================================================================
# 9. SPECIAL POWER CONVERSIONS (BODY PAYLOAD CONVERTED TO POST)
# ==============================================================================
echo -e "\n\n[9. Power Conversion Endpoints]"
test_post_api $RIG_BASE_URL/power/to_factor "(Convert power in mW to power factor)" '{"power_mW":"50000", "frequency":"145500000", "mode":"USB"}'
test_post_api $RIG_BASE_URL/power/to_mw "(Convert power factor to power in mW)" '{"power_factor":"0.50", "frequency":"145500000", "mode":"USB"}'

# ==============================================================================
# 10. RAW COMMANDS (USE WITH CAUTION)
# ==============================================================================
#echo -e "\n\n[10. Raw Commands]"
# test_post_api $RIG_BASE_URL/raw "(Send raw command: f)" '{"raw_command":"f"}'
# test_post_api $RIG_BASE_URL/raw_rx "(Send raw command: F 145500000)" '{"raw_command":"F 145500000", "number_of_expected_rx_bytes":"0"}'

echo -e "\n\n=== TESTS COMPLETE ==="