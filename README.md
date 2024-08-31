# go-monolith-arch-template
Template Monolith App On Golang (ToDo)

## running test commands
### running unit tests
- run all unit tests
```
    go test ./...
```
- run unit test by name
```
    # move test dir and run
    go test -run <name>
```
- run same name tests
```
    go test -run CreateTodo
```
### running bechmark tests
- run all benchmark tests
```
    go test -bench=.
```
- run all benchmark with poiners
``````
    go test -bench=. -benchmem
``````
