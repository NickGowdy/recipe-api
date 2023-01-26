# Recipes API 

**This API provides all that you need to create and manage a list of recipes**

<img src="./images/swedish-chef.gif" width="450" height="250" alt="happy-chef"/>

This API is built with Golang and Postgres. It can be run locally using docker which can be download [here](https://www.docker.com/) 

- Once Docker is installed, to build the API run from the root of the project  `docker-compose up --build -d` or `docker-compose up --build` to help with debugging.

- Run `docker-compose down` to stop the API and Postgres database or `docker-compose down --volumes` to also remove the database volumes.

**List of endpoints*

To get a list of your recipes use:
```
curl http://localhost:8080/recipes

```

To get a specific recipe use:
```
curl http://localhost:8080/recipes/{id}

```

More to come later....





