kubectl apply -f /Users/tharsanan/Documents/github/forked/apk/gateway/mediation/deployment/quickstart.yaml -n default
sleep 5
kubectl wait --for=condition=Ready pods --all -n envoy-gateway-system --timeout=300s

export ENVOY_SERVICE=$(kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-namespace=default,gateway.envoyproxy.io/owning-gateway-name=eg -o jsonpath='{.items[0].metadata.name}')

kubectl -n envoy-gateway-system port-forward service/${ENVOY_SERVICE} 8888:80 &

curl --verbose --header "Host: www.example.com" http://localhost:8888/get
