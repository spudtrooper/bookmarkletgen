package bookmarkletgen

// genopts --outfile=bookmarkletgen/options.go 'outfileHTML:string' 'outfileMD:string' 'baseSourceURL:string' 'footerHTML:string'

type Option func(*optionImpl)

type Options interface {
	OutfileHTML() string
	OutfileMD() string
	BaseSourceURL() string
	FooterHTML() string
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

type optionImpl struct {
	outfileHTML   string
	outfileMD     string
	baseSourceURL string
	footerHTML    string
}

func (o *optionImpl) OutfileHTML() string   { return o.outfileHTML }
func (o *optionImpl) OutfileMD() string     { return o.outfileMD }
func (o *optionImpl) BaseSourceURL() string { return o.baseSourceURL }
func (o *optionImpl) FooterHTML() string    { return o.footerHTML }

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
