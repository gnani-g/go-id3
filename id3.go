package id3

import (
	"errors"
	"log"

	"github.com/sbinet/go-python"
)

var (
	_mutagen *python.PyObject
	// _id3     *python.PyObject
)

func init() {

	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}

	/*
		gopath, _ := os.LookupEnv(`GOPATH`)
		impPath := filepath.Join(gopath, `src/github.com/gnani-g/go-id3/`)

		sys := python.PyImport_ImportModule("sys")
		if sys == nil {
			panic(`could not import 'sys'`)
		}

		// sys.path.append('/path/to/ptdraft/')
		sysPath := sys.GetAttrString(`path`)
		if sysPath == nil {
			panic(`could not get 'path'`)
		}

		//pathAdded := sysPath.CallMethod(`append`, `/usr/local/gopath/src/github.com/lotusfivestar/etv/utils/gopytest/`)
		pathAdded := sysPath.CallMethod(`append`, impPath)
		if pathAdded == nil {
			panic(`add path failed`)
		}
	*/

	_mutagen = python.PyImport_ImportModule("mutagen.id3")
	if _mutagen == nil {
		panic(`could not import 'mutagen.id3'`)
	}

	/*
		_id3 = python.PyImport_ImportModule("id3wrapper")
		if _id3 == nil {
			panic(`could not import 'id3wrapper'`)
		}
	*/

}

type ID3 struct {
	pyID3     *python.PyObject
	V2version int
}

func Open(path string, v2_version int) (id3 *ID3, err error) {

	id3 = &ID3{V2version: v2_version}

	_ID3 := _mutagen.GetAttrString(`ID3`)
	if _ID3 == nil {
		return nil, errors.New(`Unable to open file`)
	}

	kw := python.PyDict_New()
	err = python.PyDict_SetItem(kw, python.PyString_FromString(`v2_version`), python.PyInt_FromLong(id3.V2version))
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	args := python.PyTuple_New(1)
	err = python.PyTuple_SetItem(args, 0, python.PyString_FromString(path))
	if err != nil {
		return nil, errors.New(`Unable to set to list`)
	}

	id3.pyID3 = _ID3.Call(args, kw)
	if id3.pyID3 == nil {
		return nil, errors.New(`Unable to open file`)
	}

	/*
		_ID3 := _id3.GetAttrString(`ID3`)
		if _ID3 == nil {
			return nil, errors.New(`Unable to open file`)
		}

		id3.pyID3 = _ID3.CallFunction(path, python.PyInt_FromLong(v2_version))
		if id3.pyID3 == nil {
			return nil, errors.New(`Unable to open file`)
		}
	*/

	return

}

func (id3 *ID3) Frame(tagID TagID) (fr Framer, err error) {
	fr, err = initFrame(tagID)
	if err != nil {
		return
	}

	fr.setPyFrame(id3.pyID3.CallMethod(`getall`, python.PyString_FromString(string(tagID))))
	if fr.pyFrame() == nil {
		return nil, errors.New(`failed to retrieve tag`)
	}

	return
}

func (id3 *ID3) Close() (err error) {

	kw := python.PyDict_New()
	err = python.PyDict_SetItem(kw, python.PyString_FromString(`v2_version`), python.PyInt_FromLong(id3.V2version))
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	args := python.PyTuple_New(0)

	saveFunc := id3.pyID3.GetAttrString(`save`)
	if saveFunc == nil {
		return errors.New(`Unable to get save func`)
	}

	out := saveFunc.Call(args, kw)
	/*
		out := id3.pyID3.CallMethod(`save`, python.PyInt_FromLong(id3.V2version))
	*/
	if out == nil {
		return errors.New(`failed to save ID3 info`)
	}

	return
}

func (id3 *ID3) SetFrame(fr Framer) (err error) {

	out := id3.pyID3.CallMethod(`setall`, python.PyString_FromString(string(fr.TagID())), fr.pyFrame())
	if out == nil {
		return errors.New(`set failed`)
	}

	return

}

func (id3 *ID3) DeleteFrames() (err error) {

	out := id3.pyID3.CallMethod(`delete`)
	if out == nil {
		return errors.New(`delete frames failed`)
	}

	return

}
