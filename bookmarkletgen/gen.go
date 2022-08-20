package bookmarkletgen

import (
	"bytes"
	"text/template"
)

func genIndexHTML(titledJSs []titledJS, footerHTML, headerHTML, templateHTML string) ([]byte, error) {
	t := templateHTML
	if t == "" {
		t = `
	<html>
	<body>	
		<h1>Bookmarklets</h1>
		{{if .Header}}
		<p>
			{{.Header}}
		</p>
		{{end}}
		<p>
			To install, drag the links not titled 'src' to your toolbar.
		</p>
		<p>
			<ul>
			{{range $index, $p := .TitledJSs}}
					<li>
						<a href="javascript:{{$p.JS}}">{{$p.Title}}</a>
						{{if $p.HasSource }}
							(<a target="_" href="{{$p.Link}}">src</a>)
						{{end}}
						- {{$p.Description}}
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
	}
	tmpl, err := template.New("html").Parse(t)
	if err != nil {
		return nil, err
	}
	var data = struct {
		TitledJSs []titledJS
		Footer    string
		Header    string
	}{
		TitledJSs: titledJSs,
		Footer:    footerHTML,
		Header:    headerHTML,
	}
	var out bytes.Buffer
	if err = tmpl.Execute(&out, data); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// TODO: Remove, markdown doesn't work.
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
