package protoconvert

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProtoConvertSuite struct {
	suite.Suite
	ctx context.Context
	protoConvert *ProtoConvert
}

func (p *ProtoConvertSuite) SetupTest() {
	p.ctx = context.Background()
	p.protoConvert = NewProtoConvert()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ProtoConvertSuite))
}