package compress

import (
	"bytes"
	"compress/zlib"
)

//进行zlib压缩
//和php中的gzcompress一致，都是压缩的zlib格式
func ZlibCompress(src []byte) ([]byte, error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, err := w.Write(src)
	if err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return in.Bytes(), nil
}
