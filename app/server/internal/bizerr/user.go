package bizerr

import "github.com/go-kratos/kratos/v2/errors"

var ErrUserNotFound = errors.New(0, "", "未找到用户")
