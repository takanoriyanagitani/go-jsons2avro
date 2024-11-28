package jsons2avro

import (
	_ "embed"
)

type JsonArray []any

//go:embed flat-array.avsc
var FlatArraySchema string

//go:embed flat-map.avsc
var FlatMapSchema string

type JsonMap map[string]any
