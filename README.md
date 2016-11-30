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
+ Macaroon

```
go get gopkg.in/macaron.v1

```
+ Gin Web Framework

```
go get github.com/gin-gonic/gin

```
+ Go-MySQL-Driver

```
go get github.com/go-sql-driver/mysql

```
+ Copy and paste the contents of Library.sql into a local MySQL console
+ Change the value of line 23 in server.go to your username and password of MySQL
+ Run the Go files from the command line seperatly

```
go run webapp.go

```
or

```
go run server.go

```
### Functionality
+ To access server.go go to http://127.0.0.1:5000/books or for a single book go to http://127.0.0.1:5000/book/98-7955-855-3 using the books isbn.
+ To access webpage.go to http://127.0.0.1:4001/



Library Page Screen Shot.

<img src="/public/Images/Webappscreen.png" alt="Library" width="900" height="600"/>

## How to use this site
To navaigate through our website you will use the nav bar at the top of the page. This nav bar will bring you anywhere in our website.

<img src="/public/Images/navBar.png" alt="Nav" width="1000" height="150"/>

The nav bar consists of the following 

+ Home Page

On the home page is a carosel with links to library resources such as timtables, exam papers, journes etc.

<img src="/public/Images/Webappscreen.png" alt="home" width="900" height="300"/>

+ Book a Room

A page which alows users to book a room in the library.

<img src="/public/Images/bookaroom.png" alt="bookaroom" width="900" height="300"/>

+ About Us

A page consisting of information about the library such as staff information and a map of the librarys location.

<img src="/public/Images/aboutus.png" alt="about" width="900" height="300"/>

+ Opening Hours

A page with tables of what day each library is open and between what time

<img src="/public/Images/openhours.png" alt="open" width="900" height="300"/>

+ Contact Us 

A page which consists of information of how to get into contact with the library including links to socail media.

<img src="/public/Images/contactus.png" alt="contact" width="900" height="300"/>

+ Books Drop Down Menu

A drop down menu allowing users to select a category of books and display a table containing all available books. also included in the drop down is a form to add books.

<img src="/public/Images/booksadd.png" alt="add" width="900" height="300"/>

+ Sign in button

A button whihc activates a modal containg a form which would allow users to enter their details.

<img src="/public/Images/signinmodal.png" alt="modal" width="900" height="300"/>


## Database
We designed the database to contain information on all the books availble in the library. We wanted the user to be able to display the desired book in a table and even add new books into the database.


The variables in the database are as follows:
+ ISBN
+ Title
+ Author
+ Category
+ Copies
+ Loan Limit

To see the info in the database run server.go in the command line with the following command

```
 go run server.go
 
```

Server.go Screen Shot.

<img src="/public/Images/server.png" alt="Server" width="1000" height="300"/>

## Documention
We have extensive documentation on everything we set out to do in our project. It can all be found and read in the documentation folder in this repostory.


