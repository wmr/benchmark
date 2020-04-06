package main

import (
	"crypto/tls"
	"crypto/x509"
	// "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	client := &http.Client{}

	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	// switch *httpVersion {
	// case 1:
	// 	client.Transport = &http.Transport{
	// 		TLSClientConfig: tlsConfig,
	// 	}
	// case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	// }

  pr, _ := io.Pipe()

  
	// Perform the request
  url := "https://localhost:10000"
  
  req, err := http.NewRequest(http.MethodGet, url, ioutil.NopCloser(pr))
	if err != nil {
		log.Fatal(err)
  }

  resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
  
	if err != nil {
		log.Fatalf("Failed get: %s", err)
  }
  
  defer resp.Body.Close()
  
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	// fmt.Printf(
	// 	"Got response %d: %s %s\n",
	// 	resp.StatusCode, resp.Proto, string(body))

	// // Create a pipe - an object that implements `io.Reader` and `io.Writer`.
	// // Whatever is written to the writer part will be read by the reader part.
	// pr, _ := io.Pipe()

	// // Create an `http.Request` and set its body as the reader part of the
	// // pipe - after sending the request, whatever will be written to the pipe,
	// // will be sent as the request body.
	// // This makes the request content dynamic, so we don't need to define it
	// // before sending the request.
	// req, err := http.NewRequest(http.MethodGet, "https://localhost:10000/", ioutil.NopCloser(pr))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Send the request
	// resp, err = http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Got: %d", resp.StatusCode)

	// // Run a loop which writes every second to the writer part of the pipe
	// // the current time.

	// // Copy the server's response to stdout.

	// _, err = ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = io.Copy(os.Stdout, res.Body)
	// log.Fatal(err)
}
