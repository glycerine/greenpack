package msgp

import (
	"fmt"
	"reflect"
)

// Methods for deduplicating repeated occurances of the same pointer.
//
// When writing, we track the sequence of pointers written.
// When we see a duplicate pointer, we write the special
// extension "duplicate" value along with the pointer's
// occurance order in the serialization.
//
// As we read back, we keep a count that increments for every
// pointer we read, and we save a map from the count to the pointer.
// When we encounter a value that is the special value indicating reuse, then
// we refer back to the pointer k (k being found within the special extension value)
// and we plug in the k-th pointer instead.

// writer then reader methods

// ===============================
// ===============================
// Writer methods
// ===============================
// ===============================

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
func (mw *Writer) WriteIsDup(v interface{}) (res bool, err error) {
	defer func() {
		if recover() != nil {
			return
		}
	}()
	if v == nil || reflect.ValueOf(v).IsNil() {
		return false, nil
	}
	k, dup := mw.ptrWrit[v]
	if !dup {
		mw.ptrWrit[v] = mw.ptrCountNext
		fmt.Printf("\n\n $$$ write %p  -> k=%v / %#v\n\n", v, mw.ptrCountNext, v)
		mw.ptrCountNext++
		return false, nil
	} else {
		fmt.Printf("\n\n $$$ DUP write %p  -> k=%v / %#v\n\n", v, k, v)
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

// =============================
// =============================
// Reader side
// =============================
// =============================

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
	m.dedupPointers = append(m.dedupPointers, ptr)
	fmt.Printf("\n\n *** Reader.IndexEachPtrForDedup sees ptr '%#v', as sequence k=%v\n\n", ptr, len(m.dedupPointers)-1)
}

func (m *Reader) ReadIsDup() (interface{}, bool) {

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
