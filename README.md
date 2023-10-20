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
`env`  - enviroment tag
`connection-string` - path to sqlite3 db file