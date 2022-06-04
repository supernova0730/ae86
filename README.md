# AE86

> Telegram bot for delivery service

* Compile

```shell
make build
```

#### Global help
```shell
NAME:
   ae86 - delivery service application

USAGE:
   ae86 [global options] command [command options] [arguments...]

COMMANDS:
   config   generate initial config
   start    start server
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false) 
```

#### Initial config generation
```shell
NAME:
   ae86 config - generate initial config

USAGE:
   ae86 config [command options] [arguments...]

OPTIONS:
   --path value  config filepath (default: "${HOME}/.ae86/config.yaml")
```

#### Start the application
```shell
NAME:
   ae86 start - start server

USAGE:
   ae86 start [command options] [arguments...]

OPTIONS:
   --config value  filepath to config.yaml (default: "${HOME}/.ae86/config.yaml")
```