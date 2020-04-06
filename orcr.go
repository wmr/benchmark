package main

import (
	"github.com/fxamacker/cbor" // imports as package "cbor"
	"github.com/julienschmidt/httprouter"
  "github.com/scritchley/orc"
	"io"
	"log"
	"net"
	"net/http"
)

type Server struct {
	router *httprouter.Router
}

func (s *Server) Initialize() error {
	s.router = httprouter.New()
	s.router.GET("/", s.handler)

	//Creates the http server
	server := &http.Server{
		Handler: s.router,
	}

	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		return err
	}

	log.Println("HTTP server is listening..")
	return server.ServeTLS(listener, "./server.crt", "./server.key")
}

func FlushChunk(w http.ResponseWriter, p []byte) (n int, err error) {
  n, err = w.Write(p)
  if f, ok := w.(http.Flusher); ok {
    f.Flush();
  }
  return 
}

func (s *Server) handler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// We only accept HTTP/2!
	// (Normally it's quite common to accept HTTP/1.- and HTTP/2 together.)
	if req.ProtoMajor != 2 {
		log.Println("Not a HTTP/2 request, rejected!")
		w.WriteHeader(http.StatusInternalServerError)
		return
  }
  
  if f, ok := w.(http.Flusher); ok {
    f.Flush();
  }

	r, err := orc.Open("./data.orc")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	c := r.Select("col0", "col1", "col2", "col3", "col4")

	for c.Stripes() {
    var chunk []interface{}
    idx := 0
		for c.Next() {
      idx = idx + 1
      chunk = append(chunk, c.Row())
      if idx % 10000 == 0 {
        d, err := cbor.Marshal(chunk)
        FlushChunk(w, d)

        chunk = chunk[:0]
        if err != nil {
          log.Fatal(err)
        }
      }
    }
    d, err := cbor.Marshal(chunk)
    w.Write(d)
    chunk = chunk[:0]
    if err != nil {
      log.Fatal(err)
    }
    
		if err != nil {
			if err == io.EOF {
				w.Header().Set("Status", "200 OK")
				req.Body.Close()
			}
			break
		}
	}

	if err := c.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	waitc := make(chan bool)

	// HTTP2 Server
	server := new(Server)
	err := server.Initialize()
	if err != nil {
		log.Println(err)
		return
	}

	// Waits forever
	<-waitc

}
