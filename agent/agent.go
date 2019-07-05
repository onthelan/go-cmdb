package main

import (
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, req *http.Request) {
	const tpl = `<!DOCTYPE html>
              <html>
              <head>
                <meta charset="UTF-8">
                <title>{{.Title}}</title>
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
                <script async src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
              </head>
                <body>
                  <nav class="navbar navbar-default">
                    <div class="container">
                      <div class="navbar-header">
                            {{range .Items}}<a class="navbar-brand" href="/">{{ . }}</a>{{else}}<div><strong>Data Missing</strong></div>{{end}}
                      </div>
                    </div>
                  </nav>
                </body>
              </html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "Go CMDB - Agent",
		Items: []string{
			"System",
			"Services",
			"Settings",
		},
	}
	err = t.Execute(w, data)
	check(err)

}

func main() {

	http.HandleFunc("/", handler)

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/")
	err := http.ListenAndServeTLS(":8443", "MyCertificate.crt", "MyKey.key", nil)
	log.Fatal(err)
}
