## Definition

A golang pattern that i have been seeing in many project. this allow for easy configuration of complex object

```
// CREATE A config Struct
type GrpcServerConfig struct {
	Port int
}

// (optional) but create a configure function
type Config func(*GrpcServerConfig)

// create initializer function with this pattern

func NewServer(configs ...Config) *GrpcServer {
	g:= &GrpcServerConfig{}
	for _ ,c := range configs{
		c(g)
	}
	
	// .... omitted

}
```


## Glossary

## Reference
* [Golang - Cafe(https://golang.cafe/blog/golang-functional-options-pattern.html)