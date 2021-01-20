
# tmpl

a simple but powerful binder between JSON data and text templates. `tmpl` acts as a wrapper around go's [text/template](https://golang.org/pkg/text/template/) package and allows powerful templating capabilities from a simple command line interface.

## Install

### [Download the latest binary](https://github.com/jsmithdenverdev/tmpl/releases/latest)

### Go Get:
```
GO111MODULE=on go get github.com/jsmithdenverdev/tmpl
```

## Features
- Written in go, so you can download a dependency free binary
- Defaults to reading from `stdin` and writing to `stdout` allowing it to be used with unix pipes
- Supports using string templates or template files
- Supports multiple templates in a single run (cannot combine string and file templates)

## Usage

```
Usage of tmpl:
  -input string
        Input file path (default is stdin)
  -output string
        Output file path (default is stdout)
  -templateFile value
        Template file(s)
  -templateString value
        Template string(s)
  -tf value
        Template file(s)
  -ts value
        Template string(s)
```

## Examples

### Simple example

In this example we read data from `stdin`, apply a string template (using the `-ts` arg) and write to `stdout`.

`echo "{\"name\": \"World\"}" | tmpl -ts "Hello, {{.name}}"`

`> Hello, World`
### Less simple example

In this example we read data from a file named `user.json` using the `-input` flag. We then apply a file based template named `user-profile.tmpl`. Finally we write the data from the template to a file named `user-profile.html`.

`tmpl --input ./user.json --tf ./user-profile.tmpl --output user-profile.html`

user.json
```json
{
  "name": {
    "first": "Jake",
    "last": "Smith"
  },
  "hobbies": ["coding", "climbing", "backpacking"],
  "images": {
    "profile": "/static/img/jake.png"
  }
}
```

`user-profile.tmpl`
```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{.name.first}} {{.name.last}}</title>
  </head>
  <body>
    <div>
      <img src="{{.images.profile}}" alt="{{.name.first}}-{{.name.last}}" />
      <h1>{{.name.first}} {{.name.last}}</h1>
      <ul>
        {{range .hobbies -}}
        <li>{{.}}</li>
        {{end -}}
      </ul>
    </div>
  </body>
</html>
```

`user-profile.html`
```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Jake Smith</title>
  </head>
  <body>
    <div>
      <img src="/static/img/jake.png" alt="Jake-Smith" />
      <h1>Jake Smith</h1>
      <ul>
        <li>coding</li>
        <li>climbing</li>
        <li>backpacking</li>
      </ul>
    </div>
  </body>
</html>
```
