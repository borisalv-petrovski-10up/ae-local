# App engine on local machine

## Get started

### Prerequisites
You need to have the following software installed:
- Optional [Golang 1.15.10](https://go.dev/dl/)
- Required [Docker](https://docs.docker.com/engine/install/)
- Required [docker-compose](https://docker-docs.netlify.app/compose/install/)
- Required [Postman](https://www.postman.com/downloads/)

### Starting the application
Run `make start` - this will start the application on localhost:8080
in isolated App Engine Standard Environment with all required dependencies

## Running tests

### Unit tests
To run unit tests, run the `make tests` command. Tests will be executed in the isolated environment
with the datastore emulator and all required dependencies. Once tests are run, all containers will be
stopped automatically.

### Testing with Postman collection
There is Postman collection in the root directory of the project
which can be imported and then used to test the application.
To import that collection open Postman then search for a button called
`Import` which should be in the left corner next to `New` button.
Once the collection is imported, you should see a couple of create and read
requests that you can use to check how the application works. Don't forget to
start the App Engine emulator using the `make start` command before sending
request with the imported Postman collection.

## Datastore GUI
When the App Engine emulator is started, you can use the Datastore GUI located at [http://localhost:3000](http://localhost:3000)
to view datastore entities created locally. You can also use it to see data created by sample
requests from the Postman collection.