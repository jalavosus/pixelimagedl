package internal

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
