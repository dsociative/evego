package api

import (
	"encoding/xml"
	"errors"
)

func Parse(raw []byte, model Model) error {
	var err error = xml.Unmarshal(raw, model)

	if err == nil {
		apiError := model.GetError()
		if apiError.Code != "" {
			err = errors.New("Api request error code:" + apiError.Code)
		}
	}
	return err
}
