package config

var (
	// Filepath holds the mapping file path
	Filepath = "github.com/saltside/gunjan01/source/search/books.json"

	// IndexName is the name of the book index.
	IndexName = "literary_books"

	// Port is the server port.
	Port = ":8080"

	// ElasticURL is the URL on which the elastic cluster is running.
	ElasticURL = "http://localhost:9200"
)
