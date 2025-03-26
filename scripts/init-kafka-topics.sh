#!/bin/bash

# Wait for Kafka to be ready
echo "Waiting for Kafka to be ready..."
sleep 10

KAFKA_HOST="localhost:9092"

TOPICS=(
  "add-booking"
  "confirm-booking"
)

for topic in "${TOPICS[@]}"; do
  echo "Creating topic: $topic"
  kafka-topics.sh --create --if-not-exists --bootstrap-server $KAFKA_HOST --replication-factor 1 --partitions 1 --topic "$topic"
done

echo "Topic initialization complete."