#!/bin/bash
hurl --test successful/api_tests_frequency.hurl --variable tmp_frequency="145$((100 + RANDOM % 900))000"
hurl --test successful/api_tests_mode.hurl


hurl --test failing/api_tests_failing.hurl 
