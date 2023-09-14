#!/bin/bash
kubectl get configmap $POD_METRICS_ARCHIVE_CONFIGMAP -o jsonpath='{.binaryData.data}' -n $APK_RELEASE_NAMESPACE | {
    read -r base64_encoded_data
    if [[ -n "$base64_encoded_data" ]]; then
        echo "$base64_encoded_data" | base64 -d > existing_data
        gunzip -c existing_data > data.csv
    fi
}

kubectl get configmap $POD_METRICS_CONFIGMAP -o yaml -n $APK_RELEASE_NAMESPACE | grep -E '^[[:space:]]+"' | sed 's/^[[:space:]]\+"\(.*\)": "\(.*\)"$/\1:\2/g' > new_data
kubectl patch configmap $POD_METRICS_CONFIGMAP -n $APK_RELEASE_NAMESPACE --type='json' -p='[{"op": "add", "path": "/data", "value": {}}]'
cat new_data | sed -E "s/: '\{/,/g; s/\}',/\\n/g; s/}'//g; s/\". \"/\",\"/g" | sed "s/ //g" | sed 's/{/"/g' | sed 's/}/"/g' > new_data.csv

cat new_data.csv >> data.csv

gzip -c data.csv > data
kubectl delete configmap $POD_METRICS_ARCHIVE_CONFIGMAP -n $APK_RELEASE_NAMESPACE
kubectl create configmap $POD_METRICS_ARCHIVE_CONFIGMAP --from-file=data -n $APK_RELEASE_NAMESPACE
