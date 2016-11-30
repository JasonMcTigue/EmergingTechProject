# EmergingTechProject
This is the repository for our group project in Emerging Technologies. It is a group project where we have to set up a single-page web application written in the programming language Go.

###### Gary Mc Hugh, G00308668
###### Alan Murphy G00312295
###### Jason McTigue G00312233
###### Darren Fitzpatrick G00311853 

###Technologies Used:
+ Html, Java Script, Css
+ Bootstrap,
+ JQuery
+ Go
+ MySQL
+ Macaron Framework

### Prerequisites
+ SET GOPATH=C:/Users/garym/EmergingTechProject
+ Download the following packages in the comand line
+ Macaroon (go get gopkg.in/macaron.v1)
+ Gin Web Framework (go get github.com/gin-gonic/gin)
+ Go-MySQL-Driver (go get github.com/go-sql-driver/mysql)
+ Copy and paste the contents of Library.sql into a local MySQL console
+ Change the value of line 23 in server.go to your username and password of MySQL
+ Run the Go files from the command line seperatly (go run webapp.go) or (go run server.go)

### Functionality
+ To access server.go go to http://127.0.0.1:5000/books or for a single book go to http://127.0.0.1:5000/book/98-7955-855-3 using the books isbn.
+ To access webpage.go to http://127.0.0.1:4001/

Library Page Screen Shot.
<img src="/public/Images/Webappscreens.png" alt="Library" width="600" height="600"/>
 


