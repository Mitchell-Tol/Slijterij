# Slijterij
## Getting Started
### REST API
The API for Slijterij is written in Go, so having Go installed on your system is a must. Version 1.23 *works on my machine*, but change the version in **go.mod** to yours if necessary. 

### Server
#### MySQL
The database is built using MySQL, which means that that needs to be installed and configured so the API can use it. 

* On Linux (so also MacOS) you can install the MySQL server by using the command `sudo apt-get install mysql-server`
* Installing it on Windows can best be done using the installer you download for installing the MySQL Workbench.

In order to connect to the client, on Windows you can just use the MySQL Workbench, and on Linux you can use the `sudo mysql` command.

The next step is to check if the server is running. On Windows you can open the 'Services' program on your system and search for the MYSQL80 process. On Linux based systems you run `sudo systemctl status mysql` and that will show you the status of the server. If the server is not on, use `sudo systemctl start mysql` to start the server.\
To actually connect and use the server, you need to set your password. That can be done by firing the query below into the client.

```sql
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
```

For 'password' you can enter your own password as a string. Or just keep it 'password' idc it's your machine.

#### Create Database
The database can be created by simply by copying and pasting the queries from [this file](create-database.sql).
If the query throws an error, delete your local database and run again (of course make a backup if you care about development data).

#### Environment Variables
For the API to be able to connect to the server, it needs to be able to log in. It retrieves the username and password from the environment variables. How you set these differs per operating system.
In general, open your terminal / command prompt, navigate to the Slijterij project and enter the following commands:

**On Linux (so also MacOS)**
```bash
export DBUSER=username
export DBPASS=password
```

**On Windows**
```bat
set DBUSER=username
set DBPASS=password
```

If you have not updated the username after installation, this will be 'root'. The password will be the same as you entered as the one you entered earlier in the SQL query.

This step can be skipped by inserting a certain file into the project, ask me for more info about "env.go".

## The Fun Part
Once the API runs by using `go run .`, it works. The API will be running on **localhost:8080** and will accept HTTP requests. Note: If it doesn't actually work, send me some screens of the errors you get, I might still need to do some CORS stuff. 

Here follows a list of the endpoints that are to be found in Slijterij. These consist of the URL, what shape of JSON it accepts in the body, what it responds with and the current implementation status.

### Base
```
/
```
This endpoint doesn't really do anything useful. It's just useful for testing if the API is running.

#### GET
Input: None

Output:
* Status Code: 200
    * Body: "Success"

Status: IMPLEMENTED

### Bar
```
/bar
```
This endpoint is for retrieving and managing bars.

#### GET
Input: None

Output:
* Status Code: 200
    * `[{ "id": string, "name": string, "token": string }, { "id": string, "name": string, "token": string }]`
* Status Code: 500
    * `{ "message": string }`

Status: IMPLEMENTED

#### POST
Input: `{ "name": string, "password": string }`

Output:
* Status Code: 200
    * "Row X created" where X is the Row's ID (not the bar's ID)
* Status Code: 400
    * `{ "message": string }`
* Status Code: 500
    * `{ "message": string }`

Status: IMPLEMENTED

#### PUT
For updating names, passwords and tokens\
Input: `{ "id": string, "name": string, "password": string, "token": string }`

Output:
* Status Code: 200
    * `{ "id": string, "name": string, "token": string }`
* Status Code: 400
    * `{ "message": string }`
* Status Code: 404
    * `{ "message": string }`
* Status Code: 409
    * `{ "message": string }`

Status: TO BE IMPLEMENTED

#### DELETE
Input: `{ "id": string }`

Output:
* Status Code: 200
    * `null`
* Status Code: 400
    * `{ "message": string }`
* Status Code: 404
    * `{ "message": string }`

Status: TO BE IMPLEMENTED

### Login
```
/bar/login
```
For logging in
#### POST
Input: `{ "name": string, "password": string }`

Output:
* Status Code: 200
    * `{ "id": string, "name": string, "token": string }`
* Status Code: 400
    * `{ "message": string }`
* Status Code :401
    * `{ "message": string }`
* Status Code: 500
    * `{ "message": string }`

Status: IMPLEMENTED

### Drinks
```
/drinks
```
This endpoint is for retrieving and managing drinks.

#### GET
Query parameter: `barId`

Output:
* Status Code: 200
    * `[ { "id": string, "name": string, "bar_id": string, "start_price": float, "current_price": float, "multiplier": float, "tag" : string } ]`
* Status Code: 400
    * `{ "message": string }`
* Status Code: 500
    * `{ "message": string }`

Status: IMPLEMENTED

#### POST
Input: `{ "id": string, "name": string, "bar_id": string, "start_price": float, "current_price": float, "multiplier": float, "tag" : string }`

Output: 
* Status Code: 201
    * `{ "id": string, "name": string, "bar_id", string, "start_price": float, "current_price": float, "multiplier": float, "tag" : string }`
* Status Code: 400
    * `{ "message": string }`
* Status Code: 409
    * `{ "message": string }`
* Status Code: 500
    * `{ "message": string }`

Status: IMPLEMENTED

#### PUT
Input: `{ "id": string, "name": string, "bar_id": string, "start_price": float, "current_price": float, "multiplier": float, "tag" : string }`

Output:
* Status Code: 200
    * `{ "id": string, "name": string, "bar_id": string, "start_price": float, "current_price": float, "multiplier": float, "tag" : string }`
* Status Code: 400
    * `{ "message": string }`
* Status Code: 500
    * `{ "message": string }`

Status: IMPLEMENTED

#### DELETE
Input: `{ "id": string }`

Output:
* Status Code: 200
    * `null`
* Status Code: 400
    * `{ "message": string }`
* Status Code: 500
    * `{ "message": string }`

Status: IMPLEMENTED
