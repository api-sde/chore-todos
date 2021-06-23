Go Implementation using Fiber (https://github.com/gofiber/fiber)

go version go1.15.6 darwin/amd64

go run main.go

- Implemented: 
    * Simple REST API CRUD logic and validation
    * Interaction with Redis via repository
    * Model dynamic
    * Get all users: /api/user
    * Get user by email
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
    * As a user, I can create a unique account with an email
    * As a user, I can see my profile
    * As a user, I can reate a note/todo
    * As a user, I can create, edit, delete a note/todo
    * As a user, I can see all my notes/todos (personal), and group notes
    * As a user, I can create a group of user to share
    * As a user, I can search for other users by name / email, send and accept invites into a group


