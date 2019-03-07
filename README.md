Simple http server in Go
========================

# Install

```
$ go build
$ go install
```

# Usage

### Basic run

```
$ shs
```

### Set port

```
$ shs 4000
2019/03/08 02:56:30 Serving HTTP on http://localhost:4000
```

### Custom 404 page

If `404.html` file exist, shs will serve it as a custom `404 Not Found` response.


# License

MIT
