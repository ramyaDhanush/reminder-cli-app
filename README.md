# reminder-cli-app
A reminder application built in Golang


# Notifier Service

 A HTTP server which receives a input & sends OS notifications

**Dependencies**

1. Node.js
2. Yarn or NPM
3. express
4. body-parser (Middleware to transform request body into JSON to use it inside controller)
5. node-notifier

**Endpoints**

1. GET /health
2. POST /notify

