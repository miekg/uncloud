package caddy

import (
	"context"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/psviderski/uncloud/internal/corrosion"
	"github.com/psviderski/uncloud/internal/machine/store"
)

func (c *Corrosion) Store(ctx context.Context, key string, value []byte) error {
	if c.s == nil {
		if err := c.init(); err != nil {
			return err
		}
	}

	return c.s.Put(ctx, key, value)
}

func (c *Corrosion) Load(ctx context.Context, key string) ([]byte, error) {
	if c.s == nil {
		if err := c.init(); err != nil {
			return nil, err
		}
	}

	value := []byte{}
	err := c.s.Get(ctx, key, &value)
	return value, err
}

func (c *Corrosion) Delete(ctx context.Context, key string) error {
	if c.s == nil {
		if err := c.init(); err != nil {
			return err
		}
	}

	return c.s.Delete(ctx, key)
}

func (c *Corrosion) Exists(ctx context.Context, key string) bool {
	if c.s == nil {
		if err := c.init(); err != nil {
			return false
		}
	}

	_, err := c.Load(ctx, key)
	return err == nil
}

func (c *Corrosion) Stat(ctx context.Context, key string) (certmagic.KeyInfo, error) {
	if c.s == nil {
		if err := c.init(); err != nil {
			return certmagic.KeyInfo{}, err
		}
	}

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
	if c.s == nil {
		if err := c.init(); err != nil {
			return nil, err
		}
	}

	list, err := c.s.List(ctx, prefix)
	if err != nil {
		return nil, err
	}
	// TODO(miek): handle recurse
	return list, nil
}

func (s *Corrosion) Lock(ctx context.Context, key string) error   { return nil }
func (c *Corrosion) Unlock(ctx context.Context, key string) error { return nil }

func (c *Corrosion) init() error {
	corro, err := corrosion.NewAPIClient(c.addr)
	if err != nil {
		return err
	}
	c.s = store.New(corro)
	return nil
}
