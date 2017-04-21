// mixin
package vapor

import (
	"strings"
)

var mixins map[string]string

func isMixinInitializer(s string) bool {
	return strings.HasPrefix(s, "mixin ")
}
