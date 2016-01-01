package id3

import "strconv"

func (id3 *ID3) SetTitle(title string) (err error) {
	fr, err := id3.Frame(V23_TIT2)
	if err != nil {
		return err
	}

	err = fr.(*TextFrame).SetValues(title)
	if err != nil {
		return err
	}

	return id3.SetFrame(fr)
}

func (id3 *ID3) Title() (string, error) {
	fr, err := id3.Frame(V23_TIT2)
	if err != nil {
		return ``, err
	}

	return fr.(*TextFrame).Val()
}

func (id3 *ID3) SetAlbum(val string) (err error) {
	fr, err := id3.Frame(V23_TALB)
	if err != nil {
		return err
	}

	err = fr.(*TextFrame).SetValues(val)
	if err != nil {
		return err
	}

	return id3.SetFrame(fr)
}

func (id3 *ID3) Album() (string, error) {
	fr, err := id3.Frame(V23_TALB)
	if err != nil {
		return ``, err
	}

	return fr.(*TextFrame).Val()
}

func (id3 *ID3) SetYear(val int) (err error) {
	fr, err := id3.Frame(V23_TYER)
	if err != nil {
		return err
	}

	err = fr.(*TextFrame).SetValues(strconv.Itoa(val))
	if err != nil {
		return err
	}

	return id3.SetFrame(fr)
}

func (id3 *ID3) Year() (int, error) {
	fr, err := id3.Frame(V23_TYER)
	if err != nil {
		return 0, err
	}

	strVal, err := fr.(*TextFrame).Val()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strVal)
}

func (id3 *ID3) SetArtist(val string) (err error) {
	fr, err := id3.Frame(V23_TPE1)
	if err != nil {
		return err
	}

	err = fr.(*TextFrame).SetValues(val)
	if err != nil {
		return err
	}

	return id3.SetFrame(fr)
}

func (id3 *ID3) Artist() (string, error) {
	fr, err := id3.Frame(V23_TPE1)
	if err != nil {
		return ``, err
	}

	return fr.(*TextFrame).Val()
}

func (id3 *ID3) SetComment(val string) (err error) {
	fr, err := id3.Frame(V23_COMM)
	if err != nil {
		return err
	}

	err = fr.(*COMMFrame).SetValues(val)
	if err != nil {
		return err
	}

	return id3.SetFrame(fr)
}

func (id3 *ID3) Comment() (string, error) {
	fr, err := id3.Frame(V23_COMM)
	if err != nil {
		return ``, err
	}

	return fr.(*COMMFrame).Val()
}

func (id3 *ID3) SetCoverImage(mime string, desc string, data []byte) (err error) {
	fr, err := id3.Frame(V23_APIC)
	if err != nil {
		return err
	}

	err = fr.(*ImageFrame).SetValues(&Image{3, mime, desc, data})
	if err != nil {
		return err
	}

	return id3.SetFrame(fr)
}

func (id3 *ID3) CoverImage() (*Image, error) {
	fr, err := id3.Frame(V23_APIC)
	if err != nil {
		return nil, err
	}

	return fr.(*ImageFrame).ImageByType(3)
}

func (id3 *ID3) SetID(owner, id string) (err error) {
	fr, err := id3.Frame(V23_UFID)
	if err != nil {
		return err
	}

	err = fr.(*UFIDFrame).SetID(owner, id)
	if err != nil {
		return err
	}

	return id3.SetFrame(fr)
}

func (id3 *ID3) ID() (string, error) {
	fr, err := id3.Frame(V23_UFID)
	if err != nil {
		return ``, err
	}

	return fr.(*UFIDFrame).ID()
}

func (id3 *ID3) SetLanguage(val string) (err error) {
	fr, err := id3.Frame(V23_TLAN)
	if err != nil {
		return err
	}

	err = fr.(*TextFrame).SetValues(val)
	if err != nil {
		return err
	}

	return id3.SetFrame(fr)
}

func (id3 *ID3) Language() (string, error) {
	fr, err := id3.Frame(V23_TLAN)
	if err != nil {
		return ``, err
	}

	return fr.(*TextFrame).Val()
}
