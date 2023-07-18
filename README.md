**This is a Go based CRUD API to store and process movies**

# API Documentation

### 1. GET ALL MOVIES
- Base URL: `https://go-crud-api.cleverapps.io`
- Endpoint: `/movies`
- Method: `GET`
- Example Response:
  ```
  200 OK
  [
    {
        "id": "2",
        "isbn": "831479",
        "title": "Inception",
        "director": {
            "firstName": "Christopher",
            "lastName": "Nolan"
        }
    },
    {
        "id": "3",
        "isbn": "15740661",
        "title": "Avatar",
        "director": {
            "firstName": "James",
            "lastName": "Cameron"
        }
    }
  ]
  ```
<br>
  
### 2. GET MOVIE BY ID
- Base URL: `https://go-crud-api.cleverapps.io`
- Endpoint: `/movies/[id: numeric string]`
- Method: `GET`
- Example Request: `/movies/2`
- Example Successful Response:
  ```
  200 OK
  {
    "id": "2",
    "isbn": "831479",
    "title": "Inception",
    "director": {
        "firstName": "Christopher",
        "lastName": "Nolan"
    }
  }
  ```
- Example Unsuccessful Response: `404 NOT FOUND`
<br>



### 3. CREATE MOVIE
- Base URL: `https://go-crud-api.cleverapps.io`
- Endpoint: `/movies`
- Method: `POST`
- Example Request Body:
  ```
  {
    "title": "movie_name",
    "director": {
        "firstName": "director_first_name",
        "lastName": "director_last_name"
    }
  }
  ```
- Example Successful Response:
  ```
  200 OK
  {
    "id": "3", // plain counter
    "isbn": "831479", //random
    "title": "movie_name",
    "director": {
        "firstName": "director_first_name",
        "lastName": "director_last_name"
    }
  }
  ```
<br>



### 4. UPDATE MOVIE
- Base URL: `https://go-crud-api.cleverapps.io`
- Endpoint: `/movies/[id: numeric string]`
- Method: `PUT`
- Exmaple Request Query: `/movies/4`
- Example Request Body:
  ```
  {
    "title": "updated_movie_name",
    "director": {
        "firstName": "updated_director_first_name",
        "lastName": "updated_director_last_name"
    }
  }
  ```
- Example Successful Response:
  ```
  200 OK
  {
    "id": "4", // unchanged
    "isbn": "831479", // unchanged
    "title": "updated_movie_name",
    "director": {
        "firstName": "updated_director_first_name",
        "lastName": "updated_director_last_name"
    }
  }
  ```
- Example Unsuccessful Response: `404 NOT FOUND`
<br>



### 5. DELETE MOVIE
- Base URL: `https://go-crud-api.cleverapps.io`
- Endpoint: `/movies/[id: numeric string]`
- Method: `DELETE`
- Exmaple Request Query: `/movies/4`
- Example Successful Response: `200 OK`
- Example Unsuccessful Response: `404 NOT FOUND`
<br>
