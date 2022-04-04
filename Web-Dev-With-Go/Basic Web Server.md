# Web Dev With Go 

My notes from Jon Calhoun's course.

I will work with Jon's course since I believe it will help to complete some potential projects I have in mind.

<br/>

## Basic Web Server

Here is the source code for the server:

```
package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to my site </h1>")
}

func main() {

	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)

}
```

Here we run it:
```
❯ go run main.go
Starting the server on :3000...
```


We can go either to the browser and verify or for now just run a faster verification with curl:
```
❯ curl localhost:3000
<h1> Welcome to my site </h1>%
```


