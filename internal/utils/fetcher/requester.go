package fetcher

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type StateCode int

const (
	READY   = 1
	RUNNING = 2
)

type Work struct {
	FetcherHandler Fetcher // multiple
	State          StateCode
	FileName       string // TODO remove?
	Length         int64
	FileHandler    *os.File
}

func (w *Work) parse(cmd Command) {
	length, err := probe(cmd.URL)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.FileName = parseFileName(cmd.URL)
	log.Println("Downloading file", cmd.URL, w.FileName)

	file, err := os.Create(w.FileName)
	if err != nil {
		log.Fatal(err)
	}
	// defer file.Close()
	w.FileHandler = file

	w.FetcherHandler = Fetcher{
		URL: cmd.URL,
	}

	if length > 0 {
		// break-point downloading
		w.Length = length
		var pieces []RangeHeader
		rangeSize := splitSize(int64(w.Length))
		amount := int(int64(w.Length) / rangeSize)
		if int64(w.Length)%rangeSize != 0 {
			amount = amount + 1
		}

		for i := 0; i < amount; i++ {
			if i == amount-1 {
				pieces = append(pieces, RangeHeader{
					StartPos: int64(i) * RangeSize,
					EndPos:   int64(w.Length - 1),
				})
			} else {
				pieces = append(pieces, RangeHeader{
					StartPos: int64(i) * RangeSize,
					EndPos:   int64(i)*RangeSize + RangeSize - 1,
				})
			}
		}
		w.FetcherHandler.Pieces = pieces
	}

	w.State = READY
}

func (w *Work) run() {
	w.State = RUNNING
	if w.Length == 0 {
		// download directly
		if w.Length == 0 {
			_, err := w.FetcherHandler.retrieveAll(w.FileHandler)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	// support for range request
	amount := len(w.FetcherHandler.Pieces)
	var wg sync.WaitGroup
	for i := 0; i < amount; i++ {
		wg.Add(1)
		log.Println("downloading", i)
		go func(pieceN int) {
			defer wg.Done()
			_, err := w.FetcherHandler.retrievePartial(pieceN, w.FileHandler)
			if err != nil {
				log.Fatal("Error in downloading piece ", pieceN, err)
			}
		}(i)
	}

	wg.Wait()
}

