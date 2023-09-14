#!/bin/bash

convert_cpu_to_milli_cores() {

    input=$1
    if [[ $input == *m ]]; then
        echo $(echo "$input" | sed 's/m//g')
    else
        echo $(($(echo "$input") * 1000))
    fi
}

convert_memory_to_bytes() {

    input=$1
    if [[ $input == *Ki ]]; then
        echo $(($(echo "$input" | sed 's/Ki//') * 1024))
    elif [[ $input == *Mi ]]; then
        echo $(($(echo "$input" | sed 's/Mi//') * 1024 * 1024))
    elif [[ $input == *Gi ]]; then
        echo $(($(echo "$input" | sed 's/Gi//') * 1024 * 1024 * 1024))
    elif [[ $input == *k ]]; then
        echo $(($(echo "$input" | sed 's/k//') * 1000))
    elif [[ $input == *M ]]; then
        echo $(($(echo "$input" | sed 's/M//') * 1000000))
    elif [[ $input == *G ]]; then
        echo $(($(echo "$input" | sed 's/G//') * 1000000000))
    else
        echo $input
    fi
}

pods=$(kubectl get pods --selector=app.kubernetes.io/release=$APK_RELEASE_NAME -n $APK_RELEASE_NAMESPACE --no-headers | awk '{print $1 ":" $2}' | tr '\n' ','| sed 's/,$//')
IFS=',' read -ra pod_array <<< "$pods"
result=""
declare -A podCountArray
declare -A podCpuArray
declare -A podMemoryArray
for pod in "${pod_array[@]}"; do
    cleaned_output=$(echo "$pod" | sed 's/:[0-9]\/[0-9]//g' | sed 's/:/,/g')
    count=$(kubectl get pod $cleaned_output -n "$APK_RELEASE_NAMESPACE" -o=jsonpath='{.spec.containers}' | jq length)
    for ((i=0; i<count; i++)); do
        pod_metrics=$(kubectl get pod $cleaned_output -n "$APK_RELEASE_NAMESPACE" -o=jsonpath="{.spec.containers[$i]}")
        pod_name=$(echo $pod_metrics | jq '.name' | sed 's/"//g')
        cpu=$(convert_cpu_to_milli_cores $(echo $pod_metrics | jq '.resources.limits.cpu' | sed 's/"//g'))
        memory=$(convert_memory_to_bytes $(echo $pod_metrics | jq '.resources.limits.memory' | sed 's/"//g'))
        if [ -z "${podCountArray[$pod_name]}" ]; then
            podCountArray[$pod_name]=1
            podCpuArray[$pod_name]=$cpu
            podMemoryArray[$pod_name]=$memory
        else
            podCountArray[$pod_name]=$((podCountArray[$pod_name] + 1))
            podCpuArray[$pod_name]=$((podCpuArray[$pod_name] + $cpu))
            podMemoryArray[$pod_name]=$((podMemoryArray[$pod_name] + $memory))
        fi
    done
done

for pod_name in "${!podCountArray[@]}"; do
    if [ -z "${podCountArray[$pod_name]}" ]; then
        result+="{},"
    else
        result+="{$pod_name,${podCountArray[$pod_name]},$(echo "scale=1; ${podCpuArray[$pod_name]}/1000" | bc),$(echo "scale=1; ${podMemoryArray[$pod_name]}/(1024*1024)" | bc)},"
    fi
done

filtered_pod_metrics=${result%,}
echo $filtered_pod_metrics
data="{$filtered_pod_metrics}"
timestamp=$(date +%Y%m%d%H%M%S)
kubectl patch configmap $POD_METRICS_CONFIGMAP --patch "{\"data\": {\"$timestamp\": \"$data\"}}" -n $APK_RELEASE_NAMESPACE
