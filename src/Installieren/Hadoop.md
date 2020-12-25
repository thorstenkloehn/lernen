# Hadoop

## Ubuntu

```bash

wget https://corretto.aws/downloads/latest/amazon-corretto-8-x64-linux-jdk.tar.gz
mkdir /usr/java
sudo tar -C /usr/java -xzf amazon-corretto-8-x64-linux-jdk.tar.gz

sudo nano /etc/bash.bashrc
_______________
export HDFS_NAMENODE_USER="root"
export HDFS_DATANODE_USER="root"
export HDFS_SECONDARYNAMENODE_USER="root"
export YARN_RESOURCEMANAGER_USER="root"
export YARN_NODEMANAGER_USER="root"
export JAVA_HOME = 
export PATH=$GOPATH/bin:$GOROOT/bin:$JAVA_HOME/bin:$PATH
_______________

sudo apt-get install ssh
sudo apt-get install pdsh
wget https://downloads.apache.org/hadoop/common/hadoop-3.2.1/hadoop-3.2.1.tar.gz
mkdir /hadoop
sudo tar -C /hadoop -xzf hadoop-3.2.1.tar.gz
cd /hadoop/hadoop-3.2.1


```
