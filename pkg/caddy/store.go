package caddy

import (
	"context"
	"time"

	"github.com/caddyserver/certmagic"
)

func (c *Corrosion) Store(ctx context.Context, key string, value []byte) error {
	return c.s.Put(ctx, key, value)
}

func (c *Corrosion) Load(ctx context.Context, key string) ([]byte, error) {
	value := []byte{}
	err := c.s.Get(ctx, key, &value)
	return value, err
}

func (c *Corrosion) Delete(ctx context.Context, key string) error { return c.s.Delete(ctx, key) }

func (c *Corrosion) Exists(ctx context.Context, key string) bool {
	_, err := c.Load(ctx, key)
	return err == nil
}

func (c *Corrosion) Stat(ctx context.Context, key string) (certmagic.KeyInfo, error) {
	value, err := c.Load(ctx, key)
	if err != nil {
		return certmagic.KeyInfo{}, err
	}

	return certmagic.KeyInfo{
		Key:        key,
		Modified:   time.Now().UTC(),
		Size:       int64(len(value)),
		IsTerminal: true,
	}, nil
}

func (c *Corrosion) List(ctx context.Context, prefix string, recursive bool) ([]string, error) {
	list, err := c.s.List(ctx, prefix)
	if err != nil {
		return nil, err
	}
	// TODO(miek): handle recurse
	return list, nil
}

func (s *Corrosion) Lock(ctx context.Context, key string) error   { return nil }
func (c *Corrosion) Unlock(ctx context.Context, key string) error { return nil }
