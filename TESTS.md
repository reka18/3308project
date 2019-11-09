### Project Title: Social Media Site
### Team Members: Regan Karnes, Rodrigo Garcia, Dominick Graham

<br/><br/>

# Automated Test Cases and Details

Once the steps in the <b>readme.md</b> have been completed, the automated tests can then be run.  This document will guide the user in running
the tests, describe what the tests mean, and how they work.

<br/><br/>

## Getting started

Again, this document assumes that you have successfully completed the steps outlined in the <b>readme.md</b>.  If you have not already done so,
please refer back to this document before proceeding.

To run the unit tests, make sure the site is built
```bash
go build && ./SocialMediaSite
```

then run
```bash
go test -v
```

Your output should appear similar to this
```
=== RUN   TestDatabase
2019/10/29 18:57:31 Database connection established.
--- PASS: TestDatabase (0.00s)
    package_test.go:14: TestDatabase pass!
=== RUN   TestCreateDatabase
2019/10/29 18:57:31 Database connection established.
2019/10/29 18:57:31 Successfully created database.
--- PASS: TestCreateDatabase (0.05s)
    package_test.go:22: Warning: Problem creating database: pq: database "socialmediasite" already exists
    package_test.go:23: Continuing...
    package_test.go:28: TestCreateDatabase pass!
=== RUN   TestInitializeDatabase
2019/10/29 18:57:31 Database connection established.
2019/10/29 18:57:31 Successfully dropped tables: user_account
2019/10/29 18:57:31 Unable to create enum: pq: type "gender" already exists
2019/10/29 18:57:31 'user_account table' created successfully.
--- PASS: TestInitializeDatabase (0.02s)
    package_test.go:40: TestInitializeDatabase pass!
=== RUN   TestEncrypt
--- PASS: TestEncrypt (0.00s)
    package_test.go:56: TestEncrypt pass!
=== RUN   TestAddNewUserAccount
2019/10/29 18:57:31 Database connection established.
2019/10/29 18:57:31 Successfully added user <rigo.garcia@colorado.edu> to Database.
2019/10/29 18:57:31 Successfully added user <reagan.karnes@colorado.edu> to Database.
--- PASS: TestAddNewUserAccount (0.01s)
    package_test.go:76: TestAddNewUserAccount pass!
=== RUN   TestLoginUserAccount
2019/10/29 18:57:31 Database connection established.
--- PASS: TestLoginUserAccount (0.01s)
    package_test.go:101: TestLoginUserAccount pass!
PASS
ok      SocialMediaSite 0.106s
```

<br/><br/>

## Line by Line Explaination

<br/><br/>

### === RUN   TestDatabase
As its name implies this function tests the postgresql database connectivity.<br/><br/>

On succesful connection you should recieve:<br/>
```TestDatabase pass!```
<br/><br/>
On a failed connection attempt the test will show:<br/>
```Unable to open connection.```

<br/>
<br/>

### === RUN   TestCreateDataBase
The create database ensures that the socialmediasite schema has been created in postgresql.  Since the database should already exist if the site is properly built, the response is a bit misleadig<br/><br/>

On succesful connection you should recieve:<br/>
```Warning: Problem creating database:```<br/>
```pq: database "socialmediasite" already exists```<br/>
```package_test.go:23: Continuing...```<br/>
```package_test.go:28: TestCreateDatabase pass!```<br/>

<br/>

If the schema has not already been built, then you the unit test will notify you but create the schema and the test will pass:<br/>
```Warning: Problem creating database: pq: database "socialmediasite2" does not exist```<br/>
```package_test.go:23: Continuing...```<br/>
```package_test.go:28: TestCreateDatabase pass!```<br/>

<br/>
<br/>

### === RUN   TestInitializeDatabase

This test creates a gender typ as ENUM, as well as creating a user_account table.  If a user_account table already exits it will drop it an recreate it.  If the ENUM gender type already exists, the test will notify the user but no further action will be taken.<br/><br/>

If none of the types or tables currently exist in the DB they are created:<br/>
```Database connection established.```<br/>
```Successfully dropped tables: user_account```<br/>
```'gender' enum created successfully.```<br/>
```'user_account table' created successfully.```<br/>
```TestInitializeDatabase pass!```<br/>


<br/>
<br/>

If the tables and/or the ENUM type exists, the test notices and passes:<br/>
```Database connection established.```<br/>
```Successfully dropped tables: user_account```<br/>
```Unable to create enum: pq: type "gender" already exists```<br/>
```'user_account table' created successfully.```<br/>
```TestInitializeDatabase pass!```<br/>


<br/>
<br/>

### === RUN   TestEncrypt

TestEncrypt tests the hashing function used to encrypt user passwords.  The system uses sha256 to hash passwords.    If the hashed intenal password matches the previously stored value of the hash, then the test passes.<br/><br/>

On successfully hashing:<br/>
```Hashing successful.```<br/>
```TestEncrypt pass!```<br/>

<br/>

On failure:<br/>
```Unknown hashing error:...```<br/>



<br/>
<br/>


### === RUN   TestAddNewUserAccount

TestAddNewUserAccount will ensure that the user_account table is formatted, and functioning properly, two users are added to the user_account table complete with hashed passwords, and email adresses.<br/><br/>

On successfully adding the users:<br/>
 ```Database connection established.```<br/>
 ```Hashing successful.```<br/>
 ```Successfully added user <rigo.garcia@colorado.edu> to Database.```<br/>
 ```Hashing successful.```<br/>
 ```Successfully added user <reagan.karnes@colorado.edu> to Database.```<br/>
 ```--- PASS: TestAddNewUserAccount (0.01s)```<br/>


<br/>

On failure:<br/>
```"Unable to add user:"...```<br/>
<br/>
<br/>


### === RUN   TestLoginUserAccount

The final test ensures that previously created test user accounts are functioning by using the generated credentials to simulate a user login event.<br/>


On successful user login:<br/>
 ```Database connection established.```<br/>
 ```Hashing successful.```<br/>
 ```Login successful ```<br/>
 ```TestLoginUserAccount pass!```<br/>
 
 <br/>

On failure:<br/>
```"Login data mismatch."...```<br/>
<br/>
<br/>


The final message should indicate that all the tests have passed<br/>
```PASS ```<br/>
```ok      SocialMediaSite 0.600s ```<br/>

<br/><br/>
# User Acceptance Testing 
<br/><br/>
## Login Use Case

<ol>
   <li>User navigates to login site</li>
   <li>User supplies email address for login</li>
   <li>User supplies password associated with login account</li>
   <li>On incorrect Username</li>
    <ol>
        <li>text or popup will appear on the login screen notifying the user of erroneous login credentials</li>
    </ol>
    <li>On incorrect Password</li>
    <ol>
        <li>text or popup will appear on the login screen notifying the user of erroneous password credentials</li>
    </ol>
    <li>If user forgets password</li>
    <ol>
        <li>There will be a password recovery link on the login page</li>
        <li>clicking the link will send a password recovery email to the user's email address</li>
        <li>User must supply a valid email login for password recovery to work</li>
    </ol>
    <li>On successful Login</li>
    <ol>
        <li>User it taken to their main account page</li>
        <li>User session information is activated and logged in the database</li>
    </ol>
</ol>
