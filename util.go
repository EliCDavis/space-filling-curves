package main

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/EliCDavis/mesh"
)

func save(mesh mesh.Model, name string) error {
	defer timeTrack(time.Now(), "Saving Model")
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	err = mesh.Save(w)
	if err != nil {
		return err
	}
	return w.Flush()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
