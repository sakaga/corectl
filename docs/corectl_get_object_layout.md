## corectl get object layout

Evaluates the hypercube layout of an generic object

### Synopsis

Evaluates the hypercube layout of an generic object. Example: corectl get object layout OBJECT-ID --app my-app.qvf

```
corectl get object layout <object-id> [flags]
```

### Options

```
  -h, --help   help for layout
```

### Options inherited from parent commands

```
  -a, --app string               App name, if no app is specified a session app is used instead.
  -c, --config string            path/to/config.yml where parameters can be set instead of on the command line
  -e, --engine string            URL to engine (default "localhost:9076")
      --headers stringToString   Headers to use when connecting to qix engine (default [])
      --ttl string               Engine session time to live in seconds (default "30")
  -v, --verbose                  Logs extra information
```

### SEE ALSO

* [corectl get object](corectl_get_object.md)	 - Shows content of an generic object

