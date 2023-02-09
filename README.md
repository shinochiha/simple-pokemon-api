# Simple Pokemon API

## Getting Started
1. Make sure you have [Go](https://go.dev) and Postgresql installed .
2. Clone the repo
```bash
git clone 
```
3. Go to the directory and run go mod tidy to add missing requirements and to drop unused requirements
```bash
cd api-pokemon && go mod tidy
```
3. Setup your .env file
```bash
cp .env-example .env && change dbname
```
4. Start
```bash
go run main.go
```
## Build for production
1. Compile packages and dependencies
```bash
go build -o api-pokemon main.go
```
2. Setup .env file for production
```bash
cp .env-example .env && vi .env
```
3. Run executable file with systemd, supervisor, pm2 or other process manager
```bash
./api-pokemon
```

## How to use 
1. Login User with
```
{
    "email":"test@gmail.com",
    "password":"123456"
}
```

2. Get List Pokemon
```
GET http://localhost:4000/api/v1/pokemons
```
Note : Wait a few minutes until the process is complete, until the data of 1279 Pokemon is successfully entered into the database.

3. Get Detail Pokemon
```
GET http://localhost:4000/api/v1/pokemons/{id} // Select the ID of the Pokemon you want to get detail

4. Capture Pokemon
```
POST http://localhost:4000/api/v1/capture_pokemon/{id} // Select the ID of the Pokemon you want to catch.
```

5. Save My Pokemon and Provide the name of the Pokemon.
```
PUT http://localhost:4000/api/v1/save_pokemon/1
```
```
{
    "nickname":"My-ivysaur"
}
```

6. Get My pokomen
```
GET http://localhost:4000/api/v1/mypokemons
```

7. Change the name of the pokemon you have
```
http://localhost:4000/api/v1/change_name_pokemon/1
```
```
{
    "nickname":"My-bulbasaur"
}
```

8. Realease my pokemon
```
DELETE http://localhost:4000/api/v1/release_pokemon/1
```

Finish