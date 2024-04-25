# Goxios

<p align="center">

[![GoReport](https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B?style=flat-square)](https://goreportcard.com/report/github.com/Simplou/goxios)

</p>
  <em><b>Goxios</b> is an HTTP client written in Go, inspired by <a href="https://github.com/axios/axios">Axios</a>
</p>

## **Introduction:**

Goxios is a powerful HTTP client library for Go developers, heavily inspired by the renowned Axios library in JavaScript. It aims to simplify the process of making HTTP requests in Go, offering a clean and intuitive API while focusing on performance and memory efficiency.

## **Key Features:**

- **Axios-inspired API:** Goxios provides an API that will feel familiar to developers who have experience with Axios, making it easy to transition between JavaScript and Go projects.

- **Fast Development:** By abstracting away low-level details and providing a high-level interface, Goxios accelerates development, allowing you to focus on building your application logic without getting bogged down by HTTP client intricacies.

## **Getting Started:**

To start using Goxios in your project, simply import it and begin making requests:

```go
package main

import (
	"context"
	"fmt"
	"io"

	"github.com/Simplou/goxios"
)

func main() {
    // Create a new Goxios client
    client := goxios.New(context.Background())

    // Make a GET request
    requestOpts := &goxios.RequestOpts{Headers: []goxios.Header{}}
    resp, err := client.Get("https://api.sampleapis.com/codingresources/codingResources", requestOpts)
    if err != nil {
       panic(err)
    }
    defer resp.Body.Close()
    // Save the response body bytes in 'b'
    b, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b)) // Print the content of 'b' as a string
}
```

For more advanced usage and customization options, refer to the [documentation](https://pkg.go.dev/github.com/Simplou/goxios#pkg-overview) on GoDoc.
.

## **License:**

Goxios is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## **Acknowledgements:**

Special thanks to the developers of Axios for their inspiration.

## **Maintainer:**

Gabriel Luiz - [@gabrielluizdev](https://twitter.com/gabrielluizdev)

---
