## Movies CRUD API with Golang

This project implements a CRUD (Create, Read, Update, Delete) API for managing movies using Golang. It utilizes the Gorilla mux router for handling HTTP requests and JSON for data exchange.

### Features

- **GET /movies**: Retrieve all movies in the database.
- **GET /movies/{id}**: Retrieve a specific movie by its ID.
- **POST /movies**: Create a new movie entry.
- **PUT /movies/{id}**: Update an existing movie's information.
- **DELETE /movies/{id}**: Delete a movie from the database.

### Movie Structure

Each movie entry includes:
- **ID**: Unique identifier for the movie.
- **ISBN**: ISBN code of the movie.
- **Title**: Title of the movie.
- **Director**: Object containing director's first name and last name.

### Getting Started

To run the API locally:
1. Clone this repository.
2. Ensure you have Golang installed on your machine.
3. Run `go run main.go` in your terminal.
4. Access the API endpoints using tools like Postman or cURL.

### Contributions

Contributions are welcome! If you have any suggestions, improvements, or bug fixes, feel free to open an issue or submit a pull request.
