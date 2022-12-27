package types

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

const (
	// JSONFormat is JSON output format.
	JSONFormat OutputFormat = iota
	// YAMLFormat is YAML output format.
	YAMLFormat OutputFormat = iota
	// DefaultOutputFormat is default output format.
	DefaultOutputFormat = JSONFormat
)

type SimpleQuery struct {
	ProjectID  string
	OutputPath string
	OutputFmt  OutputFormat
	ImageFile  string
	ImageURI   string
}

func (q *SimpleQuery) Validate() error {
	if q.ImageFile != "" && q.ImageURI != "" {
		return errors.New("only one of image file or image URI can be specified")
	}

	return nil
}

func (q *SimpleQuery) String() string {
	return fmt.Sprintf("projectID:%s, output:%s, format:%s",
		q.ProjectID, q.OutputPath, q.OutputFmt)
}

type OutputFormat int64

func (o OutputFormat) String() string {
	switch o {
	case JSONFormat:
		return "json"
	case YAMLFormat:
		return "yaml"
	default:
		return "unknown"
	}
}

// ParseOutputFormat parses output format.
func ParseOutputFormatOrDefault(format string) OutputFormat {
	if format == "" {
		return DefaultOutputFormat
	}

	switch format {
	case "json":
		return JSONFormat
	case "yaml":
		return YAMLFormat
	default:
		log.Error().Msgf("unsupported output format: %s", format)
		return DefaultOutputFormat
	}
}
