package gcsfile

import (
	"context"

	"golang.org/x/xerrors"

	"cloud.google.com/go/storage"
)

const (
	blockSize = 4096
)

func NewFile(ctx context.Context, bucket, name string) (*File, error) {
	// TODO: refactoring...
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to new gcs client: %w", err)
	}

	handler := client.Bucket(bucket).Object(name)
	attr, err := handler.Attrs(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to get attribute: %w", err)
	}

	return &File{
		ctx:       ctx,
		size:      attr.Size,
		blockSize: blockSize,
		handler:   handler,
	}, nil
}

func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
	r, err := f.handler.NewRangeReader(f.ctx, off, f.blockSize)
	if err != nil {
		return 0, err
	}
	return r.Read(p)
}

type File struct {
	ctx       context.Context
	size      int64
	handler   *storage.ObjectHandle
	blockSize int64
}
