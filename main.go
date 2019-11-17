package main

func main() {

	curve, err := hilberCurve3D(1)
	if err != nil {
		panic(err)
	}

	err = save(curve, "out.obj")

	if err != nil {
		panic(err)
	}

}
