# What's that?

<tt>egl</tt> is a [Go](http://golang.org) package for accessing
the [EGL](http://en.wikipedia.org/wiki/EGL_(OpenGL)) (Embedded Graphics Library). EGL is the access door toward hardware accelerated graphics, through OpenGL, on many platform. The project was born with the idea of accessing the GPU of the [Raspberry PI](http://raspberrypi.org). Now has been generalized to be go installable on other platforms too. This has the benefit that you could develop Open GL ES 2.0 applications on your desktop computer using [Mesa](http://www.mesa3d.org/egl.html) and then deploy them on embedded systems like the Raspberry PI.

# Features

* Thread-safe (in progress)
* Adherent to EGL API (in progress)
* Multiplatform

# Supported platform

* Raspberry PI
* Xorg

# Install

~~~bash
$ go get github.com/remogatto/egl # bind to xorg by default
% go get github.com/remogatto/egl -tags="raspberry" # install on the raspberry
~~~

# Thanks

* Roger Roach for his [egl/opengles](https://github.com/mortdeus/egles) libraries.

# License

See [LICENSE](egl/LICENSE)
