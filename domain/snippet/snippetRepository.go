package snp

type SnippetRepository interface {
	AddSnippet(snippet Snippet) (int64, error)
	LatestSnippets() (Snippets, error)
	GetSnippet(id int64) (*Snippet, error)
	GetUserSnippet(id int64) (*UserSnippet, error)
}
