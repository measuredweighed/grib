package griblib

import (
	"bytes"
	"image/color"
	"image/png"
	"io"
	"math"

	"gopkg.in/gographics/imagick.v3/imagick"
)

type Data40 struct {
	Reference              float32 `json:"reference"`
	BinaryScale            uint16  `json:"binaryScale"`
	DecimalScale           uint16  `json:"decimalScale"`
	Bits                   uint8   `json:"bits"`
	Type                   uint8   `json:"type"`
	CompressionType        uint8   `json:"compressionType"`
	TargetCompressionRatio uint8   `json:"targetCompressionRatio"`
}

func (template Data40) getRefScale() (float64, float64) {
	bscale := math.Pow(2.0, float64(template.BinaryScale))
	dscale := math.Pow(10.0, -float64(template.DecimalScale))

	scale := bscale * dscale
	ref := dscale * float64(template.Reference)

	return ref, scale
}

func (template Data40) scaleFunc() func(uintValue int64) float64 {
	ref, scale := template.getRefScale()
	return func(value int64) float64 {
		signed := int64(value)
		return ref + float64(signed)*scale
	}
}

// ParseData40 parses data40 struct from the reader into the an array of floating-point values
func ParseData40(dataReader io.Reader, dataLength int, template *Data40) ([]float64, error) {

	if dataLength == 0 {
		return nil, nil
	}

	// Fetch raw byte stream for JPEG2000
	rawByteData := make([]byte, dataLength)
	dataReader.Read(rawByteData)

	// Initialise Imagick and use it to convert JPEG2000 to a PNG Image representation
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImageBlob(rawByteData); err != nil {
		return nil, err
	}

	mw.SetImageFormat("png")
	img, err := png.Decode(bytes.NewReader(mw.GetImageBlob()))
	if err != nil {
		return nil, err
	}

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	data := make([]float64, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y*width + x
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			data[idx] = float64(c.Y)
		}
	}

	return data, nil
}
