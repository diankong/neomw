# tier
----
tier is a middleware to try catch panic, so you won't worry about panic anymore.

Get
----
````shell
go get github.com/diankong/neomw/trier
````

````Go
import "github.com/diankong/neomw/trier"
````

Useage
----
````GO
  app := neo.App()
  //use it before any route
  app.Use(trier.Try)
  // Your Routes...
  // Start server
  app.Conf.App.Addr = ":3000"
  app.Start()
````

In your route, if there is a panic:
````GO
  app.Get("/info", func(c *neo.Ctx) (int, error) {
    panic("there is a panic")
    return 200, c.Res.Text("done")
  })
````
Your routing process won't crash, it will return status 500 and error msg to client:
````shell
curl http://127.0.0.1:3000/info
````
````Go
//-> there is a panic
````
