# url-shortener

## Config

`config/config.yaml` is path to a config file

An Example of file:
```yaml
port: "localhost:8080"
env: "local"
connection-string: "./storage/strorage.db"
```

### Fields

`port` - address to bind

`env`  - enviroment tag (affect to logging level)

`connection-string` - path to sqlite3 db file

## Build

run `make` command to build a project to executable file. After successfuly builded executable file will be located in root directory