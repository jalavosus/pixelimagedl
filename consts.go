package pixelimagedl

const (
	defaultBufSizeBytes int    = 12 * 1024 * 1024
	bufSizeEnv          string = "DL_BUF_SIZE"
)

const (
	downloadsUrl           string = "https://developers.google.com/android"
	StableFactoryImagesURL        = downloadsUrl + "/images"
	StableOTAImagesURL            = downloadsUrl + "/ota"
)

const (
	acksCookie        string = "devsite_wall_acks="
	OTAAcksCookie            = acksCookie + "nexus-ota-tos"
	FactoryAcksCookie        = acksCookie + "nexus-image-tos"
)
