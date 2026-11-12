package sendtablescs2

// fixedWidthDecoder reads the fixed8/16/32 encoders as N raw bits instead of a varint. Since the
// CS2 update of 2026-11-11 some integer fields use them (for example
// CCSPlayerController.m_recentKillQueue); read as a varint, a value >= 128 eats an extra byte and
// misaligns the rest of the packet. Signed types are sign-extended to int32 like signedDecoder.
func fixedWidthDecoder(f *field, baseType string) (fieldDecoder, bool) {
	var n uint32

	switch f.encoder {
	case "fixed8":
		n = 8
	case "fixed16":
		n = 16
	case "fixed32":
		n = 32
	default:
		return nil, false
	}

	switch baseType {
	case "int8", "int16", "int32":
		shift := 32 - n

		return func(r *reader) interface{} {
			return int32(r.readBits(n)<<shift) >> shift
		}, true
	default:
		return func(r *reader) interface{} {
			return uint64(r.readBits(n))
		}, true
	}
}
