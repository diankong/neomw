# jwt
----
jwt is a middleware for json web token

## Get
----
````shell
go get github.com/diankong/neomw/jwt
````

````Go
import "github.com/diankong/neomw/jwt"
````

## Useage
----
### Verify a token: 
````GO
  app := neo.App()
  app.Use(jwt.Jwt("your_key"))
````
If there is no valid token in Header, it will return an error msg with status 401

If there is one, then it'll parse token to user info, and store it into ctx.Session.User

So, in your route handle, you can use it like this:

````GO
  user := c.Session.User.(map[string]interface{})
  fmt.Println(user["name"])
````

### Sign a token

In your route, if there is a panic:

````GO
    //information you want to store in token, usually a user name or anything you want
    user := map[string]interface{}{"name": "your_name"}
    //the same key you used to verify it
    token, err := jwt.Sign("your_key", user)
````
Then, you can send this token to client, client need to put it into Header

````javascript
Headers["Authorization"] = "Bearer "+ token;
````

## Notice

----

In neo, middlewares are runned before your routes
So you need set Regions, otherwise every route gonna be verified by jwt

Check here: 
[http://ivpusic.github.io/neo/tutorials/2015/01/22/regions.html](http://ivpusic.github.io/neo/tutorials/2015/01/22/regions.html)

## More

----

[about json web token](http://jwt.io/)

[introduce jwt](http://jwt.io/introduction/)


