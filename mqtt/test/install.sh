# https://blog.csdn.net/junmoxi/article/details/102165598

#sudo apt-add-repository ppa:mosquitto-dev/mosquitto-ppa
#sudo apt-get update
#
#sudo apt-get install mosquitto
#sudo apt-get install mosquitto-dev
#sudo apt-get install mosquitto-clients
#
#sudo service mosquitto status 

apt-add-repository ppa:mosquitto-dev/mosquitto-ppa
apt-get update

apt-get install mosquitto -y
apt-get install mosquitto-dev -y
apt-get install mosquitto-clients -y

service mosquitto start
service mosquitto status 

