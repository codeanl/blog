package utils

import (
	"github.com/jinzhu/copier"
)

// 拷贝属性, 一般用于 vo -> po
func CopyProperties[T any](from any) (to T, err error) {
	if err := copier.Copy(&to, from); err != nil {
		return to, err
	}
	return to, nil
}
