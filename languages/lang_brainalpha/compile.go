package lang_brainalpha

import (
	"github.com/MattLimb/GoHAL/internal"
)

func compileAst(_ internal.Ast) (string, *internal.HalError) {
	return "", internal.NewCriticalHalError("brainalpha does not support transpilation", 0)
}
