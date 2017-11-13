# Go Playground

Playing around with Go. Currently it consists of...
- Simple webserver serving an HTML file @ [`http://localhost:6060`](http://localhost:6060)

*TODO*
- REST API
  - Create models & controllers
- Read/write to persistent storagei
- Add a caching layer
- Increase use of Go routines
- Add more concurrency examples
- Tests for all of the above

### Develop
```
# Get dependencies
make deps

# Start the webserver, which will live-reload and rebuild
make watch
```

### License
MIT
