package caddy

import (
	"net/netip"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/certmagic"
	"github.com/psviderski/uncloud/internal/machine/store"
)

func init() {
	caddy.RegisterModule(Corrosion{})
}

type Corrosion struct {
	endpoint netip.AddrPort
	s        *store.Store
}

// CaddyModule register the module in caddy. The returned ModuleInfo must have both a name and a constructor function. This method must not have
// any side-effects.
func (Corrosion) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "caddy.storage.uncloud",
		New: func() caddy.Module {
			return new(Corrosion)
		},
	}
}

// CertMagicStorage converts c to a certmagic.Storage instance.
func (c *Corrosion) CertMagicStorage() (certmagic.Storage, error) { return c, nil }

// UnmarshalCaddyfile sets up the storage module from Caddyfile tokens. Syntax:
//
//	uncloud <endpoint>
func (c *Corrosion) UnmarshalCaddyfile(d *caddyfile.Dispenser) (err error) {
	for d.Next() {
		if !d.NextArg() {
			return d.ArgErr()
		}
		c.endpoint, err = netip.ParseAddrPort(d.Val())
	}
	return err
}

var (
	_ caddy.StorageConverter = (*Corrosion)(nil)
	_ caddyfile.Unmarshaler  = (*Corrosion)(nil)
	_ certmagic.Storage      = (*Corrosion)(nil)
)
