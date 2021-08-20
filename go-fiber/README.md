Go Implementation using Fiber (https://github.com/gofiber/fiber)

go version go1.15.6 darwin/amd64

To start the app:
go run main.go

To run Redis locally in a container, run command:
docker run --name go-fiber-standalone -p 6379:6379 -d redis redis-server --appendonly yes

and then update repository.go with:
Addr:     "localhost:6379",

Containers:
- In root project folder, docker-compose and Dockerfile (go app) are built.
1 - Build the API image: 
docker build -t adrienbdx/go-fiber:latest .

2 - Run docker-compose up, check if the app is running and exposed at:
http://127.0.0.1:3333/hello

3 - docker-compose down

- Implemented: 
    * Simple REST API CRUD logic and validation
    * Interaction with Redis via repository
    * Model dynamic
    * Authentication logic via JWT with email, secure hashed password
    * Endpoint completed:
      - Get all users: /api/user
      - Get user by email: /api/user/{email}
      - Post new user: /api/user
      - Login existing user: /auth/login
  
  <table>
    <thead>
      <tr>
        <th>Average Response time (ms)</th>
        <th>User count</th>
        <th>Get all users: /api/user </th>
        <th>Get user by email: /api/user/email</th>
      </tr>
    </thead>
    <tbody>
        <tr>
            <td></td>
            <td>30 000</td>
            <td>~300 ms</td>
            <td><10 ms</td>
        </tr>
        <tr>
            <td></td>
            <td>60 000</td>
            <td>~500 ms</td>
            <td>5-10 ms</td>
        </tr>
    </tbody>
  </table>
  
  
- To Implement:
    * Move core logic into services
    * Testing / Performance
    * 3rd party integration
  
- Features:
    * As a user, I can create a unique account with an email - OK
    * As a user, I can see my profile
    * As a user, I can create a note/todo
    * As a user, I can create, edit, delete a note/todo
    * As a user, I can see all my notes/todos (personal), and group notes
    * As a user, I can create a group of user to share
    * As a user, I can search for other users by name / email, send and accept invites into a group


