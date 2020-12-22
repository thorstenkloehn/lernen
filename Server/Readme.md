# Installieren auf Server

```
sudo apt-get update && sudo apt-get upgrade
```
## Go Installieren
```
wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
sudo nano /etc/bash.bashrc
----Datei----

export GOROOT=/usr/local/go
export GOPATH=/Server/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
---Datei ende---
source  /etc/bash.bashrc
```

## Rust Installieren

* [Rust Downloads](https://www.rust-lang.org/tools/install)

## mdbook Installieren

```
cargo install mdbook

```

## Server Installieren

```

go get github.com/thorstenkloehn/Server
cd $GOPATH/src/github.com/thorstenkloehn/Server
bash install.sh

```

## Aktualisieren

Um Submodule zu aktualisieren, können wir verwenden

```
git submodule update --recursive --remote
```

oder

```

git pull --recurse-submodules


```

## Submodule hinzufügen

```

git submodule add https://github.com/thorstenkloehn/webassemblygo tinygo


```

## Submodule löschen

```

git submodule deinit submodule_name
git rm submodule_name


```

