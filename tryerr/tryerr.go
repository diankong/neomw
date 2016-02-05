// Middleware to handle errors in ctx.Errors
package tryerr

import (
	"fmt"
	"github.com/ivpusic/neo"
)

func Try(c *neo.Ctx, next neo.Next) {

	next()

	if c.HasErrors() {
		c.Res.Status = 500
		msg := ""
		for _, e := range c.Errors {
			msg += e.Error() + "\n"
		}
		c.Res.Text(fmt.Sprint(msg))
	}
}
