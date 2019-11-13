## SETUP DATABASE
Install postgres v9.6 locally. Make sure you also
have 9.6 version of PSQL.

### Mac
##### install, remember that your rc file is either `.bash_profile` or `.zshrc` depending on what terminal you are using
```
brew install postgresql@9.6
echo 'export PATH="/user/local/opt/postgresql@9.6/bin:$PATH"' >> <your rc file>
```
##### start service
```
brew services start postgresql@9.6
```
##### enter PSQL
```
psql postgres
    CREATE USER <your whoami> CREATEDB;
    CREATE USER sms_user CREATEDB;
    \q
```
Note this may say your username already exists. That's fine.
##### stop service
```
brew service stop postgresql@9.6
```
##### edit pg_hba.conf
```
vim /usr/local/var/postgresql@9.6/pg_hba.conf
```
##### delete the following
```
# TYPE  DATABASE        USER            ADDRESS                 METHOD

 # "local" is for Unix domain socket connections only
  local all            all                                     trust
 # IPv4 local connections:
  host  all            all            127.0.0.1/32             trust
 # IPv6 local connections:
  host  all            all            ::1/128                  trust
```
##### and add this
```
# TYPE  DATABASE        USER            ADDRESS 
  host  all             <your whoami>      127.0.0.1/32            trust
  host  all             <your whoami>      ::1/128                 trust
  local all             <your whoami>                              ident
```
##### start service
```
brew services start postgresql@9.6
```

## Create Database
Enter the postgres database
```
psql postgres
```
##### create a new database
```
CREATE DATABASE socialmediasite;
\q
```
#### TWO WAYS TO BUILD DATABASE

##### go into the new database
```
psql socialmediasite
```
##### build the user_account table
```
create type gender as enum ('M', 'F', 'O');

create extension if not exists pgcrypto;

create table user_account (
   id SERIAL PRIMARY KEY,
   age INT,
   firstName TEXT,
   lastName TEXT,
   email TEXT UNIQUE NOT NULL,
   gender gender NOT NULL,
   public BOOLEAN,
   joinDate DATE,
   active BOOLEAN,
   password TEXT
);
\q
```
#### OR
##### in the root directory
```
go build
./SocialMediaSite --create
```

## USAGE
From root directory, type `go run *.go` or `go build && ./SocialMediaSite`

To run unit test suite, build the executable and then run `go test -v`