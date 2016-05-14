# echo-server
Lightweight dockerized echo-server writed on golang.

Final docker image size ~1.5mb. It can be used for quick experiments with docker containers.

```bash
docker run -p 8080:8080 -d cortwave/echo-server
```

It will response next way:

```bash
curl -d "Hello!" localhost:8080

>Hello!
```
