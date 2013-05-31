package test

import (
	"github.com/remogatto/egl"
	pt "github.com/remogatto/prettytest"
	"testing"
)

type testInitEGLSuite struct { pt.Suite }

type testCreateEGLContextSuite struct {
	pt.Suite 
	display egl.Display
	config egl.Config
	numConfig int32
}

// Test the initialization of EGL

func (t *testInitEGLSuite) TestGetDisplay() {
	display := egl.GetDisplay(egl.DEFAULT_DISPLAY)
	t.True(display != 0)
}

func (t *testInitEGLSuite) TestInitialize() {
	display := egl.GetDisplay(egl.DEFAULT_DISPLAY)
	t.True(egl.Initialize(display, nil, nil))}

func (t *testInitEGLSuite) TestChooseConfig() {
	var (
		config egl.Config
		numConfig int32
	)
	display := egl.GetDisplay(egl.DEFAULT_DISPLAY)
	t.True(egl.Initialize(display, nil, nil))
	t.True(egl.ChooseConfig(display, configAttr, &config, 1, &numConfig))
	t.Equal(numConfig, int32(1))
	t.True(config != 0)
}


// Test the creation of an EGL context and surface

func (t *testCreateEGLContextSuite) BeforeAll() {
	t.display = egl.GetDisplay(egl.DEFAULT_DISPLAY)
	egl.Initialize(t.display, nil, nil)
	egl.ChooseConfig(t.display, configAttr, &t.config, 1, &t.numConfig)
	initPlatform()
}

func (t *testCreateEGLContextSuite) TestCreateContext() {
	context := egl.CreateContext(
		t.display,
		t.config, 
		egl.NO_CONTEXT, 
		&contextAttr[0])
	t.True(context != egl.NO_CONTEXT)
	t.True(context != egl.BAD_MATCH)
	t.True(context != egl.BAD_DISPLAY)
	t.True(context != egl.NOT_INITIALIZED)
	t.True(context != egl.BAD_CONFIG)
	t.True(context != egl.BAD_CONTEXT)
	t.True(context != egl.BAD_ATTRIBUTE)
	t.True(context != egl.BAD_ALLOC)
}

func (t *testCreateEGLContextSuite) TestWindowSurface() {
	egl.BindAPI(egl.OPENGL_ES_API)
	egl.CreateContext(
		t.display,
		t.config,
		egl.NO_CONTEXT, 
		&contextAttr[0])
	surface := egl.CreateWindowSurface(
		t.display, 
		t.config, 
		egl.NativeWindowType(uintptr(testWin.Id)), 
		nil)
	t.True(surface != egl.NO_SURFACE)
}

func (t *testCreateEGLContextSuite) TestMakeCurrent() {
	egl.BindAPI(egl.OPENGL_ES_API)
	context := egl.CreateContext(
		t.display,
		t.config,
		egl.NO_CONTEXT, 
		&contextAttr[0])
	surface := egl.CreateWindowSurface(
		t.display, 
		t.config, 
		egl.NativeWindowType(uintptr(testWin.Id)), 
		nil)
	t.True(egl.MakeCurrent(
		t.display,
		surface,
		surface,
		context,
	))
}

func TestEGL(t *testing.T) {
	pt.Run(
		t, 
		new(testInitEGLSuite),
		new(testCreateEGLContextSuite),
	)
}
