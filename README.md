# The study project. It is my first project on golang.

## Usage
To spin up the Docker container run
```
docker-compose build
docker-compose up
```
Then visit `http://localhost:8080` in your browser to view the API.

### System url
 - _"/ping"_ - Server check If the server is ready, the server will send the answer "PONG"
 - _"/info"_ - View base server info
 - _"/"_ - View base server info(the same "/info")
 
## API Documentation
#### JSON-RPC 2.0 API
 - **user**
    - [Register](documentation/api-json-rpc/user-register.md) - Create new user in system
    - [Find](documentation/api-json-rpc/find-users.md)  - Find users by params

#### REST API
 - **users**
    - [GET](documentation/api-rest/users-get.md) - Get all users
    - [POST](documentation/api-rest/users-post.md) - Create new user in system