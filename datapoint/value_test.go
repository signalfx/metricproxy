package datapoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntWire(t *testing.T) {
	iv := NewIntValue(3)
	assert.Equal(t, iv.String(), "3")
	i := iv.Int()
	assert.Equal(t, int64(3), i)
}

func TestFloatWire(t *testing.T) {
	iv := NewFloatValue(3)
	assert.Equal(t, iv.String(), "3")
	f := iv.Float()
	assert.Equal(t, 3.0, f)
}

func TestStrWire(t *testing.T) {
	iv := NewStringValue("val")
	assert.Equal(t, iv.String(), "val")
}

//
//func TestDatumValue(t *testing.T) {
//	iv := NewDatumValue(&com_signalfuse_metrics_protobuf.Datum{DoubleValue: workarounds.GolangDoesnotAllowPointerToFloat64Literal(3.0)})
//	assert.Equal(t, iv.String(), "3")
//	assert.Equal(t, iv.(FloatValue).Float(), 3.0)
//
//	iv = NewDatumValue(&com_signalfuse_metrics_protobuf.Datum{IntValue: workarounds.GolangDoesnotAllowPointerToIntLiteral(3)})
//	assert.Equal(t, iv.(IntValue).Int(), 3)
//
//	iv = NewDatumValue(&com_signalfuse_metrics_protobuf.Datum{StrValue: workarounds.GolangDoesnotAllowPointerToStringLiteral("hello")})
//	assert.Equal(t, iv.String(), "hello")
//}
