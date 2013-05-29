// +build raspberry

package egl

/*
#cgo CFLAGS: -I/opt/vc/include/interface/vmcs_host/
#include "vc_dispmanx.h"
*/
import "C"
import (
	"unsafe"
)

const (
	/* Bottom 2 bits sets the alpha mode */
	DISPMANX_FLAGS_ALPHA_FROM_SOURCE = 0
	DISPMANX_FLAGS_ALPHA_FIXED_ALL_PIXELS = 1
	DISPMANX_FLAGS_ALPHA_FIXED_NON_ZERO = 2
	DISPMANX_FLAGS_ALPHA_FIXED_EXCEED_0X07 = 3
	DISPMANX_FLAGS_ALPHA_PREMULT = 1 << 16
	DISPMANX_FLAGS_ALPHA_MIX = 1 << 17

	/* Bottom 2 bits sets the orientation */
	DISPMANX_NO_ROTATE = 0
	DISPMANX_ROTATE_90 = 1
	DISPMANX_ROTATE_180 = 2
	DISPMANX_ROTATE_270 = 3
	DISPMANX_FLIP_HRIZ = 1 << 16
	DISPMANX_FLIP_VERT = 1 << 17

	DISPMANX_PROTECTION_MAX = 0x0f
	DISPMANX_PROTECTION_NONE = 0
	DISPMANX_PROTECTION_HDCP = 11   // Derived from the WM DRM levels, 101-300

)

type VCImage C.VC_IMAGE_T

type VCRect struct {
	X, Y, Width, Height int32
}

type (
	DispmanxDisplayHandle uint32
	DispmanxElementHandle uint32
	DispmanxUpdateHandle uint32
	DispmanxResourceHandle uint32
	DispmanxProtection uint32
)

type DispmanxAlpha struct {
	Flags int
	Opacity uint32
	Mask *VCImage
}

type VCDispmanxAlpha struct {
	Flags int
	Opacity uint32
	Mask DispmanxResourceHandle
}

type DispmanxClamp struct {
  Mode int // DISPMANX_FLAGS_CLAMP_T mode;
  KeyMask int // DISPMANX_FLAGS_KEYMASK_T key_mask;
  KeyValue int // DISPMANX_CLAMP_KEYS_T key_value;
  ReplaceValue uint32 // uint32_t replace_value;
}

func VCDispmanxDisplayOpen(device uint32) DispmanxDisplayHandle {
	return DispmanxDisplayHandle(C.vc_dispmanx_display_open(C.uint32_t(device)))
}

func VCDispmanxUpdateStart(priority uint32) DispmanxUpdateHandle {
	return DispmanxUpdateHandle(C.vc_dispmanx_update_start(C.int32_t(priority)))
}

func VCDispmanxElementAdd(
	update DispmanxUpdateHandle, 
	display DispmanxDisplayHandle, 
	layer int32, 
	dstRect *VCRect,
	src DispmanxResourceHandle,
	srcRect *VCRect,
	protection DispmanxProtection,
	alpha *VCDispmanxAlpha,
	clamp *DispmanxClamp,
	transform int,
) DispmanxElementHandle {
	return DispmanxElementHandle(C.vc_dispmanx_element_add(
		C.DISPMANX_UPDATE_HANDLE_T(update), 
		C.DISPMANX_DISPLAY_HANDLE_T(display),
                C.int32_t(layer),
		(*C.VC_RECT_T)(unsafe.Pointer(dstRect)), 
		C.DISPMANX_RESOURCE_HANDLE_T(src),
                (*C.VC_RECT_T)(unsafe.Pointer(srcRect)),
		C.DISPMANX_PROTECTION_T(protection), 
                (*C.VC_DISPMANX_ALPHA_T)(unsafe.Pointer(alpha)),
                (*C.DISPMANX_CLAMP_T)(unsafe.Pointer(clamp)), 
		C.DISPMANX_TRANSFORM_T(transform)))
}

func VCDispmanxUpdateSubmitSync(update DispmanxUpdateHandle) int {
	return int(C.vc_dispmanx_update_submit_sync(
		C.DISPMANX_UPDATE_HANDLE_T(update)))
}
