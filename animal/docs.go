package animal

import (
	"fmt"
	"net/http"
	"os"
)

const html = `<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, minimum-scale=1, initial-scale=1, user-scalable=yes">
    <script src="https://unpkg.com/rapidoc@9.3.8/dist/rapidoc-min.js" integrity="sha512-0ES6eX4K9J1PrIEjIizv79dTlN5HwI2GW9Ku6ymb8dijMHF5CIplkS8N0iFJ/wl3GybCSqBJu8HDhiFkZRAf0g==" crossorigin="anonymous"></script>
  </head>
  <body>
    <rapi-doc spec-url="/doc/openapi" 
      render-style="view" 
      theme="dark"
      allow-spec-url-load="false"
      allow-spec-file-load="false"
    > </rapi-doc>
  </body>
</html>`

func GetDocumentation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusMethodNotAllowed)}, http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, html)
}

func GetOpenAPIDocumentation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusMethodNotAllowed)}, http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, os.Getenv("OPENAPI"))
}
