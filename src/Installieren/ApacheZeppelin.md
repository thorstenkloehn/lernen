# Apache Zeppelin

```bash

wget https://downloads.apache.org/zeppelin/zeppelin-0.9.0/zeppelin-0.9.0-bin-all.tgz
sudo mkdir /zeppelin
tar -C /zeppelin -xzf zeppelin-0.9.0-bin-all.tgz
cd /zeppelin/zeppelin-0.9.0-bin-all
bin/zeppelin-daemon.sh start

/etc/init/zeppelin.conf
-----------------------

description "zeppelin"

start on (local-filesystems and net-device-up IFACE!=lo)
stop on shutdown

# Respawn the process on unexpected termination
respawn

# respawn the job up to 7 times within a 5 second period.
# If the job exceeds these values, it will be stopped and marked as failed.
respawn limit 7 5

# zeppelin was installed in /usr/share/zeppelin in this example
chdir /zeppelin/zeppelin-0.9.0-bin-all
exec bin/zeppelin-daemon.sh upstart

-----------------------

sudo service zeppelin start  
sudo service zeppelin stop  
sudo service zeppelin restart
```

