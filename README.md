# Go-Form
##### NOTE: You must have go install and have exported its path.Create a user in postsql: library, Password: admin, and then run the sql file, this will create the database and the table.

### Commands 

##### Run the following command: "go run ." without the quotes.
##### Open a browser and type http://localhost:4000. If its working you see that it has "Welcome to quotebox!" in the page. 

##### To view the form navigate to http://localhost:4000/book/create. 
##### Fill the form and submit it, and boom!! you have posted the information to the database. 

### To view the data posted to the database.
##### Log into post sql with : psql --host=localhost --dbname=library --username=library, and enter 'admin' as password. 
##### Run the following command to view the data: "Select * FROM books"
