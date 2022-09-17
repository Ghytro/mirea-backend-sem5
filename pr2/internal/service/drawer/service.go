package drawer

import (
	"context"
	"errors"
)

var ErrNoSuchShape = errors.New("Нет фигуры с заданным id")

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) DrawShape(ctx context.Context, shapeID uint32) (*BoundingBox, error) {
	return parseBoundingBox(shapeID)
}

func parseBoundingBox(shapeID uint32) (*BoundingBox, error) {
	var strShape string
	switch byte((shapeID >> 24) & uint32(0xFF)) {
	case 0:
		strShape = "circle"
	case 1:
		strShape = "rectangle"
	case 2:
		strShape = "triangle"
	}
	fillID := byte((shapeID >> 16) & uint32(0xFF))
	width := int((shapeID >> 8) & uint32(0xFF))
	height := int(shapeID & uint32(0xFF))
	return NewBoundingBox(strShape, width, height, fillID)
}
