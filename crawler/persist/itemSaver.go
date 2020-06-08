package persist

import (
	"log"
)

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

//func Save(item interface{})(id string, err error){
//	client, err :=elastic.NewClient(
//		// Must turn off sniff in docker
//		elastic.SetSniff(false))
//	if err != nil {
//		return "", err
//	}
//
//	resp, err := client.Index().
//		Index("dating_profile").
//		Type("zhenai").
//		BodyJson(item).
//		Do(context.Background())
//
//	if err != nil {
//		return "", err
//	}
//	return resp.Id, nil
//
//}
