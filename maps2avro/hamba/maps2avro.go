package maps2avro

import (
	"io"
	"iter"

	ha "github.com/hamba/avro/v2"

	ja "github.com/takanoriyanagitani/go-jsons2avro"
	util "github.com/takanoriyanagitani/go-jsons2avro/util"

	eh "github.com/takanoriyanagitani/go-jsons2avro/avro/enc/hamba"
)

func MapsToWriter(
	maps iter.Seq2[ja.JsonMap, error],
	writer io.Writer,
	schema ha.Schema,
) error {
	return eh.TypedToWriter(maps, writer, schema)
}

func MapsToStdout(
	maps iter.Seq2[ja.JsonMap, error],
	schema ha.Schema,
) error {
	return eh.TypedToStdout(maps, schema)
}

func SchemaToMapsToStdout(
	s ha.Schema,
) func(iter.Seq2[ja.JsonMap, error]) util.Io[util.Void] {
	return eh.SchemaToTypedToStdout[ja.JsonMap](s)
}

var FlatSchemaString util.Io[string] = util.Of(ja.FlatMapSchema)
var FlatSchema util.Io[ha.Schema] = util.Bind(
	FlatSchemaString,
	eh.SchemaStrToSchema,
)
