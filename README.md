# serve

A simple HTTP file server that serves the current directory on the port defined by the `PORT` environment variable.

## Install

```
go get 4d63.com/serve
```

## Usage

```
serve
```

```
$ serve
2021/01/08 11:17:58 Listening on :8000
2021/01/08 11:18:04 GET /README.md 200 0.251259ms
```

```
$ PORT=3000 serve
2021/01/08 11:17:58 Listening on :3000
2021/01/08 11:18:04 GET /README.md 200 0.251259ms
```
