package bitbucket

import (
	"context"
	"fmt"
	"hash/crc32"
	"net/url"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"

	"github.com/tammio/server/internal/providers/content"
	"github.com/tammio/server/internal/providers/filter"
	"github.com/tammio/server/internal/providers/source"
)

type (
	User     string
	Password string
	Project  string
)

type Repo struct {
	User     User
	Password Password
	URL      url.URL
	filter.Repo
}

type multiRepo struct {
	log   *slog.Logger
	repos []Repo
}

func (m multiRepo) Find(owner, name string) source.Source { //nolint:ireturn
	s, ok := m.find(owner, name)
	if !ok {
		return nil
	}

	return s
}

func (m multiRepo) find(owner, name string) (sourceRepo, bool) {
	i := slices.IndexFunc(m.repos, func(r Repo) bool {
		return r.Repo.Owner == owner && r.Repo.Name == name
	})
	if i < 0 {
		return sourceRepo{}, false //nolint:exhaustruct
	}

	return sourceRepo{log: m.log, repo: m.repos[i]}, true
}

func NewMultiRepo(log *slog.Logger, repos []Repo) multiRepo {
	return multiRepo{
		log:   log,
		repos: repos,
	}
}

var _ source.Source = sourceRepo{} //nolint:exhaustruct

type sourceRepo struct {
	log  *slog.Logger
	repo Repo
}

func (r sourceRepo) ConfigHash() string {
	return fmt.Sprintf("%X", crc32.ChecksumIEEE([]byte(fmt.Sprintf("%+v", r.repo.Repo))))
}

func (r sourceRepo) Name() string { return "bitbucket proxy" }

func (r sourceRepo) GetMeta(ctx context.Context, commit string) (content.Meta, error) {
	return connect(r.log, r.repo.User, r.repo.Password, r.repo.URL.String()).
		getMeta(ctx, commit)
}

func (r sourceRepo) GetFiles(ctx context.Context, commit string) ([]content.File, error) {
	return connect(r.log, r.repo.User, r.repo.Password, r.repo.URL.String()).
		GetFiles(ctx, commit, r.repo.Repo)
}
