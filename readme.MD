## Request Counter

To start the application run 
```make start```
to start the application.

There is a configuration file in ```{project-root}/pkg/config/config.go```

By default, it serves on port: ```8081```

Application uses file as a database. If there is no database file, it creates it on start up.

File writes are atomic and shutdown safe.