package typed2avro

import (
	"bufio"
	"context"
	"errors"
	"io"
	"iter"
	"os"

	ha "github.com/hamba/avro/v2"
	ho "github.com/hamba/avro/v2/ocf"

	util "github.com/takanoriyanagitani/go-jsons2avro/util"
)

func TypedToWriter[T any](
	typed iter.Seq2[T, error],
	writer io.Writer,
	schema ha.Schema,
) error {
	var bw *bufio.Writer = bufio.NewWriter(writer)
	enc, e := ho.NewEncoderWithSchema(schema, bw)
	if nil != e {
		return e
	}

	for row, e := range typed {
		if nil != e {
			return e
		}

		var ee error = enc.Encode(row)
		if nil != ee {
			return ee
		}
	}

	return errors.Join(enc.Flush(), bw.Flush())
}

func TypedToStdout[T any](
	typed iter.Seq2[T, error],
	schema ha.Schema,
) error {
	return TypedToWriter(typed, os.Stdout, schema)
}

func SchemaToTypedToStdout[T any](
	s ha.Schema,
) func(iter.Seq2[T, error]) util.Io[util.Void] {
	return func(typed iter.Seq2[T, error]) util.Io[util.Void] {
		return func(_ context.Context) (util.Void, error) {
			return util.Empty, TypedToStdout(typed, s)
		}
	}
}

func ParseSchema(schema string) (ha.Schema, error) {
	return ha.Parse(schema)
}

var SchemaStrToSchema func(string) util.Io[ha.Schema] = util.Lift(ParseSchema)
