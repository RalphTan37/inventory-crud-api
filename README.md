# Inventory System CRUD API Application Project
I am developing a CRUD API for an inventory management system using Golang. </br>
The purpose of this personal project to deepen my understanding of web development and database iteraction. </br>
CRUD, which stands for Create, Read, Update, and Delete forms the foundation of applications that manage data. </br>

**Running the Application:** </br>
To run the application, type ```go run main.go ``` into the terminal. </br>
Open another terminal and there are several commands that can be typed. </br>

To add a new item in the inventory, type ```curl -X POST localhost:3000/inventory```. </br>
To list all items in the inventory, type ```curl localhost:3000/inventory/```. </br>
To get an item from the inventory by ID, type ```curl localhost:3000/inventory/```. </br>
To update an item in the inventory by ID, type ```curl -X PUT localhost:3000/inventory/item```. </br>
To delete an item in the inventory by ID, type ```curl -X DELETE localhost:3000/inventory/item```. </br>

*Notes:* </br>
```-X``` - specifies the HTTP method </br>
Can add ```-v``` at the end of the command line to fully view the HTTP headers </br>

In the first terminal, to terminate the server, press Crtl + C. </br>

**Third-Party Dependency:** </br>
Using Go-Chi to manage HTTP requests and the middleware package for logging HTTP requests and responses.</br>
Go-Chi is a lightweight, idiomatic and composable router for building Go HTTP services. </br>
To install Go-Chi, type: ```go get -u github.com/go-chi/chi/v5``` in the terminal. Now an entry should be added to the go.mod file. </br>
*Note:* go.sum file now exists - used to ensure consistency across package versions in case they are updated
