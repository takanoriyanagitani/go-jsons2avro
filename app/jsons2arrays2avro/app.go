package app

import (
	"iter"

	ja "github.com/takanoriyanagitani/go-jsons2avro"
	util "github.com/takanoriyanagitani/go-jsons2avro/util"
)

type App struct {
	ArraysSource util.Io[iter.Seq2[ja.JsonArray, error]]
	ArraysSink   func(iter.Seq2[ja.JsonArray, error]) util.Io[util.Void]
}

func (a App) ToSourceToSink() util.Io[util.Void] {
	return util.Bind(
		a.ArraysSource,
		a.ArraysSink,
	)
}
