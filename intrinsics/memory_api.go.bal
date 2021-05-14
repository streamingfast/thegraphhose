package intrinsics

import (
    "fmt"
    "unsafe"
)

// #include <stdlib.h>
//
// extern int32_t eosio_memcpy(void *context, int32_t dest, int32_t src, int32_t length);
// extern int32_t eosio_memmove(void *context, int32_t dest, int32_t src, int32_t length);
// extern int32_t eosio_memcmp(void *context, int32_t dest, int32_t src, int32_t length);
// extern int32_t eosio_memset(void *context, int32_t dest, int32_t value, int32_t length);
import "C"

func init() {
    intrinsicsImports.Append("memcpy", eosio_memcpy, C.eosio_memcpy)
    intrinsicsImports.Append("memmove", eosio_memmove, C.eosio_memmove)
    intrinsicsImports.Append("memcmp", eosio_memcmp, C.eosio_memcmp)
    intrinsicsImports.Append("memset", eosio_memset, C.eosio_memset)
}

//export eosio_memcpy
func eosio_memcpy(context unsafe.Pointer, dest int32, src int32, length int32) int32 {
    err := CopyMemoryRange(GetMemoryFromContext(context), dest, src, length)
    CheckError(err, "unable to perform memcpy")

    return dest
}

//export eosio_memmove
func eosio_memmove(context unsafe.Pointer, dest int32, src int32, length int32) int32 {
    fmt.Println("Called eosio_memmove")
    return 0
}

//export eosio_memcmp
func eosio_memcmp(context unsafe.Pointer, dest int32, src int32, length int32) int32 {
    memory := GetMemoryFromContext(context)

    expectedBuffer, err := ReadMemoryRangeWithLength(memory, dest, length)
    CheckError(err, "unable to read destination buffer")

    actualBuffer, err := ReadMemoryRangeWithLength(memory, src, length)
    CheckError(err, "unable to read source buffer")

    for i := int32(0); i < length; i++ {
        if expectedBuffer[i] < actualBuffer[i] {
            return -1
        }

        if expectedBuffer[i] > actualBuffer[i] {
            return 1
        }
    }

    return 0
}

//export eosio_memset
func eosio_memset(context unsafe.Pointer, dest int32, value int32, length int32) int32 {
    memory := GetMemoryFromContext(context)

    // FIXME: Implement boundaries checks!
    for i := int32(0); i < length; i++ {
        // Copied from C++ reference:
        //  > Value to be set. The value is passed as an int, but the function fills the block of memory using the unsigned char conversion of this value.
        memory[dest + i] = byte(value)
    }

    return dest
}

