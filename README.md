## Purpose

The purpose of this repository is hold files used for exploring go gin package for creating API's/webservers/webpages.

## Result of exploration

There is a simple webserver and some API's.

To see the result:

* run binary in bin/ (one for linux, one for windows)

```sh
./practice-go-gin
```

* build the application yourself (**dependencies:** go) and then run the executable.
	* for linux: 
```sh
GOOS=linux GOARCH=amd64 go build -o practice-go-gin main.go
```
	* for windows:
```sh
GOOS=windows GOARCH=amd64 go build -o practice-go-gin.exe main.go
```

After that, there are multiple ways to access the webserver/API's:

* go to `localhost:8080` in your browser
* use `curl` to query the `localhost:8080` and send GET/POST requests. Examples:

```sh
curl localhost:8080 # main page
curl localhost:8080/hi # says hi
http://localhost:8080/idwp # get idnames in html
http://localhost:8080/idapi # get idnames in json
http://localhost:8080/newid/5/test5 # send GET request to add new idname
http://localhost:8080/newid/w_params?id=5&name=test5 # send GET request to add new idname
curl localhost:8080/newid \
	--request "POST" \
	--header "Content-Type: application/json" \
	--data '{"id": 5,"name": "test5"}' # send POST request to add new idname
```

## Author

Meelis Utt
