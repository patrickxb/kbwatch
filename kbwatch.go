package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Status struct {
	Staged  bool
	Journal struct {
		UnflushedBytes int
		UnflushedPaths []string
	}
}

func main() {
	flag.Parse()
	tlf := flag.Arg(0)
	filename := filepath.Join(tlf, ".kbfs_status")
	for {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		dec := json.NewDecoder(f)
		var status Status
		if err := dec.Decode(&status); err != nil {
			log.Fatal(err)
		}
		f.Close()
		log.Printf("staged: %v", status.Staged)
		log.Printf("unflushed bytes: %d (%d MB)", status.Journal.UnflushedBytes, status.Journal.UnflushedBytes/(1024*1024))
		log.Printf("unflushed paths: %d", len(status.Journal.UnflushedPaths))
		time.Sleep(5 * time.Second)
	}

}
