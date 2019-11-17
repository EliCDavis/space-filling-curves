package main

import (
	"math"
	"time"

	"github.com/EliCDavis/mesh"
	"github.com/EliCDavis/meshedpotatoes/path"
	"github.com/EliCDavis/vector"
)

func dragonCurveRecursive(currentCurve path.Path, iterationsLeft int) path.Path {

	if iterationsLeft <= 0 {
		return currentCurve
	}

	rotated := currentCurve.
		Rotate(currentCurve[0], mesh.UnitQuaternionFromTheta(-math.Pi/2, vector.Vector3Up())).
		Reverse()

	moved := rotated.Translate(currentCurve[0].Sub(rotated[len(rotated)-1]))

	return dragonCurveRecursive(append(moved, currentCurve[1:]...), iterationsLeft-1)
}

func dragonCurve(iterationsLeft int) (mesh.Model, error) {
	defer timeTrack(time.Now(), "Building Dragon Curve")
	return dragonCurveRecursive([]vector.Vector3{
		vector.Vector3Right(),
		vector.NewVector3(.5, 0, .1),
		vector.NewVector3(.1, 0, .5),
		vector.Vector3Forward(),
	}, iterationsLeft).CreatePipe(0.1, 12)
}
