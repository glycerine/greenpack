package msgp

import (
	"fmt"
	"reflect"
)

// writer then reader

// Writer

func (mw *Writer) ResetDedup() {
	mw.ptrWrit = make(map[interface{}]int)
	mw.ptrCountNext = 0
}

// diagnostic
func (mw *Writer) PointerCount() int {
	return len(mw.ptrWrit)
}

// upon writing each pointer, first check if it is a duplicate;
// i.e. appears more than once, pointing to the same object.
func (mw *Writer) IsDup(v interface{}) (res bool, err error) {
	defer func() {
		if recover() != nil {
			// just recover from panic. these are the defaults
			//res = false
			//err = nil
			return
		}
	}()
	if v == nil || reflect.ValueOf(v).IsNil() {
		return false, nil
	}
	k, dup := mw.ptrWrit[v]
	if !dup {
		mw.ptrWrit[v] = mw.ptrCountNext
		mw.ptrCountNext++
		return false, nil
	}
	return true, mw.WriteDedupExt(k)
}

// write DedupExtension with k integer count
// of the pointer that is duplicated here. k is
// runtime appearance order.
func (mw *Writer) WriteDedupExt(k int) error {
	var by [8]byte
	kby := AppendInt(by[:0], k)
	ext := RawExtension{
		Data: kby,
		Type: DedupExtension,
	}
	return mw.WriteExtension(&ext)
}

/////////////// Reader side

func (m *Reader) ReadDedupExt() (int, error) {
	ext := RawExtension{
		Type: DedupExtension,
	}
	err := m.ReadExtension(&ext)
	if err != nil {
		return -1, err
	}
	var nbs NilBitsStack
	k, _, err := nbs.ReadIntBytes(ext.Data)
	if err != nil {
		return -1, err
	}
	return k, nil
}

func (r *Reader) ResetDedup() {
	r.dedupPointers = r.dedupPointers[:0]
}

func (r *Reader) DedupPtr(k int) interface{} {
	if k >= 0 && k < len(r.dedupPointers) {
		return r.dedupPointers[k]
	}
	panic(fmt.Sprintf("Reader.DedupPtr requested for k=%v but that was out of range! (avail=%v)", k, len(r.dedupPointers)))
	return nil
}

func (m *Reader) IndexEachPtrForDedup(ptr interface{}) {
	if ptr == nil {
		return
	}
	va := reflect.ValueOf(ptr)
	if va.IsNil() {
		return
	}
	fmt.Printf("\n Reader.IndexEachPtrForDedup sees ptr '%#v'\n", ptr)
	m.dedupPointers = append(m.dedupPointers, ptr)
}

func (m *Reader) NextIsDup() (interface{}, bool) {

	typ, err := m.peekExtensionType()
	if err != nil {
		return nil, false
	}
	if typ != DedupExtension {
		return nil, false
	}
	k, err := m.ReadDedupExt()
	if err != nil {
		return nil, false
	}
	ptr := m.DedupPtr(k)
	return ptr, true
}
