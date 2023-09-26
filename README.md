# Chronos

Chronos is a web application for SISAE

## Description

Chronos is a web application for SISAE. It should provide scheduling operations

## Getting Started

### Dependencies

* All dependencies are listed in [go.mod](https://github.com/JotaEspig/chronos/blob/main/go.mod) file
* They're installed automatically when running `go run .`
* You can also install it manually using `go mod download`

### Installing

Clone the repository:
```bash
git clone https://github.com/JotaEspig/chronos.git
cd chronos
```

Run:
```bash
go mod download
```

### Executing program

Run (it may take a while to compile in the first time):
```bash
go run .
```

## Tests

You can run the tests using:
```
go test ./tests
```

## Documentation

You have two options:

1. You can run the `godoc` program to generate dynamic documentation.
How to Install:
```bash
go install golang.org/x/tools/cmd/godoc@latest
```
Now try `godoc`. If it doesn't work, find where it was installed (probably in ~/go/bin/)
and run `<path-to-godoc>/godoc`

2. You can see the static documentation (it may be not up to date) in folder docs/

## Help

TODO

## Authors

* [João Vitor Espig](https://github.com/JotaEspig)
* [Mickael Reichert](https://github.com/mickaelrei)
* [Gustavo Tramontin Pedro](https://github.com/gustatramontin)
* [Yean Jy Chen](https://github.com/yeanjy)
* [Vítor Augusto Ueno Otto](https://github.com/vitorueno)
* João Pedro Hanisch
* Carlos Eduardo Perini Fidelis
* Guilherme de Campos

## Version History

You can see the version history in:
[CHANGELOG.md](https://github.com/JotaEspig/chronos/blob/main/CHANGELOG.md)

## License

TODO:

This project is licensed under the [NAME HERE] License - see the
[LICENSE](https://github.com/JotaEspig/chronos/blob/main/LICENSE)
file for details

## Acknowledgments

Inspiration, code snippets, etc.
* [README-Template.md](https://gist.github.com/DomPizzie/7a5ff55ffa9081f2de27c315f5018afc)
