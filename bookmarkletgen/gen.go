package bookmarkletgen

import (
	"bytes"
	"text/template"
)

func genIndexHTML(titledJSs []titledJS, footerHTML string) ([]byte, error) {
	t := `
	<html>
	<body>	
		<h1>Bookmarklets</h1>
		<p>
			To install, drag the links not titled 'src' to your toolbar.
		</p>
		<p>
			<ul>
			{{range $index, $p := .TitledJSs}}
					<li>
						<a href="javascript:{{$p.JS}}">{{$p.Title}}</a> (<a target="_" href="{{$p.Link}}">src</a>) - {{$p.Description}}
						{{if $p.Image }}
							[<a target="_" href="{{$p.Image}}">Context</a>]
						{{end}}
					</li>
			{{end}}
			</ul>
		</p>
		{{if .Footer}}
		<p>
			{{.Footer}}
		</p>
		{{end}}
	</body>
	</html>
			`
	tmpl, err := template.New("html").Parse(t)
	if err != nil {
		return nil, err
	}
	var data = struct {
		TitledJSs []titledJS
		Footer    string
	}{
		TitledJSs: titledJSs,
		Footer:    footerHTML,
	}
	var out bytes.Buffer
	if err = tmpl.Execute(&out, data); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func genIndexMD(titledJSs []titledJS) ([]byte, error) {
	tmpl, err := template.New("md").Parse(`
# Bookmarklets

To install, drag the links not titled 'src' to your toolbar.

{{range $index, $p := .TitledJSs}}
*	[{{$p.Title}}](javascript:{{$p.JS}}) ([src]({{$p.Link}})) - {{$p.Description}}
{{end}}
		`)
	if err != nil {
		return nil, err
	}
	var data = struct {
		TitledJSs []titledJS
	}{
		TitledJSs: titledJSs,
	}
	var out bytes.Buffer
	if err = tmpl.Execute(&out, data); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
