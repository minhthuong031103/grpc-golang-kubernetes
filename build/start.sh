#!/bin/bash

declare -a services=("customer" "product" "order" "fileupload" "gateway")

for service in "${services[@]}"; do
    nohup ./$service > "${service}.log" 2>&1 &
    echo "Started $service, logging to ${service}.log"
done
