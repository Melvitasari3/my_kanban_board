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
-	
# Heroku Link
https://hacktiv8-kanban-board-emf.herokuapp.com/

# How to Use Application
## User
* Login User \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/users/login \
Method    : POST \
Parameter : JSON / form \
{\
  &nbsp; "email"    : "string",\
  &nbsp; "password" : "string"\
}\
Output    : Status 200 \
{\
  &nbsp; "token" : "jwt string"\
}
* Register User \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/users/register \
Method    : POST \
Parameter : JSON / form \
{\
  &nbsp; "full_name" : "string",\
  &nbsp; "email" : "string",\
  &nbsp; "pasword" : "string"\
}\
Output : Status 201 \
{\
  &nbsp; "id" : "integer",\
  &nbsp; "full_name" : "string",\
  &nbsp; "email" : "string",\
  &nbsp; "created_at" : "date"\
}
* Update Account
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/users/update-account \
Method    : PUT \
Headers : Authorization (Bearer Token string) \
Parameter : - 
* Delete User \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/users/register \
Method    : DELETE \
Headers : Authorization (Bearer Token string) \
Parameter : - 


## Categories
* Create New Categories \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/categories \
Method    : POST \
Headers : Authorization (Bearer Token string) \
Parameter : -
* Get All Categories \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/categories \
Method    : GET \
Headers : Authorization (Bearer Token string) \
Parameter : -
* Update Category by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/categories/:categoryId \
Method    : PATCH \
Headers : Authorization (Bearer Token string) \
Parameter : 
> categoryId (integer)
* Delete Category by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/categories/:categoryId \
Method    : DELETE \
Headers : Authorization (Bearer Token string) \
Parameter : 
> categoryId (integer) \

## Task
* Create New Tasks \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks \
Method    : POST \
Headers : Authorization (Bearer Token string) \
Parameter : -
* Get All Tasks \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks \
Method    : GET \
Headers : Authorization (Bearer Token string) \
Parameter : -
* Update task by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks/:tasksId \
Method    : PUT \
Headers : Authorization (Bearer Token string) \
Parameter : 
> taskId (integer)
* Delete Category by ID \
Alamat    : https://hacktiv8-kanban-board-emf.herokuapp.com/tasks/:tasksId \
Method    : DELETE \
Headers : Authorization (Bearer Token string) \
Parameter : 
> taskId (integer) \
