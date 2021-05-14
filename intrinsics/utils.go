package intrinsics

import (
	"fmt"
)

// var UUUUUUURG_EncodedActionData []byte

type ApplyContext struct {
	Receiver int64
	Account  int64
	Action   int64
}

// func NewApplyContext(receiver, account, action string, abi *eos.ABI, actionDataJSON string) *ApplyContext {
// 	nReceiver, err := ToNameEncoded(receiver)
// 	ErrorCheck(err, fmt.Sprintf("Unable to encode receiver [%s]", receiver))
// 	nAccount, err := ToNameEncoded(account)
// 	ErrorCheck(err, fmt.Sprintf("Unable to encode account [%s]", account))
// 	nAction, err := ToNameEncoded(action)
// 	ErrorCheck(err, fmt.Sprintf("Unable to encode action [%s]", action))

// 	encodedActionData, err := abi.EncodeAction(eos.ActionName(action), []byte(actionDataJSON))
// 	ErrorCheck(err, fmt.Sprintf("unable to encode action %q", actionDataJSON))

// 	//TODO: AAAAAAAAAAAAAARG ugly hack
// 	// figure out this shit:  cgo argument has Go pointer to Go pointer
// 	UUUUUUURG_EncodedActionData = encodedActionData

// 	return &ApplyContext{
// 		Receiver: nReceiver,
// 		Account:  nAccount,
// 		Action:   nAction,
// 	}
// }

func ReadCStringAtOffset(memory []byte, startOffset int32) (string, error) {
	for i, byteValue := range memory[startOffset:] {
		if byteValue == 0 {
			return ReadMemoryRangeAsString(memory, startOffset, int32(i)+startOffset)
		}
	}

	return "", fmt.Errorf("C string terminator not found from offset %d", startOffset)
}

func ReadMemoryRange(memory []byte, startOffset int32, endOffset int32) ([]byte, error) {
	byteCount := int32(len(memory))
	if startOffset < 0 || startOffset >= byteCount || endOffset > byteCount {
		return nil, fmt.Errorf("offset [%d, %d[", startOffset, endOffset)
	}

	return memory[startOffset:endOffset], nil
}

func ReadMemoryRangeWithLength(memory []byte, startOffset int32, length int32) ([]byte, error) {
	endOffset := startOffset + length
	return ReadMemoryRange(memory, startOffset, endOffset)
}

func ReadMemoryRangeAsString(memory []byte, startOffset int32, endOffset int32) (string, error) {
	value, err := ReadMemoryRange(memory, startOffset, endOffset)
	if err != nil {
		return "", err
	}

	return string(value), nil
}

func ReadMemoryRangeAsStringOfLength(memory []byte, startOffset int32, strLength int32) (string, error) {
	value, err := ReadMemoryRangeWithLength(memory, startOffset, strLength)
	if err != nil {
		return "", err
	}

	return string(value), nil
}

func CopyMemoryRange(memory []byte, destinationOffset int32, startOffset int32, length int32) error {
	copy(memory[destinationOffset:destinationOffset+length], memory[startOffset:startOffset+length])
	return nil
}

func WriteMemoryRange(memory []byte, data []byte, destinationOffset int32, length int32) error {
	copy(memory[destinationOffset:destinationOffset+length], data)
	return nil
}
