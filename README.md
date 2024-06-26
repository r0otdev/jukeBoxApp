# Jukebox Application

Programming Language : Golang (Minimum 1.20)

Data storage : In-memory storage as of now

## How to run

go run .

## How to run Test

go test ./app

## APIs

1. GET : /v1/api/albums/ 

	Get all the albums in the system sorted by release date

2. POST : /v1/api/albums/ 

	Create a new record for an album in the system

	Sample post json body will be like;

	{
		"name": "Album",
		"genre": "POP",
		"release_date": "2023-01-05T12:00:00Z",
		"price": 120,
		"description": "Test",
		"musicians": [1001,1002]
	}


3. PATCH : /v1/api/albums/{albumID}

	Update a corresponding album, albumID should be supplied in url. Partial album request body is allowed

4. GET : /v1/api/albums/{albumID}/musicians

	GET all associated musicians of a album

5. GET : /v1/api/musicians/ 

	Get all the musicians in the system sorted by release date

6. POST : /v1/api/musicians/ 

	Create a new record for an musician in the system

	Sample post json body will be like;

	{
		"name": "Album",
		"type": "Singer"
	}


7. PATCH : /v1/api/musicians/{musicianID}

	Update a corresponding musician, musicianID should be supplied in url. Partial musician request body is allowed

8. GET : /v1/api/musicians/{musicianID}/albums

	GET all associated albums of a musician

## TODO

1. Implement Logger
2. Implement any database provision
3. Implement middlewares for request validation
