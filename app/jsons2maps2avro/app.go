package app

import (
	"iter"

	ja "github.com/takanoriyanagitani/go-jsons2avro"
	util "github.com/takanoriyanagitani/go-jsons2avro/util"
)

type App struct {
	MapsSource util.Io[iter.Seq2[ja.JsonMap, error]]
	MapsSink   func(iter.Seq2[ja.JsonMap, error]) util.Io[util.Void]
}

func (a App) ToSourceToSink() util.Io[util.Void] {
	return util.Bind(
		a.MapsSource,
		a.MapsSink,
	)
}
