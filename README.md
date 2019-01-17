# Business Strategy Generator written in Go

This was inspired by [a blog post by Simon Wardley](http://blog.gardeviance.org/2014/07/a-quick-route-to-building-strategy.html)

It also doubles as a Go programming guide and hence has more comments than actual lines of code

Content in the code is governed by the [Creative Commons Attribution-Share Alike 3.0 License](http://creativecommons.org/licenses/by-sa/3.0/)

It's a silly project but you can submit PRs to it if you'd like

## To run

```
git clone <this repo>
cd strategy_generator
go run main.go
```

## To create a binary executable

```
git clone <this repo>
cd strategy_generator
go build
```

This creates an executable `strategy_generator` and a `go.sum` which is a checksum for all the files that went into the binary executable.

## Go setup

It used to be that you needed to make a 'workspace' directory and then set an environment variable called `GOPATH` to that directory. All the user defined packages would have to be created relative to that path. This project uses go modules (available from go 1.11).

The file itself was created with `go mod init strategy_generator`. The versioning of the `require` line happened automatically after running `go build`
