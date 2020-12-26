# Hadoop

## Ubuntu

```bash

wget https://corretto.aws/downloads/latest/amazon-corretto-8-x64-linux-jdk.tar.gz
mkdir /usr/java
sudo tar -C /usr/java -xzf amazon-corretto-8-x64-linux-jdk.tar.gz

sudo apt install openssh-server openssh-client -y
sudo adduser hdoop
su - hdoop
ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys
ssh localhost
wget https://downloads.apache.org/hadoop/common/hadoop-3.2.1/hadoop-3.2.1.tar.gz
mkdir hadoop
sudo tar -C hadoop -xzf hadoop-3.2.1.tar.gz
cd hadoop/hadoop-3.2.1

sudo nano /etc/bash.bashrc
_______________
export PDSH_RCMD_TYPE=ssh
#Hadoop Related Options
export HADOOP_HOME=/hadoop/hadoop-3.2.1
export PATH=$GOPATH/bin:$GOROOT/bin:$JAVA_HOME/bin:$PATH
_______________


```
