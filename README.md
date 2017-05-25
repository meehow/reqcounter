Request counter
===============

It's very simple http request counter.
I made to debug proxy server, but feel free to use it for whatever you need.

There is deployed "cloud version" on http://reqcounter.herokuapp.com/

How to use it?
==============

Just post any json with element `id` inside:

```
$ curl -i --data '{"id": "test"}' http://reqcounter.herokuapp.com/ 
HTTP/1.1 202 Accepted
Server: Cowboy
Connection: keep-alive
Date: Thu, 25 May 2017 11:49:15 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
Via: 1.1 vegur
```

and you can check how many requests successfully arrived by giving the same `id` as a parameter to GET request:

```
$ curl 'http://reqcounter.herokuapp.com/?id=test' 
1
```

You can also use it with [ApacheBench](https://en.wikipedia.org/wiki/ApacheBench)

[![asciicast](https://asciinema.org/a/c9vg9nhkhm6xc6d8zp8h62d5o.png)](https://asciinema.org/a/c9vg9nhkhm6xc6d8zp8h62d5o)
