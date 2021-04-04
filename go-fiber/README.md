Go Implementation using Fiber (https://github.com/gofiber/fiber)

go version go1.15.6 darwin/amd64

go run main.go

- Implemented: 
    * Simple REST API CRUD logic and validation
    * Interaction with Redis via repository
    * Model dynamic
    * Get all users: /api/user: ~300 ms with 30000
    * Get user by email: /api/user/email <10ms with 30000

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


