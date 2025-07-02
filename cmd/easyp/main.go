package main

import (
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/slog"

	"github.com/easyp-tech/server/cmd/easyp/internal/config"
	"github.com/easyp-tech/server/cmd/easyp/internal/config/cachetype"
	"github.com/easyp-tech/server/internal/connect"
	"github.com/easyp-tech/server/internal/https"
	"github.com/easyp-tech/server/internal/providers/bitbucket"
	"github.com/easyp-tech/server/internal/providers/cache"
	"github.com/easyp-tech/server/internal/providers/cache/artifactory"
	"github.com/easyp-tech/server/internal/providers/filter"
	"github.com/easyp-tech/server/internal/providers/github"
	"github.com/easyp-tech/server/internal/providers/localgit"
	"github.com/easyp-tech/server/internal/providers/localgit/namedlocks"
	"github.com/easyp-tech/server/internal/providers/multisource"
)

//nolint:gochecknoglobals
var (
	cfgFile = flag.String("cfg", "./local.config.yml", "path to Config file")
)

const (
	minNumberOfRepos = 128
)

func main() {
	flag.Parse()

	var (
		cfg      = must(config.ReadYaml[config.Config](*cfgFile))
		log      = newLogger(cfg.Log.Level)
		nameLock = namedlocks.New(minNumberOfRepos)
		cache    = buildCache(log, cfg.Cache)
		storage  = multisource.New(
			log,
			cache,
			localgit.New(cfg.Local.Storage, filterRepos(cfg.Local.Repos), nameLock),
			bbProxy(log, cfg.Proxy.BitBucket),
			githubProxy(log, cfg.Proxy.Github),
		)
		handler = connect.New(log, storage, cfg.Domain)
		serve   = func() error { return http.ListenAndServe(cfg.Listen.String(), loggingMiddleware(log, handler)) } //nolint:gosec
	)

	log.Debug("started", slog.Any("config", cfg))

	if cfg.TLS.CertFile != "" {
		serve = func() error {
			return https.ListenAndServe(cfg.Listen, loggingMiddleware(log, handler), cfg.TLS.CertFile, cfg.TLS.KeyFile, cfg.TLS.CACertFile)
		}
	}

	if err := serve(); err != nil {
		log.Error("shutdown", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func newLogger(level string) *slog.Logger {
	var logLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn", "warning":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}

func loggingMiddleware(log *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		headers := r.Header.Clone()

		maskSensitiveHeaders(headers)

		log.Debug("request details",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Any("headers", headers),
		)

		lrw := &loggingResponseWriter{ResponseWriter: w}

		next.ServeHTTP(lrw, r)

		duration := time.Since(start)
		log.Info("request completed",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", lrw.status),
			slog.Duration("duration", duration),
		)
	})
}

func maskSensitiveHeaders(headers http.Header) {
	for key := range headers {
		if isSensitiveHeader(key) {
			headers.Set(key, "***")
		}
	}
}

func isSensitiveHeader(key string) bool {
	key = strings.ToLower(key)
	return key == "authorization" ||
		key == "cookie" ||
		key == "x-api-key" ||
		key == "token"
}

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

func githubProxy(log *slog.Logger, defs []config.GithubRepo) multisource.Source { //nolint:ireturn
	repos := make([]github.Repo, 0, len(defs))
	for _, def := range defs {
		repos = append(
			repos,
			github.Repo{
				Token: def.AccessToken,
				Repo: filter.Repo{
					Owner:    def.Repo.Owner,
					Name:     def.Repo.Name,
					Prefixes: def.Repo.Prefixes,
					Paths:    def.Repo.Paths,
				},
			},
		)
	}

	return github.NewMultiRepo(log, repos)
}

func bbProxy(log *slog.Logger, defs []config.BitBucketRepo) multisource.Source { //nolint:ireturn
	repos := make([]bitbucket.Repo, 0, len(defs))
	for _, def := range defs {
		repos = append(
			repos,
			bitbucket.Repo{
				User:     bitbucket.User(def.User),
				Password: bitbucket.Password(def.AccessToken),
				URL:      def.BaseURL.URL,
				Repo: filter.Repo{
					Owner:    def.Repo.Owner,
					Name:     def.Repo.Name,
					Prefixes: def.Repo.Prefixes,
					Paths:    def.Repo.Paths,
				},
			},
		)
	}

	return bitbucket.NewMultiRepo(log, repos)
}

func filterRepos(defs []config.Repo) []filter.Repo { //nolint:ireturn
	repos := make([]filter.Repo, 0, len(defs))
	for _, def := range defs {
		repos = append(
			repos,
			filter.Repo{
				Owner:    def.Owner,
				Name:     def.Name,
				Prefixes: def.Prefixes,
				Paths:    def.Paths,
			},
		)
	}

	return repos
}

func buildCache(log *slog.Logger, cfg config.Cache) multisource.Cache { //nolint:ireturn
	switch cfg.Type {
	case cachetype.None:
		return cache.Noop{}
	case cachetype.Local:
		return cache.Local{Dir: cfg.Local.Dir}
	case cachetype.Artifactory:
		return artifactory.New(
			log,
			cfg.Artifactory.BaseURL.String(),
			cfg.Artifactory.User,
			cfg.Artifactory.AccessToken,
		)
	default:
		panic("unreachable reached")
	}
}
