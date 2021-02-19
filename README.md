# Test Task Service.

# Topics
* [Prerequisites](#Prerequisites)
* [Set up and run service](#set-up-and-run-service)
* [Environment Variables](#environment-variables)
* [Endpoints](#endpoints) 
  * [Auth Endpoints](#server-endpoints)
    * [POST /auth/signin](#auth-signin)
    * [POST /auth/signup]
    * [GET /auth/signout]
  * [Member Endpoints](#member-endpoints)
    * [GET /members]
  * [Service Endpoints](#service-endpoints)
    * [POST /service/status]
    * [GET /service/info]

* [Appendix A. Response codes](#appendix-a-response-codes)
* [Appendix B. Request headers](#appendix-b-request-headers)

# Prerequisites
- Golang
- MySQL
- Docker
- Make

# Set up and run service

- clone the repository from GitHub.
```
$ git clone https://github.com/dsuhinin/suhinin-backend-1.git
```
- go to the project folder root.
- run `make lint` to run linter over the code.
- run `make docker_build_image service_start` to build and run everything in Docker.
- open in a browser `http://localhost:8081`

# Environment Variables

| Name          		| Required | Example		| Description	| 
| ----------------------------- | --------- | ----------------- | ------------- |		 
| JWT_KEY       		| Y 	    | A8F5F167F44F  	| The Key to sign JWT tokens.|
| MYSQL_USER     		| Y         | root 		| Database user.|
| MYSQL_PASS 			| Y         | password    	| Database password.|
| MYSQL_HOST 			| Y         | localhost 	| Database host.|
| MYSQL_PORT 			| Y         | 3306 		| Database port.|
| MYSQL_DB_NAME 		| Y         | database_name 	| Database name.|
| SERVER_ADDRESS 		| N         | 127.0.0.1:8080 	| Server address and port where service is listening for incoming requests. If not provided then `127.0.0.1:8080` will be used.|
| LOG_LEVEL 			| Y         | debug 		| Logger level. Possible values are: debug, info, warn, error.|
| CORS_ENABLED 			| N         | true 		| Enables CORS.|
| CORS_ENABLED_DEBUG 	| N         | true 		| Enables CORS debug info.|
| CORS_ALLOWED_ORIGINS 	| N         | http://localhost:8081 		| Provides allowed origings.|
| CORS_ALLOWED_METHODS 	| N         | POST,PUT,PATCH,DELETE 		| Provides allowed http methods.|
| CORS_ALLOWED_HEADERS 	| N         | Content-Type,Authorization 	| Provides allowed http headers.|

# Endpoints
# Auth Endpoints
## POST /auth/signin
Endpoint provides `signin` functionality.

**Request info**

```
HTTP Request method    POST
Request URL            https://localhost:8080/auth/signin
```

**Request body**

```
{
	"email": "user.email@gmail.com",
	"password": "password"
}
```

**Response body**

```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTM3NDgxNDAsImp0aSI6IjEiLCJpc3MiOiJzdWhpbmluLmRtaXRyaXlAZ21haWwuY29tIn0.9HUQFb4MlcQVHvXH6XiJ82FhQ3JylmXFHKHL52VuIj4"
}
```


## POST /auth/signup
Endpoint provides `signup` functionality.

**Request info**

```
HTTP Request method    POST
Request URL            https://localhost:8080/auth/signup
```

**Request body**

```
{
    "email": "user.email@gmail.com",
	"password": "password",
	"confirm_password": "password"  
}
```

**Response body**

```
{}
```

## GET /auth/signout
Endpoint provides `signout` functionality.

**Request info**

```
HTTP Request method    GET
Request URL            https://localhost:8080/auth/signout
Headers		           Authorization Bearer token
```

**Request body**

```
{}
```
**Response body**

```
{}
```

# Member Endpoints
## GET /members
Endpoint to get `members` data.

**Request info**

```
HTTP Request method    GET
Request URL            https://localhost:8080/members
Headers		           Authorization Bearer token
```
**Note**
For more information about `Authorization` header [Appendix B. Request headers](#appendix-b-request-headers)

**Request body**

```
{}
```

**Response body**

```
{}
```

# Service Endpoints
## GET /service/info
The endpoint to get service info.

**Request info**

```
HTTP Request method    GET
Request URL            https://localhost:8080/service/info
```
**Request body**

```
{}
```

**Response body**

```
{
  "dependencies": {
    "mysql": {
      "status": 200,
      "latency": 0.002757021
    }
  }
}
```
**Note**
Endpoint give the full picture about `service` current state. Endpoint mostly needs for Kubernetes readiness probe.

## GET /service/status
The endpoint to get service status.

**Request info**

```
HTTP Request method    GET
Request URL            https://localhost:8080/service/status
```

**Request body**

```
{}
```

**Response body**
```
{}
```
**Note**
Endpoint give the full picture about `service` current state. Endpoint mostly needs for Kubernetes liveness probe.

# Appendix A. Response codes

### HTTP error codes

Application uses standard HTTP response codes:
```
200 - Success
201 - Created
400 - Request error
401 - Authentication error
404 - Entity not found
500 - Internal Server error
```

Additional information about the error is returned as JSON-object like:
```
{
    "code": "{error-code}",
    "message": "{error-message}"
}
```
# Appendix B. Request headers
HTTP header required for API calls:
- `Authorization` : `Bearer token` - authorization based on `JWT` token.
