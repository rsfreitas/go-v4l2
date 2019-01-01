package v4l2

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
)

type V4l2Image struct {
	img    image.Image
	width  uint
	height uint
	size   uint
}

func (i *V4l2Image) ToJpeg(filename string) error {
	buf := &bytes.Buffer{}

	if err := jpeg.Encode(buf, i.img, nil); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}

func (i *V4l2Image) Width() uint {
	return i.width
}

func (i *V4l2Image) Height() uint {
	return i.height
}

func (i *V4l2Image) Size() uint {
	return i.size
}

func newV4l2Image(frame []byte, width, height int, format V4l2ImageFormat) (*V4l2Image, error) {
	var img image.Image

	switch format {
	case ImageFormatBGR24:
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		img.Pix = frame

	case ImageFormatYUYV:
		yuyv := image.NewYCbCr(image.Rect(0, 0, width, height), image.YCbCrSubsampleRatio422)

		for i := range yuyv.Cb {
			ii := i * 4
			yuyv.Y[i*2] = frame[ii]
			yuyv.Y[i*2+1] = frame[ii+2]
			yuyv.Cb[i] = frame[ii+1]
			yuyv.Cr[i] = frame[ii+3]

		}

		img = yuyv
	}

	return &V4l2Image{
		img:    img,
		width:  uint(width),
		height: uint(height),
		size:   uint(len(frame)),
	}, nil
}
