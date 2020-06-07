package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Save: Got item "+"#%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}
