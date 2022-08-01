package utils

import (
    "strings"
    "bytes"
    "strconv"
    "encoding/hex"

    log "github.com/sirupsen/logrus"
    "github.com/Arachnid/evmdis"
)

const swarmHashLength = 43

var swarmHashProgramTrailer = [...]byte{0x00, 0x29}
var swarmHashHeader = [...]byte{0xa1, 0x65}

func DissasembleWithoutPushValues(address string, data string) string {
    processedData := bytes.TrimSpace([]byte(data))
    bytecode := make([]byte, hex.DecodedLen(len(processedData)))

    _, err := hex.Decode(bytecode, processedData)
    if err != nil {
        log.Panic("Could not decode hex string: %v", err)
    }

    // detect swarm hash and remove it from bytecode

	bytecodeLength := len(bytecode)
	if bytecode[bytecodeLength-1] == swarmHashProgramTrailer[1] &&
		bytecode[bytecodeLength-2] == swarmHashProgramTrailer[0] &&
		bytecode[bytecodeLength-43] == swarmHashHeader[0] &&
		bytecode[bytecodeLength-42] == swarmHashHeader[1] {
		bytecodeLength -= swarmHashLength // remove swarm part
	}

    // process opcodes
    var opcodes []string
    i := 0
    for i < bytecodeLength { 
        op := evmdis.OpCode(bytecode[i])
        if op.IsPush() {
            skip, _ := strconv.Atoi(strings.ReplaceAll(op.String(), "PUSH", ""))
            i += skip
        }
        opcode := op.String()
        if !strings.HasPrefix(opcode, "Missing opcode") {
            opcodes = append(opcodes, opcode)
        }
        i++
    }

    return strings.Join(opcodes, " ")
}
