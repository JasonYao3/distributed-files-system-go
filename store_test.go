package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedFileName := "71056ad8aa24742ea41ea36fa2e3452a31636e82"
	expectedPathName := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != expectedFileName {
		t.Errorf("have %s want %s", pathKey.Filename, expectedFileName)
	}

}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	id := generateID()
	key := "specials"
	data := []byte("some jpg bytes")

	if _, err := s.writeStream(id, key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(id, key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	s := newStore()
	id := generateID()
	defer teardown(t, s)

	for i := 0; i < 50; i++ {

		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")

		if _, err := s.writeStream(id, key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); !ok {
			t.Errorf("expected to have key %s", key)
		}

		_, r, err := s.Read(id, key)
		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(r)
		fmt.Println(string(b))

		if string(b) != string(data) {
			t.Errorf("want %s have %s", data, b)
		}

		if err := s.Delete(id, key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); ok {
			t.Errorf("expected to NOT have key %s", key)
		}
	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
