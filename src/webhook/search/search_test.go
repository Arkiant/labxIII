package search

import (
	"context"
	"io"
	"reflect"
	"testing"

	"github.com/Arkiant/labxIII/src/conversation"
	"github.com/Arkiant/labxIII/src/webhook/transaction"
)

func TestSearchService_Run(t *testing.T) {
	type fields struct {
		transactioner transaction.Service
		rq            conversation.Criteria
	}
	type args struct {
		ctx    context.Context
		bodyRQ io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SearchService{
				transactioner: tt.fields.transactioner,
				rq:            tt.fields.rq,
			}
			if got := s.Run(tt.args.ctx, tt.args.bodyRQ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchService.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
