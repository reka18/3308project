## SETUP DATABASE
Install postgres v9.6 locally. Make sure you also
have 9.6 version of PSQL.

### Mac
##### install
```
brew install postgresql@9.6 libpq
echo 'export PATH="/user/local/opt/postgresql@9.6/bin:$PATH"' >> <your rc file>
```
##### start service
```
brew services start postgresql@9.6
```
##### enter PSQL
```
psql postgres
    CREATE USER <your user name> CREATEDB;
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
Delete the following
```
# TYPE  DATABASE        USER            ADDRESS                 METHOD

 # "local" is for Unix domain socket connections only
  local all            all                                     trust
 # IPv4 local connections:
  host  all            all            127.0.0.1/32            trust
 # IPv6 local connections:
  host  all            all            ::1/128                 trust
```
And add this
```
# TYPE  DATABASE        USER            ADDRESS 
  host  all             <username>      127.0.0.1/32            trust
  host  all             <username>      ::1/128                 trust
  local all             <username>                                ident
```
##### start service
```
brew services start postgresql@9.6
```

## CREATE TABLES
Create a local postgres database and execute the following SQL code
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
```

## USAGE
From root directory, type `go run *.go`