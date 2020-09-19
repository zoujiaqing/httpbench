package main

import (
    "github.com/valyala/fasthttp"
    "fmt"
    "runtime"
    "os"
)

const helloWorldStr = "Hello, World!"

func main() {
    fmt.Printf("ENV says GOMAXPROCS=%s\n", os.Getenv("GOMAXPROCS"))
    fmt.Printf("GOMAXPROCS says GOMAXPROCS=%d\n", runtime.GOMAXPROCS(0))
    fmt.Printf("runtime says MAXPROCS=%d\n", runtime.NumCPU())

    server := &fasthttp.Server{
        Name:                          "Go01234567890123456789",
        Handler:                       Plaintext,
        DisableHeaderNamesNormalizing: true,
    }

    if err := server.ListenAndServe(":8080"); err != nil {
        panic(err)
    }
}

func Plaintext(ctx *fasthttp.RequestCtx) {
    // ctx.Response.Header.Set("Keep-Alive", "timeout=10")
    ctx.Response.SetBodyString(helloWorldStr)
}
