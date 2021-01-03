# Server

```
sudo apt-get update 
sudo apt-get upgrade
apt-get install  gnupg

```
### JAVA 

```

wget https://corretto.aws/downloads/latest/amazon-corretto-8-x64-linux-jdk.tar.gz
mkdir /usr/java
sudo tar -C /usr/java -xzf amazon-corretto-8-x64-linux-jdk.tar.gz
```
### Go

```
wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz

```

### bash.bashrc

```
sudo nano /etc/bash.bashrc
----Datei----
export PDSH_RCMD_TYPE=ssh
export JAVA_HOME=/usr/java/amazon-corretto-8.275.01.1-linux-x64
export GOROOT=/usr/local/go
export GOPATH=/go
export PATH=$GOPATH/bin:$GOROOT/bin:$JAVA_HOME/bin:$PATH
---Datei ende---
source  /etc/bash.bashrc

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
