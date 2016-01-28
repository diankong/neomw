// Middleware to try catch panic
// So you don't need to worry about panic anymore
package trier

import (
	"fmt"
	"github.com/ivpusic/neo"
	"github.com/manucorporat/try"
)

func Try(ctx *neo.Ctx, next neo.Next) {

	try.This(func() {
		next()
	}).Catch(func(e try.E) {
		// Print error
		ctx.Res.Status = 500
		ctx.Res.Text(fmt.Sprint(e))
	})

}
