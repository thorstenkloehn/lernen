# Linux Desktop

## lubuntu Installieren

* EFI System 200mb efi fat32
* rest mit / ext4


## Browser

* [Google Chrome](https://www.google.de/chrome)
* [Firefox](https://www.mozilla.org/de/firefox/developer)


## Git

```

sudo apt-get install git-core
sudo apt-get install curl

```

## Programmiersprachen


* [Nodjs](https://github.com/nodesource/distributions/blob/master/README.md)

### Go

```
wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
sudo nano /etc/bash.bashrc
----Datei----

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
---Datei ende---
source  /etc/bash.bashrc

```

### C und C++

``` 
sudo apt update
sudo apt upgrade
sudo apt-get install cmake gcc clang gdb build-essential git-core


```





## IDE

```

sudo snap install clion --classic
sudo snap install intellij-idea-ultimate --classic


```

## Datenbank
### Postgresql

```

sudo apt-get install postgresql postgresql-client postgis


```

* [Weiter Lesen](https://wiki.ubuntuusers.de/PostgreSQL)


### Sqlite

```
sudo apt-get install sqlite3   sqlite3-doc 


```