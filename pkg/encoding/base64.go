package encoding

import "encoding/base64"

func B64Encode(data string, padding ...rune) string {
	enc := maybeApplyPadding(base64.StdEncoding, padding...)
	return enc.EncodeToString([]byte(data))
}

func B64URLEncode(data string, padding ...rune) string {
	enc := maybeApplyPadding(base64.URLEncoding, padding...)
	return enc.EncodeToString([]byte(data))
}

func B64RawEncode(data string, padding ...rune) string {
	enc := maybeApplyPadding(base64.RawStdEncoding, padding...)
	return enc.EncodeToString([]byte(data))
}

func B64URLRawEncode(data string, padding ...rune) string {
	enc := maybeApplyPadding(base64.RawURLEncoding, padding...)
	return enc.EncodeToString([]byte(data))
}

func B64EncodeBytes(data []byte, padding ...rune) string {
	enc := maybeApplyPadding(base64.StdEncoding, padding...)
	return enc.EncodeToString(data)
}

func B64URLEncodeBytes(data []byte, padding ...rune) string {
	enc := maybeApplyPadding(base64.URLEncoding, padding...)
	return enc.EncodeToString(data)
}

func B64RawEncodeBytes(data []byte, padding ...rune) string {
	enc := maybeApplyPadding(base64.RawStdEncoding, padding...)
	return enc.EncodeToString(data)
}

func B64URLRawEncodeBytes(data []byte, padding ...rune) string {
	enc := maybeApplyPadding(base64.RawURLEncoding, padding...)
	return enc.EncodeToString(data)
}

func B64EncodeToBytes(data string, padding ...rune) []byte {
	enc := maybeApplyPadding(base64.StdEncoding, padding...)
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, []byte(data))
	return buf
}

func B64URLEncodeToBytes(data string, padding ...rune) []byte {
	enc := maybeApplyPadding(base64.URLEncoding, padding...)
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, []byte(data))
	return buf
}

func B64RawEncodeToBytes(data string, padding ...rune) []byte {
	enc := maybeApplyPadding(base64.RawStdEncoding, padding...)
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, []byte(data))
	return buf
}

func B64URLRawEncodeToBytes(data string, padding ...rune) []byte {
	enc := maybeApplyPadding(base64.RawURLEncoding, padding...)
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, []byte(data))
	return buf
}

func B64EncodeBytesToBytes(data []byte, padding ...rune) []byte {
	enc := maybeApplyPadding(base64.StdEncoding, padding...)
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, data)
	return buf
}

func B64URLEncodeBytesToBytes(data []byte, padding ...rune) []byte {
	enc := maybeApplyPadding(base64.URLEncoding, padding...)
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, data)
	return buf
}

func B64RawEncodeBytesToBytes(data []byte, padding ...rune) []byte {
	enc := maybeApplyPadding(base64.RawStdEncoding, padding...)
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, data)
	return buf
}

func B64URLRawEncodeBytesToBytes(data []byte, padding ...rune) []byte {
	enc := maybeApplyPadding(base64.RawURLEncoding, padding...)
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, data)
	return buf
}

func B64Decode(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func B64URLDecode(data string) (string, error) {
	b, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func B64RawDecode(data string) (string, error) {
	b, err := base64.RawStdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func B64URLRawDecode(data string) (string, error) {
	b, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func B64DecodeToBytes(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func B64URLDecodeToBytes(data string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(data)
}

func B64RawDecodeToBytes(data string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(data)
}

func B64URLRawDecodeToBytes(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}

func B64DecodeBytes(data []byte) (string, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(buf, data)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func B64URLDecodeBytes(data []byte) (string, error) {
	buf := make([]byte, base64.URLEncoding.DecodedLen(len(data)))
	n, err := base64.URLEncoding.Decode(buf, data)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func B64RawDecodeBytes(data []byte) (string, error) {
	buf := make([]byte, base64.RawStdEncoding.DecodedLen(len(data)))
	n, err := base64.RawStdEncoding.Decode(buf, data)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func B64URLRawDecodeBytes(data []byte) (string, error) {
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(data)))
	n, err := base64.RawURLEncoding.Decode(buf, data)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func B64DecodeBytesToBytes(data []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(buf, data)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func B64URLDecodeBytesToBytes(data []byte) ([]byte, error) {
	buf := make([]byte, base64.URLEncoding.DecodedLen(len(data)))
	n, err := base64.URLEncoding.Decode(buf, data)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func B64RawDecodeBytesToBytes(data []byte) ([]byte, error) {
	buf := make([]byte, base64.RawStdEncoding.DecodedLen(len(data)))
	n, err := base64.RawStdEncoding.Decode(buf, data)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func B64URLRawDecodeBytesToBytes(data []byte) ([]byte, error) {
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(data)))
	n, err := base64.RawURLEncoding.Decode(buf, data)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func MustB64Decode(data string) string {
	result, err := B64Decode(data)
	if err != nil {
		panic(err)
	}
	return result
}

func MustB64URLDecode(data string) string {
	result, err := B64URLDecode(data)
	if err != nil {
		panic(err)
	}
	return result
}

func MustB64RawDecode(data string) string {
	result, err := B64RawDecode(data)
	if err != nil {
		panic(err)
	}
	return result
}

func MustB64URLRawDecode(data string) string {
	result, err := B64URLRawDecode(data)
	if err != nil {
		panic(err)
	}
	return result
}

func maybeApplyPadding(enc *base64.Encoding, padding ...rune) *base64.Encoding {
	if len(padding) > 0 {
		return enc.WithPadding(padding[0])
	}
	return enc
}
