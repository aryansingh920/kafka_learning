# Create topic 
docker exec -it kafka-learning /opt/kafka/bin/kafka-topics.sh \
  --create --topic test-topic \
  --bootstrap-server localhost:9092 \
  --partitions 1 \
  --replication-factor 1

#  Producer sending data
docker exec -it kafka-learning /opt/kafka/bin/kafka-console-producer.sh \
  --topic test-topic \
  --bootstrap-server localhost:9092

#  Consumer receiving data
docker exec -it kafka-learning /opt/kafka/bin/kafka-console-consumer.sh \
  --topic test-topic \
  --from-beginning \
  --bootstrap-server localhost:9092
