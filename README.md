# EXM

The austere search engine; in other terms: Grep with a Web API.

### EXM Server

Fetch data from known (networked) sources.

### EXM API

Exposes an API to search trough the data.

### Configuration

```json
{
  "Directory": "/foo/bar",
  "ServerAddress": ":1800",
  "CacheSize" : 16,
  "Sources": [
    {
      "Type": "http/tcp/udp/local",
      "Address": "",
      "Timeout": 0,
      "Object": "lorem.html"
    }
  ]
}
```
