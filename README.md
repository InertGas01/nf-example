# NF-Example

This repository is a sample NF for a simple HTTP service!
Try adding more services and learn how to collaborate using GitHub.

## Compile & Run

```sh
make
./bin/nf -c config/nfcfg.yaml
```

## Try Service

```sh
> curl -X GET http://127.0.0.163:8000/default/
"Hello free5GC!"

> curl -X GET http://127.0.0.163:8000/spyfamily/
"Hello SPYxFAMILY!"

> curl -X GET http://127.0.0.163:8000/spyfamily/character/Loid
"Character: Loid Forger"

> curl -X GET http://127.0.0.163:8000/notebook/User_Guide
> curl -X PUT http://127.0.0.163:8000/notebook/new_note_title/content
> curl -X POST http://127.0.0.163:8000/notebook/new_note_title/Content
> curl -X POST http://127.0.0.163:8000/notebook/new_note_title/append/can't
> curl -X POST http://127.0.0.163:8000/notebook/new_note_title/append/contain
> curl -X POST http://127.0.0.163:8000/notebook/new_note_title/append/whitespaces.
> curl -X GET http://127.0.0.163:8000/notebook/new_note_title
```

## Go Test

```sh
> go test -v ./...
```
