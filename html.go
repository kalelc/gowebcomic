package main

import "html/template"

var templ = template.Must(template.New("issuelist").Parse(`
<div align='center'>
<h1>{{.Num}} - {{.Title}}</h1>
<div><img src='{{.Img}}'></div>
<div><img src='{{.Year}}'></div>
</div>
`))
