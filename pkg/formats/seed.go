package formats

type ImageFormat string

const (
	Gif      ImageFormat = "image/gif"
	Jpeg     ImageFormat = "image/jpeg"
	Png      ImageFormat = "image/png"
	Png8     ImageFormat = "image/png8"
	JpegPng  ImageFormat = "image/vnd.jpeg-png"
	JpegPng8 ImageFormat = "image/vnd.jpeg-png8"
)
