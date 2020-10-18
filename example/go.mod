module function

go 1.15

replace function/handler => ./handler

replace github.com/alexellis/go-modules-test/github-go-tester/pkg => ./handler/pkg

require (
	github.com/alexellis/go-modules-test/github-go-tester/pkg v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.1
)
