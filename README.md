##golang-task-restful

A RESTful Service that Supports the Maintenance of A Simple List of TODO’s

###Installation

- After cloning this repository into your local, make sure `$GOROOT` and `$GOPATH` is set properly, by adding the follwing lines into your ~/.bashrc or ~/.profile (assume you've installed Go at `/usr/local/go` and cloned the repository to ~/Code/golang-task-restful)
```
$export GOROOT=/usr/local/go
$export GOPATH=~/Code/golang-task-restful
```
then `source ~/.bashrc` or `source ~/.profile` or restart your terminal

- golang-task-restful uses sqlite3 as backend storage, make sure you have sqlite3 installed on your machine. There's issue to `go get github.com/mattn/go-sqlite3` on Mac plateform using default copmiler, which is the command line tool of Xcode, please refer [here](https://github.com/mattn/go-sqlite3/issues/45). So Mac is not supported for current vertion.
```
$sudo apt-get install sqlite3
```

- golang-task-restful uses [Martini](http://martini.codegangsta.io) framework to route URLs, as well as [martini-contrib/binding](https://github.com/codegangsta/martini-contrib/tree/master/binding) and [martini-contrib/render](https://github.com/codegangsta/martini-contrib/tree/master/render) to receive parameters of URL and generate returning JSON respetively.
Execute the following command when you're in your local cloned repository
```
$go get github.com/codegangsta/martini
$go get github.com/codegangsta/martini-contrib/binding
$go get github.com/codegangsta/martini-contrib/render
```
Then execute the following one for sqlite3 driver
```
$go get github.com/mattn/go-sqlite3
```
- `$go run main.go` to enjoy it!

### Usage Examples


- To create new task:
```
$curl -X POST -H "Content-Type: application/json" -d '{"description":"test_des","due":"2020-01-01 10:00:00","completed":false}' http://localhost:3000/task/add/data.json
```
Then you will get `{"status":"success"}` or `{"status":"error", "message": MESSAGE}`, in which MESSAGE represents the actual error message.

- To query all the tasks:
```
$curl -X GET http://localhost:3000/task/list
```
Then you will get output similar to:
```
{"results":[{"description":"text","due":"2008-01-01 10:00:00","completed":false},{"description":"second task","due":"2014-04-22 10:51:00","completed":true},{"description":"cont","due":"2018-01-01 10:00:00","completed":false},{"description":"cont2","due":"2028-01-01 10:00:00","completed":false},{"description":"","due":"2050-01-01 10:00:00","completed":false},{"description":"test_des","due":"2050-01-01 10:00:00","completed":false},{"description":"test_des","due":"2020-01-01 10:00:00","completed":false}],"status":"success"}
```
The value of `"results"` is a JSON list. Or you may get `{"status":"error", "message": MESSAGE}` with error message.

- To fetch certain task by denoting its id (here use 1 as example)
```
$curl -X GET http://localhost:3000/task/1
```
Then you will get:
```
{"results":{"description":"text","due":"2008-01-01 10:00:00","completed":false},"status":"success"}
```
or `{"status":"error", "message": MESSAGE}` with error message.

- To update certain task by denoting its id (here also use 1 as example)
```
$curl -X POST -H "Content-Type: application/json" -d '{"description":"test_des","due":"2050-01-01 10:00:00","completed":false}' http://localhost:3000/task/add/data.json
```
Then you will get `{"status":"success"}` or `{"status":"error", "message": MESSAGE}` with error message.


- To delete certain task by denoting its id (here also use a as example)
```
$curl -X DELETE http://localhost:3000/task/delete/1
```
Then you will get `{"status":"success"}` or `{"status":"error", "message": MESSAGE}` with error message.


###References
- [JSend](http://labs.omniti.com/labs/jsend) A specification that lays down some rules for how JSON responses from server should be formatted.
- [Vim plugins for golang](golang.org/misc/vim/readme.txt)


###Author

**[Wei Lin](http://www.github.com/ivanlw)**

