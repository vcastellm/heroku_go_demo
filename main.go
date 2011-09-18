package main

import (
	"http"
	"io"
	"log"
	"os"
)

var port = func() string {
	tmpport := os.Getenv("PORT")
	if tmpport == "" {
		tmpport = "5000"
	}

	return tmpport
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, html)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":" + port(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.String())
	}
}

const html = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
	"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">

<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>

	<title>Heroku Go/Golang Demo @ GitHub</title>
	<link href='http://fonts.googleapis.com/css?family=Lato:regular,bold' rel='stylesheet' type='text/css'>
	<style type="text/css">
		body {
  		margin-top: 1.0em;
  		background: url("http://victorcoder.github.com/tuna/images/bg-body.jpg") repeat scroll 0 0 transparent;
		  font-family: 'Lato', arial, serif;
  		color: #000000;
    }
    #container {
      margin: 0 auto;
      width: 700px;
    }
		h1 { font-size: 3.8em; margin-bottom: 3px; }
		h1 .small { font-size: 0.4em; }
		h1 a { text-decoration: none }
		h2 { font-size: 1.5em;}
    h3 { text-align: center;}
    a { color:#C83535; }
    .description { font-size: 1.2em; margin-bottom: 30px; margin-top: 30px; font-style: italic;}
    .download { float: right; }
		pre { background: #000; color: #fff; padding: 15px;}
    hr { border: 0; width: 80%; border-bottom: 1px solid #aaa}
    .footer { text-align:center; padding-top:30px; font-style: italic; }
	</style>
	
</head>

<body>
  <div id="container">

    <h1><a href="http://github.com/victorcoder/heroku_go_demo">Heroku Go/golang demo</a>
      <span class="small">by <a href="http://github.com/victorcoder">victorcoder</a></span></h1>

    <div class="description">
      Hello from a Go app running in Heroku!
    </div>

	<h2>Code</h2>
    <p>You can clone the project with <a href="http://git-scm.com">Git</a>
      by running:
      <pre>$ git clone git://github.com/victorcoder/heroku_go_demo</pre>
    </p>

    <div class="footer">
      get the source code on GitHub : <a href="http://github.com/victorcoder/heroku_go_demo">victorcoder/heroku_go_demo</a>
    </div>

  </div>

  
</body>
</html>
`

