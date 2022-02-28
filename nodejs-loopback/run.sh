#!/bin/bash
set -ex

# export ELASTIC_APM_ACTIVE=true
export ELASTIC_APM_SERVICE_NAME=nodejs-loopback
export ELASTIC_APM_ENVIRONMENT=dev
export ELASTIC_APM_DISABLE_METRICS='system*,transaction*,span*'
export ELASTIC_APM_SERVER_URL=http://127.0.0.1:8201
export ELASTIC_APM_METRICS_INTERVAL='0s'
export ELASTIC_APM_BREAKDOWN_METRICS=false

node server.js
