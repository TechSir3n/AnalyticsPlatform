# Запуск второго микросервиса
Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_services/t1; go run main.go service_grpc.go"

# Запуск первого микросервиса
#Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_services/t2; go run main.go service_grpc.gp"

# Запуск первого клиента 
Start-Sleep -Seconds 2
Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_clients/t1; go run main.go grpc_client.go"

# Запуск второго клиента 
#Start-Sleep -Seconds 2
#Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd grpc_clients/t2; go run main.go grpc_client.go"

Start-Process "powershell" -ArgumentList "-NoExit", "-Command", "cd kafka/consumer; go run main.go consumer.go"

# -NoExit let doesn't close console(output) after run 
# -Command for write process which need to run 