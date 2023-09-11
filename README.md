# logger

this is a logger that extends slog with a `logger` name and add stack traces on error

## example

```go
l := log.NewLogger(&log.HandlerOptions{Name: "wire-controller", AddSource: false})
l = l.WithGroup("group").With("k11", "v11", "k12", "v12")
l.Info("test",
		"k1", "v1")

l.Info("")
```

output

```json
{
  "time": "2023-09-11T04:38:19.23852-05:00",
  "level": "INFO",
  "message": "test",
  "logger": "wire-controller",
  "data": {
    "group": {
      "k11": "v11",
      "k12": "v12",
      "k1": "v1"
    }
  }
}
{
  "time": "2023-09-11T04:38:19.238526-05:00",
  "level": "INFO",
  "message": "",
  "logger": "wire-controller",
  "data": {
    "group": {
      "k11": "v11",
      "k12": "v12"
    }
  }
}
```