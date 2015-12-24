# go-modhandler

A library that provides a ("middleware") HTTP handler to deal with
conditional GETs by sending out a "Last-Modified" HTTP response header,
and properly dealing with a "If-Modified-Since" HTTP request header,
for the Go programming language.


## Example

For example:
```
func ServerHTTP(w http.ResponseWriter, r *http.Request) {

    modtime := ... //@TODO: This is where you figure out the date time for this.
                   //       For example: the modification date of a file, or
                   //       the "when_updated" column in a database, or something.

    modhandler.LastModifiedServeHTTP(w, r, modtime, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

        //@TODO: Put your "real" ServerHTTP() code in here.
        //
        // Note that a "Last-Modified" header will automagically have been added
        // to the 'w' parameter in this func.

    }))

}
```


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-modhandler

[![GoDoc](https://godoc.org/github.com/reiver/go-modhandler?status.svg)](https://godoc.org/github.com/reiver/go-modhandler)
