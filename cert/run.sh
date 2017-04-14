#!/usr/bin/env bash

go build && ./cert -crt ./data/admin-crt.pem -key ./data/admin-key.pem -ca-crt ./data/ca.crt  -ca-key ./data/ca.key
