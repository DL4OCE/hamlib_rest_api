# hamlib_rest_api

To install ***become root*** and use

```
cd /tmp
git clone https://github.com/deinname/hamlib-rest-api.git
cd hamlib-rest-api

# Modify rigctld.config according to your needs:
editor rigctld.json

chmod +x install.sh update_rigctld_services.sh
./install.sh
./update_rigctld_services.sh

```