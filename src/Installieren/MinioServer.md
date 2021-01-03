# Minio Server

## Installieren 

```bash

mkdir minio
cd minio
wget https://dl.min.io/server/minio/release/linux-amd64/minio
chmod +x minio

```

### Einstellung

```bash

nano start.sh

-----------------------------------

MINIO_ACCESS_KEY=test MINIO_SECRET_KEY=test ./minio server /data


-----------------------------------
```

### Starten

```bash

./start.sh
```