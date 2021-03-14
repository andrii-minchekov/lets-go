package cfg

import (
	"flag"
	"github.com/andrii-minchekov/lets-go/domain"
	"sync"
)

type flagConfig struct {
	addr      *string
	dsn       *string
	secret    *string
	htmlDir   *string
	staticDir *string
	init      *sync.Once
}

func (r flagConfig) lazyInit() {
	r.init.Do(flag.Parse)
}

var FlagConfig = newFlagConfig()

func newFlagConfig() cfg.Config {
	config := flagConfig{
		addr:      flag.String("addr", ":4000", "HTTP network address"),
		dsn:       flag.String("dsn", "user=postgres dbname=snippetbox password=postgres sslmode=disable", "Postgres Datasource Name"),
		secret:    flag.String("secret", "s6Nd%+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key"),
		htmlDir:   flag.String("html-dir", "./ui/html", "Path to HTML templates"),
		staticDir: flag.String("static-dir", "./ui/static", "Path to static assets"),
		init:      &sync.Once{},
	}
	return config
}

func (r flagConfig) DSN() string {
	r.lazyInit()
	return *r.dsn
}

func (r flagConfig) HTMLDir() string {
	r.lazyInit()
	return *r.htmlDir
}

func (r flagConfig) StaticDir() string {
	r.lazyInit()
	return *r.staticDir
}

func (r flagConfig) Secret() string {
	r.lazyInit()
	return *r.secret
}

func (r flagConfig) Addr() string {
	r.lazyInit()
	return *r.addr
}
