# Hadoop

## Ubuntu

```bash

wget https://corretto.aws/downloads/latest/amazon-corretto-8-x64-linux-jdk.tar.gz
mkdir /usr/java
sudo tar -C /usr/java -xzf amazon-corretto-8-x64-linux-jdk.tar.gz
sudo nano /etc/bash.bashrc
_______________
export PDSH_RCMD_TYPE=ssh
export JAVA_HOME=/usr/java/amazon-corretto-8.275.01.1-linux-x64
export PATH=$GOPATH/bin:$GOROOT/bin:$JAVA_HOME/bin:$PATH
_______________
sudo apt-get install ssh
sudo apt-get install pdsh
adduser hadoop
su - hadoop
ssh-keygen -t rsa
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys
su - hadoop
wget https://downloads.apache.org/hadoop/common/hadoop-3.2.1/hadoop-3.2.1.tar.gz
mkdir hadoop
sudo tar -C hadoop -xzf hadoop-3.2.1.tar.gz
cd hadoop/hadoop-3.2.1


```
