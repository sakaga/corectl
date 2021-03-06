## corectl build

Reloads and saves the app after updating connections, dimensions, measures, objects and the script

### Synopsis

Builds the app. Example: corectl build --connections ./myconnections.yml --script ./myscript.qvs
			


```
corectl build [flags]
```

### Options

```
  -a, --app string               App name, if no app is specified a session app is used instead.
  -c, --config string            path/to/config.yml where parameters can be set instead of on the command line
      --connections string       path/to/connections.yml that contains connections that are used in the reload. Note that when specifying connections in the config file they are specified inline, not as a file reference!
      --dimensions string        A list of generic dimension json paths
  -e, --engine string            URL to engine (default "localhost:9076")
      --headers stringToString   Headers to use when connecting to qix engine (default [])
  -h, --help                     help for build
      --measures string          A list of generic measures json paths
      --objects string           A list of generic object json paths
      --script string            path/to/reload-script.qvs that contains a qlik reload script. If omitted the last specified reload script for the current app is reloaded
      --silent                   Do not log reload progress
      --ttl string               Engine session time to live in seconds (default "30")
```

### Options inherited from parent commands

```
  -v, --verbose   Logs extra information
```

### SEE ALSO

* [corectl](corectl.md)	 - 

