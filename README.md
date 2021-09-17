## How to run the application

To run the application, you need to install docker and docker-compose on your computer. Once your docker-compose is installed run the following command:

```
    $ docker-compose up -d
```

Your application should then start. You can make changes in the src/ directory, build using normal golang commands. Much of the work has already been done such as writing to the database, but you will need to create
the api endpoint to query and return the messages when a websocket client connects and sends a message to the server.

To build the application you can do the following:

```
   $ docker exec -it golang_chatapp_1 /bin/bash
   $ go build cmd/chatapp/main.go
   $ ./main
```

## How you will be graded

[] Can you create an api endpoint http://localhost:3000/logs that properly returns the log text in json format?
[] Can you create some kind of golang test to show your application works?
[] Can you create a script that uploads data to the mongo database and test your function?


## Download Mongo Database
You can download the MongoDB config from this link:
https://drive.google.com/file/d/1cUpoX8Hdn31gdl1md53EWmvX9Xnd_4aW/view?usp=sharing