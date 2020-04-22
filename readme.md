Go, React and MySQL Boilerplate

## Get Started

### 1. Prerequisites

- [Go](https://golang.org/) - Go programming language
- [NodeJs](https://nodejs.org/en/)
- [NPM](https://npmjs.org/) - Node package manager
- [MySQL](https://www.mysql.com/downloads/) - Relational database management system (RDBMS)

### 2. Installation

On the command prompt run the following commands:

``` 
 $ cd workspace
 $ git clone https://github.com/Bikranshu/go-react-boilerplate.git
 $ cd go-react-boilerplate
 $ cd src
 $ go get github.com/{package-name}  # Download a particular remote package. See 'go help packages' for details
 $ cp env.yaml.example env.yaml # Edit it with your secret key and database information
 $ swag init [Optional] # Run this command if you add entity CRUD opration
 ```
 Finally, start and build the application:
 
 ```
 $ go run main.go
 $ go build 
```

Import `users.sql` for default user:
- username: admin@admin.com
- password: 1234

### 3. Usage

URL : http://localhost:3000/

Navigate to http://localhost:3000/swagger/ for the API documentation.

### 4. Useful Link
- Object Relational Mapping(ORM) - [GORM](http://gorm.io)
- Environment Configuration - [Viper](https://github.com/spf13/viper)
- Structured Logger - [Logrus](https://github.com/sirupsen/logrus)
- JSON Web Tokens - [JWT](https://github.com/dgrijalva/jwt-go)
- Password Hash & Salt - [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)
- Cross Origin Resource Sharing - [CORS](https://github.com/rs/cors)
- URL router and dispatcher - [gorilla mux](https://github.com/gorilla/mux)
- Redis Client - [redis](https://github.com/go-redis/redis)
- API Documentation - [swag](https://github.com/swaggo/swag) and [http-swagger](https://github.com/swaggo/http-swagger)
- Modern build setup with no configuration - [Create React App](https://create-react-app.dev/)
- JavaScript library for building UI - [React](https://facebook.github.io/react/)
- Predictable state container - [Redux](http://redux.js.org/)
- Declarative routing for React - [React-Router](https://reacttraining.com/react-router/)
- Promise based HTTP client - [Axios](https://github.com/mzabriskie/axios)
- React UI library - [Ant Design](https://ant.design/)
- Pluggable JavaScript Linter - [ESLint](http://eslint.org/)
- Opinionated Code Formatter - [Prettier](https://www.npmjs.com/package/prettier)