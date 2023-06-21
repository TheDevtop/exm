# API Documentation

### API Table

| URL             | Function                          |
|-----------------|-----------------------------------|
| /search/object  | Search for matches in object      |
| /search/global  | Search for matches in all objects |
| /index/object   | Return the dictionary of object   |
| /index/global   | Return the list of known objects  |
| /meta/object    | Return the metadata of objects    |
| /ping           | Check if we are online            |

### API Forms

Request form:
```json
{
    "Object" : "",
    "Regex" : ""
}
```
Result form:
```json
{
    "Error" : "",
    "Count" : 0,
    "Results" : [""]
}
```
Metadata form:
```json
{
    "Error": "",
    "Object": "",
    "Type": "",
    "Size": 0,
    "Source": "",
    "LastModified": ""
}
```
