##golang-task-restful



a RESTful service that will support the maintenance of a simple list of todo’s

###Installation

- After clone this repository into your local, make sure `$GOROOT` and `$GOPATH` is set properly, by adding the follwing lines into your ~/.bashrc or ~/.profile (assume you've installed Go at `/usr/local/go` and cloned the repository to ~/Code/golang-task-restful
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


