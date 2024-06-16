#!/bin/bash

declare -a services=("customer" "product" "order" "gateway")

for service in "${services[@]}"; do
    pid=$(ps aux | grep "./$service" | grep -v grep | awk '{print $2}')
    if [ ! -z "$pid" ]; then
        kill $pid
        echo "Stopped $service with PID $pid"
    else
        echo "No running process found for $service"
    fi
done
