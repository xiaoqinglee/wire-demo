package pkgBar

import "github.com/google/wire"

var PkgBarProviderSet = wire.NewSet(NewTypeC, NewTypeD)
