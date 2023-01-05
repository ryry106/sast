package usecase

import "sast/presenter"

type previewOption struct {
	csvDir       string
	port         int
	templateName string
	startDate    string
}

func NewPreviewOption(csvDir string, port int, templateName string, startDate string) *previewOption {
	return &previewOption{csvDir: csvDir, port: port, templateName: templateName, startDate: startDate}
}

func (po *previewOption) Do() {
	s := presenter.NewServ(po.csvDir, po.port, po.templateName, po.startDate)
	s.Up()
}
