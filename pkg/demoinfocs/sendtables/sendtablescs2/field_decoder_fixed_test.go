package sendtablescs2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedWidthDecoderUnsigned(t *testing.T) {
	f := field{varName: "m_recentKillQueue", encoder: "fixed8", fieldType: newFieldType("uint8")}
	dec := findDecoder(&f)
	r := newReader([]byte{0xFF, 0x01, 0x80})

	assert.Equal(t, uint64(0xFF), dec(r))
	assert.Equal(t, uint64(0x01), dec(r))
	assert.Equal(t, uint64(0x80), dec(r))
	assert.Zero(t, r.remBytes())

	f32 := field{encoder: "fixed32", fieldType: newFieldType("uint32")}
	assert.Equal(t, uint64(0xDEADBEEF), findDecoder(&f32)(newReader([]byte{0xEF, 0xBE, 0xAD, 0xDE})))
}

func TestFixedWidthDecoderSignedStaysInt32(t *testing.T) {
	f8 := field{encoder: "fixed8", fieldType: newFieldType("int8")}
	assert.Equal(t, int32(-1), findDecoder(&f8)(newReader([]byte{0xFF})))

	f16 := field{encoder: "fixed16", fieldType: newFieldType("int16")}
	assert.Equal(t, int32(-2), findDecoder(&f16)(newReader([]byte{0xFE, 0xFF})))

	f32 := field{encoder: "fixed32", fieldType: newFieldType("int32")}
	assert.Equal(t, int32(11), findDecoder(&f32)(newReader([]byte{0x0B, 0, 0, 0})))
}

func TestUint8WithoutEncoderStaysVarint(t *testing.T) {
	f := field{varName: "m_x", fieldType: newFieldType("uint8")}
	assert.Equal(t, uint64(300), findDecoder(&f)(newReader([]byte{0xAC, 0x02})))
}
