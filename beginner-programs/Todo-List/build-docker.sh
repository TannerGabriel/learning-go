# Trying to automatically detect external ip address
export EXTERNAL_IP=$(ifconfig | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p' |tail -n 1)

# Uncomment the follwoing line and insert your server ip address if it was not detected correctly
export EXTERNAL_IP=10.0.0.8
echo Detected external IP: $EXTERNAL_IP
echo "if incorrect, edit in build-docker.sh and rebuild !!"
read -p "press [ENTER] to continue"

docker-compose build --build-arg VUE_APP_API_URL=$EXTERNAL_IP