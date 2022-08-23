package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"go.kuoruan.net/v8go-polyfills/base64"
	"go.kuoruan.net/v8go-polyfills/console"
	"go.kuoruan.net/v8go-polyfills/fetch"
	v8 "rogchap.com/v8go"
)

type HookOutputWriter struct{}

func (h *HookOutputWriter) Write(p []byte) (n int, err error) {
	log.Println(string(p))
	return 0, nil
}

func main() {
	var stdin []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	isolate := v8.NewIsolate()
	global := v8.NewObjectTemplate(isolate)
	if err := fetch.InjectTo(isolate, global); err != nil {
		panic(err)
	}
	if err := base64.InjectTo(isolate, global); err != nil {
		panic(err)
	}
	ctx := v8.NewContext(isolate, global)

	writer := HookOutputWriter{}
	if err := console.InjectTo(ctx, console.WithOutput(&writer)); err != nil {
		panic(err)
	}
	v, err := ctx.RunScript(string(stdin), "stdin")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", v)
}
