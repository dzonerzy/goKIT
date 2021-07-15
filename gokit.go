package goKIT

// #cgo LDFLAGS: -ladvapi32
// #include "kit.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

// Use KIT v0.2-alpha
const (
	// Globals
	KIT_DEFAULT_ID = "Local\\KIT"
	// Types
	KIT_TYPE_BIND       = 0x1
	KIT_TYPE_CONNECT    = 0x2
	KIT_TYPE_DISCONNECT = 0x3
	KIT_TYPE_ACCEPT     = 0x4
	KIT_TYPE_HANDSHAKE  = 0x5
	KIT_TYPE_DATA       = 0x6
	// Flags
	KIT_FLAG_DEFAULT          = 0x2
	KIT_FLAG_BINDED           = 0x4
	KIT_FLAG_CLOSED           = 0x8
	KIT_FLAG_CONNECTED        = 0x10
	KIT_FLAG_ACCEPTED         = 0x20
	KIT_FLAG_CLIENT_HANDSHAKE = 0x40
	KIT_FLAG_SERVER_HANDSHAKE = 0x80
	KIT_FLAG_DISCONNECTED     = 0x100
	KIT_FLAG_RESERVED1        = 0x10000
	KIT_FLAG_RESERVED2        = 0x20000
	KIT_FLAG_RESERVED3        = 0x40000
	KIT_FLAG_RESERVED4        = 0x80000
	// Errors
	KIT_OK                              = 0x0
	KIT_ERR_CREATE_FILE_MAPPING         = 0x8000
	KIT_ERR_MAP_VIEW_OF_FILE            = 0x8001
	KIT_ERR_CREATE_PACKET               = 0x8002
	KIT_ERR_WRITE_MAP                   = 0x8003
	KIT_BIND_FAILED                     = 0x8004
	KIT_INVALID_PARAMETER               = 0x8005
	KIT_CONNECT_FAILED                  = 0x8006
	KIT_MEMORY_READ_ERROR               = 0x8007
	KIT_TIMEOUT_ERROR                   = 0x8008
	KIT_KEY_GENERATION_FAILED           = 0x8009
	KIT_PACKET_MISMATCH_CRC32           = 0x800A
	KIT_SHARED_SECRET_GENERATION_FAILED = 0x800B
	KIT_KDF_FAILED                      = 0x800C
	KIT_INITIALIZATION_FAILED           = 0x800D
	KIT_CRYPTO_ERROR                    = 0x800E
	KIT_MEMORY_RESIZE_ERROR             = 0x800F
	KIT_TIMER_ERROR                     = 0x8010
	KIT_NO_MORE_SLOT                    = 0x8011
	// Data types
	KIT_DATA_TEXT   = 0x1
	KIT_DATA_BINARY = 0x2
	KIT_DATA_NONE   = 0x4
)

// Packet represent a kit packet
type Packet struct {
	_pkt C.pkpacket
}

// Return the packet content as a byte array
func (p *Packet) Content() []byte {
	return p._pkt.body.anon0[:p._pkt.body.length]
}

// ClientInfo represent a connected client
type ClientInfo struct {
	info C.pkclientinfo
}

// Return the connected client ID
func (ci *ClientInfo) ClientID() int {
	return int(ci.info.clientid)
}

// KIT represent the main interface to the kit library
type KIT struct {
	_instance C.kinstance
}

// Init initialize the library, call kit_init under the hood
func (k *KIT) Init() bool {
	return int(C.kit_init()) != 0
}

// Connect to a KIT server
func (k *KIT) Connect(id string) bool {
	var kid *C.char
	if id != "" {
		kid = C.CString(id)
	} else {
		kid = C.CString(KIT_DEFAULT_ID)
	}
	ret := C.kit_connect(kid, &k._instance)
	C.free(unsafe.Pointer(kid))
	return ret != 0
}

// Bind a KIT server over a named shared memory
func (k *KIT) Bind(id string) bool {
	var kid *C.char
	if id != "" {
		kid = C.CString(id)
	} else {
		kid = C.CString(KIT_DEFAULT_ID)
	}
	ret := C.kit_bind(kid, &k._instance)
	C.free(unsafe.Pointer(kid))
	return ret != 0
}

// ListenAndAccept listen and accept new clients
func (k *KIT) ListenAndAccept() (*ClientInfo, *KIT) {
	info := C.kit_listen_and_accept(&k._instance)
	if info != nil {
		return &ClientInfo{info: info}, &KIT{_instance: info.instance}
	}
	return nil, nil
}

// Write data over the shared memory
func (k *KIT) Write(data interface{}) bool {
	switch v := data.(type) {
	case string:
		var d = C.CString(v)
		ret := C.kit_write(&k._instance, (*C.uchar)(unsafe.Pointer(d)), C.ulonglong(len(v)))
		C.free(unsafe.Pointer(d))
		return ret != 0
	case []byte:
		return C.kit_write(&k._instance, (*C.uchar)(unsafe.Pointer(&v[0])), C.ulonglong(len(v))) != 0
	case int:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return C.kit_write(&k._instance, (*C.uchar)(unsafe.Pointer(&bs[0])), C.ulonglong(len(bs))) != 0
	default:
		return false
	}
}

// Read data on the shared memory
func (k *KIT) Read() *Packet {
	var pkt = C.kit_read(&k._instance)
	if pkt == nil {
		return nil
	}
	return &Packet{_pkt: pkt}
}

// IsDisconnect check if a packet is a disconnect packet
func (k *KIT) IsDisconnect(pkt *Packet) bool {
	if pkt._pkt.header._type == KIT_TYPE_DISCONNECT && pkt._pkt.header.flags == KIT_FLAG_DISCONNECTED {
		C.free(unsafe.Pointer(pkt._pkt))
		return true
	}
	return false
}

// NotifyDisconnect notify the underlying server when a client disconnect
func (k *KIT) NotifyDisconnect(info *ClientInfo) {
	C.kit_notify_disconnect(info.info)
}

// Disconnect close the connection
func (k *KIT) Disconnect() bool {
	return C.kit_disconnect(&k._instance) != 0
}

// Error return a human readable error
func (k *KIT) Error() error {
	return fmt.Errorf(C.GoString(C.kit_human_error()))
}

func (k *KIT) Select() int {
	panic("not implemented")
}

func (k *KIT) ErrorNum() int {
	panic("not implemented")
}
