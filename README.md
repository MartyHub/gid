# gid

`gid` is a small lib to convert a string to a valid [Go identifier](https://go.dev/ref/spec#Identifiers).

* checking for [keywords](https://go.dev/ref/spec#Keywords)
  and [predeclared identifiers](https://go.dev/ref/spec#Predeclared_identifiers)
* checking for [initialisms](https://github.com/golang/go/wiki/CodeReviewComments#initialisms)

## TLDR;

```go
import "github.com/MartyHub/gid"

tok := gid.Default()

tok.ExportID("id") // ID 
tok.ExportID("my_id") // MyID 
tok.ExportID("json_id") // JSONId

tok.UnexportID("ID") // id 
tok.UnexportID("my_id") // myID 
tok.UnexportID("json_id") // jsonID
```
