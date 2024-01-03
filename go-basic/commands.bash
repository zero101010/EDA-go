go env
go help fmt
go install 
go help doc
go test
got test -v 
go test -bench Soma
go test -cover
go test -coverprofile c.out && go tool cover -html=c.out
gofmt -d filename
golint filename
govet
