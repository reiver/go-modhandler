package modhandler


import (
	"net/http"
	"time"
)


const (
	ifModifiedSinceHeaderName = "If-Modified-Since"
	lastModifiedHeaderName    = "Last-Modified"
)


// LastModified replies to the HTTP request in such a way that it deals with
// conditional GETs by sending out a "Last-Modified" HTTP response header,
// and properly dealing with a "If-Modified-Since" HTTP request header.
//
// Note that this request an http.Handler. You will still need to call
// the returned handler's ServerHTTP method.
func LastModified(modtime time.Time, subhandler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Generate value for Last-Modified HTTP response header.
		lastModifiedHeaderValue := modtime.UTC().Format(time.RFC1123)

		// If the HTTP request has a If-Modified-Since request header
		// and it matches the value of the "Last-Modified" HTTP response
		// header we would have sent, then try to deal with it.
		//
		// (If we do indeed deal with this, then this func will return
		// at this point, and will not continue with the rest of
		// the code.)
		ifModifiedSinceHeaderValue := r.Header.Get(ifModifiedSinceHeaderName)
		if "" != ifModifiedSinceHeaderValue && lastModifiedHeaderValue == ifModifiedSinceHeaderValue {

			headers := w.Header()
			headers.Del("Content-Length")
			headers.Del("Content-Type")

			w.WriteHeader(http.StatusNotModified)

			return
		}

		// Set the value for the Last-Modified HTTP response header.
		w.Header().Add(lastModifiedHeaderName, lastModifiedHeaderValue)

		// Pass to sub-handler.
		subhandler.ServeHTTP(w, r)
	})
}


// LastModifiedServeHTTP is similar to the LastModified func, expect that instead of
// returning a handler, it creates the handle and then immediately calls its ServerHTTP
// method.
func LastModifiedServeHTTP(w http.ResponseWriter, r *http.Request, modtime time.Time, subhandler http.Handler) {
	LastModified(modtime, subhandler).ServeHTTP(w, r)
}
