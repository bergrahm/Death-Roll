# Go Death-Roller

A simple script for command-line death-rolls. If you've configured your GOROOT, GOPATH and GOBIN correctly you can build the script and run it in your terminal by typing following command into your terminal.

```
go run Death-Roll [int]
```

# Dockerfile

Build the script and run it through docker.

```
docker build -t death-roll .
```

```
docker run -it --rm death-roll
```
