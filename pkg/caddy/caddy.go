// Package caddy implement a caddy "uncloud" storage module, that allows it to store certificates in corrosion.
// This is used to build a custom uncloud caddy together with the L4 app. This code in this package is a thing
// wrapper over internal/machine/store and is used from uncloud/caddy which has the other (Go) files to make
// it an offical - ready to use - caddy module.
//
// The global caddy section just needs to reference the corrosion endpoint for this to work.
//
//	storage uncloud <api-endpoint>
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
	addr netip.AddrPort
	s    *store.Store
}

// CaddyModule register the module in caddy. The returned ModuleInfo must have both a name and a constructor function.
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
//	storage uncloud <api-endpoint>
func (c *Corrosion) UnmarshalCaddyfile(d *caddyfile.Dispenser) (err error) {
	for d.Next() {
		if !d.NextArg() {
			return d.ArgErr()
		}
		c.addr, err = netip.ParseAddrPort(d.Val())
	}
	return err
}

var (
	_ caddy.StorageConverter = (*Corrosion)(nil)
	_ caddyfile.Unmarshaler  = (*Corrosion)(nil)
	_ certmagic.Storage      = (*Corrosion)(nil)
)
