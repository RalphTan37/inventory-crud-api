# Inventory System CRUD API Application Project
I am developing a CRUD API for an inventory management system using Golang. </br>
The purpose of this personal project to deepen my understanding of web development and database iteraction. </br>
CRUD, which stands for Create, Read, Update, and Delete forms the foundation of applications that manage data. </br>

**Running the Application:** </br>
To run the application, type ```go run main.go``` into the terminal. </br>
Open another terminal and there are several commands that can be typed. </br>

To add a new item in the inventory, type ```curl -X POST localhost:3000/inventory -d '{ "item_ID": "123e4567-e89b-12d3-a456-426614174000", "name": "Example Item", "category": "Category A", "quantity": 10, "price": 19.99, "supplier": "Supplier Inc.", "location": "Warehouse 1", "status": "Available", "expiration_date": "2025-12-31T00:00:00Z" }'```. </br>
To list all items in the inventory, type ```curl localhost:3000/inventory/```. </br>
To get an item from the inventory by ID, type ```curl localhost:3000/inventory/```. </br>
To update an item in the inventory by ID, type ```curl -X PUT localhost:3000/inventory/123e4567-e89b-12d3-a456-426614174000```. </br>
To delete an item in the inventory by ID, type ```curl -X DELETE localhost:3000/inventory/123e4567-e89b-12d3-a456-426614174000```. </br>

*Notes:* </br>
```-X``` - specifies the HTTP method </br>
Can add ```-v``` at the end of the command line to fully view the HTTP headers </br>

In the first terminal, to terminate the server, press Crtl + C. </br>

**Third-Party Dependency:** </br>
Using Go-Chi to manage HTTP requests and the middleware package for logging HTTP requests and responses.</br>
Go-Chi is a lightweight, idiomatic and composable router for building Go HTTP services. </br>
To install Go-Chi, type: ```go get -u github.com/go-chi/chi/v5``` in the terminal. An entry should be added to the go.mod file. </br>
*Note:* go.sum file now exists - used to ensure consistency across package versions in case they are updated </br>

Using Go-Redis as primary data storage. </br>
Go-Redis is an in-memory data structure store - which makes it fast. The downside if that the data stored only usually lasts for a short time. Not as safe as PostgreSQL (probably the better chocie for this project). </br>
To install Go-Redis, type: ```go get github.com/redis/go-redis/v9``` in the terminal. </br>

*Using Docker to install Redis. </br>
Docker is an open-source platform that automates the deployment, scaling, and management of applications using containerization - lightweight, standalone, and executable software packages. </br>

```docker ps``` lists the running containers on Docker host. </br>
```sudo service redis-server stop``` stops any current Redis servers running. </br>
```docker run -p 6379:6379 redis:latest``` will download and run the latest Redis img and bing the system 6379 port to the docker container 6379 port. To check the redis is running, type ```redis-cli``` in a second terminal and ```127.0.0.1:6379``` prompt is returned- can enter ```KEYS *``` to retrieve all keys in selected database. </br>

Using Google's UUID Package for assigning unique IDs to items. <br>
To install, type: ```go get github.com/google/uuid```.