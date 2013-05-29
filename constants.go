package egl

/*
#cgo pkg-config: egl
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
*/
import "C"

const (
	/* EGL Versioning */
	VERSION_1_0 = 1
	VERSION_1_1 = 1
	VERSION_1_2 = 1
	VERSION_1_3 = 1
	VERSION_1_4 = 1

	/* EGL Enumerants. Bitmasks and other exceptional cases aside, most
	 * enums are assigned unique values starting at 0x3000.
	 */
	/* EGL aliases */

	//FALSE = 0
	//TRUE  = 1

	/* Out-of-band handle values */
	DEFAULT_DISPLAY = 0
	NO_CONTEXT      = 0
	NO_DISPLAY      = 0
	NO_SURFACE      = 0

	/* Out-of-band attribute value */
	DONT_CARE = -1

	/* Errors / GetError return values */
	SUCCESS             = 0x3000
	NOT_INITIALIZED     = 0x3001
	BAD_ACCESS          = 0x3002
	BAD_ALLOC           = 0x3003
	BAD_ATTRIBUTE       = 0x3004
	BAD_CONFIG          = 0x3005
	BAD_CONTEXT         = 0x3006
	BAD_CURRENT_SURFACE = 0x3007
	BAD_DISPLAY         = 0x3008
	BAD_MATCH           = 0x3009
	BAD_NATIVE_PIXMAP   = 0x300A
	BAD_NATIVE_WINDOW   = 0x300B
	BAD_PARAMETER       = 0x300C
	BAD_SURFACE         = 0x300D
	CONTEXT_LOST        = 0x300E /* EGL 1.1 - IMG_power_management */

	/* Reserved= 0x300F=-0x301F for additional errors */

	/* Config attributes */
	BUFFER_SIZE             = 0x3020
	ALPHA_SIZE              = 0x3021
	BLUE_SIZE               = 0x3022
	GREEN_SIZE              = 0x3023
	RED_SIZE                = 0x3024
	DEPTH_SIZE              = 0x3025
	STENCIL_SIZE            = 0x3026
	CONFIG_CAVEAT           = 0x3027
	CONFIG_ID               = 0x3028
	LEVEL                   = 0x3029
	MAX_PBUFFER_HEIGHT      = 0x302A
	MAX_PBUFFER_PIXELS      = 0x302B
	MAX_PBUFFER_WIDTH       = 0x302C
	NATIVE_RENDERABLE       = 0x302D
	NATIVE_VISUAL_ID        = 0x302E
	NATIVE_VISUAL_TYPE      = 0x302F
	SAMPLES                 = 0x3031
	SAMPLE_BUFFERS          = 0x3032
	SURFACE_TYPE            = 0x3033
	TRANSPARENT_TYPE        = 0x3034
	TRANSPARENT_BLUE_VALUE  = 0x3035
	TRANSPARENT_GREEN_VALUE = 0x3036
	TRANSPARENT_RED_VALUE   = 0x3037
	NONE                    = 0x3038 /* Attrib list terminator */
	BIND_TO_TEXTURE_RGB     = 0x3039
	BIND_TO_TEXTURE_RGBA    = 0x303A
	MIN_SWAP_INTERVAL       = 0x303B
	MAX_SWAP_INTERVAL       = 0x303C
	LUMINANCE_SIZE          = 0x303D
	ALPHA_MASK_SIZE         = 0x303E
	COLOR_BUFFER_TYPE       = 0x303F
	RENDERABLE_TYPE         = 0x3040
	MATCH_NATIVE_PIXMAP     = 0x3041 /* Pseudo-attribute (not queryable) */
	CONFORMANT              = 0x3042

	/* Reserved= 0x3041=-0x304F for additional config attributes */

	/* Config attribute values */
	SLOW_CONFIG           = 0x3050 /* CONFIG_CAVEAT value */
	NON_CONFORMANT_CONFIG = 0x3051 /* CONFIG_CAVEAT value */
	TRANSPARENT_RGB       = 0x3052 /* TRANSPARENT_TYPE value */
	RGB_BUFFER            = 0x308E /* COLOR_BUFFER_TYPE value */
	LUMINANCE_BUFFER      = 0x308F /* COLOR_BUFFER_TYPE value */

	/* More config attribute values, for TEXTURE_FORMAT */
	NO_TEXTURE   = 0x305C
	TEXTURE_RGB  = 0x305D
	TEXTURE_RGBA = 0x305E
	TEXTURE_2D   = 0x305F

	/* Config attribute mask bits */
	PBUFFER_BIT                 = 0x0001 /* SURFACE_TYPE mask bits */
	PIXMAP_BIT                  = 0x0002 /* SURFACE_TYPE mask bits */
	WINDOW_BIT                  = 0x0004 /* SURFACE_TYPE mask bits */
	VG_COLORSPACE_LINEAR_BIT    = 0x0020 /* SURFACE_TYPE mask bits */
	VG_ALPHA_FORMAT_PRE_BIT     = 0x0040 /* SURFACE_TYPE mask bits */
	MULTISAMPLE_RESOLVE_BOX_BIT = 0x0200 /* SURFACE_TYPE mask bits */
	SWAP_BEHAVIOR_PRESERVED_BIT = 0x0400 /* SURFACE_TYPE mask bits */

	OPENGL_ES_BIT  = 0x0001 /* RENDERABLE_TYPE mask bits */
	OPENVG_BIT     = 0x0002 /* RENDERABLE_TYPE mask bits */
	OPENGL_ES2_BIT = 0x0004 /* RENDERABLE_TYPE mask bits */
	OPENGL_BIT     = 0x0008 /* RENDERABLE_TYPE mask bits */

	/* QueryString targets */
	VENDOR      = 0x3053
	VERSION     = 0x3054
	EXTENSIONS  = 0x3055
	CLIENT_APIS = 0x308D

	/* QuerySurface / SurfaceAttrib / CreatePbufferSurface targets */
	HEIGHT                = 0x3056
	WIDTH                 = 0x3057
	LARGEST_PBUFFER       = 0x3058
	TEXTURE_FORMAT        = 0x3080
	TEXTURE_TARGET        = 0x3081
	MIPMAP_TEXTURE        = 0x3082
	MIPMAP_LEVEL          = 0x3083
	RENDER_BUFFER         = 0x3086
	VG_COLORSPACE         = 0x3087
	VG_ALPHA_FORMAT       = 0x3088
	HORIZONTAL_RESOLUTION = 0x3090
	VERTICAL_RESOLUTION   = 0x3091
	PIXEL_ASPECT_RATIO    = 0x3092
	SWAP_BEHAVIOR         = 0x3093
	MULTISAMPLE_RESOLVE   = 0x3099

	/* RENDER_BUFFER values / BindTexImage / ReleaseTexImage buffer targets */
	BACK_BUFFER   = 0x3084
	SINGLE_BUFFER = 0x3085

	/* OpenVG color spaces */
	VG_COLORSPACE_sRGB   = 0x3089 /* VG_COLORSPACE value */
	VG_COLORSPACE_LINEAR = 0x308A /* VG_COLORSPACE value */

	/* OpenVG alpha formats */
	VG_ALPHA_FORMAT_NONPRE = 0x308B /* ALPHA_FORMAT value */
	VG_ALPHA_FORMAT_PRE    = 0x308C /* ALPHA_FORMAT value */

	/* Constant scale factor by which fractional display resolutions &
	 * aspect ratio are scaled when queried as integer values.
	 */
	DISPLAY_SCALING = 10000
	/* Back buffer swap behaviors */
	BUFFER_PRESERVED = 0x3094 /* SWAP_BEHAVIOR value */
	BUFFER_DESTROYED = 0x3095 /* SWAP_BEHAVIOR value */

	/* CreatePbufferFromClientBuffer buffer types */
	OPENVG_IMAGE = 0x3096

	/* QueryContext targets */
	CONTEXT_CLIENT_TYPE = 0x3097

	/* CreateContext attributes */
	CONTEXT_CLIENT_VERSION = 0x3098

	/* Multisample resolution behaviors */
	MULTISAMPLE_RESOLVE_DEFAULT = 0x309A /* MULTISAMPLE_RESOLVE value */
	MULTISAMPLE_RESOLVE_BOX     = 0x309B /* MULTISAMPLE_RESOLVE value */

	/* BindAPI/QueryAPI targets */
	OPENGL_ES_API = 0x30A0
	OPENVG_API    = 0x30A1
	OPENGL_API    = 0x30A2

	/* GetCurrentSurface targets */
	DRAW = 0x3059
	READ = 0x305A

	/* WaitNative engines */
	CORE_NATIVE_ENGINE = 0x305B

	/* EGL 1.2 tokens renamed for consistency in EGL 1.3 */
	COLORSPACE          = VG_COLORSPACE
	ALPHA_FORMAT        = VG_ALPHA_FORMAT
	COLORSPACE_sRGB     = VG_COLORSPACE_sRGB
	COLORSPACE_LINEAR   = VG_COLORSPACE_LINEAR
	ALPHA_FORMAT_NONPRE = VG_ALPHA_FORMAT_NONPRE
	ALPHA_FORMAT_PRE    = VG_ALPHA_FORMAT_PRE
)
