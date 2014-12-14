/*
Copyright 2014 Hajime Hoshi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ebiten

import (
	"github.com/go-gl/gl"
	"image"
)

// IsKeyPressed returns a boolean indicating whether key is pressed.
func IsKeyPressed(key Key) bool {
	return currentUI.input.isKeyPressed(key)
}

// CursorPosition returns a position of a mouse cursor.
func CursorPosition() (x, y int) {
	return currentUI.input.cursorPosition()
}

// IsMouseButtonPressed returns a boolean indicating whether mouseButton is pressed.
func IsMouseButtonPressed(mouseButton MouseButton) bool {
	return currentUI.input.isMouseButtonPressed(mouseButton)
}

func glFilter(f Filter) int {
	switch f {
	case FilterNearest:
		return gl.NEAREST
	case FilterLinear:
		return gl.LINEAR
	default:
		panic("not reached")
	}
}

// NewRenderTargetID returns a new RenderTargetID.
func NewRenderTargetID(width, height int, filter Filter) (RenderTargetID, error) {
	return currentUI.newRenderTargetID(width, height, glFilter(filter))
}

// NewTextureID returns a new TextureID.
func NewTextureID(img image.Image, filter Filter) (TextureID, error) {
	return currentUI.newTextureID(img, glFilter(filter))
}
