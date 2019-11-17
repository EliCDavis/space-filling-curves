package main

import (
	"math"
	"time"

	"github.com/EliCDavis/mesh"
	"github.com/EliCDavis/meshedpotatoes/path"
	"github.com/EliCDavis/vector"
)

func hilbertCurveRecursive3D(currentCurve path.Path, iterationsLeft int) path.Path {

	if iterationsLeft <= 0 {
		return currentCurve
	}

	bottomLeft := currentCurve.
		Rotate(currentCurve[0], mesh.UnitQuaternionFromTheta(math.Pi/2, vector.Vector3Up())).
		Reverse()

	topLeft := currentCurve.
		Translate(bottomLeft[len(bottomLeft)-1].Sub(currentCurve[0]).Add(vector.Vector3Forward()))

	topRight := currentCurve.
		Translate(topLeft[len(topLeft)-1].Sub(currentCurve[0]).Add(vector.Vector3Right()))

	rot := currentCurve.
		Reverse().
		Rotate(currentCurve[0], mesh.UnitQuaternionFromTheta(-math.Pi/2, vector.Vector3Up()))

	bottomRight := rot.Translate(topRight[len(topRight)-1].Sub(vector.Vector3Forward()).Sub(rot[0]))

	return hilbertCurveRecursive3D(bottomLeft.Combine(topLeft).Combine(topRight).Combine(bottomRight), iterationsLeft-1)
}

func hilberCurve3D(iterationsLeft int) (mesh.Model, error) {
	defer timeTrack(time.Now(), "Building 3D Hilbert Curve")
	path := hilbertCurveRecursive3D([]vector.Vector3{
		vector.Vector3Zero(),
		vector.Vector3Forward(),
		vector.Vector3Forward().Add(vector.Vector3Up()),
		vector.Vector3Forward().Add(vector.Vector3Up()).Add(vector.Vector3Right()),
		vector.Vector3Forward().Add(vector.Vector3Right()),
		vector.Vector3Right(),
	}, iterationsLeft)
	return path.CreatePipe(0.1, 12)
}
