

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

```
# Verification command for dazzling-remora
docker exec dazzling-remora /opt/kafka/bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```


## Create a Topic
Following command creates a topic:
```
docker exec dazzling-remora /opt/kafka/bin/kafka-topics.sh --create --topic events --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

### Output
```
Created topic events.
```


## Describe that Topic
The following command creates a topic:
```
docker exec dazzling-remora /opt/kafka/bin/kafka-topics.sh --describe --topic events --bootstrap-server localhost:9092
```
### Output
```
Topic: events   TopicId: xNBVu3MYSvG9WkbexUOW2w PartitionCount: 1       ReplicationFactor: 1    Configs: min.insync.replicas=1
        Topic: events   Partition: 0    Leader: 1       Replicas: 1     Isr: 1  Elr:    LastKnownElr: 
```

