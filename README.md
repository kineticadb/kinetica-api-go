<h3 align="center" style="margin:0px">
 <img width="200" src="https://www.kinetica.com/wp-content/uploads/2018/08/kinetica_logo.svg" alt="Kinetica Logo"/>
</h3>
<h5 align="center" style="margin:0px">
	<a href="https://www.kinetica.com/">Website</a>
	|
	<a href="https://docs.kinetica.com/7.2/">Docs</a>
	|
	<a href="https://join.slack.com/t/kinetica-community/shared_invite/zt-1bt9x3mvr-uMKrXlSDXfy3oU~sKi84qg">Community Slack</a>   
</h5>


# Kinetica GOLANG API

- [Overview](#overview)
- [Support](#support)
- [Contact Us](#contact-us)


## Overview

This project contains the source code of the Golang Kinetica API.

The documentation can be found at <https://docs.kinetica.com/7.2/>.

### Usage

For using this API in a `GO` project add the lines

```GO
require (
 github.com/kineticadb/gpudb-api-go v0.0.1
)
```

to the `go.mod` file of your project.

For changes to the client-side API, please refer to
[CHANGELOG.md](CHANGELOG.md).

### Starter example

```GO

package main

import (
 "context"
 "fmt"
 "os"
 "sync"
 "time"

 "github.com/kineticadb/gpudb-api-go/example"
 "github.com/kineticadb/gpudb-api-go/kinetica"
 "go.uber.org/multierr"
 "go.uber.org/zap"
)

func main() {
 endpoint := os.Args[1]
 username := os.Args[2]
 password := os.Args[3]

 ctx := context.TODO()
 options := kinetica.KineticaOptions{Username: username, Password: password}
 // fmt.Println("Options", options)
 dbInst := kinetica.NewWithOptions(ctx, endpoint, &options)
}

```

### Logging

The logging is done using Uber zap package and `lumberjack` for rotating files based on size. Time based rotation is not supported yet. The configuration for `lumberjack` can be found here - <https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2>

#### Default Log Config file

This is included and will be used in case a user defined config file is not found. The name of the file is `log_config.yaml`.

```yaml
level: 'info'
development: true
disableCaller: false
disableStacktrace: false
encoding: 'console'
encoderConfig:
  messageKey: 'msg'
  levelKey: 'level'
  timeKey: 'ts'
  nameKey: 'logger'
  callerKey: 'caller'
  functionKey: 'function'
  stacktraceKey: 'stacktrace'
  skipLineEnding: false
  lineEnding: "\n"
  levelEncoder: 'capital'
  timeEncoder: 'iso8601'
  durationEncoder: 'string'
  callerEncoder: 'full'
  nameEncoder: 'full'
  consoleSeparator: ' | '
outputPaths:
  # Implements logging to the console
  - 'stdout'
  # Implements rolling logs using lumberjack logger; config parameters are supplied as
  # query params. Here maxSize is 10MB after which the logger rolls over; maximum
  # number of backups (maxBackups) kept is 5 and maxAge is 10 days.
  # The name of the log file in this case is "logs/gpudb-api.log" where the
  # "logs" directory is under the current directory on the local machine.
  - 'lumberjack://localhost/logs/gpudb-api.log?maxSize=10&maxBackups=5&maxAge=10'
errorOutputPaths:
  - 'stderr'
  - './logs/error_logs'
initialFields:
  app: 'gpudb-api'
```


## Support

For bugs, please submit an
[issue on Github](https://github.com/kineticadb/kinetica-api-go/issues).

For support, you can post on
[stackoverflow](https://stackoverflow.com/questions/tagged/kinetica) under the
``kinetica`` tag or
[Slack](https://join.slack.com/t/kinetica-community/shared_invite/zt-1bt9x3mvr-uMKrXlSDXfy3oU~sKi84qg).

## Contact Us

- Ask a question on Slack:
  [Slack](https://join.slack.com/t/kinetica-community/shared_invite/zt-1bt9x3mvr-uMKrXlSDXfy3oU~sKi84qg)
- Follow on GitHub:
  [Follow @kineticadb](https://github.com/kineticadb)
- Email us:  <support@kinetica.com>
- Visit:  <https://www.kinetica.com/contact/>
