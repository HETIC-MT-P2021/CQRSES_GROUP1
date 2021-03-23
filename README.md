# Our school project  🎢

The goal of this project is to create an optimized API using two patterns: CQRS & Event sourcing.
### Patterns used 
CQRS separates reads and writes into different models, using commands to update data, and queries to read data.

The Event Sourcing pattern defines an approach to handling operations on data that's driven by a sequence of events, each of which is recorded in an append-only store. 
Application code sends a series of events that imperatively describe each action that has occurred on the data to the event store, where they're persisted. Each event represents a set of changes to the data. 
### Consumer repository
The consumer wait to receive messages from the message broker host in the main project 

[Consumer link](https://github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer)
## Features 🌈

- User authentication with JWT.
- User management.
- Article management.


## Starting the project 🚀

After cloning the repo, `cd` into the project, create the .env according to .env.example, and run the following commands

```bash
docker-compose up --build
```

After that : 
Go to http://localhost:8091/

Enter the env COUCH_USER/COUCH_PASSWORD

Create Two Buckets as CouchBase:

- event-store
- read-models

Restart the docker-compose
```bash
docker-compose restart
```
### Endpoints  🔀
Healthcheck ❤️

```http request
GET /
```
Login 🔒
```http request
POST /login
```

#### Users 👦
Create a new user
```http request
POST /users
```
Get all users
```http request
GET /users
```
Get a user by id 
```http request
GET /users/{id}
```
Update a user
```http request
PUT /users/{id}
```
Delete a user
```http request
DELETE /users/{id}
```
#### POSTS 📰
Create a new post
```http request
POST /posts
```
Get all posts
```http request
GET /posts
```
Get a post by id 
```http request
GET /posts/{id}
```
Update a post
```http request
PUT /posts/{id}
```
Delete a post
```http request
DELETE /posts/{id}
```
### API Demo ✨
The API doc is available [here](https://documenter.getpostman.com/view/14693906/TzCFhqn8)


### Documentation

You can find the technical and functional documentation in French and English version by clicking on the link below:

[LINK DRIVE](https://drive.google.com/drive/folders/1ZGYXEniZO2mb2mk9MUTHwAJNXuh1KYUC?usp=sharing)

### Technical Choices 🔧

Feel free to discuss with any contributor about the technical choices that were made:


- Go Version: `1.15`
- PostgreSQL database: `13`
- CouchBase database: `5.5.x`


## Contributing 💡
Share with us your ideas to improve our project !

See [contributing](https://github.com/HETIC-MT-P2021/CQRSES_GROUP1/blob/main/CONTRIBUTING.MD) guidelines.

### Authors 🏄 

- [Tsabot](https://github.com/Tsabot)
- [myouuu](https://github.com/myouuu)
- [acauchois](https://github.com/acauchois)
- [gensjaak](https://github.com/gensjaak)

### Tasks management 🎨
We like to iterate quickly, setting up an agile board helps everyone stay on task.  With Trello, we Created lists for backlogged items, what’s being worked on in the current sprint, and (most importantly) what’s completed.

You can take a look at our organization of tasks on [Trello](https://trello.com/b/uY6KOh4i/go-cqrs)


### License 🔖

The code is available under the MIT license.
