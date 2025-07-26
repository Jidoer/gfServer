package protorpc

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	// "fmt"
)

const (
    PROTORPCName       = "MRPC"
    PROTORPCVersion    = 1
    PROTORPCHeadLength = 12
)

type ProtoRPCHead struct {
    Protocol [4]byte
    Version  byte
    Flags    byte
    Reserved [2]byte
    Length   uint32
}

type ProtoRPCMessage struct {
    Head ProtoRPCHead
    Body []byte
}

// 计算包的总长度
func ProtorpcPackageLength(head *ProtoRPCHead) int {
    return PROTORPCHeadLength + int(head.Length)
}

// 初始化头部
func ProtorpcHeadInit(head *ProtoRPCHead) {
    copy(head.Protocol[:], PROTORPCName)
    head.Version = PROTORPCVersion
    head.Reserved[0] = 0
    head.Reserved[1] = 0
    head.Length = 0
}

// 初始化消息
func ProtorpcMessageInit(msg *ProtoRPCMessage) {
    ProtorpcHeadInit(&msg.Head)
    msg.Body = nil
}

// 检查头部是否合法
func ProtorpcHeadCheck(head *ProtoRPCHead) error {
    if head.Protocol != [4]byte{77, 82, 80, 67} { // "MRPC" 的 ASCII 值
        return errors.New("invalid protocol")
    }
    if head.Version != PROTORPCVersion {
        return errors.New("invalid version")
    }
    return nil
}

// 打包函数
func ProtorpcPack(msg *ProtoRPCMessage, buf *[]byte) (int, error) {
    if msg == nil || buf == nil || len(*buf) == 0 {
        return -1, errors.New("invalid input")
    }
    packlen := ProtorpcPackageLength(&msg.Head)
    if len(*buf) < packlen {
        return -2, errors.New("buffer too small")
    }

    p := 0
    copy((*buf)[p:], msg.Head.Protocol[:])
    p += 4
    (*buf)[p] = msg.Head.Version
    p++
    (*buf)[p] = msg.Head.Flags
    p++
    copy((*buf)[p:], msg.Head.Reserved[:])
    p += 2

    // hton length
    binary.BigEndian.PutUint32((*buf)[p:], msg.Head.Length)
    p += 4

    // memcpy body
    if msg.Body != nil && msg.Head.Length > 0 {
        copy((*buf)[p:], msg.Body)
    }
    fmt.Printf("msg body: %s\n", hex.Dump(msg.Body))
    //dump hex
    fmt.Printf("output_.Hex: %s\n", hex.Dump((*buf)[:packlen]))

    return packlen, nil
}

// 解包函数
func ProtorpcUnpack(msg *ProtoRPCMessage, buf []byte) (int, error) {
    if msg == nil || buf == nil || len(buf) == 0 {
        return -1, errors.New("invalid input")
    }
    if len(buf) < PROTORPCHeadLength {
        return -2, errors.New("buffer too small")
    }

    p := 0
    copy(msg.Head.Protocol[:], buf[p:p+4])
    p += 4
    msg.Head.Version = buf[p]
    p++
    msg.Head.Flags = buf[p]
    p++
    copy(msg.Head.Reserved[:], buf[p:p+2])
    p += 2

    // ntoh length
    msg.Head.Length = binary.BigEndian.Uint32(buf[p : p+4])
    p += 4

    packlen := ProtorpcPackageLength(&msg.Head)
    if len(buf) < packlen {
        return -3, errors.New("buffer too small")
    }

    // NOTE: just shadow copy
    if len(buf) > PROTORPCHeadLength {
        msg.Body = buf[PROTORPCHeadLength:len(buf)]
    }

    return packlen, nil
}