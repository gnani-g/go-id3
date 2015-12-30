package id3

import "errors"

type TagID string

const (
	V23_TIT2 = `TIT2`
	V23_TALB = `TALB`
	V23_TYER = `TYER`
	V23_TPE1 = `TPE1`
	V23_COMM = `COMM`
	V23_APIC = `APIC`
)

func initFrame(tagID TagID) (fr Framer, err error) {
	switch tagID {
	case V23_TIT2, V23_TALB, V23_TYER, V23_TPE1:
		return &TextFrame{frame: frame{tagID: tagID}}, nil
	case V23_COMM:
		return &COMMFrame{frame: frame{tagID: tagID}}, nil
	case V23_APIC:
		return &ImageFrame{frame: frame{tagID: tagID}}, nil
	default:
		return nil, errors.New(`Invalid tag id`)
	}
}
