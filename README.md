# TODO

## How to

## Getting Started

### Prerequisites
- Install [Go](https://golang.org/doc/install)

### Installing
```
$ cd my_projects/
$ git clone https://github.com/dibikhin/TODO.git
$ cd TODO/
$ go mod download
$ cp example.env .env
```

### Testing
```
$ cd TODO/
$ make test
Testing client...
...
PASS
...
Testing server...
...
PASS
...
```

## Running
Run as is, no compilation neaded. Running locally assumed.

### Server
Open first terminal, then:
```
$ cd TODO/
$ make serve
Server:
2022/04/19 11:15:24 Starting...
2022/04/19 11:15:24 Started
...
```

NOTE: Hit `ctrl+c` to exit.

### Client
Open second terminal, then:
```
$ cd TODO/
$ make connect
Client:
2022/04/19 11:17:56 Starting...
...
2022/04/19 11:17:56 Started

Welcome to TODO
...
```

NOTE: Hit `ctrl+c` to exit.

## Internals

### Project Structure
- `/cmd` — Entry points
- `/pkg` — The packages
- `.env` — config
- `example.env` — config example
- etc.

### Features
- The UI is CLI
- Client handles loosing connection well

## Authors
- [Roman Dibikhin](https://github.com/dibikhin)

## License
This project is licensed under the MIT License — see the [LICENSE](./LICENSE) file for details.

## Acknowledgments
Thanks to:
- TODO
