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
	"math"
)

// GeometryMatrixDim is a dimension of a GeometryMatrix.
const GeometryMatrixDim = 3

// A GeometryMatrix represents a matrix to transform geometry when rendering a texture or a render target.
type GeometryMatrix struct {
	Elements [GeometryMatrixDim - 1][GeometryMatrixDim]float64
}

// GeometryMatrixI returns an identity geometry matrix.
func GeometryMatrixI() GeometryMatrix {
	return GeometryMatrix{
		[GeometryMatrixDim - 1][GeometryMatrixDim]float64{
			{1, 0, 0},
			{0, 1, 0},
		},
	}
}

func (g *GeometryMatrix) dim() int {
	return GeometryMatrixDim
}

// Element returns a value of a matrix at (i, j).
func (g *GeometryMatrix) Element(i, j int) float64 {
	return g.Elements[i][j]
}

// Concat multiplies a geometry matrix with the other geometry matrix.
func (g *GeometryMatrix) Concat(other GeometryMatrix) {
	result := GeometryMatrix{}
	mul(&other, g, &result)
	*g = result
}

// IsIdentity returns a boolean indicating whether the geometry matrix is an identity.
func (g *GeometryMatrix) IsIdentity() bool {
	return isIdentity(g)
}

func (g *GeometryMatrix) setElement(i, j int, element float64) {
	g.Elements[i][j] = element
}

// Translate translates the geometry matrix by (tx, ty).
func (g *GeometryMatrix) Translate(tx, ty float64) {
	g.Elements[0][2] += tx
	g.Elements[1][2] += ty
}

// Scale scales the geometry matrix by (x, y).
func (g *GeometryMatrix) Scale(x, y float64) {
	g.Elements[0][0] *= x
	g.Elements[0][1] *= x
	g.Elements[0][2] *= x
	g.Elements[1][0] *= y
	g.Elements[1][1] *= y
	g.Elements[1][2] *= y
}

// Rotate rotates the geometry matrix by theta.
func (g *GeometryMatrix) Rotate(theta float64) {
	sin, cos := math.Sincos(theta)
	rotate := GeometryMatrix{
		[2][3]float64{
			{cos, -sin, 0},
			{sin, cos, 0},
		},
	}
	g.Concat(rotate)
}
