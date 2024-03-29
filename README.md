# my_kanban_board
Back End App to create and modify kanban board

# Collaborator
1. Erico
2. Melvita
3. Firman

# List Library
- "fmt"
- log
- "net/http"
- "encoding/json"
- "database/sql"
- "strconv"
- "gorm.io/driver/postgres"
-	"gorm.io/gorm"
-	"github.com/gin-gonic/gin"
-	"golang.org/x/crypto/bcrypt"
-	"github.com/dgrijalva/jwt-go"
-	"github.com/asaskevich/govalidator"

# Heroku Link
https://hacktiv8-kanban-board-emf.herokuapp.com/

# Postman Link
https://documenter.getpostman.com/view/18861248/UVeGs6U6

# How to Use Application
## User
* Login User \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/users/login \
Method    : POST \
Parameter : JSON / form \
&nbsp; {\
  &nbsp; &nbsp; &nbsp; "email"    : "string",\
  &nbsp; &nbsp; &nbsp; "password" : "string"\
&nbsp; }\
Response    : Status 200 \
&nbsp; {\
  &nbsp; &nbsp; &nbsp; "token" : "jwt string"\
&nbsp; }

* Register User \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/users/register \
Method    : POST \
Parameter : JSON / form \
&nbsp; {\
  &nbsp; &nbsp; &nbsp; "full_name" : "string",\
  &nbsp; &nbsp; &nbsp; "email" : "string",\
  &nbsp; &nbsp; &nbsp; "pasword" : "string"\
&nbsp; }\
Response : Status 201 \
&nbsp; {\
  &nbsp; &nbsp; &nbsp; "id" : "integer",\
  &nbsp; &nbsp; &nbsp; "full_name" : "string",\
  &nbsp; &nbsp; &nbsp; "email" : "string",\
  &nbsp; &nbsp; &nbsp; "created_at" : "date"\
&nbsp; }

* Update Account
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/users/update-account \
Method    : PUT \
Headers : Authorization (Bearer Token string) \
Request : JSON / form \
&nbsp; {\
  &nbsp; &nbsp; &nbsp; "full_name" : "string",\
  &nbsp; &nbsp; &nbsp; "email" : "string",\
&nbsp; }\
Response : Status 200 \
&nbsp; {\
  &nbsp; &nbsp; &nbsp; "id" : "integer",\
  &nbsp; &nbsp; &nbsp; "full_name" : "string",\
  &nbsp; &nbsp; &nbsp; "email" : "string",\
  &nbsp; &nbsp; &nbsp; "updated_at" : "date"\
&nbsp; }


* Delete User \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/users/register \
Method    : DELETE \
Headers : Authorization (Bearer Token string) \
Parameter : - \
Response : Status 200 \
&nbsp; {\
  &nbsp; &nbsp; &nbsp; "message" : "Your account has been seccessfully deleted" \
&nbsp; }


## Categories
* Create New Categories \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/categories \
Method    : POST \
Headers : Authorization (Bearer Token string) \
Parameter : -
Request : JSON / form \
&nbsp; { \
  &nbsp; &nbsp; &nbsp; "type" : "string" \
&nbsp; } \
Response : Status 201 \
&nbsp; { \
  &nbsp; &nbsp; &nbsp; "id" : "integer" \
  &nbsp; &nbsp; &nbsp; "type" : "string" \
  &nbsp; &nbsp; &nbsp; "created_at" : "date" \
&nbsp; }  

* Get All Categories \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/categories \
Method    : GET \
Headers : Authorization (Bearer Token string) \
Parameter : - \
Response : Status 200 \
[ \
  &nbsp; { \
    &nbsp; &nbsp; &nbsp; &nbsp; "id" : "integer", \
    &nbsp; &nbsp; &nbsp; &nbsp; "type" : "string", \
    &nbsp; &nbsp; &nbsp; &nbsp; "updated_at : "date", \
    &nbsp; &nbsp; &nbsp; &nbsp; "created_at : "date", \
    &nbsp; &nbsp; &nbsp; &nbsp; "Tasks":[ \
      &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; { \
        &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "id" : "integer", \
        &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "title" : "string", \
        &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "description" : "string", \
        &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "user_id" : "integer", \
        &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "category_id": "date", \
        &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; "updated_at":"date" \
      &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; }  \
    &nbsp; &nbsp; &nbsp; &nbsp; ]  \
  &nbsp; }  \
]

* Update Category by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/categories/:categoryId \
Method    : PATCH \
Headers : Authorization (Bearer Token string) \
Parameter : categoryId (integer) \
Request : JSON / form \
&nbsp; { \
  &nbsp; &nbsp; &nbsp; "type" : "string"  \
&nbsp; } \
Response : Status 200 \
&nbsp; {  \
  &nbsp; &nbsp; &nbsp; "id" : "integer", \
  &nbsp; &nbsp; &nbsp; "type" : "string", \
  &nbsp; &nbsp; &nbsp; "updated_at" : "date" \
&nbsp; } 

