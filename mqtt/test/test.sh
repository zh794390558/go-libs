#!/bin/bash
mosquitto_pub -h localhost -t "mqtt" -m "Hello MQTT" 
mosquitto_sub -h localhost -t "mqtt" -v 

