// Package bookmarkletgen generates an HTML index from the javascript bookmarklets in js.
package bookmarkletgen

import (
	"io/ioutil"
	"log"
	"sort"

	"github.com/pkg/errors"
	"github.com/yosssi/gohtml"
)

// GenerateIndexFiles generates bookmarklet index files for the given JS files.
func GenerateIndexFiles(jsFiles []string, os ...Option) error {
	opts := MakeOptions(os...)
	titledJSs, err := inspectFiles(jsFiles, opts.BaseSourceURL())
	if err != nil {
		return err
	}

	sort.Slice(titledJSs, func(i, j int) bool {
		return titledJSs[i].Title < titledJSs[j].Title
	})

	if opts.OutfileHTML() != "" {
		var footerHTML string
		if opts.FooterHTML() != "" {
			footerHTML = opts.FooterHTML()
		} else if opts.FooterHTMLFile() != "" {
			b, err := ioutil.ReadFile(opts.FooterHTMLFile())
			if err != nil {
				return errors.Errorf("reading footer HTML: %v", err)
			}
			footerHTML = string(b)
		}

		var headerHTML string
		if opts.HeaderHTML() != "" {
			headerHTML = opts.HeaderHTML()
		} else if opts.HeaderHTMLFile() != "" {
			b, err := ioutil.ReadFile(opts.HeaderHTMLFile())
			if err != nil {
				return errors.Errorf("reading header HTML: %v", err)
			}
			headerHTML = string(b)
		}

		var templateHTML string
		if opts.TemplateHTML() != "" {
			templateHTML = opts.TemplateHTML()
		} else if opts.TemplateHTMLFile() != "" {
			b, err := ioutil.ReadFile(opts.TemplateHTMLFile())
			if err != nil {
				return errors.Errorf("reading template HTML: %v", err)
			}
			templateHTML = string(b)
		}

		out, err := genIndexHTML(titledJSs, footerHTML, headerHTML, templateHTML)
		if err != nil {
			return err
		}
		log.Printf("Writing to %s\n", opts.OutfileHTML())
		formatted := gohtml.Format(string(out))
		if err := ioutil.WriteFile(opts.OutfileHTML(), []byte(formatted), 0755); err != nil {
			return err
		}
	}

	if opts.OutfileMD() != "" {
		out, err := genIndexMD(titledJSs)
		if err != nil {
			return err
		}
		log.Printf("Writing to %s\n", opts.OutfileMD())
		if err := ioutil.WriteFile(opts.OutfileMD(), out, 0755); err != nil {
			return err
		}
	}

	return nil
}
