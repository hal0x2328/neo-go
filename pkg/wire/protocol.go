package wire

type ProtocolVersion uint32

const (
	defaultVersion ProtocolVersion = 0
	cmdSize                        = 12
	UserAgent                      = "NEO-GO-V2" // TODO: This may be relocated to a config file
)

// ServiceFlag indicates the services provided by the node. 1 = P2P Full Node
type ServiceFlag uint64

const (
	NodePeerService ServiceFlag = 1
	// BloomFilerService ServiceFlag = 2 // Not implemented
	// PrunedNode        ServiceFlag = 3 // Not implemented
	// LightNode         ServiceFlag = 4 // Not implemented

)

// Magic is the network that NEO is running on
type Magic uint32

const (
	Production Magic = 0x00746e41
	Test       Magic = 0x74746e41
)