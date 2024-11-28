package jsons2maps

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

func ReaderToJsonMaps(r io.Reader) iter.Seq2[ja.JsonMap, error] {
	return func(yield func(ja.JsonMap, error) bool) {
		var dec *json.Decoder = json.NewDecoder(r)
		var err error = nil
		var buf map[string]any

		for {
			clear(buf)

			err = dec.Decode(&buf)

			if io.EOF == err {
				return
			}

			if !yield(ja.JsonMap(buf), err) {
				return
			}
		}
	}
}

func StdinToJsonMaps() util.Io[iter.Seq2[ja.JsonMap, error]] {
	return func(_ context.Context) (iter.Seq2[ja.JsonMap, error], error) {
		var br io.Reader = bufio.NewReader(os.Stdin)
		return ReaderToJsonMaps(br), nil
	}
}
