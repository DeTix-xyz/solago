package sdk

import (
	"bytes"
)

type Instruction struct {
	ProgramIDIndex        SerializableUInt8
	AccountAddressIndexes *CompactArray[SerializableUInt8]
	Data                  *CompactArray[Serializable]
}

func (instruction *Instruction) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	instruction.ProgramIDIndex.Serialize(buffer)
	instruction.AccountAddressIndexes.Serialize(buffer)
	instruction.Data.Serialize(buffer)

	return buffer
}

// type InstructionData struct {
// 	Data Serializable
// }

// func (instructionData *InstructionData) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
// 	// we may simply be passed a byte array, just write it to the buffer
// 	bytes, ok := instructionData.Data.([]byte)

// 	if ok {
// 		buffer.Write(bytes)
// 		return buffer
// 	}

// 	// otherwise parse the struct
// 	structValues := reflect.ValueOf(instructionData.Data)

// 	for i := 0; i < structValues.NumField(); i++ {
// 		binary.Write(buffer, binary.LittleEndian, structValues.Field(i).Interface())
// 	}

// 	return buffer
// }
