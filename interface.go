package plugin

import (
	"github.com/livexy/plugin/cacher"
	"github.com/livexy/plugin/dber"
	"github.com/livexy/plugin/exceler"
	"github.com/livexy/plugin/ider"
	"github.com/livexy/plugin/jsoner"
	"github.com/livexy/plugin/osser"
	"github.com/livexy/plugin/pdfer"
	"github.com/livexy/plugin/smser"
	"github.com/livexy/plugin/streamer"
	"github.com/livexy/plugin/worder"

	"go.uber.org/zap"
)

type DBer interface {
	New() dber.Dber
}

type IDer interface {
	New(cfg ider.IDConfig) string
}

type Cacher interface {
	New(cfg cacher.CacheConfig, logger *zap.Logger) (cacher.Cacher, error)
}

type OSSer interface {
	New(cfg osser.OSSConfig) (osser.OSSer, error)
}

type Exceler interface {
	New(cfg exceler.ExcelConfig) exceler.Exceler
}

type Worder interface {
	New(cfg worder.WordConfig) worder.Worder
}

type Jsoner interface {
	New(cfg jsoner.JsonConfig) jsoner.Jsoner
}

type Streamer interface {
	New(cfg streamer.StreamConfig) streamer.Streamer
}

type SMSer interface {
	New(cfg smser.SMSConfig) smser.SMSer
}

type PDFer interface {
	New(cfg pdfer.PDFConfig) pdfer.PDFer
}
