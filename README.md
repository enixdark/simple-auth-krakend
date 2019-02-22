# simple-auth-krakend

simple Basic authentication midleware for the [KrakenD framework](https://github.com/devopsfaith/krakend).

** This repo 's based repo krakend-http-auth of kpacha, but with update code to can intergated into latest version of krakend

## Usage

Using this component is as easy as:

1. Add the static credentials to your config file. You must define on every endpoint to intercept.

```
"extra_config": {
    "github.com/enixdark/simple-auth-krakend": {
        "pass": "test",
        "user": "test"
    }
}
```

2. Import the right package for you. There is a `HandlerFactory` implementation for the basic routing lib: `gin`.

```
import(
	auth "github.com/enixdark/simple-auth-krakend"
)
```

3. Decorate your `HandlerFactory` with the `auth.HandlerFactory`

Check the dummy implementation in the `example` dir. It contains a simple config file and a KrakenD api-gateway with the auth `HandlerFactory` wrapping the default endpoint factory.

```
krakendgin.Config{
	Engine:         gin.Default(),
	ProxyFactory:   proxy.DefaultFactory(logger),
	Middlewares:    []gin.HandlerFunc{},
	Logger:         logger,
	HandlerFactory: auth.HandlerFactory(krakendgin.EndpointHandler),
}
```

## Playing with the example

Build it

```
$ make all
```

And run it

```
$ ./auth -l DEBUG -c example/krakend.json
```

From a new terminal, try to acces the private endpoint with and without the auth header

```
$ curl -i http://127.0.0.1:8080/show/1
HTTP/1.1 403 Forbidden
Content-Type: text/plain; charset=utf-8
Date: Sun, 01 Oct 2017 17:47:18 GMT
Content-Length: 17

{"Unauthorized"}

$ curl -i -u test:test http://127.0.0.1:8080/show/1
HTTP/1.1 200 OK
Cache-Control: public, max-age=0
Content-Type: application/json; charset=utf-8
X-Krakend: Version undefined
Date: Sun, 01 Oct 2017 17:48:09 GMT
Content-Length: 159

```
