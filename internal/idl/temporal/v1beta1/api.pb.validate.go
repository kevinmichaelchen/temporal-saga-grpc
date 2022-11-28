// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: temporal/v1beta1/api.proto

package temporalv1beta1

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

// Validate checks the field values on WorkflowOptions with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WorkflowOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkflowOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkflowOptionsMultiError, or nil if none found.
func (m *WorkflowOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkflowOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WorkflowId

	if len(errors) > 0 {
		return WorkflowOptionsMultiError(errors)
	}

	return nil
}

// WorkflowOptionsMultiError is an error wrapping multiple validation errors
// returned by WorkflowOptions.ValidateAll() if the designated constraints
// aren't met.
type WorkflowOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkflowOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkflowOptionsMultiError) AllErrors() []error { return m }

// WorkflowOptionsValidationError is the validation error returned by
// WorkflowOptions.Validate if the designated constraints aren't met.
type WorkflowOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowOptionsValidationError) ErrorName() string { return "WorkflowOptionsValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowOptionsValidationError{}

// Validate checks the field values on CreateLicenseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateLicenseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLicenseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLicenseRequestMultiError, or nil if none found.
func (m *CreateLicenseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLicenseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWorkflowOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateLicenseRequestValidationError{
					field:  "WorkflowOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateLicenseRequestValidationError{
					field:  "WorkflowOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWorkflowOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateLicenseRequestValidationError{
				field:  "WorkflowOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLicense()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateLicenseRequestValidationError{
					field:  "License",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateLicenseRequestValidationError{
					field:  "License",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLicense()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateLicenseRequestValidationError{
				field:  "License",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOrg()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateLicenseRequestValidationError{
					field:  "Org",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateLicenseRequestValidationError{
					field:  "Org",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrg()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateLicenseRequestValidationError{
				field:  "Org",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetProfile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateLicenseRequestValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateLicenseRequestValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateLicenseRequestValidationError{
				field:  "Profile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateLicenseRequestMultiError(errors)
	}

	return nil
}

// CreateLicenseRequestMultiError is an error wrapping multiple validation
// errors returned by CreateLicenseRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateLicenseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLicenseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLicenseRequestMultiError) AllErrors() []error { return m }

// CreateLicenseRequestValidationError is the validation error returned by
// CreateLicenseRequest.Validate if the designated constraints aren't met.
type CreateLicenseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLicenseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLicenseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLicenseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLicenseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLicenseRequestValidationError) ErrorName() string {
	return "CreateLicenseRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateLicenseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLicenseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLicenseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLicenseRequestValidationError{}

// Validate checks the field values on CreateLicenseResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateLicenseResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLicenseResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLicenseResponseMultiError, or nil if none found.
func (m *CreateLicenseResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLicenseResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateLicenseResponseMultiError(errors)
	}

	return nil
}

// CreateLicenseResponseMultiError is an error wrapping multiple validation
// errors returned by CreateLicenseResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateLicenseResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLicenseResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLicenseResponseMultiError) AllErrors() []error { return m }

// CreateLicenseResponseValidationError is the validation error returned by
// CreateLicenseResponse.Validate if the designated constraints aren't met.
type CreateLicenseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLicenseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLicenseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLicenseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLicenseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLicenseResponseValidationError) ErrorName() string {
	return "CreateLicenseResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateLicenseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLicenseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLicenseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLicenseResponseValidationError{}

// Validate checks the field values on License with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *License) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on License with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LicenseMultiError, or nil if none found.
func (m *License) ValidateAll() error {
	return m.validate(true)
}

func (m *License) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return LicenseMultiError(errors)
	}

	return nil
}

// LicenseMultiError is an error wrapping multiple validation errors returned
// by License.ValidateAll() if the designated constraints aren't met.
type LicenseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LicenseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LicenseMultiError) AllErrors() []error { return m }

// LicenseValidationError is the validation error returned by License.Validate
// if the designated constraints aren't met.
type LicenseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LicenseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LicenseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LicenseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LicenseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LicenseValidationError) ErrorName() string { return "LicenseValidationError" }

// Error satisfies the builtin error interface
func (e LicenseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLicense.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LicenseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LicenseValidationError{}

// Validate checks the field values on Org with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Org) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Org with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OrgMultiError, or nil if none found.
func (m *Org) ValidateAll() error {
	return m.validate(true)
}

func (m *Org) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return OrgMultiError(errors)
	}

	return nil
}

// OrgMultiError is an error wrapping multiple validation errors returned by
// Org.ValidateAll() if the designated constraints aren't met.
type OrgMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrgMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrgMultiError) AllErrors() []error { return m }

// OrgValidationError is the validation error returned by Org.Validate if the
// designated constraints aren't met.
type OrgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrgValidationError) ErrorName() string { return "OrgValidationError" }

// Error satisfies the builtin error interface
func (e OrgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrgValidationError{}

// Validate checks the field values on Profile with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Profile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Profile with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProfileMultiError, or nil if none found.
func (m *Profile) ValidateAll() error {
	return m.validate(true)
}

func (m *Profile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return ProfileMultiError(errors)
	}

	return nil
}

// ProfileMultiError is an error wrapping multiple validation errors returned
// by Profile.ValidateAll() if the designated constraints aren't met.
type ProfileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProfileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProfileMultiError) AllErrors() []error { return m }

// ProfileValidationError is the validation error returned by Profile.Validate
// if the designated constraints aren't met.
type ProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileValidationError) ErrorName() string { return "ProfileValidationError" }

// Error satisfies the builtin error interface
func (e ProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileValidationError{}