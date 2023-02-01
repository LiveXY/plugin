package pdfer

type PDFConfig struct {
	Path   string `yaml:"path"`   // 路径
	Driver string `yaml:"driver"` // 驱动
}

type PDFer interface {
	Html2Pdf(htmlpath, pdfpath string) error
}
