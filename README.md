# Inventory System CRUD API Application Project
I am developing a CRUD API for an inventory management system using Golang. </br>
The purpose of this personal project to deepen my understanding of web development and database iteraction. </br>
CRUD, which stands for Create, Read, Update, and Delete forms the foundation of applications that manage data. </br>

**Running the Application:**
To run the application, type ```go run main.go ``` into the terminal. Open another terminal and type ```curl localhost:3000/inventory``` - a GET request. 
The output should result: Inventory System Project.  It'll mean the server is properly working. </br>
If ```curl localhost:3000/```, it'll return 202 page not found because there is no handler for the / route. Go-Chi automatically handles it. </br>
In the first terminal, to terminate the server, press Crtl + C. </br>

For a POST request, type ```curl -X POST localhost:3000/inventory -v``` - which specifys the HTTP method and fully view the HTTP headers.

**Third-Party Dependency:**
Using Go-Chi to manage HTTP requests, along with the middleware package for logging HTTP requests and responses.</br>
Go-Chi is a lightweight, idiomatic and composable router for building Go HTTP services. </nr>
To install Go-Chi, type: ```go get -u github.com/go-chi/chi/v5``` in the terminal. Now an entry should be added to the go.mod file. </br>
Note: go.sum file now exists - used to ensure consistency across package versions in case they are updated