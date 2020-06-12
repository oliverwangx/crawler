package persist

import (
	"context"
	"crawler/crawler/engine"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Save: Got item "+"#%d: %v", itemCount, item)
			itemCount++

			err := Save(item)
			if err != nil {
				log.Printf("Item Saver: error "+"saving item %v: %v", item, err)
			}

		}
	}()
	return out
}

func Save(
	item engine.Item) error {

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	if item.Type == "" {
		return "", errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
