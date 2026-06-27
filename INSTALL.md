# hamlib_rest_api


```
git clone https://github.com/DL4OCE/hamlib_rest_api.git /tmp/hamlib_rest_api
cd /tmp/hamlib_rest_api

# Modify rigctld.config according to your needs:
editor config/rigctld.json
editor config/rotctld.json

sudo chmod +x install.sh update_hamlib_services.sh
sudo ./install.sh
# use
# sudo ./update_hamlib_services.sh
# whenever you updated the config/rigctld.json!

```

After installation, you can find a GUI to start / stop services at http://localhost:8080/gui/control_services.html

