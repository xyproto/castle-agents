package main

import (
	"errors"
	"math"

	"github.com/fogleman/fauxgl"
	pp "github.com/xyproto/pixelpusher"
)

const pi2 = math.Pi * 2.0

var (
	x, y         = 160, 100
	meshHexColor = "#ffffff"
	cameraAngle  float32
	mesh         *fauxgl.Mesh
)

func onDraw(canvas *pp.Canvas) error {
	if mesh == nil {
		m, err := LoadMeshOBJ("bevelcube/bevelcube.obj")
		if err != nil {
			return err
		}
		mesh = m
	}
	return DrawMesh(canvas, mesh, cameraAngle, meshHexColor)
}

func onPress(left, right, up, down, space, enter, esc bool) error {
	if left || down {
		cameraAngle += 0.1
		if cameraAngle > pi2 {
			cameraAngle -= pi2
		}
	} else if right || up {
		cameraAngle -= 0.1
		if cameraAngle < 0 {
			cameraAngle += pi2
		}
	}
	if space {
		if meshHexColor != "#ff0000" {
			meshHexColor = "#ff0000"
		} else {
			meshHexColor = "#ffffff"
		}
	}
	if esc {
		return errors.New("quit")
	}
	return nil
}

func main() {
	pp.New("Rotate Cube").Run(onDraw, onPress, nil, nil)
}
