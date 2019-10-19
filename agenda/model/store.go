package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Store struct {
	Path string
}

func (s *Store)Load(data interface{}) error {
	file, err := os.OpenFile(s.Path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("open file error when store load data %v", err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	if err := dec.Decode(data); err != nil {
		return err
	}
	return nil
}

func (s *Store)Persist(data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(s.Path, bs, os.ModePerm)
}