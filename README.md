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

This creates an executable `strategy_generator`.

## Go setup

It used to be that you needed to make a 'workspace' directory and then set an environment variable called `GOPATH` to that directory. All the user defined packages would have to be created relative to that path. This project uses go modules (available from go 1.11).

The file itself was created with `go mod init github.com/nishakm/strategy_generator`. If the project were to import another project, running `go build` will automatically version all of the dependencies either with a version number or a pseudo-version containing the git sha at which the dependencies were downloaded.

The advantage of using go modules is that you can clone this project and build the code as is without dealing with your environment.

However, if you are doing this from scratch i.e. not within the context of a github project, make sure all your work is in a dedicated directory. I have all my work in a directory called `go`.
