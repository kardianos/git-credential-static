package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

func get(r io.Reader, w io.Writer) error {
	kv, err := read(r)
	if err != nil {
		return err
	}
	rw, err := openKVStore(false)
	if err != nil {
		return err
	}

	list, err := readKVStoreRead(rw)
	rw.Close()
	if err != nil {
		return err
	}

	found, _ := find(kv, list)
	if found == nil {
		return nil
	}
	for key, value := range found {
		fmt.Fprintf(w, "%s=%s\n", key, value)
	}
	return nil
}

func store(r io.Reader) error {
	kv, err := read(r)
	if err != nil {
		return err
	}
	rw, err := openKVStore(false)
	if err != nil {
		return err
	}

	list, err := readKVStoreRead(rw)
	rw.Close()

	found, n := find(kv, list)
	if found == nil {
		list = append(list, kv)
	} else {
		list[n] = kv
	}
	rw, err = openKVStore(true)
	if err != nil {
		return err
	}
	err = readKVStoreWrite(rw, list)
	ferr := rw.Close()
	if ferr != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

const storeName = ".gitstaticstore"

type KVList []KV

func openKVStore(trunc bool) (io.ReadWriteCloser, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	storePath := filepath.Join(u.HomeDir, storeName)
	mode := os.O_CREATE | os.O_RDWR
	if trunc {
		mode |= os.O_TRUNC
	}
	return os.OpenFile(storePath, mode, 0600)
}

func readKVStoreRead(r io.Reader) (KVList, error) {
	var list KVList
	coder := json.NewDecoder(r)
	err := coder.Decode(&list)
	if err == io.EOF {
		return list, nil
	}
	return list, err
}
func readKVStoreWrite(w io.Writer, list KVList) error {
	coder := json.NewEncoder(w)
	return coder.Encode(list)
}

func find(kv KV, list KVList) (KV, int) {
outer:
	for i, check := range list {
		for key, value := range kv {
			if key == "password" {
				continue
			}
			ckValue, has := check[key]
			if !has {
				continue outer
			}
			if ckValue != value {
				continue outer
			}
		}
		return check, i
	}
	return nil, -1
}
