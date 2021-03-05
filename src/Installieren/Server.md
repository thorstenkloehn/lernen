# Server

```
sudo apt-get update 
sudo apt-get upgrade
apt-get install  gnupg

```


## Go

```
wget https://golang.org/dl/go1.15.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.15.7.linux-amd64.tar.gz

```

## bash.bashrc

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
## MySQL

```
wget -c https://dev.mysql.com/get/mysql-apt-config_0.8.16-1_all.deb
sudo dpkg -i mysql-apt-config*
sudo apt update
sudo apt-get install mysql-server

```


## Apache Zeppelin 

```
mkdir /zeppelin
sudo apt-get install openjdk-8-jdk-headless 
wget https://downloads.apache.org/zeppelin/zeppelin-0.9.0/zeppelin-0.9.0-bin-all.tgz
sudo tar -C /zeppelin -xzf zeppelin-0.9.0-bin-all.tgz
cd /zeppelin/zeppelin-0.9.0-bin-all
cp conf/shiro.ini.template conf/shiro.ini
cd conf
nano shiro.ini
cd ..
bin/zeppelin-daemon.sh start


```

## Python Installieren

```
sudo apt install python3 python3-dev git curl python-is-python3 

```
## JupyterHub

```
curl -L https://tljh.jupyter.org/bootstrap.py | sudo -E python3 - --admin user:password
sudo tljh-config set http.port 8088
sudo tljh-config reload proxy


```

## PostgreSQL

```
sudo apt-get install postgresql postgresql-client postgis
```

### monodb

* [Download](https://https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/)

### Neo4j

* [Download](https://neo4j.com/download-center/?ref=web-product-database/#community)

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
