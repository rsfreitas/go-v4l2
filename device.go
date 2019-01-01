package v4l2

//
// To compile you must export the following:
// export CGO_CFLAGS_ALLOW=".*"
//

/*
#cgo CFLAGS: -I/usr/local/include -fgnu89-inline
#cgo LDFLAGS: -L/usr/local/lib -lv4l2 -lcollections
#include <stdlib.h>
#include <v4l2/libv4l2.h>
#include <collections/collections.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// Options is the structure used to pass the arguments when opening a
// video4linux2 device.
type Options struct {
	Device   string
	Width    int
	Height   int
	Format   V4l2ImageFormat
	Model    V4l2Model
	Channel  V4l2Channel
	LoopBack string
}

// V4l2 is the representation of a video4linux2 device.
type V4l2 struct {
	options Options
	v4l2    unsafe.Pointer
}

// Close closes the device.
func (v *V4l2) Close() {
	C.v4l2_close(v.v4l2)
}

// Capture captures an image from the video4linux2 device.
func (v *V4l2) Capture() (*V4l2Image, error) {
	img := C.v4l2_grab_image(v.v4l2, false)

	if err := C.v4l2_get_last_error(); err != 0 {
		return nil, errors.New(C.GoString(C.v4l2_strerror(err)))
	}

	data := C.GoBytes(unsafe.Pointer(C.v4l2_image_data(img)),
		C.int(C.v4l2_image_size(img)))

	v4l2Image, err := newV4l2Image(data, int(C.v4l2_image_width(img)),
		int(C.v4l2_image_height(img)), v.options.Format)

	if err != nil {
		return nil, err
	}

	return v4l2Image, nil
}

// SetSetting adjusts a specific setting in the capture device.
func (v *V4l2) SetSetting(setting V4l2Setting, value int) error {
	if ret := C.v4l2_set_setting(v.v4l2, setting.toCint(), C.int(value)); ret != 0 {
		err := C.v4l2_get_last_error()
		return errors.New(C.GoString(C.v4l2_strerror(err)))
	}

	return nil
}

// GetSetting gets the current value of a setting of a device.
func (v *V4l2) GetSetting(setting V4l2Setting) (int, error) {
	value := C.v4l2_get_setting(v.v4l2, setting.toCint())

	if value < 0 {
		err := C.v4l2_get_last_error()
		return 0, errors.New(C.GoString(C.v4l2_strerror(err)))
	}

	return int(value), nil
}

// Open opens a video4linux2 device to capture frames.
func Open(options Options) (*V4l2, error) {
	v4l2 := C.v4l2_open(C.CString(options.Device), C.int(options.Width),
		C.int(options.Height), options.Format.toCint(),
		options.Model.toCint(), options.Channel.toCint())

	if err := C.v4l2_get_last_error(); err != 0 {
		return nil, errors.New(C.GoString(C.v4l2_strerror(err)))
	}

	return &V4l2{
		options: options,
		v4l2:    v4l2,
	}, nil
}

func NewWebcam(options Options) (*V4l2, error) {
	return nil, nil
}

func NewRPI(options Options) (*V4l2, error) {
	return nil, nil
}

func NewLoopback(device string, source *V4l2) (*V4l2, error) {
	loopback := C.v4l2_loopback_open(C.CString(device), source.v4l2, true)

	if err := C.v4l2_get_last_error(); err != 0 {
		return nil, errors.New(C.GoString(C.v4l2_strerror(err)))
	}

	return &V4l2{
		v4l2: loopback,
	}, nil
}

func init() {
	// We must initialize libcollections before using libv4l2 API.
	C.cl_init(nil)
}
