// Package bookmarkletgen generates an HTML index from the javascript bookmarklets in js.
package bookmarkletgen

import (
	"io/ioutil"
	"log"
	"sort"
)

// GenerateIndexFiles generates bookmarklet index files for the given JS files.
func GenerateIndexFiles(jsFiles []string, os ...Option) error {
	opts := makeOptionImpl(os...)
	titledJSs, err := inspectFiles(jsFiles, opts.baseSourceURL)
	if err != nil {
		return err
	}

	sort.Slice(titledJSs, func(i, j int) bool {
		return titledJSs[i].Title < titledJSs[j].Title
	})

	if opts.outfileHTML != "" {
		out, err := genIndexHTML(titledJSs, opts.footerHTML)
		if err != nil {
			return err
		}
		log.Printf("Writing to %s\n", opts.outfileHTML)
		if err := ioutil.WriteFile(opts.outfileHTML, out, 0755); err != nil {
			return err
		}
	}

	if opts.outfileMD != "" {
		out, err := genIndexMD(titledJSs)
		if err != nil {
			return err
		}
		log.Printf("Writing to %s\n", opts.outfileMD)
		if err := ioutil.WriteFile(opts.outfileMD, out, 0755); err != nil {
			return err
		}
	}

	return nil
}
