# Server

```
sudo apt-get update 
sudo apt-get upgrade
apt-get install  gnupg


```

## Datenbank

### Postgresql
```

sudo apt-get install postgresql postgresql-client postgis

```
### Mysql

```
wget -c https://dev.mysql.com/get/mysql-apt-config_0.8.13-1_all.deb
apt-get install  gnupg
sudo dpkg -i mysql-apt-config*
sudo apt update
sudo apt-get install mysql-server
sudo systemctl enable mysql
sudo nano /etc/mysql/mysql.conf.d/mysqld.cnf
-------------------Datei mysqld.cnf --------------
[mysqld]
bind-address = 0.0.0.0
-------------------Datei mysqd ende----------------
mysql -u root -p
use mysql;
update user set host='%' where user='eigene Benutzername';
update db set host='%' where user='euer_benutzer';
sudo service mysql restart

```
### monodb 

* [Download](https://https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/)

### Cassandra

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
