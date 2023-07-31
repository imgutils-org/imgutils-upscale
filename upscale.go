// Package upscale provides image upscaling utilities.
package upscale

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/draw"
)

// Algorithm specifies the upscaling algorithm to use.
type Algorithm int

const (
	// NearestNeighbor is fastest but produces blocky results.
	NearestNeighbor Algorithm = iota
	// Bilinear provides smooth results with decent performance.
	Bilinear
	// CatmullRom provides high quality results.
	CatmullRom
)

// Options configures the upscale operation.
type Options struct {
	Algorithm Algorithm
}

// DefaultOptions returns sensible defaults.
func DefaultOptions() Options {
	return Options{
		Algorithm: CatmullRom,
	}
}

// ByFactor upscales an image by a multiplication factor.
func ByFactor(src image.Image, factor float64, opts Options) image.Image {
	if factor <= 0 {
		factor = 1
	}

	bounds := src.Bounds()
	newW := int(float64(bounds.Dx()) * factor)
	newH := int(float64(bounds.Dy()) * factor)

	return ToSize(src, newW, newH, opts)
}

// ToSize upscales an image to the specified dimensions.
func ToSize(src image.Image, width, height int, opts Options) image.Image {
	if width <= 0 || height <= 0 {
		return src
	}

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	var scaler draw.Scaler
	switch opts.Algorithm {
	case NearestNeighbor:
		scaler = draw.NearestNeighbor
	case Bilinear:
		scaler = draw.BiLinear
	default:
		scaler = draw.CatmullRom
	}

	scaler.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}

// ToWidth upscales an image to a specific width, maintaining aspect ratio.
func ToWidth(src image.Image, width int, opts Options) image.Image {
	bounds := src.Bounds()
	ratio := float64(bounds.Dy()) / float64(bounds.Dx())
	height := int(float64(width) * ratio)
	return ToSize(src, width, height, opts)
}

// ToHeight upscales an image to a specific height, maintaining aspect ratio.
func ToHeight(src image.Image, height int, opts Options) image.Image {
	bounds := src.Bounds()
	ratio := float64(bounds.Dx()) / float64(bounds.Dy())
	width := int(float64(height) * ratio)
	return ToSize(src, width, height, opts)
}

// Double doubles the size of an image.
func Double(src image.Image) image.Image {
	return ByFactor(src, 2, DefaultOptions())
}

// Triple triples the size of an image.
func Triple(src image.Image) image.Image {
	return ByFactor(src, 3, DefaultOptions())
}

// Quadruple quadruples the size of an image.
func Quadruple(src image.Image) image.Image {
	return ByFactor(src, 4, DefaultOptions())
}

// UpscaleFromFile reads an image file and upscales it.
func UpscaleFromFile(path string, factor float64, opts Options) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return ByFactor(src, factor, opts), nil
}

// SaveJPEG saves the upscaled image as JPEG.
func SaveJPEG(img image.Image, w io.Writer, quality int) error {
	if quality <= 0 || quality > 100 {
		quality = 85
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
}

// SavePNG saves the upscaled image as PNG.
func SavePNG(img image.Image, w io.Writer) error {
	return png.Encode(w, img)
}
