package pkgFoo

import "github.com/google/wire"

var PkgFooProviderSet = wire.NewSet(NewTypeA, NewTypeB)
