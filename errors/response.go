package errors

import (
	"errors"
	"net/http"
)

// Response error response
type Response struct {
	Error       error
	ErrorCode   int
	Description string
	URI         string
	StatusCode  int
	Header      http.Header
}

// NewResponse create the response pointer
func NewResponse(err error, statusCode int) *Response {
	return &Response{
		Error:      err,
		StatusCode: statusCode,
	}
}

// SetHeader sets the header entries associated with key to
// the single element value.
func (r *Response) SetHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)
}

// https://tools.ietf.org/html/rfc6749#section-5.2
var (
	ErrInvalidRequest                 = errors.New("invalid_request")
	ErrUnauthorizedClient             = errors.New("unauthorized_client")
	ErrAccessDenied                   = errors.New("access_denied")
	ErrUnsupportedResponseType        = errors.New("unsupported_response_type")
	ErrInvalidScope                   = errors.New("invalid_scope")
	ErrServerError                    = errors.New("server_error")
	ErrTemporarilyUnavailable         = errors.New("temporarily_unavailable")
	ErrInvalidClient                  = errors.New("invalid_client")
	ErrInvalidClient1                 = errors.New("invalid_client1")
	ErrInvalidClient2                 = errors.New("invalid_client2")
	ErrInvalidClient3                 = errors.New("invalid_client3")
	ErrInvalidClient4                 = errors.New("invalid_client4")
	ErrInvalidClient5                 = errors.New("invalid_client5")
	ErrInvalidClient6                 = errors.New("invalid_client6")
	ErrInvalidGrant                   = errors.New("invalid_grant")
	ErrUnsupportedGrantType           = errors.New("unsupported_grant_type")
	ErrCodeChallengeRquired           = errors.New("invalid_request")
	ErrUnsupportedCodeChallengeMethod = errors.New("invalid_request")
	ErrInvalidCodeChallengeLen        = errors.New("invalid_request")
)

// Descriptions error description
var Descriptions = map[error]string{
	ErrInvalidRequest:                 "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed",
	ErrUnauthorizedClient:             "The client is not authorized to request an authorization code using this method",
	ErrAccessDenied:                   "The resource owner or authorization server denied the request",
	ErrUnsupportedResponseType:        "The authorization server does not support obtaining an authorization code using this method",
	ErrInvalidScope:                   "The requested scope is invalid, unknown, or malformed",
	ErrServerError:                    "The authorization server encountered an unexpected condition that prevented it from fulfilling the request",
	ErrTemporarilyUnavailable:         "The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server",
	ErrInvalidClient:                  "Client authentication failed",
	ErrInvalidClient1:                 "Client authentication failed1",
	ErrInvalidClient2:                 "Client authentication failed2",
	ErrInvalidClient3:                 "Client authentication failed3",
	ErrInvalidClient4:                 "Client authentication failed4",
	ErrInvalidClient5:                 "Client authentication failed5",
	ErrInvalidClient6:                 "Client authentication failed6",
	ErrInvalidGrant:                   "The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client",
	ErrUnsupportedGrantType:           "The authorization grant type is not supported by the authorization server",
	ErrCodeChallengeRquired:           "PKCE is required. code_challenge is missing",
	ErrUnsupportedCodeChallengeMethod: "Selected code_challenge_method not supported",
	ErrInvalidCodeChallengeLen:        "Code challenge length must be between 43 and 128 charachters long",
}

// StatusCodes response error HTTP status code
var StatusCodes = map[error]int{
	ErrInvalidRequest:                 400,
	ErrUnauthorizedClient:             401,
	ErrAccessDenied:                   403,
	ErrUnsupportedResponseType:        401,
	ErrInvalidScope:                   400,
	ErrServerError:                    500,
	ErrTemporarilyUnavailable:         503,
	ErrInvalidClient:                  401,
	ErrInvalidClient1:                 401,
	ErrInvalidClient2:                 401,
	ErrInvalidClient3:                 401,
	ErrInvalidClient4:                 401,
	ErrInvalidClient5:                 401,
	ErrInvalidClient6:                 401,
	ErrInvalidGrant:                   401,
	ErrUnsupportedGrantType:           401,
	ErrCodeChallengeRquired:           400,
	ErrUnsupportedCodeChallengeMethod: 400,
	ErrInvalidCodeChallengeLen:        400,
}
