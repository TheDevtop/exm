# Documentation: API

The API server listens on port 1800.

### Routes:

- /ping         (Check if EXM is online)
- /search       (Search for patterns in objects)
- /replace      (Replace pattern with mapping, return entire object)
- /map/reduce   (Replace pattern with mapping, return mapped only)

### JSON Forms:

**SearchForm**

Used by: /search
```json
{
  "Object": "",
  "Regex": ""
}
```

**ReplaceForm**

Used by: /replace, /map/reduce
```json
{
  "Object": "",
  "Regex": "",
  "Mapping" : ""
}
```

**ResultForm**

Used by everyone
```json
{
  "Route": "",
  "Error": "",
  "Result": [
    ""
  ]
}
```
