# imgutils-upscale

[![Go Reference](https://pkg.go.dev/badge/github.com/imgutils-org/imgutils-upscale.svg)](https://pkg.go.dev/github.com/imgutils-org/imgutils-upscale)
[![Go Report Card](https://goreportcard.com/badge/github.com/imgutils-org/imgutils-upscale)](https://goreportcard.com/report/github.com/imgutils-org/imgutils-upscale)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library for upscaling images with high-quality interpolation. Part of the [imgutils](https://github.com/imgutils-org) collection.

## Features

- Factor-based upscaling (2x, 3x, 4x, etc.)
- Target dimension upscaling
- Multiple interpolation algorithms
- Aspect ratio preservation
- Convenience functions for common operations

## Installation

```bash
go get github.com/imgutils-org/imgutils-upscale
```

## Quick Start

```go
package main

import (
    "image"
    "os"

    "github.com/imgutils-org/imgutils-upscale"
)

func main() {
    // Open small image
    file, _ := os.Open("small.jpg")
    defer file.Close()
    src, _, _ := image.Decode(file)

    // Double the size
    large := upscale.Double(src)

    // Save result
    out, _ := os.Create("large.jpg")
    defer out.Close()
    upscale.SaveJPEG(large, out, 90)
}
```

## Usage Examples

### Quick Upscale

```go
// Double the size (2x)
doubled := upscale.Double(src)

// Triple the size (3x)
tripled := upscale.Triple(src)

// Quadruple the size (4x)
quadrupled := upscale.Quadruple(src)
```

### Custom Factor

```go
// Upscale by 1.5x
opts := upscale.DefaultOptions()
result := upscale.ByFactor(src, 1.5, opts)

// Upscale by 2.5x
result := upscale.ByFactor(src, 2.5, opts)
```

### To Specific Dimensions

```go
// Upscale to exact size
result := upscale.ToSize(src, 1920, 1080, upscale.DefaultOptions())

// Upscale to width (height calculated to maintain ratio)
result := upscale.ToWidth(src, 1920, upscale.DefaultOptions())

// Upscale to height (width calculated to maintain ratio)
result := upscale.ToHeight(src, 1080, upscale.DefaultOptions())
```

### Interpolation Algorithms

```go
// Fastest (blocky results, good for pixel art)
opts := upscale.Options{Algorithm: upscale.NearestNeighbor}
result := upscale.Double(src)

// Smooth results
opts := upscale.Options{Algorithm: upscale.Bilinear}
result := upscale.ByFactor(src, 2, opts)

// Highest quality (default)
opts := upscale.Options{Algorithm: upscale.CatmullRom}
result := upscale.ByFactor(src, 2, opts)
```

### From File

```go
// Load and upscale in one step
large, err := upscale.UpscaleFromFile("small.jpg", 2.0, upscale.DefaultOptions())
if err != nil {
    log.Fatal(err)
}
```

## API Reference

### Types

#### Algorithm

```go
type Algorithm int

const (
    NearestNeighbor Algorithm = iota // Fastest, blocky
    Bilinear                         // Smooth, fast
    CatmullRom                       // High quality
)
```

#### Options

```go
type Options struct {
    Algorithm Algorithm
}
```

### Functions

| Function | Description |
|----------|-------------|
| `DefaultOptions()` | Returns defaults (CatmullRom) |
| `ByFactor(src, factor, opts)` | Upscale by multiplication factor |
| `ToSize(src, w, h, opts)` | Upscale to exact dimensions |
| `ToWidth(src, w, opts)` | Upscale to width, maintain ratio |
| `ToHeight(src, h, opts)` | Upscale to height, maintain ratio |
| `Double(src)` | 2x upscale with default options |
| `Triple(src)` | 3x upscale with default options |
| `Quadruple(src)` | 4x upscale with default options |
| `UpscaleFromFile(path, factor, opts)` | Load and upscale from file |
| `SaveJPEG(img, w, quality)` | Save as JPEG |
| `SavePNG(img, w)` | Save as PNG |

## Algorithm Comparison

| Algorithm | Quality | Speed | Best For |
|-----------|---------|-------|----------|
| NearestNeighbor | Low | Fastest | Pixel art, retro graphics |
| Bilinear | Medium | Fast | General use, previews |
| CatmullRom | High | Slower | Photos, final output |

## Use Cases

### Print Preparation

```go
// Upscale for high-DPI printing (300 DPI from 72 DPI)
factor := 300.0 / 72.0 // ~4.17x
printReady := upscale.ByFactor(src, factor, upscale.DefaultOptions())
```

### Retro Game Graphics

```go
// Pixel-perfect upscaling for retro games
opts := upscale.Options{Algorithm: upscale.NearestNeighbor}
pixelArt := upscale.ByFactor(sprite, 4, opts)
```

### Thumbnail to Full Size

```go
// Expand thumbnail for preview
preview := upscale.ToSize(thumbnail, 800, 600, upscale.DefaultOptions())
```

## Requirements

- Go 1.16 or later

## Related Packages

- [imgutils-resize](https://github.com/imgutils-org/imgutils-resize) - General resizing (up and down)
- [imgutils-thumbnail](https://github.com/imgutils-org/imgutils-thumbnail) - Thumbnail generation
- [imgutils-sdk](https://github.com/imgutils-org/imgutils-sdk) - Unified SDK

## License

MIT License - see [LICENSE](LICENSE) for details.
