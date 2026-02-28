#  Producer sending data
docker exec -it kafka-learning /opt/kafka/bin/kafka-console-producer.sh \
  --topic test-topic \
  --bootstrap-server localhost:9092
