package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/mocks/mock_mysql"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutocomplete(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMysql := mock_mysql.NewMockClienter(ctrl)

	ctx := context.Background()
	handler := NewAdminHandler(&Params{
		DB: mockMysql,
	})

	var tts = []struct {
		caseName     string
		params       *dtos.AutocompleteRequest
		expectations func()
		results      func(response *dtos.AutocompleteResponse, err error)
	}{
		{
			caseName: "db error",
			params: &dtos.AutocompleteRequest{
				Prefix: "lala",
			},
			expectations: func() {
				err := errors.New("any")

				mockMysql.EXPECT().Autocomplete(ctx, "lala").Return(nil, err)

			},
			results: func(response *dtos.AutocompleteResponse, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "no result found",
			params: &dtos.AutocompleteRequest{
				Prefix: "lala",
			},
			expectations: func() {

				mockMysql.EXPECT().Autocomplete(ctx, "lala").Return(nil, errs.ErrNoResultFound)

			},
			results: func(response *dtos.AutocompleteResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errs.ErrNoResultFound, err)
			},
		},
		{
			caseName: "successful",
			params: &dtos.AutocompleteRequest{
				Prefix: "lala",
			},
			expectations: func() {

				mockMysql.EXPECT().Autocomplete(ctx, "lala").Return([]string{"lalalu"}, nil)

			},
			results: func(response *dtos.AutocompleteResponse, err error) {
				assert.Equal(t, response.Result, []string{"lalalu"})
			},
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		tt.expectations()

		tt.results(handler.Autocomplete(ctx, tt.params))
	}

}
