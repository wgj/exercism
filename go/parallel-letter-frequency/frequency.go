package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	var wg sync.WaitGroup
	masterMap := FreqMap{}
	for _, v := range s {
		wg.Add(1)
		// Store `v` before dropping into the goroutine.
		str := v
		// This feels pretty dirty. It would have been nicer to rewrite Frequency() with a channel, and make ConcurrentFrequency() keep one map, instead of stitching together three.
		go func() {
			m := Frequency(str)
			// Combine goroutine's map into masterMap.
			for k, v := range m {
				masterMap[k] += v
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return masterMap
}
