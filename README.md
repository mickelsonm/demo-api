# Demo API

This is a demonstration API using MySQL as a backend RDMS.

# How to build the application?

Using sraight up Go:

    go get

    go run main.go

Using Docker:

    docker build -t demoapi .

    docker run -d -p 9000:8080 --env-file=$(pwd)/.appenv --restart=always --name demoapi demoapi

    Note: .appenv would be a file with all the database settings, etc.

# How to set up the database?

You will need to take care of the MySQL setup (Docker works great).

Use the `setup-db-schema.sql` script to utilize the schema followed within the API.


# Example Interactions via curl

GET All Tickets:

    curl -X GET http://192.168.99.100:9000/tickets

GET a specific ticket

    curl -X GET http://192.168.99.100:9000/tickets/1

Create a ticket

    curl -H "Content-Type: application/json" -d '{"short_desc":"fixing things","long_desc":"we are working on getting things fixed"}' http://192.168.99.100:9000/tickets

Update a ticket

    curl -H "Content-Type: application/json" -X PUT -d '{"id": 2,"short_desc":"fixing computer","long_desc":"we are working on getting the computer fixed"}' http://192.168.99.100:9000/tickets

Delete a ticket

    curl -H "Content-Type: application/json" -X DELETE -d '{"id": 2,"short_desc":"fixing computer","long_desc":"we are working on getting the computer fixed"}' http://192.168.99.100:9000/tickets

# License

MIT - Not that it matters for a demo app right? :)
