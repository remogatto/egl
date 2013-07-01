# What's that?

<tt>egl</tt> is a [Go](http://golang.org) package for accessing the
[EGL](http://en.wikipedia.org/wiki/EGL_(OpenGL)) (Embedded Graphics
Library). EGL is the access door toward hardware accelerated graphics,
through OpenGL, on many embedded devices. The project was born for
accessing the GPU of the [Raspberry PI](http://raspberrypi.org). Now
has been generalized to be go installable on other platforms too. This
has the benefit that you could develop Open GL ES 2.0 application on
your desktop computer using [Mesa](http://www.mesa3d.org/egl.html) and
deploy them on embedded systems like the Raspberry PI.

# Currently supported platform

* Raspberry PI
* Xorg

# Install

The package aims to be multiplatform. To achive this result two
approaches are used: [build
constraints](http://golang.org/pkg/go/build) and
per-platform/per-implementation initialization [boilerplate
code](platform/). By default egl will use the xorg implementation.

~~~bash
$ go get github.com/remogatto/egl # use xorg by default
~~~

To build egl against a particular implementation use the specific
build constraint, for example:

~~~bash
% go get github.com/remogatto/egl -tags="raspberry" # install on the raspberry
~~~

# Usage

Please refer to the [examples](examples/).

# To Do

* Add support for other platforms (e.g. android)
* Add tests

# Credits

Thanks to Roger Roach for his [egl/opengles](https://github.com/mortdeus/egles) libraries. I stole a lot from his repository!

# License

See [LICENSE](LICENSE)
