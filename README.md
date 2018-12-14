# Wait For PSQL

This is a container to use that waits for your Postgresql database server to start up being stopping.

Idea and base code for this is influenced by a repo that does the same try wait for MySQL [here](https://github.com/jimmysawczuk/wait-for-mysql).

## Run Outside Container

This is how you can run this locally outside of a container.

```
go build -o waitforpsql main.go
./waitforpsql 'host=[host] dbname=[name] sslmode=disable user=root password=[password]'
```