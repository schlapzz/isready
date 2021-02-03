isready
===
isready is a powerful tool to check if a service is ready within a single comman
d. This tool support a various collection of service to test with

```
Usage:
  is-ready [command]

Available Commands:
  curl        checks if a service available with a http request
  deployment  Wait until kubernetes deployment is ready
  help        Help about any command
  kafka       A brief description of your command
  mongo       A brief description of your command
  mysql       A brief description of your command
  psql        Check if postgresql database is ready
  version     Show version of 'isready'

Flags:
  -h, --help             help for is-ready
      --retries int32    number of retries before abort (default 3)
      --timeout string   timeout for connection (default "30s")

Use "is-ready [command] --help" for more information about a command.
```