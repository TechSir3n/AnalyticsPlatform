# Запуск Kafka producer
Start-Process "powershell" -ArgumentList "-NoExit",  "-Command", "cd kafka/producer; go run main.go kafka_producer.go"

# Запуск Kafka admin 
Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd kafka/admin; go run main.go kafka_admin.go"

# # Запуск Kafka consumer
Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd kafka/consumer; go run main.go kafka_consumer.go"

# Запуск gRPC service1
 Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_services/service_t1; go run main.go service_grpc_t1.go"

# Запуск gRPC service2
 Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_services/service_t2; go run main.go service_grpc_t2.go"

# Запуск data processing
Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd data_processing; go run data_processor.go"

# -NoExit let doesn't close console(output) after run 
# -Command for write process which need to run 