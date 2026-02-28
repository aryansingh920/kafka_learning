# Create topic 
docker exec -it kafka-learning /opt/kafka/bin/kafka-topics.sh \
  --create --topic test-topic-4 \
  --bootstrap-server localhost:9092 \
  --partitions 4 \
  --replication-factor 1

#List all topics 
docker exec -it kafka-learning /opt/kafka/bin/kafka-topics.sh \
  --list \
  --bootstrap-server localhost:9092

# Describe the topic made
docker exec -it kafka-learning /opt/kafka/bin/kafka-topics.sh \
  --describe --topic test-topic-4 \
  --bootstrap-server localhost:9092 

# Delete a topic 
docker exec -it kafka-learning /opt/kafka/bin/kafka-topics.sh \
  --delete \
  --topic test-topic \
  --bootstrap-server localhost:9092
