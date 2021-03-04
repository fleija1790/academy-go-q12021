# Golang Bootcamp
## The Challenge


## First Deliverable (due March 4th 23:59PM)

This deliverable covers:

- Create an API
- Add an endpoint to read from a CSV file
- The CSV should have any information, for example:
- The items in the CSV must have an ID element (int value)
- The endpoint should get information from the CSV by some field 
- The result should be displayed as a response
- Clean architecture proposal
- Use best practices
- Basic Handling of Errors
- Create a client to consume an external API
- Add an endpoint to consume the external API client
- The information obtained should be stored in the CSV file
- Update the endpoint made in the first deliverable to display the result as a JSON
- Refactor if needed

Endpoints URLs:
- List all Jokes: `localhost:3000/api/v1/jokes`
- Get Joke by ID: `localhost:3000/api/v1/jokes/{jokeID}` 
- Get more jokes from external API: `localhost:3000/api/v1/updade-jokes`