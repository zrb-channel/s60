package s60

import (
	s60 "github.com/zrb-channel/s60/schema"
)

func LoginRedirect(req *s60.LoginRedirect) (string, error) {

	params, err := req.Params()
	if err != nil {
		return "", err
	}
	return loginRedirectAddr + params, nil
}
