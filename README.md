# Pack API

## Implementation
- Reading pack size from `.env`(Alongside with log level and server port). 
It will use default values(Including docker-compose and Makefile)
- UI is implemented via Makefile(you can test it with `make get-packs-number items=12001` command)
- Clean architecture
- Unit and Integration tests
- Docker and docker-compose setup
- Makefile for simplicity
- Dependency injection

# TBD
- CI/CD
- Linter

## How to
To use config - copy `.env.example`, rename it to `.env` and fill it with your data.
It is not necessary, app will use the default values in `.env` file is absent.

All commands can be found at [Makefile](Makefile).
- To start the service, just run:
```make
make start
```

To get needed packs run:
```make
make get-packs-number items=12001
```
You can replace 12001 with any other number
P.S. If you use your custom 

To stop the service:
```make
make stop
```

## Testing
To run all tests:
```make
make tests
```
