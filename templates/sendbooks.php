<!--
Author: Gary McHugh
Date: November 29th, 2016
Adapted from: http://www.tutorialrepublic.com/php-tutorial/php-mysql-insert-query.php
-->

<?php
/* Attempt MySQL server connection. Assuming you are running MySQL
server with default setting (user 'root' with no password) */
$link = mysqli_connect("localhost", "root", "hello", "Library");
 
// Check connection
if($link === false){
    die("ERROR: Could not connect. " . mysqli_connect_error());
}
 
// Escape user inputs for security
$isbn = mysqli_real_escape_string($link, $_POST['ISBN']);
$title = mysqli_real_escape_string($link, $_POST['title']);
$author = mysqli_real_escape_string($link, $_POST['author']);
$category = mysqli_real_escape_string($link, $_POST['category']);
$copies = mysqli_real_escape_string($link, $_POST['copies']);
$loanlimit = mysqli_real_escape_string($link, $_POST['loan_limit']);
 
// attempt insert query execution
$sql = "INSERT INTO books (ISBN, title, author, category, copies, loan_limit) VALUES ('$ISBN', '$title', '$author', '$categories', '$copies', '$loan_limit')";
if(mysqli_query($link, $sql)){
    echo "Records added successfully.";
} else{
    echo "ERROR: Could not able to execute $sql. " . mysqli_error($link);
}
 
// close connection
mysqli_close($link);
?>  