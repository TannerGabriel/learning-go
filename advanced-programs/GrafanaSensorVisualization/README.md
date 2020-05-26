# Grafana InfluxDB Sensor Visualization

Grafana dashboard for visualizing sensor data. Data will be aquired using an ESP8266 and send over MQTT to Telegraf which saves it to the InfluxDB database. 

## Requirements

* ESP8266
* TMP006 Temperature sensor (You can use the `mqtt.go` script instead to send fake data)
* PC with Docker installation for running MQTT, Telegraf, InfluxDB and Grafana

## Technology Stack

* Mosquitto is our MQTT broker which accepts data coming from sensors
* Telegraf subscribes to the MQTT topic, where sensors publish the data. Telegraf then saves the data into InfluxDB
* Grafana reads the data in InfluxDB and visualizes them in a dashboard

## Getting started

### Configuring telegram

Telegraf needs to be configured to read the data of your MQTT topic. For that, you will need to modify the following lines in the `telegraf.conf` files.

```conf
servers = ["tcp://mqtt_server_ip:1883"]
topics = [
  "sensors"
]
data_format = "influx"
```

Here you need to insert the IP of your MQTT server and the topics you want to be saved in the database (This will be localhost if you are using the included docker-compose file). Then you can continue by modifing the following lines under the `outputs.influxdb` tag.

```conf
[[outputs.influxdb]]
  urls = ["http://influxdb_server_ip:8086"]
  skip_database_creation = true
  username = "telegraf"
  password = "telegraf"
```

Here you need to insert the IP Adress of your influxdb server and the username and password of the database (This will also be localhost if you use the included docker-compose file).

### Deploying Grafana, InfluxDB, Telegram and Mosquitto

Now that you have configured telegram you can deploy it using the included docker-compose file.

```bash 
docker-compose up -d
```

This will create all containers needed for the project. You now only need to upload the `SensorControl` script to your ESP8266 and configure your Grafana datasource and query. If you don't have a ESP8266 or TMP066 temperature sensor you can use the `mqtt.go` script instead which just sends fake sensor data over MQTT.