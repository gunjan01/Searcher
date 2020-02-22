package search

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	elastic "gopkg.in/olivere/elastic.v7"
)

// Es is a wrapper to elasticsearch client.
type Es struct {
	Client *elastic.Client
	ctx    context.Context
}

// NewES returns an elastic client of type Es.
func NewES() (*Es, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		return nil, err
	}

	return &Es{
		Client: client,
		ctx:    context.Background(),
	}, err
}

// CreateNewIndex creates a new index.
func (c *Es) CreateNewIndex(indexName string, mappings string) error {
	createIndex, err := c.Client.CreateIndex(indexName).BodyString(mappings).Do(c.ctx)
	if err != nil {
		logrus.WithError(err).Error("Unable to create index")
		return err
	}

	if !createIndex.Acknowledged {
		err = fmt.Errorf("Index creation - %s not acknowledged", indexName)
		logrus.WithError(err).Error("Created index not acknowledged")
		return err
	}

	return nil
}

// EnsureIndex ensures an index is present. If not, it creates it.
func (c *Es) EnsureIndex(indexName string, mappings string) error {
	exists, err := c.Client.IndexExists(indexName).Do(c.ctx)
	if err != nil {
		return err
	}
	if exists {
		logrus.Info("Index already exists")
		return nil
	}

	return c.CreateNewIndex(indexName, mappings)
}
