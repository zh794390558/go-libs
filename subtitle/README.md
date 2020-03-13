# Build
go build

# Monitor
nohup ./rsmonitor &

# test
curl -H 'token:54321'  -H "Content-Type:application/json" -X POST --data '{"room_id": "001", "alert_type": "alarm"}' http://127.0.0.1:8082/push
