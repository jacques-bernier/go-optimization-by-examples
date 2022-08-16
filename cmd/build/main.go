package main

import (
	"html/template"
	"log"
	"os"
)

type page struct {
	Title    string
	Examples []example
}

type example struct {
	Name string
	Code string
}

func main() {
	home := `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Examples}}<div>{{ .Name }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>
`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(home)
	check(err)

	data := page{
		Title: "Go optimization by examples",
		Examples: []example{
			{
				Name: "Example 1",
				Code: "",
			},
			{
				Name: "Example 2",
				Code: "",
			},
		},
	}
	file, err := os.Create("examples/index.html")
	check(err)
	defer file.Close()

	err = t.Execute(file, data)
	check(err)
}