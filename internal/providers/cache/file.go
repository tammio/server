package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/easyp-tech/server/internal/providers/content"
)

type Local struct {
	Dir string
}

func (c Local) Get(_ context.Context, owner, repoName, commit, configHash string) ([]content.File, error) {
	if c.Dir == "" {
		return nil, nil
	}

	fullName := path.Join(c.Dir, owner, repoName, configHash, commit+".json")

	data, err := os.ReadFile(fullName)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}

		return nil, fmt.Errorf("reading %q: %w", fullName, err)
	}

	var out []content.File

	if err = json.Unmarshal(data, &out); err != nil { //nolint:musttag
		return nil, fmt.Errorf("decoding %q: %w", fullName, err)
	}

	return out, nil
}

func (c Local) Put(_ context.Context, owner, repoName, commit, configHash string, in []content.File) error {
	if c.Dir == "" {
		return nil
	}

	fullDir := path.Join(c.Dir, owner, repoName, configHash)

	err := os.MkdirAll(fullDir, 0o750) //nolint:gomnd
	if err != nil {
		return fmt.Errorf("creating dir %q: %w", fullDir, err)
	}

	var (
		fileName = path.Join(fullDir, commit+".json")
		tmpName  = fileName + ".tmp"
	)

	file, err := os.Create(tmpName)
	if err != nil {
		return fmt.Errorf("creating %q: %w", tmpName, err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err = encoder.Encode(in); err != nil { //nolint:musttag
		return fmt.Errorf("writing %q: %w", tmpName, err)
	}

	if err = os.Rename(tmpName, fileName); err != nil {
		return fmt.Errorf("renaming %q to %q: %w", tmpName, fileName, err)
	}

	return nil
}

func (c Local) Ping(_ context.Context) error {
	if _, err := os.Stat(c.Dir); err != nil {
		return fmt.Errorf("local cache dir inaccessible: %w", err)
	}
	return nil
}
