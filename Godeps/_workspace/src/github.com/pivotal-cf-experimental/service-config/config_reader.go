package service_config

import (
	"errors"
	"fmt"

	"github.com/fraenkel/candiedyaml"
	"github.com/imdario/mergo"
)

type Reader struct {
	configBytes []byte
}

func NewReader(configBytes []byte) *Reader {
	return &Reader{
		configBytes: configBytes,
	}
}

func (r Reader) Read(model interface{}) error {
	err := candiedyaml.Unmarshal(r.configBytes, model)
	if err != nil {
		return errors.New(fmt.Sprintf("Unmarshaling config: %s", err.Error()))
	}

	return nil
}

func (r Reader) ReadWithDefaults(model interface{}, defaults interface{}) error {

	err := r.Read(model)
	if err != nil {
		return err
	}

	return mergo.Merge(model, defaults)
}
