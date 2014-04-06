# gowiki - go Webdev Tutorial

## Getting Started

Install go:
http://golang.org/doc/install

Tutorial here:
http://golang.org/doc/articles/wiki/

Version I'm developing with:

```
D:\dev\test\go\gowiki>go version
go version go1.2.1 windows/amd64
```

## My thoughts

Initial setup on win7 was simple.
Intellij needed a 3rd party plugin to support golang, but the latest version is only supported on Intellij v13. Grr.
SublimeText3 has syntax highlighting support for golang.
There's vim-go for Ivan ;)

Go slices are like collections.
Functions can return multiple values. I guess that means less wrapping constructs?
Underscore used for "throw away" values.

### Go's error reporting

Ye-olde stack trace - I had to read it a bit to find the line-number in my program - `wiki.go:43`:

```
2014/03/30 08:14:34 http: panic serving [::1]:53521: runtime error: invalid memory address or nil pointer dereference
goroutine 3 [running]:
net/http.func·009()
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/net/http/server.go:1093 +0xb1
runtime.panic(0x5ff940, 0x89190f)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/runtime/panic.c:248 +0x11b
main.viewHandler(0x132500, 0xc08400b5a0, 0xc08401d270)
        D:/dev/test/go/gowiki/wiki.go:43 +0x205
net/http.HandlerFunc.ServeHTTP(0x6b4870, 0x132500, 0xc08400b5a0, 0xc08401d270)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/net/http/server.go:1220 +0x43
net/http.(*ServeMux).ServeHTTP(0xc08403f0f0, 0x132500, 0xc08400b5a0, 0xc08401d270)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/net/http/server.go:1496 +0x166
net/http.serverHandler.ServeHTTP(0xc0840076e0, 0x132500, 0xc08400b5a0, 0xc08401d270)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/net/http/server.go:1597 +0x171
net/http.(*conn).serve(0xc084040280)
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/net/http/server.go:1167 +0x7ba
created by net/http.(*Server).Serve
        C:/Users/ADMINI~1/AppData/Local/Temp/2/makerelease250988475/go/src/pkg/net/http/server.go:1644 +0x28e
```

## Krzysztof Kowalczyk's Parting thoughts

(of Sumatra PDF fame)

I think writing non-trivial web services is a sweet spot for Go.

Most of the needed functionality is part of standard library. For almost everything else there are 3rd party libraries.

Writing in Go is almost as fast and fluent as writing in Python but the code is order of magnitude faster and uses less memory.

## Read more

Useful Golang website pages:

* http://golang.org/doc/
* http://golang.org/doc/code.html
* http://golang.org/doc/effective_go.html
* http://golang.org/doc/faq#ancestors

Golang package references:

* http://golang.org/pkg/net/http/

http://blog.kowalczyk.info/article/uvw2/Thoughts-on-Go-after-writing-3-websites.html
http://blog.golang.org/go-slices-usage-and-internals

* https://github.com/fatih/vim-go
* https://github.com/spf13/hugo

Useful packages
* http://www.gorillatoolkit.org/