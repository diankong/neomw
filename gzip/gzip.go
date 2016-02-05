// Middleware for gzip
// This mw basically is copied from echo's middleware
// check more for echo: https://github.com/labstack/echo
// gin's gzip middleware is simple, but it need ResponseWrite can be handled in middleware, which both echo and neo can't
// (Because gin return a ResponseWrite directly, but echo and neo get RespWriter by function: resp.Writer() )
package gzip

import (
	"bufio"
	"compress/gzip"
	// "fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/ivpusic/neo"
)

//build a gzip writer, will use it to replace the resp writer in ctx
type (
	gzipWriter struct {
		io.Writer
		http.ResponseWriter
	}
)

func (w gzipWriter) Write(b []byte) (int, error) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}

func (w gzipWriter) Flush() error {
	return w.Writer.(*gzip.Writer).Flush()
}

//Hijack lets the caller take over the connection
//After a call to Hijack(), the HTTP server library	 will not do anything else with the connection
//It becomes the caller's responsibility to manage and close the connection.
func (w gzipWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func (w *gzipWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

var writerPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(ioutil.Discard)
	},
}

// Gzip returns a middleware which compresses HTTP response using gzip compression scheme.
func Gzip(c *neo.Ctx, next neo.Next) {
	scheme := "gzip"

	c.Res.Header().Set("Vary", "Accept-Encoding")
	// c.Res.Header().Set("Content-Encoding", "gzip")
	if strings.Contains(c.Res.Header().Get("Accept-Encoding"), scheme) {
		w := writerPool.Get().(*gzip.Writer)
		//replace the rep in ctx
		w.Reset(c.Res.Writer())
		defer func() {
			w.Close()
			writerPool.Put(w)
		}()
		gw := gzipWriter{Writer: w, ResponseWriter: c.Res.Writer()}
		c.Res.Header().Set("Content-Encoding", scheme)
		//need a set function for c.Res
		c.Response().SetWriter(gw)
	}
	next()
	return

}
