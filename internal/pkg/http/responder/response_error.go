package responder

import (
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"encoding/json"
	"errors"
	"net/http"
)

func CheckErrorType(err error, target ...error) bool {
	for _, t := range target {
		if errors.Is(err, t) {
			return true
		}
	}
	return false
}

func errToHttpFormat(err error) (int, interface{}) {
	switch {
	case CheckErrorType(err, errs.ErrAuth):
		return http.StatusUnauthorized, CommonResponse{
			Status:      1,
			Description: "unauthorized request",
		}
	case CheckErrorType(
		err,
		errs.ErrEmptyBodyRequest,
		errs.ErrInvalidCountry,
		errs.ErrInvalidRequest,
		errs.ErrNoResultFound,
	):
		return http.StatusBadRequest, CommonResponse{
			Status:      1,
			Description: err.Error(),
		}
	default:
		return http.StatusInternalServerError, CommonResponse{
			Status:      1,
			Description: "internal server error",
		}
	}
}

func ResponseError(w http.ResponseWriter, err error, additionalData ...interface{}) error {
	httpCode, resp := errToHttpFormat(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	advanceResponse, isAdvanceCommonResponseType := resp.(AdvanceCommonResponse)
	if len(additionalData) > 0 && isAdvanceCommonResponseType {
		advanceResponse.Data = additionalData[0]
		return json.NewEncoder(w).Encode(advanceResponse)
	}

	return json.NewEncoder(w).Encode(resp)
}
