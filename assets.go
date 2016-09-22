//assets.go
package main

var assets = map[string]string{}

func init() {
	assets["index.html"] = `
<html>
<body>
	{{range .Actions}}
		<div>
			<h3><a href="./{{.Id}}">{{.Id}}</a></h3>
			{{.Description}}
		</div>
	{{end}}
</body>
</html>`
}
