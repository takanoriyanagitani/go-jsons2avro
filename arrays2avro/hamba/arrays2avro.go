package arrays2avro

import (
	"io"
	"iter"

	ha "github.com/hamba/avro/v2"

	ja "github.com/takanoriyanagitani/go-jsons2avro"
	util "github.com/takanoriyanagitani/go-jsons2avro/util"

	eh "github.com/takanoriyanagitani/go-jsons2avro/avro/enc/hamba"
)

func ArraysToWriter(
	arrays iter.Seq2[ja.JsonArray, error],
	writer io.Writer,
	schema ha.Schema,
) error {
	return eh.TypedToWriter(arrays, writer, schema)
}

func ArraysToStdout(
	arrays iter.Seq2[ja.JsonArray, error],
	schema ha.Schema,
) error {
	return eh.TypedToStdout(arrays, schema)
}

func SchemaToArraysToStdout(
	s ha.Schema,
) func(iter.Seq2[ja.JsonArray, error]) util.Io[util.Void] {
	return eh.SchemaToTypedToStdout[ja.JsonArray](s)
}

var FlatSchemaString util.Io[string] = util.Of(ja.FlatArraySchema)
var FlatSchema util.Io[ha.Schema] = util.Bind(
	FlatSchemaString,
	eh.SchemaStrToSchema,
)
