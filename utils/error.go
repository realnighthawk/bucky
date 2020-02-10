package utils

import "github.com/kumarabd/appkit/errors"

func ErrUnmarshal(key string, err error) error {
	return errors.New("ERR.UNMARSHAL", "Unmarshal error for key: "+key+", error: "+err.Error())
}

func ErrGetBool(key string, err error) error {
	return errors.New("ERR.GETBOOL", "Error while getting Boolean value for key: "+key+", error: "+err.Error())
}
