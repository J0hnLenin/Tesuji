package bootstrap

import "github.com/J0hnLenin/Tesuji/internal/services/protoconvert"

func InitProtoCnvert() *protoconvert.ProtoConvert {
	return protoconvert.NewProtoConvert()
}