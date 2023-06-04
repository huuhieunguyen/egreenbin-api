# eGreenBin

# Golang Gin REST API

This is a simple REST API built with Golang (Gin framework) that allows users to create, read, update, and delete collection in database.

## Dependencies

This project requires the following dependencies to be installed:

- Golang ([https://go.dev/](https://go.dev/))
- Gin framewrok (for handling HTTP requests) ([https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin))
- MongoDB Go driver (for connecting to MongoDB database) ([https://github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver))

## Installation

1. Clone this repository to your local machine
2. Install the dependencies using go modules. Run the following command in the root directory of the project:
    
    ```go
    go get ./...
    ```
    

## Configuration

1. Create a `.env` file in the root directory of the project and add the following variables:
    
    ```go
    DB_DATABASE=<your_mongodb_database>
    DB_USERNAME=<your_mongodb_username>
    DB_PASSWORD=<your_mongodb_password>
    MONGO_URI=<your_mongodb_uri>
    PORT=<port_number>
    ```
    
    Replace `<your_mongodb_uri>` with the URI of your MongoDB database, and `<port_number>` with your desired port number.
    
2. Whenever you make a `.env` file update, use the following command to activate the latest values:
    
    ```go
    source .env
    ```
    

## Usage

1. Run the server by running the following command in the root directory of the project:
    
    ```go
    go run main.go
    ```
    
2. Create a student by sending a POST request to `http://localhost:<port_number>/api/students`
    
    Sample request:
    
    ```json
    {
        "code": "20520514",
        "name": "eGreenBin",
        "numOfCorrect": 25,
        "numOfWrong": 19,
        "imageAvatarUrl": "https://afamilycdn.com/2019/9/25/photo-1-15693889404421259167070.jpg?fbclid=IwAR2YuYMfdc_RazmNjtgWKej14GDwFMn4xnjzu-cWmy5lRN2eLhXEgp-SkQc",
        "parentEmail": "phanvanminh1234567890@gmail.com",
        "Note": ""
    }
    
    ```
    
    Sample response:
    
    ```json
    {
        "data": {
            "id": "6416c301131dc564c6c32r11",
            "code": "20520514",
            "name": "eGreenBin",
            "numOfCorrect": 25,
            "numOfWrong": 19,
            "imageAvatarUrl": "https://afamilycdn.com/2019/9/25/photo-1-15693889404421259167070.jpg?fbclid=IwAR2YuYMfdc_RazmNjtgWKej14GDwFMn4xnjzu-cWmy5lRN2eLhXEgp-SkQc",
            "parentEmail": "phanvanminh1234567890@gmail.com",
            "note": ""
        }
    }
    ```
    
3. Retrieve a student by sending a GET request to `http://localhost:<port_number>/api/students/{id}`
    
    Sample response:
    
    ```json
    {
        "data": {
            "id": "6416c301131dc564c6c32r11",
            "code": "20520514",
            "name": "eGreenBin",
            "numOfCorrect": 25,
            "numOfWrong": 19,
            "imageAvatarUrl": "https://afamilycdn.com/2019/9/25/photo-1-15693889404421259167070.jpg?fbclid=IwAR2YuYMfdc_RazmNjtgWKej14GDwFMn4xnjzu-cWmy5lRN2eLhXEgp-SkQc",
            "parentEmail": "phanvanminh1234567890@gmail.com",
            "note": ""
        }
    }
    ```
    
4. Update a student by sending a PUT request to `http://localhost:<port_number>/api/students/{id}`
    
    Sample request:
    
    ```go
    {
        "code": "205201",
        "name": "eGreenBin",
        "numOfCorrect": 25,
        "numOfWrong": 19,
        "imageAvatarUrl": "https://afamilycdn.com/2019/9/25/photo-1-15693889404421259167070.jpg?fbclid=IwAR2YuYMfdc_RazmNjtgWKej14GDwFMn4xnjzu-cWmy5lRN2eLhXEgp-SkQc",
        "parentEmail": "phanvanminh1234567890@gmail.com",
        "Note": ""
    }
    
    ```
    
5. Delete a student by sending a DELETE request to `http://localhost:<port_number>/api/students/{id}`

## Contributing

If you would like to contribute to this project, please fork the repository and make a pull request.

## Authors

- ******************************Nguyen Huu Hieu******************************
