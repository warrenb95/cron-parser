# Cron Parser

This is a simple cron parser written with the go standard library.

## Running the Application

To run the Cron Parser CLI, execute the following command from the project root:

```bash
go run main.go "*/15 0 1,15 * 1-5 /usr/bin/find"
```

Will produce output like:

```bash
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```

## Running tests

The application has unit tests and they can be run with the following command:

```bash
go test ./...
```

## TODO

* [ ] separate validation from the cron parser logic.
* [ ] use interfaces instead of passing in structs; more separation of layers.
* [ ] add a `--help` flag and explaining usage; alternative to the readme for usage instructions.
