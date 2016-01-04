package id3

import (
	"errors"
	"fmt"

	"github.com/sbinet/go-python"
	"golang.org/x/text/encoding/unicode"
)

var (
	utf16 = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
)

type Framer interface {
	TagID() TagID
	setPyFrame(*python.PyObject)
	pyFrame() *python.PyObject
}

type frame struct {
	//id3 *ID3

	tagID TagID
	pyFr  *python.PyObject
}

func (f *frame) TagID() TagID {
	return f.tagID
}

func (f *frame) setPyFrame(pyFr *python.PyObject) {
	f.pyFr = pyFr
}

func (f *frame) pyFrame() *python.PyObject {
	return f.pyFr
}

type TextFrame struct {
	frame
}

func (fr *TextFrame) Val() (string, error) {
	vals, err := fr.Vals()
	if err != nil {
		return ``, err
	}

	if len(vals) > 0 {
		return vals[0], nil
	}

	return ``, nil
}

func (fr *TextFrame) Vals() (vals []string, err error) {

	for i := 0; i < python.PyList_GET_SIZE(fr.pyFr); i++ {
		item := python.PyList_GetItem(fr.pyFr, 0)
		if item == nil {
			return nil, errors.New(`Unable to read title`)
		}

		pyVals := item.GetAttrString(`text`)
		if pyVals == nil {
			return nil, errors.New(`Unable to read title`)
		}

		val := ``

		for j := 0; j < python.PyList_GET_SIZE(pyVals); j++ {
			pyVal := python.PyList_GetItem(pyVals, 0)
			if pyVal == nil {
				return nil, errors.New(`Unable to read title`)
			}

			v, err := utf16.NewDecoder().String(python.PyString_AsString(pyVal))
			if err != nil {
				return nil, err
			}

			val += ` ` + v
		}

		if len(val) > 0 {
			val = val[1:]
		}

		vals = append(vals, val)
	}

	return
}

func (fr *TextFrame) SetValues(vals ...string) (err error) {

	fr.pyFr = python.PyList_New(0)
	if fr.pyFr == nil {
		return errors.New(`list creation failed`)
	}

	return fr.AppendValues(vals...)

}

func (fr *TextFrame) AppendValues(vals ...string) (err error) {

	position := python.PyList_GET_SIZE(fr.pyFr)

	for _, v := range vals {

		v, err = utf16.NewEncoder().String(v)
		if err != nil {
			return
		}

		pyVal := python.PyList_New(0)
		if pyVal == nil {
			return errors.New(`list creation failed`)
		}
		python.PyList_Insert(pyVal, 0, python.PyString_FromString(v))

		frClass := _mutagen.GetAttrString(string(fr.tagID))
		if frClass == nil {
			return errors.New(`Unable to get class`)
		}

		tagIns := frClass.CallFunction() //python.PyLong_FromLong(1), n_title)
		if tagIns == nil {
			return errors.New(`instanciation failed`)
		}

		tagIns.SetAttrString(`encoding`, python.PyInt_FromLong(1))
		tagIns.SetAttrString(`text`, pyVal)

		python.PyList_Insert(fr.pyFr, position, tagIns)

		position++

	}

	return

}

type COMMFrame struct {
	frame
}

func (fr *COMMFrame) Val() (string, error) {
	vals, err := fr.Vals()
	if err != nil {
		return ``, err
	}

	if len(vals) > 0 {
		return vals[0], nil
	}

	return ``, nil
}

func (fr *COMMFrame) Vals() (vals []string, err error) {

	for i := 0; i < python.PyList_GET_SIZE(fr.pyFr); i++ {
		item := python.PyList_GetItem(fr.pyFr, 0)
		if item == nil {
			return nil, errors.New(`Unable to read title`)
		}

		pyVals := item.GetAttrString(`text`)
		if pyVals == nil {
			return nil, errors.New(`Unable to read title`)
		}

		val := ``

		for j := 0; j < python.PyList_GET_SIZE(pyVals); j++ {
			pyVal := python.PyList_GetItem(pyVals, 0)
			if pyVal == nil {
				return nil, errors.New(`Unable to read title`)
			}

			v, err := utf16.NewDecoder().String(python.PyString_AsString(pyVal))
			if err != nil {
				return nil, err
			}

			val += ` ` + v
		}

		if len(val) > 0 {
			val = val[1:]
		}

		vals = append(vals, val)
	}

	return
}

func (fr *COMMFrame) SetValues(vals ...string) (err error) {

	fr.pyFr = python.PyList_New(0)
	if fr.pyFr == nil {
		return errors.New(`list creation failed`)
	}

	return fr.AppendValues(vals...)

}

func (fr *COMMFrame) AppendValues(vals ...string) (err error) {

	position := python.PyList_GET_SIZE(fr.pyFr)

	for _, v := range vals {

		v, err = utf16.NewEncoder().String(v)
		if err != nil {
			return
		}

		pyVal := python.PyList_New(0)
		if pyVal == nil {
			return errors.New(`list creation failed`)
		}
		python.PyList_Insert(pyVal, 0, python.PyString_FromString(v))

		frClass := _mutagen.GetAttrString(string(fr.tagID))
		if frClass == nil {
			return errors.New(`Unable to get class`)
		}

		tagIns := frClass.CallFunction() //python.PyLong_FromLong(1), n_title)
		if tagIns == nil {
			return errors.New(`instanciation failed`)
		}

		tagIns.SetAttrString(`encoding`, python.PyInt_FromLong(1))
		tagIns.SetAttrString(`lang`, python.PyString_FromString(`eng`))
		tagIns.SetAttrString(`text`, pyVal)

		python.PyList_Insert(fr.pyFr, position, tagIns)

		position++

	}

	return

}

