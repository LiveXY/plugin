package dber

type DBConfig struct {
	Path         string   `yaml:"path"` // 路径
	Driver       string   `yaml:"driver"`
	Sources      []string `yaml:"sources"`
	Replicas     []string `yaml:"replicas"`
	MaxIdleConns int      `yaml:"maxIdleConns"`
	MaxOpenConns int      `yaml:"maxOpenConns"`
}

type Dber interface {
	Init(logname string, cfg DBConfig, logger any) (any, error)
	IfNull() string
	ExAdd(field string, val any) any
	If() string
	GroupConcat(field string) string
	GetSlots() int
	GetCreateID(value any, table, pk string) int64
	ClobScan(clob *Clob, v any) error
}
