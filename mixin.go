// mixin
package vapor

import (
	"strings"
)

var mixins map[string]vaporTree

type mixin struct {
	*block
}

func isMixinInitializer(s string) bool {
	return strings.HasPrefix(s, "mixin ")
}

func newMixin() (*mixin, *vaporError) {
	return nil, nil
}
