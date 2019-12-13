## URL TO USE LIVE
http://bit.ly/cu-boulder-socialmediasite

## SETUP DATABASE
Install postgres v12 locally. Make sure you also
have 12 version of PSQL.

### Mac
##### install, remember that your rc file is either `.bash_profile` or `.zshrc` depending on what terminal you are using
```
brew install postgresql
echo 'export PATH="/user/local/opt/postgresql/bin:$PATH"' >> <your rc file>
```
##### start service
```
brew services start postgresql
```
##### enter PSQL
```
psql postgres
    CREATE USER <your whoami> CREATEDB;
    \q
```
Note this may say your username already exists. That's fine.
##### stop service
```
brew service stop postgresql
```
##### edit pg_hba.conf
```
vim /usr/local/var/postgresql/pg_hba.conf
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
brew services start postgresql
```

## Create Database
In project root, run the script `./build.sh`. This will generate an executable
in the root directory. Next, to create the database, run `./app --init` to set
up a clean database
#### BUILD DATABASE

##### in the root directory
```
go build
./SocialMediaSite --create
```

## USAGE
From root directory, type `go run *.go` or `go build && ./SocialMediaSite`

To run unit test suite, build the executable and then run `go test -v`
