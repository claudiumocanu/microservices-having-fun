### Sometimes I'm lonely...  
- Frontend in Angular 
- Most of the backend components in golang
- Maybe add some node, java, python components just for fun, larter on
- RabbitMQ for communication between some of the microservices
- A relational DB, maybe postgres
- MongoDB to perist some messages
- Shall see if will make sense to add a Redis instance as well

---

### DONE short-list
- Sketch frontend
- Initial config of rabbitMQ
- Initial config of PostgresDB
- Simple REST API in golang (Angular->go->postgres)
- Postman collection to hit the go backend
- Move angular and the go-microservice in containers
- Integrate frontend with the go backend microservice(sensor list only)

### TODO short-list
- write a sensor
- bring up mongo
- write a coordinator
- write microservice to fetch sensor data from mongo

### DONE Big picture

### TODO Big picture
- Create go message/event emitters -> rabbit
- Figure out an auto-discovery mechanism
  - sensors must emit an event when they come online
  - some coordinator must auto-update the postgres db
  - a mongo sensor message consumer must listen and persist the messages in mongo