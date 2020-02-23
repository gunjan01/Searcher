package search

import (
	"fmt"

	"github.com/gunjan01/searcher/source/config"
	elastic "gopkg.in/olivere/elastic.v7"
)

// Searcher will parse the ES result and construct
// the response.
type Searcher struct {
	request GetLiteraryBooksRequest
}

// NewSearcher returns a searcher
func newSearcher(request GetLiteraryBooksRequest) Searcher {
	return Searcher{
		request: request,
	}
}

// buildLocationQuery finds a book based on its location.
func (p *Searcher) buildLocationQuery() elastic.Query {
	if p.request.Location != nil {
		return elastic.NewTermQuery("location", p.request.Location)
	}
	return nil
}

// buildAuthorQuery finds a book based on the author.
func (p *Searcher) buildAuthorQuery() elastic.Query {
	if p.request.Author != nil {
		author := string(*p.request.Author)
		query := elastic.NewMatchQuery("author", author).
			Operator("AND").
			Boost(3)

		if len(author) > 3 {
			query = query.Fuzziness("AUTO").
				MaxExpansions(20).
				PrefixLength(2)
		}
		return query
	}
	return nil
}

// buildTitleQuery finds a book based on the title.
func (p *Searcher) buildTitleQuery() elastic.Query {
	if p.request.Title != nil {
		query := elastic.NewMatchQuery("title", p.request.Title).
			Operator("AND").
			Boost(3)

		if len(p.request.Query) > 3 {
			query = query.Fuzziness("AUTO").
				MaxExpansions(20).
				PrefixLength(2)
		}
		return query
	}
	return nil
}

// buildTextQuery finds a books based on its contents.
func (p *Searcher) buildTextQuery() elastic.Query {
	query := elastic.NewMatchQuery("text", p.request.Query).
		Operator("AND").
		Boost(3)

	if len(p.request.Query) > 3 {
		query = query.Fuzziness("AUTO").
			MaxExpansions(20).
			PrefixLength(2)
	}

	return query
}

// SearchQuery builds the elastic search query.
func (c *Es) SearchQuery(request GetLiteraryBooksRequest) (*elastic.SearchSource, error) {
	queries := []elastic.Query{}

	searchSource := elastic.NewSearchSource()
	searcher := newSearcher(request)
	queries = append(queries, searcher.buildAuthorQuery())
	queries = append(queries, searcher.buildTitleQuery())
	queries = append(queries, searcher.buildLocationQuery())

	query := elastic.NewBoolQuery()
	query.Filter(queries...)

	searchSource.Query(query)

	return searchSource, nil
}

// ExtractBooks extracts the ES response and returns the results.
func (c *Es) ExtractBooks(searchSource *elastic.SearchSource) ([]elastic.SearchHits, error) {
	source, err := searchSource.Source()
	if err != nil {
		return nil, err
	}

	search := c.Client.Search(config.IndexName).Source(source)

	result, err := search.Do(c.ctx)
	if err != nil {
		return nil, err
	}

	books := result.Hits.Hits
	fmt.Println(books)

	return nil, nil
}
