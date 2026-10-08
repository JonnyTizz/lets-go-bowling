# 👯‍♂️ Pair-programming interview (Go)

A bowling scoring kata, built test-first while pairing.

## Setup

- Go 1.21+ installed (`go version`)
- Install the test watcher:

  ```sh
  go install gotest.tools/gotestsum@latest
  ```

## Running the tests

Open `bowling.go` and `bowling_test.go` side by side, then run the watcher in a terminal next to the editor:

```sh
gotestsum --watch
```

It re-runs the tests on save (like `jest --watchAll`). Without the watcher:

```sh
go test ./...
```

## Example solutions

Interviewer reference implementations live under `solutions/`. Run one with e.g. `go test ./solutions/frame/`.
