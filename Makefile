up:
	docker compose up -d
log:
	docker compose logs -f broker
topic-list:
	docker compose exec broker /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server localhost:9092 --list
create-topic:
		docker compose exec broker /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server localhost:9092 \
		--create --topic orders --partitions 3 --replication-factor 1
describe-topic:
	docker compose exec broker /opt/kafka/bin/kafka-topics.sh \
		--bootstrap-server localhost:9092 \
		--describe --topic orders
produce:
	docker compose exec broker /opt/kafka/bin/kafka-console-producer.sh \
		--bootstrap-server localhost:9092 \
		--topic orders
consume:
	docker compose exec broker /opt/kafka/bin/kafka-console-consumer.sh \
		--bootstrap-server localhost:9092 \
		--topic orders --from-beginning
