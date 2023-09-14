#!/bin/bash

if [ "$1" == "retrieve" ]; then
  exec bash /scripts/retrieve-pod-metrics.sh
elif [ "$1" == "archive" ]; then
  exec bash /scripts/archive-pod-metrics.sh
else
  echo "Unknown script specified"
fi