type ImageFrame struct {
	frame
}

type Image struct {
	Type int
	Mime string
	Desc string
	Data []byte
}

func (fr *ImageFrame) ImageByType(t int) (*Image, error) {
	allImgs, err := fr.All()
	if err != nil {
		return nil, err
	}

	for _, img := range allImgs {
		if img.Type == t {
			return img, nil
		}
	}

	return nil, nil
}

func (fr *ImageFrame) All() (imgs []*Image, err error) {

	for i := 0; i < python.PyList_GET_SIZE(fr.pyFr); i++ {
		item := python.PyList_GetItem(fr.pyFr, 0)
		if item == nil {
			return nil, errors.New(`Unable to read title`)
		}

		pyType := item.GetAttrString(`type`)
		if pyType == nil {
			return nil, errors.New(`Unable to read img type`)
		}

		pyMime := item.GetAttrString(`mime`)
		if pyMime == nil {
			return nil, errors.New(`Unable to read image mime`)
		}

		pyDesc := item.GetAttrString(`desc`)
		if pyDesc == nil {
			return nil, errors.New(`Unable to read image desc`)
		}

		pyData := item.GetAttrString(`data`)
		if pyData == nil {
			return nil, errors.New(`Unable to read image data`)
		}

		bArr := python.PyByteArray_FromObject(pyData)

		imgs = append(imgs, &Image{
			python.PyInt_AsLong(pyType),
			python.PyString_AsString(pyMime),
			python.PyString_AsString(pyDesc),
			python.PyByteArray_AsBytes(bArr),
		})

	}
	fmt.Println(len(imgs[0].Data))

	return
}

func (fr *ImageFrame) SetValues(imgs ...*Image) (err error) {

	fr.pyFr = python.PyList_New(0)
	if fr.pyFr == nil {
		return errors.New(`list creation failed`)
	}

	return fr.AppendValues(imgs...)

}

func (fr *ImageFrame) AppendValues(imgs ...*Image) (err error) {

	position := python.PyList_GET_SIZE(fr.pyFr)

	for _, img := range imgs {

		frClass := _mutagen.GetAttrString(string(fr.tagID))
		if frClass == nil {
			return errors.New(`Unable to get class`)
		}

		tagIns := frClass.CallFunction()
		if tagIns == nil {
			return errors.New(`instanciation failed`)
		}

		tagIns.SetAttrString(`encoding`, python.PyInt_FromLong(1))
		tagIns.SetAttrString(`type`, python.PyInt_FromLong(img.Type))
		tagIns.SetAttrString(`mime`, python.PyString_FromString(img.Mime))
		tagIns.SetAttrString(`desc`, python.PyString_FromString(img.Desc))
		tagIns.SetAttrString(`data`, python.PyString_FromStringAndSize(string(img.Data), len(img.Data)))

		python.PyList_Insert(fr.pyFr, position, tagIns)

		position++

	}

	return

}

type UFIDFrame struct {
	frame
}

func (fr *UFIDFrame) ID() (string, error) {
	vals, err := fr.IDs()
	if err != nil {
		return ``, err
	}

	if len(vals) > 1 {
		return ``, errors.New(`multiple key-val pairs. use IDbyOwner`)
	}

	for _, v := range vals {
		return v, nil
	}

	return ``, nil
}

func (fr *UFIDFrame) Owners() (vals []string, err error) {
	ids, err := fr.IDs()
	if err != nil {
		return
	}

	for k, _ := range ids {
		vals = append(vals, k)
	}

	return
}

func (fr *UFIDFrame) IDs() (vals map[string]string, err error) {

	vals = make(map[string]string)
	for i := 0; i < python.PyList_GET_SIZE(fr.pyFr); i++ {
		item := python.PyList_GetItem(fr.pyFr, 0)
		if item == nil {
			return nil, errors.New(`Unable to read title`)
		}

		pyKey := item.GetAttrString(`owner`)
		if pyKey == nil {
			return nil, errors.New(`Unable to read owner`)
		}

		pyData := item.GetAttrString(`data`)
		if pyData == nil {
			return nil, errors.New(`Unable to read title`)
		}

		vals[python.PyString_AsString(pyKey)] = python.PyString_AsString(pyData)
	}

	return
}

func (fr *UFIDFrame) SetID(owner, id string) (err error) {
	vals := make(map[string]string)
	vals[owner] = id

	return fr.SetValues(vals)
}

func (fr *UFIDFrame) SetValues(vals map[string]string) (err error) {

	fr.pyFr = python.PyList_New(0)
	if fr.pyFr == nil {
		return errors.New(`list creation failed`)
	}

	return fr.AppendValues(vals)

}

func (fr *UFIDFrame) AppendValues(vals map[string]string) (err error) {

	position := python.PyList_GET_SIZE(fr.pyFr)

	for k, v := range vals {

		frClass := _mutagen.GetAttrString(string(fr.tagID))
		if frClass == nil {
			return errors.New(`Unable to get class`)
		}

		tagIns := frClass.CallFunction() //python.PyLong_FromLong(1), n_title)
		if tagIns == nil {
			return errors.New(`instanciation failed`)
		}

		tagIns.SetAttrString(`owner`, python.PyString_FromString(k))
		tagIns.SetAttrString(`data`, python.PyString_FromString(v))

		python.PyList_Insert(fr.pyFr, position, tagIns)

		position++

	}

	return

}
