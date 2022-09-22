# Middleware

## Resume Materi

### Middleware

### Middleware

Middleware is an entity that hooks into server's request/response processing. Middleware allow us to implement various functionalities around incoming HTTP requests to your server and outgoing responses.

### Setup Middleware Echo

Echo#Pre(): Executed before router process the request
- HTTPSRedirect
- HTTPSWWWRedirect
- WWWRedirect
- AddTrailingSlash
- RemoveTrailingSlash
- MethodOverride
- Rewrite

Echo#Use(): Executed after router processes the request and has full access to echo.ContextAPI
- BodyLimit
- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- Secure
- CORS
- Static

### Implementation

1. Log Middleware
    - Logs the information HTTP request
    - As a footprint
    - Datasource for analytics
2. Auth Middleware
    - User identification
    - Secure data and information

    Authentication Middleware:
    1. Basic authentication
        Basic Auth is an http authentication technique request, this method requires username and password information to be inserted in the request header.
    2. JSON Web Token
        JWT authentication is a token-based stateless authentication mechanism. It is popularly used as a client-side-based stateless session, this means the server doesn’t have to completely rely on a data store (or) database to save session information