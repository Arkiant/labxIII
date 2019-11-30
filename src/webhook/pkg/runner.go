package pkg

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type RunnerFactory interface {
	NewRunner() Runner
}

type Runner interface {
	Run(context.Context, io.Reader) interface{}
}

func NewRunnerHandle(rf RunnerFactory) http.HandlerFunc {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rs := rf.NewRunner().Run(ctx, r.Body)
		response, err := json.Marshal(rs)
		if err != nil {
			println(err.Error())
			response = []byte(`{error:["` + err.Error() + `"]`)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("yo_que_Ser", "1231231")
		_, err = w.Write(response)
		if err != nil {
			println(err.Error())
			panic(err)
		}
	})
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, rs interface{}) {

}
