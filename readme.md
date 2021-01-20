# Tmpl

Simple binder between json data and one or more text templates.

## Usage

### Single template

`tmpl --input ./api.json --output ./api.md --template ./api.tmpl`

### Multiple templates

`tmpl --input ./api.json --output ./api.md --template ./api.base.tmpl --template ./api.resources.tmpl`

### Template glob

`tmpl --input ./api.json --output ./api.md --templateGlob ./templates/*.tmpl`

### Raw template text

`tmpl --input ./api.json --output ./api.md --templateRaw "# {{.Service}}"`

### Multiple raw template texts

`tmpl --input ./api.json --output ./api.md --templateRaw "# {{.Service}} {{template \"author\" .}}" --templateRaw "{{define \"author\"}}Author : {{.Author}}{{end}}"`

### Reading from stdin

`cat ./api.json | tmpl --format json --output ./api.md --template ./api.tmpl`

### Writing to stdout

`tmpl --input ./api.json --stdout --template ./api.tmpl`
