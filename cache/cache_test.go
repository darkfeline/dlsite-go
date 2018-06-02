package cache

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"go.felesatra.moe/dlsite"
)

func TestGetMissingWork(t *testing.T) {
	t.Parallel()
	d, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Error making temp dir: %s", err)
	}
	defer os.RemoveAll(d)
	p := filepath.Join(d, "tmp.db")
	c, err := Open(p)
	if err != nil {
		t.Fatalf("Error opening cache: %s", err)
	}
	defer c.Close()
	_, err = c.Get(dlsite.RJCode("RJ1234"))
	if err == nil {
		t.Errorf("Got nil error when getting missing key")
	}
}

func TestPutAndGetWork(t *testing.T) {
	t.Parallel()
	d, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Error making temp dir: %s", err)
	}
	defer os.RemoveAll(d)
	p := filepath.Join(d, "tmp.db")
	c, err := Open(p)
	if err != nil {
		t.Fatalf("Error opening cache: %s", err)
	}
	defer c.Close()
	r := dlsite.RJCode("RJ1234")
	w := &dlsite.Work{
		RJCode: r,
		Name:   "foobar",
	}
	err = c.Put(w)
	if err != nil {
		t.Fatalf("Error putting work: %s", err)
	}
	got, err := c.Get(r)
	if err != nil {
		t.Fatalf("Error getting work: %s", err)
	}
	if !reflect.DeepEqual(w, got) {
		t.Errorf("Expected %#v, got %#v", w, got)
	}
}
