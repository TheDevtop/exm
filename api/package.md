# Documentation: API

The API server listens on port 1800.

### Routes:

- /ping           (Check if EXM is online)
- /search         (Search for patterns in objects)
- /search/all     (Search for patterns for all objects)
- /replace        (Replace pattern with mapping, return entire object)
- /replace/all    (Replace for all objects)
- /reduce         (Reduce the object to its dictionary)
- /reduce/all     (Reduce for all objects)
- /map/reduce     (Replace pattern with mapping, return mapped only)
- /map/reduce/all (MapReduce for all objects)

### JSON Forms:

**ObjectForm**

Used by: /reduce
```json
{
  "Object": ""
}
```

**SearchForm**

Used by: /search, /search/all
```json
{
  "Object": "",
  "Regex": ""
}
```

**ReplaceForm**

Used by: /replace, /replace/all, /map/reduce, /map/reduce/all
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

**MultiResultForm**

Used by: /*/all
```json
{
  "Route": "",
  "Error": null,
  "Results": {
    "": [""]
  }
}
```
