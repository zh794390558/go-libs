#!/usr/bin/env bash

go build main.go && ./main -crt admin-crt.pem -key admin-key.pem -ca-crt ca.crt  -ca-key ca.key
