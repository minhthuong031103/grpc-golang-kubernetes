// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: order/order.proto

package order

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

// Validate checks the field values on Order with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Order) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Order with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OrderMultiError, or nil if none found.
func (m *Order) ValidateAll() error {
	return m.validate(true)
}

func (m *Order) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for Total

	// no validation rules for OrderDate

	// no validation rules for Email

	// no validation rules for ShippingAddress

	// no validation rules for Status

	if len(errors) > 0 {
		return OrderMultiError(errors)
	}

	return nil
}

// OrderMultiError is an error wrapping multiple validation errors returned by
// Order.ValidateAll() if the designated constraints aren't met.
type OrderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderMultiError) AllErrors() []error { return m }

// OrderValidationError is the validation error returned by Order.Validate if
// the designated constraints aren't met.
type OrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderValidationError) ErrorName() string { return "OrderValidationError" }

// Error satisfies the builtin error interface
func (e OrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderValidationError{}

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on PayloadWithSingleOrder with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PayloadWithSingleOrder) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayloadWithSingleOrder with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PayloadWithSingleOrderMultiError, or nil if none found.
func (m *PayloadWithSingleOrder) ValidateAll() error {
	return m.validate(true)
}

func (m *PayloadWithSingleOrder) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOrder()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PayloadWithSingleOrderValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PayloadWithSingleOrderValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PayloadWithSingleOrderValidationError{
				field:  "Order",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PayloadWithSingleOrderMultiError(errors)
	}

	return nil
}

// PayloadWithSingleOrderMultiError is an error wrapping multiple validation
// errors returned by PayloadWithSingleOrder.ValidateAll() if the designated
// constraints aren't met.
type PayloadWithSingleOrderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayloadWithSingleOrderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayloadWithSingleOrderMultiError) AllErrors() []error { return m }

// PayloadWithSingleOrderValidationError is the validation error returned by
// PayloadWithSingleOrder.Validate if the designated constraints aren't met.
type PayloadWithSingleOrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayloadWithSingleOrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayloadWithSingleOrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayloadWithSingleOrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayloadWithSingleOrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayloadWithSingleOrderValidationError) ErrorName() string {
	return "PayloadWithSingleOrderValidationError"
}

// Error satisfies the builtin error interface
func (e PayloadWithSingleOrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayloadWithSingleOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayloadWithSingleOrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayloadWithSingleOrderValidationError{}

// Validate checks the field values on PayloadWithOrderID with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PayloadWithOrderID) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayloadWithOrderID with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PayloadWithOrderIDMultiError, or nil if none found.
func (m *PayloadWithOrderID) ValidateAll() error {
	return m.validate(true)
}

func (m *PayloadWithOrderID) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return PayloadWithOrderIDMultiError(errors)
	}

	return nil
}

// PayloadWithOrderIDMultiError is an error wrapping multiple validation errors
// returned by PayloadWithOrderID.ValidateAll() if the designated constraints
// aren't met.
type PayloadWithOrderIDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayloadWithOrderIDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayloadWithOrderIDMultiError) AllErrors() []error { return m }

// PayloadWithOrderIDValidationError is the validation error returned by
// PayloadWithOrderID.Validate if the designated constraints aren't met.
type PayloadWithOrderIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayloadWithOrderIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayloadWithOrderIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayloadWithOrderIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayloadWithOrderIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayloadWithOrderIDValidationError) ErrorName() string {
	return "PayloadWithOrderIDValidationError"
}

// Error satisfies the builtin error interface
func (e PayloadWithOrderIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayloadWithOrderID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayloadWithOrderIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayloadWithOrderIDValidationError{}

// Validate checks the field values on PayloadWithMultipleOrders with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PayloadWithMultipleOrders) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayloadWithMultipleOrders with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PayloadWithMultipleOrdersMultiError, or nil if none found.
func (m *PayloadWithMultipleOrders) ValidateAll() error {
	return m.validate(true)
}

func (m *PayloadWithMultipleOrders) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetOrders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PayloadWithMultipleOrdersValidationError{
						field:  fmt.Sprintf("Orders[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PayloadWithMultipleOrdersValidationError{
						field:  fmt.Sprintf("Orders[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PayloadWithMultipleOrdersValidationError{
					field:  fmt.Sprintf("Orders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PayloadWithMultipleOrdersMultiError(errors)
	}

	return nil
}

// PayloadWithMultipleOrdersMultiError is an error wrapping multiple validation
// errors returned by PayloadWithMultipleOrders.ValidateAll() if the
// designated constraints aren't met.
type PayloadWithMultipleOrdersMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayloadWithMultipleOrdersMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayloadWithMultipleOrdersMultiError) AllErrors() []error { return m }

// PayloadWithMultipleOrdersValidationError is the validation error returned by
// PayloadWithMultipleOrders.Validate if the designated constraints aren't met.
type PayloadWithMultipleOrdersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayloadWithMultipleOrdersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayloadWithMultipleOrdersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayloadWithMultipleOrdersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayloadWithMultipleOrdersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayloadWithMultipleOrdersValidationError) ErrorName() string {
	return "PayloadWithMultipleOrdersValidationError"
}

// Error satisfies the builtin error interface
func (e PayloadWithMultipleOrdersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayloadWithMultipleOrders.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayloadWithMultipleOrdersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayloadWithMultipleOrdersValidationError{}