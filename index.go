// index.go
package main

var templete_index = `


<html>
<body>
	{{range .Actions}}
		<div>
			<h3><a href="./{{.Id}}">{{.Id}}</a></h3>
			{{.Description}}
		</div>
	{{end}}
</body>
</html>


`
