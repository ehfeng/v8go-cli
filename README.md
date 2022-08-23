# v8go-cli

tool for testing v8go via command line stdin. To use it, clone repo, `go install`

```sh
cat 'console.log("Hello, world!"); "Returned string"' | v8go-cli
```
