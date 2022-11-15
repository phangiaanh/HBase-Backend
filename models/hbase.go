package models

import "github.com/tsuna/gohbase"

var (
	hbaseClient gohbase.Client
)

func init() {
	hbaseClient = gohbase.NewClient("hbase-docker")
}
