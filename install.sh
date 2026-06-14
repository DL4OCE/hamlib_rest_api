#!/bin/bash

apt-get update && sudo apt-get install -y libhamlib-utils jq
cp rigctld@.service /etc/systemd/system/
cp instances.json /etc/hamlib_rest_api
ACTIVE_SERVICES=$(systemctl list-units --all --plain --no-legend "rigctld@*" | awk '{print $1}')
for service in $ACTIVE_SERVICES; do
    echo "  -> Stoppe und deaktiviere Dienst: $service"
    systemctl stop "$service"
    systemctl disable "$service"
done

cp hamlib_rest_api.service /etc/systemd/system
systemctl enable --now hamlib_rest_api.service
echo "Please modify /etc/hamlib_rest_api/rigctld.config according to your needs and run update_rigctld_instances.sh"

