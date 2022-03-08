package bookmarkletgen

//go:generate genopts --outfile=options.go "outfileHTML:string" "outfileMD:string" "baseSourceURL:string" "footerHTML:string" "footerHTMLFile:string" "headerHTML:string" "headerHTMLFile:string" "templateHTML:string" "templateHTMLFile:string"

type Option func(*optionImpl)

type Options interface {
	OutfileHTML() string
	OutfileMD() string
	BaseSourceURL() string
	FooterHTML() string
	FooterHTMLFile() string
	HeaderHTML() string
	HeaderHTMLFile() string
	TemplateHTML() string
	TemplateHTMLFile() string
}

func OutfileHTML(outfileHTML string) Option {
	return func(opts *optionImpl) {
		opts.outfileHTML = outfileHTML
	}
}

func OutfileMD(outfileMD string) Option {
	return func(opts *optionImpl) {
		opts.outfileMD = outfileMD
	}
}

func BaseSourceURL(baseSourceURL string) Option {
	return func(opts *optionImpl) {
		opts.baseSourceURL = baseSourceURL
	}
}

func FooterHTML(footerHTML string) Option {
	return func(opts *optionImpl) {
		opts.footerHTML = footerHTML
	}
}

func FooterHTMLFile(footerHTMLFile string) Option {
	return func(opts *optionImpl) {
		opts.footerHTMLFile = footerHTMLFile
	}
}

func HeaderHTML(headerHTML string) Option {
	return func(opts *optionImpl) {
		opts.headerHTML = headerHTML
	}
}

func HeaderHTMLFile(headerHTMLFile string) Option {
	return func(opts *optionImpl) {
		opts.headerHTMLFile = headerHTMLFile
	}
}

func TemplateHTML(templateHTML string) Option {
	return func(opts *optionImpl) {
		opts.templateHTML = templateHTML
	}
}

func TemplateHTMLFile(templateHTMLFile string) Option {
	return func(opts *optionImpl) {
		opts.templateHTMLFile = templateHTMLFile
	}
}

type optionImpl struct {
	outfileHTML      string
	outfileMD        string
	baseSourceURL    string
	footerHTML       string
	footerHTMLFile   string
	headerHTML       string
	headerHTMLFile   string
	templateHTML     string
	templateHTMLFile string
}

func (o *optionImpl) OutfileHTML() string      { return o.outfileHTML }
func (o *optionImpl) OutfileMD() string        { return o.outfileMD }
func (o *optionImpl) BaseSourceURL() string    { return o.baseSourceURL }
func (o *optionImpl) FooterHTML() string       { return o.footerHTML }
func (o *optionImpl) FooterHTMLFile() string   { return o.footerHTMLFile }
func (o *optionImpl) HeaderHTML() string       { return o.headerHTML }
func (o *optionImpl) HeaderHTMLFile() string   { return o.headerHTMLFile }
func (o *optionImpl) TemplateHTML() string     { return o.templateHTML }
func (o *optionImpl) TemplateHTMLFile() string { return o.templateHTMLFile }

func makeOptionImpl(opts ...Option) *optionImpl {
	res := &optionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeOptions(opts ...Option) Options {
	return makeOptionImpl(opts...)
}
