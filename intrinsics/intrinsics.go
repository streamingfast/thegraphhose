package intrinsics

import (
	"github.com/pkg/errors"
)

// const eosioIntrinsicsNamespace = "env"

// var intrinsicsImports *wasmer.Imports = wasmer.NewImports().Namespace(eosioIntrinsicsNamespace)

// func GetIntrinsicsImports() *wasmer.Imports {
// 	return intrinsicsImports
// }

// func GetMemoryFromContext(context unsafe.Pointer) []byte {
// 	instanceContext := wasmer.IntoInstanceContext(context)

// 	return instanceContext.Memory().Data()
// }

func CheckError(err error, message string, arguments ...interface{}) {
	if err != nil {
		panic(errors.Wrapf(err, message, arguments...))
	}
}
