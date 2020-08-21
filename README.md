# rest-websocket-tester [![Build Status](https://api.travis-ci.com/b00lduck/rest-websocket-tester.svg)](https://travis-ci.com/github/b00lduck/rest-websocket-tester)

Testserver which sends incoming HTTP requests over active websocket connections

## Try it out

First run
```
go run cmd/app/main.go
```

Then go to [https://www.websocket.org/echo.html] and connect to ws://localhost:8080


You may also run this to send a message to the websocket:
```
curl -d "HAHA" localhost:8080
```