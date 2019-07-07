# Wait For Postgres

This is a container to use that waits for your Postgresql database server to start up being stopping.

Idea and base code for this is influenced by a repo that does the same try wait for MySQL [here](https://github.com/jimmysawczuk/wait-for-mysql).

## Run With Container

```
docker run --network host mmcken3/wait-for-psql 'host=[hostname] dbname=[name] sslmode=disable user=[username] password=[password]'
```

## Run Outside Container

This is how you can run this locally outside of a container.

```
go build -o waitforpsql main.go
./waitforpsql 'host=[host] dbname=[name] sslmode=disable user=root password=[password]'
```

## Pushing to Docker Hub

[Docker Hub](https://hub.docker.com/r/mmcken3/wait-for-psql)

Get the current github short tag using this command:

```
git describe --abbrev=0 --tags
```

Build, Tag, and push this Docker image to Docker Hub using that tag.

```
docker build -t mmcken3/wait-for-psql .
docker tag mmcken3/wait-for-psql:latest mmcken3/wait-for-psql:[github-tag]
docker push mmcken3/wait-for-psql:latest
docker push mmcken3/wait-for-psql:[github-tag]
```

## Go Dep

This tool is now set up to be using go mods with the vendors committed as of v1.1. If you would like to use the version of this tool that uses dep it was switched over after release v1.0. Here are steps for how to get yourself in a new branch with the tool that was using dep.

git clone https://github.com/mmcken3/wait-for-psql.git
cd wait-for-psql
git checkout tags/v1.0 -b wait-with-dep
