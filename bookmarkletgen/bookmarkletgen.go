package bookmarkletgen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/html"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

type fileMetadata struct {
	title, description, image string
}

type minifiedJS struct {
	js string
	md fileMetadata
}

func findFileMetadata(b []byte) fileMetadata {
	var title string
	var descr string
	var image string
	lines := strings.Split(string(b), "\n")
	for _, s := range lines {
		if m := titleRE.FindStringSubmatch(s); m != nil && len(m) == 2 {
			title = strings.TrimSpace(m[1])
		} else if m := descrRE.FindStringSubmatch(s); m != nil && len(m) == 2 {
			descr = strings.TrimSpace(m[1])
		} else if m := imageRE.FindStringSubmatch(s); m != nil && len(m) == 2 {
			image = strings.TrimSpace(m[1])
		}
	}
	res := fileMetadata{
		title:       title,
		description: descr,
		image:       image,
	}
	return res
}

func minifyAndFindFileMetadataFromFile(f string) (*minifiedJS, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, errors.Errorf("ReadFile(%q): %v", f, err)
	}
	res, err := minifyAndFindFileMetadataFromBytes(b)
	if err != nil {
		return nil, errors.Errorf("minifyAndFindFileMetadataFromBytes: %v", err)
	}
	if res.md.title == "" {
		res.md.title = titleFromFileName(f)
	}
	return res, nil
}

func minifyAndFindFileMetadataFromString(jsString string) (*minifiedJS, error) {
	return minifyAndFindFileMetadataFromBytes([]byte(jsString))
}

func minifyAndFindFileMetadataFromBytes(b []byte) (*minifiedJS, error) {
	r := bytes.NewBuffer(b)
	var w bytes.Buffer
	m := minify.New()
	var o js.Minifier
	if err := o.Minify(m, &w, r, nil); err != nil {
		return nil, err
	}
	js := w.String()
	md := findFileMetadata(b)
	res := &minifiedJS{
		js: js,
		md: md,
	}
	return res, nil
}

var (
	titleRE = regexp.MustCompile(`.*@Title:(.*)$`)
	descrRE = regexp.MustCompile(`.*@Description:(.*)$`)
	imageRE = regexp.MustCompile(`.*@Image:(.*)$`)
)

func titleFromFileName(f string) string {
	ext := path.Ext(f)
	base := path.Base(f)
	titleBase := base[0 : len(base)-len(ext)]
	parts := strings.Split(titleBase, "-")
	output := []string{}
	for _, p := range parts {
		cap := strings.Title(strings.ToLower(p))
		output = append(output, cap)

	}
	return strings.Join(output, " ")
}

func inspectFiles(jsFiles []string, baseSourceURL string) ([]titledJS, error) {
	hasSource := baseSourceURL != ""
	var titledJSs []titledJS
	for _, f := range jsFiles {
		minified, err := minifyAndFindFileMetadataFromFile(f)
		if err != nil {
			return nil, err
		}
		js := minified.js
		link := fmt.Sprintf(baseSourceURL+"/%s", path.Base(f))
		title := minified.md.title
		descr := html.Linkify(minified.md.description)
		image := minified.md.image
		t := titledJS{
			Title:       title,
			Description: descr,
			Image:       image,
			JS:          js,
			Link:        link,
			HasSource:   hasSource,
		}
		titledJSs = append(titledJSs, t)
	}

	return titledJSs, nil
}
