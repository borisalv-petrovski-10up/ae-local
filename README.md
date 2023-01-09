# App engine on local machine

## Get started

### Prerequisites
You need to have the following software installed:
- [Golang 1.15.10](https://go.dev/dl/)
- [Docker](https://docs.docker.com/engine/install/)
- [docker-compose](https://docker-docs.netlify.app/compose/install/)
- [Postman](https://www.postman.com/downloads/)

### Starting the application
Run `make start` - this will start the application on localhost:8080
in isolated App Engine Standard Environment with all required dependencies

## Running tests

### Unit tests
Simply run this command `make tests` and it will run all tests in isolated environment 
with all required dependencies

### Testing with Postman collection
There is Postman collection in the root directory of the project
which can be imported and then used to test the application.
To import that collection open Postman then search for a button called
`Import` which should be in the left corner next to `New` button.
You can now navigate to the current directory and select the file under
`postman` folder. A new collection in Postman should be imported.
Once `make start` is run, we can send request with the imported
Postman collection.

## Datastore GUI
When some requests are send trough Postman collection we can
navigate to `http://localhost:3000`. This will open the Datastore GUI
where we can see all stored data from Postman requests.