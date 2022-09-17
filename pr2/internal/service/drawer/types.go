package drawer

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type Shape interface {
	HTMLString() string
}

type BaseShape struct {
	Fill byte
}

type BoundingBox struct {
	Width, Height int
	Shape         Shape
}

func NewBoundingBox(shape string, width, height int, shapeFill byte) (*BoundingBox, error) {
	result := &BoundingBox{
		Width:  width,
		Height: height,
	}
	baseShape := BaseShape{Fill: shapeFill}
	switch shape {
	case "circle":
		result.Shape = &Circle{
			BaseShape: baseShape,
			Center: Point{
				X: width / 2,
				Y: height / 2,
			},
			Radius: int(math.Min(float64(width), float64(height))) / 2,
		}
	case "rectangle":
		result.Shape = &Rectangle{
			BaseShape: baseShape,
			Width:     width,
			Height:    height,
		}
	case "triangle":
		result.Shape = &Triangle{
			BaseShape: baseShape,
			Points: [3]Point{
				{X: 0, Y: height},
				{X: width / 2, Y: 0},
				{X: width, Y: height},
			},
		}
	default:
		return nil, ErrNoSuchShape
	}
	return result, nil
}

func (b *BoundingBox) HTMLString() string {
	return fmt.Sprintf(
		`<svg width="%d" height="%d">%s</svg>`,
		b.Width,
		b.Height,
		b.Shape.HTMLString(),
	)
}

type Circle struct {
	BaseShape
	Center Point
	Radius int
}

func (c *Circle) HTMLString() string {
	return fmt.Sprintf(`
		<circle cx="%d" cy="%d" r="%d" stroke="black" stroke-width="4" fill=%q />`,
		c.Center.X,
		c.Center.Y,
		c.Radius,
		fillToString(c.Fill),
	)
}

type Rectangle struct {
	BaseShape
	Width, Height int
}

func (r *Rectangle) HTMLString() string {
	return fmt.Sprintf(`
		<rect width="%d" height="%d" style="fill:%s;stroke-width:10;stroke:rgb(0,0,0)" />`,
		r.Width,
		r.Height,
		fillToString(r.Fill),
	)
}

type Triangle struct {
	BaseShape
	Points [3]Point
}

func (r *Triangle) HTMLString() string {
	return fmt.Sprintf(`
	<polygon points="%d,%d %d,%d %d,%d"
	style="fill:%s;stroke:black;stroke-width:5;fill-rule:evenodd;" />`,
		r.Points[0].X, r.Points[0].Y,
		r.Points[1].X, r.Points[1].Y,
		r.Points[2].X, r.Points[2].Y,
		fillToString(r.Fill),
	)
}

func fillToString(fill byte) string {
	switch fill {
	case 0:
		return "red"
	case 1:
		return "green"
	case 2:
		return "blue"
	}
	return "white"
}
