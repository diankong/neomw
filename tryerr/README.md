# tryerr
----
tryerr is a middleware to handle errors you appended to ctx.Errors

After your route handle, it'll check if there is any error in ctx.Errors. 

If there is one (or more), it will set statusCode to 500 and return all those error messages to client.

Get
----
````shell
go get github.com/diankong/neomw/tryerr
````

````Go
import "github.com/diankong/neomw/tryerr"
````

Useage
----
````GO
  app := neo.App()
  //use it before any route
  app.Use(tryerr.Try)
  // Your Routes...
  // Start server
  app.Conf.App.Addr = ":3000"
  app.Start()
````

In your route, if you append a error to ctx:
````GO
  app.Get("/info", func(c *neo.Ctx) (int, error) {
    c.Error(fmt.Errorf("error: %s", "a new error"))
    return 200, c.Res.Text("done")
  })
````
After your routing process, it will return status 500 and error msg to client:
````shell
curl http://127.0.0.1:3000/info
````
````Go
//-> error: a new error
````
