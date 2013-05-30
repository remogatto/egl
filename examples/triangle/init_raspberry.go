// +build raspberry

package main

import (
	"flag"
	"fmt"
	"github.com/mortdeus/mathgl"
	"github.com/remogatto/egl"
	"github.com/remogatto/egl/platform/raspberry"
	gl "github.com/remogatto/opengles2"
	"log"
	"math"
	"time"
)

const (
	INITIAL_WINDOW_WIDTH  = 1920
	INITIAL_WINDOW_HEIGHT = 1080
)

func initialize() {
	egl.BCMHostInit()
	platform.Initialize(platform.DefaultConfigAttributes, platform.DefaultContextAttributes)
}
