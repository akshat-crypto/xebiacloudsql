package resources

import (
	"context"

	"github.com/akshat-crypto/cq-source-testcqplugin/client"
	"github.com/akshat-crypto/cq-source-testcqplugin/internal/xkcd"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"golang.org/x/sync/errgroup"
)

func Comics() *schema.Table {
	return &schema.Table{
		Name:      "xkcd_comics",
		Resolver:  fetchComics,
		Transform: transformers.TransformWithStruct(&xkcd.Comic{}),
	}
}

func fetchComics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	comic, err := c.XKCD.GetLatestComic(0)
	if err != nil {
		return err
	}
	res <- comic
	g := errgroup.Group{}
	g.SetLimit(10)
	for i := 1; i < comic.Num; i++ {
		i := i
		g.Go(func() error {
			comic, err := c.XKCD.GetComic(i)
			if err != nil {
				c.Logger.Error().Err(err).Msgf("failed to fetch comic %d", i)
				return err
			}
			res <- comic
			return nil
		})
	}
	return g.Wait()
}

// func SampleTable() *schema.Table {
// 	return &schema.Table{
// 		Name:     "testcqplugin_sample_table",
// 		Resolver: fetchSampleTable,
// 		Columns: []schema.Column{
// 			{
// 				Name: "column",
// 				Type: schema.TypeString,
// 			},
// 		},
// 	}
// }

// func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
// 	return fmt.Errorf("not implemented")
// }
