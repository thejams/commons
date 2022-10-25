package commons

import (
	"testing"
)

func Test_DecodeError(t *testing.T) {
	t.Run("should return the correct http error messages and codes", func(t *testing.T) {
		// return "", &errors.InternalError{Message: err.Error()}
		// s, ok := i.(string)
		type cases struct {
			err    error
			msg    string
			status int
		}

		testCases := []cases{
			{&InternalError{Message: "this is an internal error"}, "InternalServerError: this is an internal error", 500},
			{&NotFoundError{Message: "this is an not found error"}, "NotFound: this is an not found error", 404},
			{&BadRequestError{Message: "this is an bad request error"}, "BadRequest: this is an bad request error", 400},
		}

		for _, x := range testCases {
			status, body := DecodeError(x.err)
			if status != x.status {
				t.Error("Expected", x.status, "Got", status)
			}
			if body.Message != x.msg {
				t.Error("Expected", x.msg, "Got", body.Message)
			}
		}
	})

}
