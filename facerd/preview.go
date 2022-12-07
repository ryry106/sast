package facerd

import (
	"ryry/prev-bdc/preview"
)

func Run(csvPath string) {
  // todo ファイルのチェックなど



	serv := &preview.Serv{CsvPath: csvPath}
	serv.Up()
}
