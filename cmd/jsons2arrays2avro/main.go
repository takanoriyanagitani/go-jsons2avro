package main

import (
	"context"
	"iter"
	"log"

	ha "github.com/hamba/avro/v2"

	ja "github.com/takanoriyanagitani/go-jsons2avro"
	util "github.com/takanoriyanagitani/go-jsons2avro/util"

	ah "github.com/takanoriyanagitani/go-jsons2avro/arrays2avro/hamba"
	js "github.com/takanoriyanagitani/go-jsons2avro/jsons2array/std"

	ap "github.com/takanoriyanagitani/go-jsons2avro/app/jsons2arrays2avro"
)

var arraysSource util.Io[iter.Seq2[ja.JsonArray, error]] = js.
	StdinToJsonArrays()

var flatSchema util.Io[ha.Schema] = ah.FlatSchema

var app util.Io[ap.App] = util.Bind(
	flatSchema,
	func(s ha.Schema) util.Io[ap.App] {
		return func(_ context.Context) (ap.App, error) {
			return ap.App{
				ArraysSource: arraysSource,
				ArraysSink:   ah.SchemaToArraysToStdout(s),
			}, nil
		}
	},
)

var source2sink util.Io[util.Void] = util.Bind(
	app,
	func(a ap.App) util.Io[util.Void] { return a.ToSourceToSink() },
)

func sub(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, e := source2sink(ctx)
	return e
}

func main() {
	e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
