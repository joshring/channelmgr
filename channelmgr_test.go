package channelmgr_test

import (
	"log"
	"testing"
	"time"

	"github.com/joshring/channelmgr"
)

func TestWithDeadline(t *testing.T) {

	mgr := channelmgr.NewManager()

	mgr.AddTask(
		func(args []any) {
			log.Printf("Sleeping for 1sec\n")
			time.Sleep(1 * time.Second)
			log.Printf("finished 1sec sleep\n")
		},
		[]any{},
	)

	mgr.AddTask(
		func(args []any) {
			log.Printf("Sleeping for 2sec\n")
			time.Sleep(2 * time.Second)
			log.Printf("finished 2sec sleep\n")
		},
		[]any{},
	)

	timeout := 3 * time.Second
	log.Printf("waiting for jobs to complete for %v\n", timeout)
	timeStart := time.Now()
	timeTaken := time.Hour

	ok := mgr.WaitWithDeadline(timeout)
	if ok {
		timeTaken = time.Since(timeStart)
		log.Printf("Complete in: %vsec", time.Since(timeStart))
	} else {
		t.Fatal("unable to wait for deadline correctly")
	}

	log.Printf("Complete in: %vsec", timeTaken)

	if timeTaken.Truncate(time.Second) != 2*time.Second {
		t.Fatal("jobs did not complete within expected timeframe")
	}

}

func TestWithoutDeadline(t *testing.T) {

	mgr := channelmgr.NewManager()

	mgr.AddTask(
		func(args []any) {
			log.Printf("Sleeping for 1sec\n")
			time.Sleep(1 * time.Second)
			log.Printf("finished 1sec sleep\n")
		},
		[]any{},
	)

	mgr.AddTask(
		func(args []any) {
			log.Printf("Sleeping for 2sec\n")
			time.Sleep(2 * time.Second)
			log.Printf("finished 2sec sleep\n")
		},
		[]any{},
	)

	log.Printf("waiting for jobs to complete without timeout\n")
	timeStart := time.Now()

	mgr.WaitWithoutDeadline()

	timeTaken := time.Since(timeStart)
	log.Printf("Complete in: %vsec", timeTaken)

	if timeTaken.Truncate(time.Second) != 2*time.Second {
		t.Fatal("jobs did not complete within expected timeframe")
	}

}
