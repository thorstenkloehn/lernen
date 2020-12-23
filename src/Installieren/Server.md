# Server

```
sudo apt-get update 
sudo apt-get upgrade
apt-get install  gnupg


```

## Datenbank

```

sudo apt-get install postgresql postgresql-client postgis

```
## systemctl / systemd richtig verwenden

### Erstellen von Dienste

* nano /etc/systemd/system/thorsten.service

* nano /etc/systemd/system/

### Dienste Aktiviereb

* sudo systemctl enable thorsten.service

### Dienste Deaktivieren

* sudo systemctl disable thorsten.service

### Status eines Services

* systemctl status thorsten.service

### Start Service

* systemctl start thorsten.service

### Stop Service

* systemctl stop thorsten.service

### Restart Service

* systemctl restart thorsten.service

### Reload

* systemctl reload thorsten.service
