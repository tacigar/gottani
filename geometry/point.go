package geometry

type Point struct {
	X int
	Y int
}

func (p *Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p *Point) Sub(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p *Point) Mul(factor float64) Point {
	return Point{
		X: int(float64(p.X) * factor),
		Y: int(float64(p.Y) * factor),
	}
}

func (p *Point) Div(divisor float64) Point {
	return Point{
		X: int(float64(p.X) / divisor),
		Y: int(float64(p.Y) / divisor),
	}
}

func (p *Point) AddAssign(other Point) *Point {
	p.X += other.X
	p.Y += other.Y
	return p
}

func (p *Point) SubAssign(other Point) *Point {
	p.X -= other.X
	p.Y -= other.Y
	return p
}

func (p *Point) MulAssign(factor float64) *Point {
	point := p.Mul(factor)
	p.X = point.X
	p.Y = point.Y
	return p
}

func (p *Point) DivAssign(divisor float64) *Point {
	point := p.Div(divisor)
	p.X = point.X
	p.Y = point.Y
	return p
}
