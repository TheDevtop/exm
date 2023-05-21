# EXM Search Engine

EXM is the cloud-native search engine from the future.
It isn't out of development, but can be considererd production-ready.

EXM stands for **Ex Mathematica**, which is a reference to the movie Ex Machina.

### Build:

1. Clone/checkout this repository.
2. Configure the Dockerfile environment variables.
3. Execute: `make docker`.
4. Execute: `docker image list` to verify the build image.

### Deploy:

1. Copy the [Docker Compose](https://github.com/TheDevtop/exm/blob/main/deploy/docker-compose.yml) file.
2. Modify the environment variables to math your environment.
3. Execute: `docker compose up`, to launch your EXM instance.

### Usage:
```
$ ./exm --help
Usage of ./exm:
  -cache int
        Specify cache size (default 8)
  -driver string
        Specify storage driver (default "none")
```