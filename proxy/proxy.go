package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const HOST = "https://test.jakecoffman.com"

func main() {
	// make logger print out file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	remote, err := url.Parse(HOST)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler(httputil.NewSingleHostReverseProxy(remote)))
	http.HandleFunc("/iframe", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(iframe))
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}

func handler(proxy *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println(r.Header)
		proxy.ServeHTTP(w, r)
		log.Println(time.Since(start))
	}
}

const iframe = `<html lang="en">
<head>
<title>IFRAMER</title>
<style>
html, body {
margin: 0;
padding: 0;
}
iframe {
width: 90%;
height: 90%;
border: 1px solid black;
border-radius: 5px;
}
.flex {
display: flex;
align-content: center;
justify-content: center;
height: 90%;
}
</style>
</head>
<body>
<h1>iframe</h1>
<div class="flex">
	<iframe src="/" allow="microphone"></iframe>
</div>
</body>
</html>
`
