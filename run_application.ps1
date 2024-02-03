# # Запуск gRPC service1
Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_services/t1; go run main.go service_grpc.go"

#Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_clients/t1; go run main.go grpc_client.go"



# Запуск Kafka admin 
# Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd kafka/admin; go run main.go kafka_admin.go"

# # # Запуск Kafka consumer
Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd kafka/consumer; go run main.go consumer.go"



# # Запуск gRPC service2
#  Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_services/t2; go run main.go service_grpc.gp"

# # Запуск data processing
# Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd data_processing; go run data_processor.go"

# -NoExit let doesn't close console(output) after run 
# -Command for write process which need to run 