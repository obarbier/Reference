## Definition
go build just compiles the executable file and moves it to the destination. go install does a little bit more. It moves the executable file to $GOPATH/bin and caches all non-main packages which are imported to $GOPATH/pkg. The cache will be used during the next compilation provided the source did not change yet.
## Glossary
* 

## Reference
* [Stackoveflow](https://stackoverflow.com/questions/24069664/what-does-go-install-do)
* [github - NanXiao](https://github.com/NanXiao/golang-101-hacks/blob/master/posts/go-build-vs-go-install.md)