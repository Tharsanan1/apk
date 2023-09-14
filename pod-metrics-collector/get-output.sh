#!/bin/bash

if [ $# -lt 1 ] || [ $# -gt 2 ]; then
    echo "Usage: $0 <configmap> [namespace]"
    exit 1
fi
configmap="$1"
if [ $# -eq 2 ]; then
    namespace="$2"
else
    namespace="default"
fi
kubectl get configmap $configmap -o jsonpath='{.binaryData.data}' -n $namespace | {
    read -r base64_encoded_data
    if [[ -n "$base64_encoded_data" ]]; then
        echo "$base64_encoded_data" | base64 -d > data
        echo "\"timestamp\",\"pod_name,pod_count,pod_cpu,pod_memory\",\"...\"" > data.csv
        gunzip -c data >> data.csv
    fi
}

read -p "Do you want to clear the ConfigMap? (y/n): " choice
if [ "$choice" == "y" ]; then
    kubectl patch configmap "$configmap" -n "$namespace" --type='json' -p='[{"op": "add", "path": "/binaryData", "value": {}}]'
    echo "ConfigMap '$configmap' in namespace '$namespace' has been cleared."
fi

