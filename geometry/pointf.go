// Copyright 2021 Tacigar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package geometry

type PointF struct {
	X float64
	Y float64
}

func (p *PointF) Add(other PointF) PointF {
	return PointF{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p *PointF) Sub(other PointF) PointF {
	return PointF{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p *PointF) Mul(factor float64) PointF {
	return PointF{
		X: p.X * factor,
		Y: p.Y * factor,
	}
}

func (p *PointF) Div(divisor float64) PointF {
	return PointF{
		X: p.X / divisor,
		Y: p.Y / divisor,
	}
}

func (p *PointF) AddAssign(other PointF) *PointF {
	p.X += other.X
	p.Y += other.Y
	return p
}

func (p *PointF) SubAssign(other PointF) *PointF {
	p.X -= other.X
	p.Y -= other.Y
	return p
}

func (p *PointF) MulAssign(factor float64) *PointF {
	point := p.Mul(factor)
	p.X = point.X
	p.Y = point.Y
	return p
}

func (p *PointF) DivAssign(divisor float64) *PointF {
	point := p.Div(divisor)
	p.X = point.X
	p.Y = point.Y
	return p
}
