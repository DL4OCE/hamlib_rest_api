# hamlib_rest_api


```
cd /tmp
git clone https://github.com/DL4OCE/hamlib_rest_api.git
cd hamlib_rest_api

# Modify rigctld.config according to your needs:
editor rigctld.json

sudo chmod +x install.sh update_rigctld_services.sh
sudo ./install.sh
# use
# sudo ./update_rigctld_services.sh
# whenever you updated the rigctld.json!

```

After installation, you can find a rigctld GUI at http://localhost:8080/gui/config_rigctld.html

