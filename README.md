# json vanity

pretty-print json from the command line

## usage

`curl http://example.com/results.json | json-vanity` outputs the following:

```json
{
 "objs": [
  {
   "id": "aaa"
  },
  {
   "id": "bbb"
  }
 ]
}
```

## developing

  1. `go get github.com/defeated/json-vanity`
  2. `go test -cover`
