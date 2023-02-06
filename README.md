# Recipes API 

**This API provides all that you need to create and manage a list of recipes**

<img src="./images/swedish-chef.gif" width="450" height="250" alt="happy-chef"/>

This API is built with Golang and Postgres. It can be run locally using docker which can be download [here](https://www.docker.com/) 

- Once Docker is installed, to build the API run from the root of the project  `docker-compose up --build -d` or `docker-compose up --build` to help with debugging.

- Run `docker-compose down` to stop the API and Postgres database or `docker-compose down --volumes` to also remove the database volumes.

**List of endpoints**

To use this API you will first need to register an account, this can be done with the following:

```
curl -X POST -H "Content-Type: application/json" -d '{"firstname": "test", "lastname": "user","email": "testuser123@gmail.com", "password": "mypassword123"}' http://localhost:8080/register
```

You can then login with:

```
curl -X POST -H "Content-Type: application/json" -d '{"email": "testuser123@gmail.com", "password": "mypassword123"}' http://localhost:8080/login
```

This will give you a JSON Web Token (JWT) that will grant you access to your recipes.

To get a list of your recipes use:

```
curl -X GET 'http://localhost:8080/recipe' \
   -H 'Content-Type: application/json' \
   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU3MDc1MTksImlhdCI6MTY3NTcwMzkxOSwicmVjaXBlX3VzZXJfaWQiOjY3fQ.0R-DR8YQKkmj9YX4JcBjamxyBIptiQv-NdRfwD3jfzg' 
```

To get a specific recipe use:

```
curl -X GET 'http://localhost:8080/recipe/{id}' \
   -H 'Content-Type: application/json' \
   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU3MDc1MTksImlhdCI6MTY3NTcwMzkxOSwicmVjaXBlX3VzZXJfaWQiOjY3fQ.0R-DR8YQKkmj9YX4JcBjamxyBIptiQv-NdRfwD3jfzg' 
```

To save a new recipe, you can use:

```
curl -X POST http://localhost:8080/recipe \
   -H 'Content-Type: application/json' \
   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU3MDc1MTksImlhdCI6MTY3NTcwMzkxOSwicmVjaXBlX3VzZXJfaWQiOjY3fQ.0R-DR8YQKkmj9YX4JcBjamxyBIptiQv-NdRfwD3jfzg' \
   -d '{"recipeName": "A new recipe", "recipeSteps": "some steps for the recipe"}'
```

You can update an existing recipe with

```
curl -X PUT http://localhost:8080/recipe/{id} \
   -H 'Content-Type: application/json' \
   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU3MDc1MTksImlhdCI6MTY3NTcwMzkxOSwicmVjaXBlX3VzZXJfaWQiOjY3fQ.0R-DR8YQKkmj9YX4JcBjamxyBIptiQv-NdRfwD3jfzg' \
   -d '{"recipeName": "A new recipe 123", "recipeSteps": "some steps for the recipe 123"}'
```

Finally a recipe can be deleted with this:

```
curl -X DELETE http://localhost:8080/recipe/{id} \
   -H 'Content-Type: application/json' \
   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU3MDc1MTksImlhdCI6MTY3NTcwMzkxOSwicmVjaXBlX3VzZXJfaWQiOjY3fQ.0R-DR8YQKkmj9YX4JcBjamxyBIptiQv-NdRfwD3jfzg' 
```



