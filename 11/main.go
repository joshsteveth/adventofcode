package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	serialNumber = 5535
	numWorker    = 10
)

type entry struct {
	x, y, val, size int
}

var grid [300][300]int

func main() {

	now := time.Now()
	defer func(t time.Time) {
		fmt.Printf("total runtime: %v\n", time.Since(t))
	}(now)

	grid = [300][300]int{}

	// calculate the value of each grid
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			grid[x][y] = calcVal(serialNumber, x+1, y+1)
		}
	}

	wg := sync.WaitGroup{}
	ch := make(chan *entry)
	workers := make(chan struct{}, numWorker)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var maxX, maxY, maxVal, maxSize int

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, ch <-chan *entry) {

		for i := 0; i < 300; i++ {

			select {
			case <-ctx.Done():
				wg.Done()
				return
			case entry := <-ch:
				fmt.Printf("got new entry with size: %d\n", entry.size)
				if entry.val > maxVal {

					fmt.Printf("new leading entry!  %+v\n", *entry)

					maxX = entry.x
					maxY = entry.y
					maxSize = entry.size
					maxVal = entry.val
				}
			}
		}

		wg.Done()

	}(ctx, &wg, ch)

	for i := 0; i < numWorker; i++ {
		workers <- struct{}{}
	}

	for size := 1; size <= 300; size++ {
		<-workers
		go func(size int, ch chan<- *entry, workers chan struct{}) {

			var maxx, maxy, maxval int

			for x := 0; x < 301-size; x++ {
				for y := 0; y < 301-size; y++ {
					X, Y := x+1, y+1

					v := calcGridVal(grid, X, Y, size)
					if v > maxval {
						maxx, maxy, maxval = X, Y, v
					}
				}
			}

			ch <- &entry{maxx, maxy, maxval, size}

			workers <- struct{}{}

		}(size, ch, workers)

	}

	wg.Wait()
	close(workers)

	// result is: (237,284,11) with maximum value: 91
	fmt.Printf("result is: (%d,%d,%d) with maximum value: %d\n", maxX, maxY, maxSize, maxVal)
}

func copyGrid(grid [300][300]int) (res [300][300]int) {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			res[i][j] = grid[i][j]
		}
	}
	return
}

func calcVal(serialNumber, x, y int) int {

	rackID := x + 10
	powerLevel := rackID * y
	increasedPowerLevel := powerLevel + serialNumber

	mult := rackID * increasedPowerLevel

	return (mult%1000)/100 - 5
}

func calcGridVal(grid [300][300]int, x, y, squareSize int) (res int) {

	for xi := x; xi < x+squareSize; xi++ {
		for yi := y; yi < y+squareSize; yi++ {
			res += grid[xi-1][yi-1]
		}
	}

	return
}