* Delete Category by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/categories/:categoryId \
Method    : DELETE \
Headers : Authorization (Bearer Token string) \
Parameter : categoryId (integer) \
Request : - \
Response : Status 200 \
&nbsp; { \
  &nbsp; &nbsp; &nbsp; "message" : "Category has been successfully deleted" \
&nbsp; }

## Task
* Create New Tasks \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks \
Method    : POST \
Headers : Authorization (Bearer Token string) \
Parameter : - \
Request : JSON / form \
&nbsp; { \
  &nbsp; &nbsp; "title" : "string", \
  &nbsp; &nbsp; "description" : "string", \
  &nbsp; &nbsp; "category_id" : "integer" \
&nbsp; } \
Response : Status 201 \
&nbsp; { \
  &nbsp; &nbsp; "id" : "integer", \
  &nbsp; &nbsp; "title" : "string", \
  &nbsp; &nbsp; "status" : "boolean", \
  &nbsp; &nbsp; "description" : "string", \
  &nbsp; &nbsp; "user_id" : "integer", \
  &nbsp; &nbsp; "category_id" : "integer", \
  &nbsp; &nbsp; "created_at" : "date" \
&nbsp; }

* Get All Tasks \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks \
Method    : GET \
Headers : Authorization (Bearer Token string) \
Parameter : - \
Request : - \
Response : Status 200 \
&nbsp; [ \
  &nbsp; &nbsp; { \
    &nbsp; &nbsp; &nbsp; "id" : "integer", \
    &nbsp; &nbsp; &nbsp; "title" : "string", \
    &nbsp; &nbsp; &nbsp; "status" : "boolean", \
    &nbsp; &nbsp; &nbsp; "description" :  "string", \
    &nbsp; &nbsp; &nbsp; "user_id" : "integer", \
    &nbsp; &nbsp; &nbsp; "category_id" : "integer", \
    &nbsp; &nbsp; &nbsp; "created_at" : "date", \
    &nbsp; &nbsp; &nbsp; "User" : { \
      &nbsp; &nbsp; &nbsp; &nbsp; "id" : "integer", \
      &nbsp; &nbsp; &nbsp; &nbsp; "email" : "string", \
      &nbsp; &nbsp; &nbsp; &nbsp; "full_name" : "string" \
    &nbsp; &nbsp; &nbsp; } \
  &nbsp; &nbsp; } \
&nbsp; ] 

* Update task by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks/:tasksId \
Method    : PUT \
Headers : Authorization (Bearer Token string) \
Parameter : taskId (integer) \
Request : JSON / form \
&nbsp; { \
  &nbsp; &nbsp; "title" : "string", \
  &nbsp; &nbsp; "description" : "string" \
&nbsp; }
Response : Status 200 \
&nbsp; { \
  &nbsp; &nbsp; "id" : "integer", \
  &nbsp; &nbsp; "title" : "string", \
  &nbsp; &nbsp; "description" :  "string", \
  &nbsp; &nbsp; "status" : "boolean", \
  &nbsp; &nbsp; "user_id" : "integer", \
  &nbsp; &nbsp; "category_id" : "integer", \
  &nbsp; &nbsp; &nbsp; "updated_at" : "date" \
&nbsp; }

* Update task status by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks/update-status/:taskId \
Method    : PATCH \
Headers : Authorization (Bearer Token string) \
Parameter : taskId (integer) \
Request : JSON /form \
&nbsp; { \
  &nbsp; &nbsp; "status" : "boolean" \
&nbsp; } \
Response : Status 200 \
&nbsp; { \
  &nbsp; &nbsp; "id" : "integer", \
  &nbsp; &nbsp; "title" : "string", \
  &nbsp; &nbsp; "description" :  "string", \
  &nbsp; &nbsp; "status" : "boolean", \
  &nbsp; &nbsp; "user_id" : "integer", \
  &nbsp; &nbsp; "category_id" : "integer", \
  &nbsp; &nbsp; "updated_at" : "date" \
&nbsp; }

* Update task category by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks/update-category/:taskId \
Method    : PATCH \
Headers : Authorization (Bearer Token string) \
Parameter : taskId (integer) \
Request : JSON /form \
&nbsp; { \
  &nbsp; &nbsp; "category_id" : "integer" \
&nbsp; } \
Response : Status 200 \
&nbsp; { \
  &nbsp; &nbsp; "id" : "integer", \
  &nbsp; &nbsp; "title" : "string", \
  &nbsp; &nbsp; "description" :  "string", \
  &nbsp; &nbsp; "status" : "boolean", \
  &nbsp; &nbsp; "user_id" : "integer", \
  &nbsp; &nbsp; "category_id" : "integer", \
  &nbsp; &nbsp; "updated_at" : "date" \
&nbsp; }

* Delete Category by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks/:tasksId \
Method    : DELETE \
Headers : Authorization (Bearer Token string) \
Parameter : taskId (integer) \
Request : - \
Response : Status 200 \
&nbsp; { \
  &nbsp; &nbsp; "message" : "Task has been successfully deleted" \
&nbsp; } 
