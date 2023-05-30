package util

const (
	// init magic number
	MAGIC_NUMBER = 101
	VERSION = 1

	// net
	NET_TYPE_REQUEST = 1
	NET_TYPE_ACK     = 1 << 1
	NET_TYPE_INIT    = 1 << 2
	NET_TYPE_INFORM  = 1 << 3
	NET_TYPE_QUIT    = 1 << 4

	// memory
	MEM_FLAG_WORKING = 1
	MEM_FLAG_COMMON  = 0
	MEM_HEADER_SIZE  = 9

	FILE_NAME_SUFFIX = ".txt"
	NONE             = 0
	CLIENT           = 1
	SERVER           = 2

	CLIENT_ONLY = 0
	SERVER_ONLY = 1
	DUPLEX      = 2
	EVENTS_SIZE = 1024
	
	TCP_CONNECTING = 0
	MEM_INIT_REQU = 1
	PROTOCL_OVER = 2

	SHACK_RESPONSE = "yes"
	
)
