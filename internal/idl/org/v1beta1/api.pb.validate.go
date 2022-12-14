// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: org/v1beta1/api.proto

package orgv1beta1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateOrgRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateOrgRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrgRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrgRequestMultiError, or nil if none found.
func (m *CreateOrgRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrgRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return CreateOrgRequestMultiError(errors)
	}

	return nil
}

// CreateOrgRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOrgRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOrgRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrgRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrgRequestMultiError) AllErrors() []error { return m }

// CreateOrgRequestValidationError is the validation error returned by
// CreateOrgRequest.Validate if the designated constraints aren't met.
type CreateOrgRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrgRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrgRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrgRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrgRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrgRequestValidationError) ErrorName() string { return "CreateOrgRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateOrgRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrgRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrgRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrgRequestValidationError{}

// Validate checks the field values on CreateOrgResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateOrgResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrgResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrgResponseMultiError, or nil if none found.
func (m *CreateOrgResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrgResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateOrgResponseMultiError(errors)
	}

	return nil
}

// CreateOrgResponseMultiError is an error wrapping multiple validation errors
// returned by CreateOrgResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateOrgResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrgResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrgResponseMultiError) AllErrors() []error { return m }

// CreateOrgResponseValidationError is the validation error returned by
// CreateOrgResponse.Validate if the designated constraints aren't met.
type CreateOrgResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrgResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrgResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrgResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrgResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrgResponseValidationError) ErrorName() string {
	return "CreateOrgResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrgResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrgResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrgResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrgResponseValidationError{}
