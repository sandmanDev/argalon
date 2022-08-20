# argalon
argalon jin sample structure code

Steps to execute the code

1- Clone this repo on your system

2- Install Docker & Docker-Compose on your system

3- cd argalon (move in to clone directory)

4- In cdm run command -> **docker-compose up**



Your container will up in few seconds on port 8088
You can test this using below curl 


curl --location --request POST '127.0.0.1:8088/argalon/divide_two' \
--header 'Content-Type: application/json' \
--data-raw '{"a":10,"b":2}'
