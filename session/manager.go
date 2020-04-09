package session

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/zeebo/errs"

	"github.com/jtolds/jam/backends"
	"github.com/jtolds/jam/blobs"
	"github.com/jtolds/jam/manifest"
	"github.com/jtolds/jam/pathdb"
	"github.com/jtolds/jam/streams"
	"github.com/jtolds/jam/utils"
)

const (
	pathPrefix = "manifests/"
	timeFormat = "2006/01/02/15-04-05.000000000"
)

func timestampToPath(timestamp time.Time) string {
	return pathPrefix + timestamp.UTC().Format(timeFormat)
}

func pathToTimestamp(path string) (time.Time, error) {
	if !strings.HasPrefix(path, pathPrefix) {
		return time.Time{}, errs.New("backend had unexpected behavior: path returned does not start with %q: %q", pathPrefix, path)
	}
	ts, err := time.ParseInLocation(timeFormat, strings.TrimPrefix(path, pathPrefix), time.UTC)
	return ts, errs.Wrap(err)
}

type Snapshot interface {
	List(ctx context.Context, prefix, delimiter string, cb func(ctx context.Context, entry *ListEntry) error) error
	Open(ctx context.Context, path string) (*manifest.Metadata, *streams.Stream, error)
	Close() error
}

type Manager struct {
	backend backends.Backend
	logger  utils.Logger
	blobs   *blobs.Store
}

func NewManager(logger utils.Logger, backend backends.Backend, blobStore *blobs.Store) *Manager {
	return &Manager{
		backend: backend,
		logger:  logger,
		blobs:   blobStore,
	}
}

func (s *Manager) ListSnapshots(ctx context.Context,
	cb func(ctx context.Context, timestamp time.Time) error) error {
	// TODO: backend.List is not ordered. we could use the fact that manifests are stored
	//		using timeFormat format and list by years and months in decreasing order to get
	// 		an order
	return errs.Wrap(s.backend.List(ctx, pathPrefix, func(ctx context.Context, path string) error {
		timestamp, err := pathToTimestamp(path)
		if err != nil {
			s.logger.Printf("invalid manifest format: %q, skipping. error: %v", path, err)
			return nil
		}
		return cb(ctx, timestamp)
	}))
}

func (s *Manager) latestTimestamp(ctx context.Context) (time.Time, error) {
	var latest time.Time
	err := s.ListSnapshots(ctx, func(ctx context.Context, timestamp time.Time) error {
		if latest.IsZero() || timestamp.After(latest) {
			latest = timestamp
		}
		return nil
	})
	if err != nil {
		return latest, err
	}
	return latest, nil
}

func (s *Manager) LatestSnapshot(ctx context.Context) (Snapshot, error) {
	latest, err := s.latestTimestamp(ctx)
	if err != nil {
		return nil, err
	}
	if latest.IsZero() {
		return nil, fmt.Errorf("no snapshots exist yet")
	}
	return s.OpenSnapshot(ctx, latest)
}

func (s *Manager) openSession(ctx context.Context, timestamp time.Time) (*Session, error) {
	rc, err := s.backend.Get(ctx, timestampToPath(timestamp), 0)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errs.Combine(err, rc.Close())
	}()

	db, err := pathdb.Open(ctx, s.backend, s.blobs, rc)
	if err != nil {
		return nil, err
	}

	return newSession(s.backend, db, s.blobs), nil
}

func (s *Manager) OpenSnapshot(ctx context.Context, timestamp time.Time) (Snapshot, error) {
	return s.openSession(ctx, timestamp)
}

func (s *Manager) NewSession(ctx context.Context) (*Session, error) {
	latest, err := s.latestTimestamp(ctx)
	if err != nil {
		return nil, err
	}
	if latest.IsZero() {
		return newSession(s.backend, pathdb.New(s.backend, s.blobs), s.blobs), nil
	} else {
		return s.openSession(ctx, latest)
	}
}