#!/bin/bash

for id in $(jq -r 'keys[]' /etc/hamlib_rest_api/rigctld.json); do
    echo "Aktiviere TRX Instanz $id..."
    sudo systemctl daemon-reload
    sudo systemctl enable --now rigctld@$id.service
done