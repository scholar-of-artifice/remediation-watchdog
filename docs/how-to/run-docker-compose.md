

Building the services with docker compose:
```bash
remediation-watchdog % docker compose up -d
```

Making sure the connection is established:

```bash
remediation-watchdog % docker exec -it bashful-yak redis-cli ping
PONG
```

Tearing things down:
```bash
docker compose down
docker compose down --volumes # for cleaning up the data too
```


