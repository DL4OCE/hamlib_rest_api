# ==============================================================================
### --- TC: get / set mode --- ###
# ==============================================================================

## --> store original values

GET http://localhost:8080/trx/1/mode

HTTP 200
[Captures]
original_mode: jsonpath "$.mode"
original_passband: jsonpath "$.passband"


## --> set mode and passband

POST http://localhost:8080/trx/1/mode
Content-Type: application/json
{
  "mode": "LSB",
  "passband": "2700"
}

HTTP 200
[Asserts]
jsonpath "$" isEmpty


## --> test if values set, correctly

GET http://localhost:8080/trx/1/mode

HTTP 200
[Asserts]
jsonpath "$.mode" == "L"
jsonpath "$.passband" == "S"


# restore original values

POST http://localhost:8080/trx/1/mode
Content-Type: application/json
{
  "mode": "{{original_mode}}",
  "passband": "{{original_passband}}"
}

HTTP 200
[Asserts]
jsonpath "$" isEmpty


## --> test if values set, correctly

GET http://localhost:8080/trx/1/mode

HTTP 200
[Asserts]
jsonpath "$.mode" == "{{original_mode}}"
jsonpath "$.passband" == "{{original_passband}}"


# ==============================================================================
# --- TC: get / set split mode ---
# ==============================================================================

## --> store original split mode values
GET http://localhost:8080/trx/1/split_mode

HTTP 200
[Captures]
original_split_mode: jsonpath "$.mode"
original_split_passband: jsonpath "$.passband"

## --> set split mode and passband
POST http://localhost:8080/trx/1/split_mode
Content-Type: application/json
{
  "mode": "USB",
  "passband": "2400"
}

HTTP 200
[Asserts]
jsonpath "$" isEmpty

## --> test if split values set, correctly
GET http://localhost:8080/trx/1/split_mode

HTTP 200
[Asserts]
jsonpath "$.mode" == "U"
jsonpath "$.passband" == "S"

## --> restore original split mode values
POST http://localhost:8080/trx/1/split_mode
Content-Type: application/json
{
  "mode": "{{original_split_mode}}",
  "passband": "{{original_split_passband}}"
}

HTTP 200
[Asserts]
jsonpath "$" isEmpty


# ==============================================================================
# --- TC: get / set split frequency and mode combined ---
# ==============================================================================

## --> store original combined values
GET http://localhost:8080/trx/1/split_frequency_mode

HTTP 200
[Captures]
orig_comb_freq: jsonpath "$.frequency"
orig_comb_mode: jsonpath "$.mode"
orig_comb_passband: jsonpath "$.passband"

## --> set combined split parameters
POST http://localhost:8080/trx/1/split_frequency_mode
Content-Type: application/json
{
  "frequency": "145500000",
  "mode": "CW",
  "passband": "500"
}

HTTP 200
[Asserts]
jsonpath "$" isEmpty

## --> test if combined values set, correctly
GET http://localhost:8080/trx/1/split_frequency_mode

HTTP 200
[Asserts]
jsonpath "$.frequency" == "145500000"
jsonpath "$.mode" == "4"
jsonpath "$.passband" == "5"

## --> restore original combined values
POST http://localhost:8080/trx/1/split_frequency_mode
Content-Type: application/json
{
  "frequency": "{{orig_comb_freq}}",
  "mode": "{{orig_comb_mode}}",
  "passband": "{{orig_comb_passband}}"
}

HTTP 200
[Asserts]
jsonpath "$" isEmpty


## --> test if combined values set, correctly
GET http://localhost:8080/trx/1/split_frequency_mode

HTTP 200
[Asserts]
jsonpath "$.frequency" == "145500000"
jsonpath "$.mode" == "4"
jsonpath "$.passband" == "5"