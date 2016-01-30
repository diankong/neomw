// Middleware for jwt
package jwt

import (
	"fmt"
	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/ivpusic/neo"
)

// key is the secret key for decoding token
// when done, info from token will be stored in ctx.Session.User
// In your route handle, use it like this:
// user := c.Session.User.(map[string]interface{})
// fmt.Println(user["name"])
func Jwt(key string) neo.Middleware {
	return func(c *neo.Ctx, next neo.Next) {

		if key == "" {
			c.Res.WriteHeader(500)
			c.Res.Text("can't find key to handle jwt")
			return
		}

		//t will be the Token object we need
		t, err := jwt_go.ParseFromRequest(c.Req.Request, func(token *jwt_go.Token) (interface{}, error) {
			return []byte(key), nil
		})

		if err != nil {
			fmt.Println(err.Error())
			c.Res.WriteHeader(401)
			c.Res.Text(err.Error())
			return
		}

		//user info is in the token t now
		if !t.Valid {
			fmt.Println("invalid token")
			c.Res.WriteHeader(401)
			c.Res.Text("invalid token")
			return
		}

		// Store token claims to ctx
		c.Session.Authenticated = true
		c.Session.User = t.Claims
		next()
		return
	}
}

// key is the secret key for encoding token
// claims is the infomation you want to store in token, usually, a user's name or more
// for example:
// user := map[string]interface{}{"name": "yourname"}
// token, err := jwt.Sign("yourkey", user)
// In neo, middlewares are runned before any route,
// so you need set Regions, otherwise every route gonna be verified by jwt
// check it here: http://ivpusic.github.io/neo/tutorials/2015/01/22/regions.html
func Sign(key string, claims map[string]interface{}) (string, error) {
	// Create the token
	token := jwt_go.New(jwt_go.SigningMethodHS256)
	// Set claims
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(key))
}
