package plugin_service

import (
	"fmt"
	"regexp"
)

type ExportInfo struct {
	Title   string
	Version string
}

func GetExportInfo(xmlContent string) (*ExportInfo, error) {
	r, err := regexp.Compile(`<(.+:)*([a-zA-Z]*\d*[a-zA-Z]*) schemeVersion=\"(.+?)\">`)
	if err != nil {
		return nil, err
	}

	results := r.FindStringSubmatch(xmlContent)
	if len(results) < 4 {
		return nil, fmt.Errorf("Wrong format export file, could not determine version and title")
	}

	return &ExportInfo{Title: results[2], Version: results[3]}, nil
}
