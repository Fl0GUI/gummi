package throne

import (
	"context"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"j322.ica/gumroad-sammi/config"
)

type Sale map[string]interface{}

type ThroneClient struct {
	sales     chan Sale
	snapshots *firestore.QuerySnapshotIterator
	ctx       context.Context
	ctxCancel context.CancelFunc
}

var client ThroneClient

func Start(c *config.ThroneConfig) error {
	ctx, cnl := context.WithCancel(context.Background())
	it, err := makeQueryIterator(c.CreatorId, ctx)
	if err != nil {
		return err
	}
	client = ThroneClient{
		make(chan Sale),
		it,
		ctx,
		cnl,
	}
	go client.pipe()
	return nil
}

func Stop() {
	client.ctxCancel()
	client.snapshots.Stop()
	close(client.sales)
}

func GetSalesChan() chan Sale {
	return client.sales
}

func makeQueryIterator(creatorId string, ctx context.Context) (*firestore.QuerySnapshotIterator, error) {
	if len(creatorId) == 0 {
		return nil, CreatorIdEmpty
	}
	c, err := firestore.NewClient(
		ctx,
		"onlywish-9d17b",
		option.WithAPIKey("AIzaSyAfHFvON8isjIYc9j5UC_UuXve-ZAKuAUg"),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create throne client: %v", err)
	}

	collRef := c.Collection("overlays")
	if collRef == nil {
		return nil, errors.New("throne collection for overlays not found")
	}
	queryIt := collRef.Where("creatorId", "==", creatorId).OrderBy("createdAt", firestore.Desc).Limit(1).Snapshots(ctx)
	return queryIt, nil
}

func (t *ThroneClient) pipe() {
	for {
		snap, err := t.snapshots.Next()
		if err := t.ctx.Err(); err != nil {
			break
		}
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Printf("Throne update failed: %s\n", err)
			continue
		}
		docs, err := snap.Documents.GetAll()
		if err != nil {
			log.Printf("Throne update failed: %s\n", err)
			continue
		}
		if len(docs) != 1 {
			panic(fmt.Sprintf("Expected amount of documents from limit(1) to be 1, not %v", len(docs)))
		}
		t.sales <- docs[0].Data()
	}
}
