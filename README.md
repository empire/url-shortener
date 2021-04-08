The projects structure is based on [project-layout](https://github.com/golang-standards/project-layout)

## Make url shorten
Shorten a url:

```bash
curl --header "Content-Type: application/json"  --request POST --data '{"age": 2, "url": "https://www.google.com/", "hash": "gl"}' localhost:9000/shorten
```

Retreive by code
```bash
curl localhost:9000/gl
```

## Build and run
Build docker images
```bash
docker-compose build
```

Run containers
```
docker-compose up
```
