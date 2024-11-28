package jsons2array

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"iter"
	"os"

	ja "github.com/takanoriyanagitani/go-jsons2avro"
	util "github.com/takanoriyanagitani/go-jsons2avro/util"
)

func ReaderToJsonArrays(r io.Reader) iter.Seq2[ja.JsonArray, error] {
	return func(yield func(ja.JsonArray, error) bool) {
		var dec *json.Decoder = json.NewDecoder(r)
		var err error = nil
		var buf []any

		for {
			clear(buf)
			buf = buf[:0]

			err = dec.Decode(&buf)

			if io.EOF == err {
				return
			}

			if !yield(ja.JsonArray(buf), err) {
				return
			}
		}
	}
}

func StdinToJsonArrays() util.Io[iter.Seq2[ja.JsonArray, error]] {
	return func(_ context.Context) (iter.Seq2[ja.JsonArray, error], error) {
		var br io.Reader = bufio.NewReader(os.Stdin)
		return ReaderToJsonArrays(br), nil
	}
}
