package opengl

// #cgo LDFLAGS: -framework OpenGL
//
// #include <OpenGL/gl.h>
// #include <stdlib.h>
import "C"
import (
	"image"
	"github.com/hajimehoshi/go-ebiten/graphics"
)

type Device struct {
	screenWidth int
	screenHeight int
	screenScale int
	graphicsContext *GraphicsContext
	offscreenTexture *Texture
	drawFunc func(graphics.GraphicsContext, graphics.Texture)
	funcs []func()
}

func NewDevice(screenWidth, screenHeight, screenScale int,
	drawFunc func(graphics.GraphicsContext, graphics.Texture)) *Device {
	device := &Device{
		screenWidth: screenWidth,
		screenHeight: screenHeight,
		screenScale: screenScale,
		graphicsContext: newGraphicsContext(screenWidth, screenHeight, screenScale),
		drawFunc: drawFunc,
		funcs: []func(){},
	}
	device.offscreenTexture = device.NewTexture(screenWidth, screenHeight).(*Texture)
	return device
}

func (device *Device) Update() {
	for _, f := range device.funcs {
		f()
	}
	device.funcs = []func(){}

	g := device.graphicsContext
	C.glEnable(C.GL_TEXTURE_2D)
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MIN_FILTER, C.GL_NEAREST)
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MAG_FILTER, C.GL_NEAREST)
	g.SetOffscreen(device.offscreenTexture)
	g.Clear()
	device.drawFunc(g, device.offscreenTexture)
	g.flush()
	
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MIN_FILTER, C.GL_LINEAR)
	C.glTexParameteri(C.GL_TEXTURE_2D, C.GL_TEXTURE_MAG_FILTER, C.GL_LINEAR)
	g.resetOffscreen()
	g.Clear()
	geometryMatrix := graphics.IdentityGeometryMatrix()
	geometryMatrix.SetA(float64(g.screenScale))
	geometryMatrix.SetD(float64(g.screenScale))
	g.DrawTexture(device.offscreenTexture,
		0, 0, device.screenWidth, device.screenHeight,
		geometryMatrix, graphics.IdentityColorMatrix())
	g.flush()
}

func (device *Device) NewTexture(width, height int) graphics.Texture {
	return createTexture(device, width, height, nil)
}

func (device *Device) NewTextureFromImage(img image.Image) graphics.Texture {
	var pix []uint8
	switch img.(type) {
	case *image.RGBA:
		pix = img.(*image.RGBA).Pix
	case *image.NRGBA:
		pix = img.(*image.NRGBA).Pix
	default:
		panic("image should be RGBA or NRGBA")
	}
	size := img.Bounds().Size()
	return createTexture(device, size.X, size.Y, pix)
}

func (device *Device) executeWhenDrawing(f func()) {
	device.funcs = append(device.funcs, f)
}
