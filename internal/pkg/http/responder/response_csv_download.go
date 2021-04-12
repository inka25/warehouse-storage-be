package responder

import (
	//"encoding/csv"
	"fmt"
	"github.com/dnlo/struct2csv"
	"net/http"
)

func ResponseCSVDownload(rw http.ResponseWriter, filename string, data interface{}) error {

	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", filename))
	rw.Header().Set("Content-Type", "text/csv")
	rw.Header().Set("Transfer-Encoding", "chunked")

	writer := struct2csv.NewWriter(rw)
	if err := writer.WriteStructs(data); err != nil {
		return err
	}

	return nil
}
