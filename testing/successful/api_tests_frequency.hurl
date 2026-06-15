### --- TC: get frequency --- ###

## --> store original value
GET http://localhost:8080/trx/1/frequency

HTTP 200
[Asserts]
jsonpath "$.freq" matches /^[0-9]+$/
[Captures]
original_frequency: jsonpath "$.freq"

## --> set random frequency
POST http://localhost:8080/trx/1/frequency
Content-Type: application/json
#[Options]
#variable: tmp_frequency = "145{{randomInt 100 999}}000"
{
  "newValue": "{{tmp_frequency}}"
}

HTTP 200

[Asserts]
jsonpath "$" isEmpty
#[Captures]
#tmp_freq: request jsonpath "$.newValue"

# has frequency been set, correctly?
GET http://localhost:8080/trx/1/frequency

HTTP 200
[Asserts]
jsonpath "$.freq" == "{{tmp_frequency}}"


## --> set original frequency
POST http://localhost:8080/trx/1/frequency
Content-Type: application/json
{
  "newValue": "{{original_frequency}}"
}

HTTP 200

[Asserts]
jsonpath "$" isEmpty

# has frequency been set, correctly?
GET http://localhost:8080/trx/1/frequency

HTTP 200
[Asserts]
jsonpath "$.freq" == "{{original_frequency}}"









# ==============================================================================
# --- TC: get / set split frequency ---
# ==============================================================================

## --> store original split value
GET http://localhost:8080/trx/1/split_frequency

HTTP 200
[Captures]
original_split_freq: jsonpath "$.qrg"

## --> set new split frequency (using the same external bash variable logic)
POST http://localhost:8080/trx/1/split_frequency
Content-Type: application/json
{
  "newValue": "{{tmp_frequency}}"
}

HTTP 200
[Asserts]
jsonpath "$" isEmpty

# has split frequency been set, correctly?
GET http://localhost:8080/trx/1/split_frequency

HTTP 200
[Asserts]
jsonpath "$.qrg" == "{{tmp_frequency}}"

## --> restore original split frequency
POST http://localhost:8080/trx/1/split_frequency
Content-Type: application/json
{
  "newValue": "{{original_split_freq}}"
}

HTTP 200
[Asserts]
jsonpath "$" isEmpty

# verify restore
GET http://localhost:8080/trx/1/split_frequency

HTTP 200
[Asserts]
jsonpath "$.qrg" == "{{original_split_freq}}"