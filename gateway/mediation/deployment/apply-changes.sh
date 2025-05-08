cd /Users/tharsanan/Documents/github/forked/apk/gateway/mediation
./gradlew build
kubectl scale deployment grpc-ext-proc --replicas=0
kubectl scale deployment grpc-ext-proc --replicas=1
sleep 3
kubectl logs -l app=grpc-ext-proc -f