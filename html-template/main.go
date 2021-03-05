package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const markdown = `

	#{{.Title}}

{{range .Items}} + {{ . }}{{else}}- Nothing in Items{{end}}

{{range .Names}} > {{ . }}{{else}}- Nothing in Names{{end}}
	
`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("markdown").Parse(markdown)
	check(err)

	data := struct {
		Title string
		Items []string
		Names []string
	}{
		Title: "Markdown template from Go",
		Items: []string{
			"Hello\n",
			"From\n",
			"Template\n",
		},
		Names: []string{
			"Trond\n",
			"Lene\n",
			"Linn\n",
			"Rune\n",
			"Morten\n",
			"Theo\n",
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)

}
