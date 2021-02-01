# Our school project

The goal of this project is to create a simple API with user creation and authentification.
We used Golang for our API, and Postgresql as a database.

## Features

- User authentication with JWT.

## Starting the project

After cloning the repo, `cd` into the project, create the .env according to .env.example, and run the following commands

```bash
docker-compose up --build
```

Go to http://localhost:8091/
Enter the env COUCH_USER/COUCH_PASSWORD
Create Two Buckets as CouchBase:

- event-store
- read-models

Restart the docker-compose

### Documentation

You can find the api doc by clicking on the link below :

[Swagger](https://app.swaggerhub.com/apis-docs/acauchois/GoTemplate/1.0.0)

### Technical Choices

Feel free to discuss with any contributor about the technical choices that were made.

- Go version: `1.15`
- PostgreSQL: `13`

## Contributing

See [CONTRIBUTING.MD](https://github.com/HETIC-MT-P2021/CQRSES_GROUP1/blob/main/CONTRIBUTING.MD)

### Authors

- [Tsabot](https://github.com/Tsabot)
- [myouuu](https://github.com/myouuu)
- [acauchois](https://github.com/acauchois)
- [gensjaak](https://github.com/gensjaak)

### License

The code is available under the MIT license.
