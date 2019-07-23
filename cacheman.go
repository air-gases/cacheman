package cacheman

import (
	"fmt"
	"strings"

	"github.com/aofei/air"
)

// GasConfig is a set of configurations for the `Gas`.
type GasConfig struct {
	MustRevalidate  bool
	NoCache         bool
	NoStore         bool
	NoTransform     bool
	Public          bool
	Private         bool
	ProxyRevalidate bool
	MaxAge          int
	SMaxAge         int

	Skippable func(*air.Request, *air.Response) bool
}

// Gas returns an `air.Gas` that is used to manage the Cache-Control header
// based on the gc.
func Gas(gc GasConfig) air.Gas {
	ds := []string{}
	if gc.MustRevalidate {
		ds = append(ds, "must-revalidate")
	}

	if gc.NoCache {
		ds = append(ds, "no-cache")
	}

	if gc.NoStore {
		ds = append(ds, "no-store")
	}

	if gc.NoTransform {
		ds = append(ds, "no-transform")
	}

	if gc.Public {
		ds = append(ds, "public")
	}

	if gc.Private {
		ds = append(ds, "private")
	}

	if gc.ProxyRevalidate {
		ds = append(ds, "proxy-revalidate")
	}

	if gc.MaxAge >= 0 {
		ds = append(ds, fmt.Sprintf("max-age=%d", gc.MaxAge))
	}

	if gc.SMaxAge >= 0 {
		ds = append(ds, fmt.Sprintf("s-maxage=%d", gc.SMaxAge))
	}

	directives := strings.Join(ds, ", ")

	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if gc.Skippable != nil && gc.Skippable(req, res) {
				return next(req, res)
			}

			res.Header.Set("Cache-Control", directives)
			err := next(req, res)
			if err != nil && !res.Written {
				res.Header.Del("Cache-Control")
			}

			return err
		}
	}
}
