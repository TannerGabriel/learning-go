#include <Arduino.h>
#include <Wire.h>
#include <SPI.h>
#include <ESP8266WiFi.h>
#include <PubSubClient.h>
#include <Adafruit_Sensor.h>
#include "Adafruit_TMP006.h"
#include <NTPClient.h>
#include <WiFiUdp.h>

Adafruit_TMP006 tmp006;

// Update these with values suitable for your network.
const char* ssid = "ssid";
const char* password = "password";
const char* mqtt_server = "10.0.0.22";

WiFiClient espClient;
PubSubClient client(espClient);

// Create variable to hold mqtt messages
#define MSG_BUFFER_SIZE	(100)
char msg[MSG_BUFFER_SIZE];

// Define NTP Client to get time
WiFiUDP ntpUDP;
NTPClient timeClient(ntpUDP);

// Variables to save date and time
uint32_t unixtime;

// Connecting to the WIFI network
void setup_wifi() {
  delay(10);
  
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  randomSeed(micros());

  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}

void reconnect() {
  // Loop until we're reconnected
  while (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    // Create a random client ID
    String clientId = "ESP8266Client-";
    clientId += String(random(0xffff), HEX);
    // Attempt to connect
    if (client.connect(clientId.c_str(), "mqtt", "")) {
      Serial.println("connected");
      // Once connected, publish an announcement...
      client.publish("outTopic", "hello world");
      // ... and resubscribe
      client.subscribe("inTopic");
    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");
      // Wait 5 seconds before retrying
      delay(5000);
    }
  }
}

void setup() {
  Serial.begin(9600);
  
  setup_wifi();

  client.setServer(mqtt_server, 1883);

  if (! tmp006.begin()) {
    Serial.println("No sensor found");
    while (1);
  }

  // Start the time client with an offset of 3600 seconds (+1 hours timezone)
  timeClient.begin();
  timeClient.setTimeOffset(3600);
}

void loop() {
  // Connect to the mqtt client
  if (!client.connected()) {
    reconnect();
  }
  client.loop();

  // Update the time client for the timestamps
  while(!timeClient.update()) {
    timeClient.forceUpdate();
  }

  // Get the unix timestamp
  unixtime = timeClient.getEpochTime();

  // Get the current object temperatur of the sensor
  float objt = tmp006.readObjTempC();

  // Create the message that will be send using mqtt
  String message = String("weather,location=us-midwest temperature="+String(objt) + " " + String(unixtime) + "0000000000");
  message.toCharArray(msg, message.length());
  Serial.println(msg);

  // Send the message on the sensors topic
  client.publish("sensors", msg);
  delay(1000); 
}
