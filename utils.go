package v4l2

/*
#cgo CFLAGS: -I/usr/local/include/v4l2
#include <libv4l2.h>
*/
import "C"

type V4l2ImageFormat int

const (
	ImageFormatUnknown V4l2ImageFormat = iota
	ImageFormatGray
	ImageFormatBGR24
	ImageFormatYUV420
	ImageFormatYUYV
)

func (f V4l2ImageFormat) toCint() uint32 {
	var i uint32

	switch f {
	case ImageFormatUnknown:
		i = C.V4L2_IMAGE_FMT_UNKNOWN

	case ImageFormatGray:
		i = C.V4L2_IMAGE_FMT_GRAY

	case ImageFormatBGR24:
		i = C.V4L2_IMAGE_FMT_BGR24

	case ImageFormatYUV420:
		i = C.V4L2_IMAGE_FMT_YUV420

	case ImageFormatYUYV:
		i = C.V4L2_IMAGE_FMT_YUYV
	}

	return i
}

type V4l2Model int

const (
	ModelUnknown V4l2Model = iota
	ModelBt878
	ModelUSBWebcam
	ModelRPICamera
)

func (v V4l2Model) toCint() uint32 {
	var i uint32

	switch v {
	case ModelUnknown:
		i = C.V4L2_MODEL_UNKNOWN

	case ModelBt878:
		i = C.V4L2_MODEL_BT878_CARD

	case ModelUSBWebcam:
		i = C.V4L2_MODEL_USB_WEBCAM

	case ModelRPICamera:
		i = C.V4L2_MODEL_RPI_CAMERA
	}

	return i
}

type V4l2Channel int

const (
	ChannelUnknown V4l2Channel = iota
	ChannelTuner
	ChannelComposite
	ChannelSVideo
)

func (v V4l2Channel) toCint() int32 {
	var i int32

	switch v {
	case ChannelUnknown:
		i = C.V4L2_CHANNEL_UNKNOWN

	case ChannelTuner:
		i = C.V4L2_CHANNEL_TUNER

	case ChannelComposite:
		i = C.V4L2_CHANNEL_COMPOSITE

	case ChannelSVideo:
		i = C.V4L2_CHANNEL_SVIDEO
	}

	return i
}

type V4l2Setting int

const (
	SettingUnknown V4l2Setting = iota
	SettingBrightness
	SettingContrast
	SettingSaturation
	SettingHue
)

func (v V4l2Setting) toCint() uint32 {
	var i uint32

	switch v {
	case SettingUnknown:
		i = C.V4L2_SETTING_UNKNOWN

	case SettingBrightness:
		i = C.V4L2_SETTING_BRIGHTNESS

	case SettingContrast:
		i = C.V4L2_SETTING_CONTRAST

	case SettingSaturation:
		i = C.V4L2_SETTING_SATURATION

	case SettingHue:
		i = C.V4L2_SETTING_HUE
	}

	return i
}

func V4l2SettingFromInt(v uint32) V4l2Setting {
	switch v {
	case C.V4L2_SETTING_BRIGHTNESS:
		return SettingBrightness

	case C.V4L2_SETTING_CONTRAST:
		return SettingContrast

	case C.V4L2_SETTING_SATURATION:
		return SettingSaturation

	case C.V4L2_SETTING_HUE:
		return SettingHue
	}

	return SettingUnknown
}
