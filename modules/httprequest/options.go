package httprequest

import (
	"net/http"
	"time"

	"github.com/imdario/mergo"
	"github.com/juju/errors"
)

var (
	DefaultOptions = Option{
		Timeout: time.Duration(30 * time.Second),
		Headers: http.Header{},
	}
)

type Option struct {
	Timeout time.Duration
	Headers http.Header
}

type Options struct {
	Option
}

func NewOptions() (*Options, error) {
	b := Options{
		Option: DefaultOptions,
	}

	return &b, nil
}

func (p *Options) Apply(opt ...Option) error {
	for _, o := range opt {
		if err := mergo.Merge(p, o, mergo.WithOverride); err != nil {
			return errors.Annotate(err, "Merge")
		}
	}

	return nil
}

func Timeout(timeout time.Duration) Option {
	return Option{
		Timeout: timeout,
	}
}

func Accept(contentType string) Option {
	header := http.Header{}
	header.Set("Accept", contentType)

	return Option{
		Headers: header,
	}
}
