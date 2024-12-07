module 4.6

go 1.22.2

require (
	github_searcher v0.0.0
	github_cruder v0.0.0

)

replace (
	github_searcher => ../4.5_JSON/Ex4.10_GitHub_Issue_Searcher
	github_cruder => ../4.5_JSON/Ex4.11_Github_Issues_CRUD
)
