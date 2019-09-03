# aplicom-go

Server listening on incomming binary data (Aplicom protocol D). Parses the binary data and sends it to the pubsub topic iot-gnss.

Aplicom protocol D: D-Protocol consists of a single binary message packet for sending device status information including GPS data to
a server application. Protocol is one-way only, that is, there are no response messages.

Project ID: einride-portal  
Deploy on GCP VM instance: "aplicom-server"  
External ip: 35.228.99.113  
Listening tcp port: 5144  

## Run and test locally
To locally run aplicom-go server and send data to it. First enable some debugging in application to log some receiving bytes then open two terminals. The test sends a binary aplicom version 3 packet to port 5144.
1. Terminal 1: Run `make go-run-server-local`
2. Termanal 2: Run `make send-local-test-data`

## Deploy manually
The aplicom-go application is deployed on a GCP VM instance "aplicom-server".

### Locally

1. Run `make deploy`