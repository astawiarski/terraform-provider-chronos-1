#!/bin/bash

TF_ACC=yes CHRONOS_URL=${CHRONOS_URL} go test ./chronos -v
