package output

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dotnetmentor/trail-digger/trail"
)

const (
	FieldAwsRegion          string = "AwsRegion"
	FieldErrorCode          string = "ErrorCode"
	FieldErrorMessage       string = "ErrorMessage"
	FieldEventID            string = "EventID"
	FieldEventName          string = "EventName"
	FieldEventSource        string = "EventSource"
	FieldEventTime          string = "EventTime"
	FieldEventType          string = "EventType"
	FieldRecipientAccountID string = "RecipientAccountID"
	FieldUserIdentityArn    string = "UserIdentity.Arn"
)

var (
	FieldAccessors map[string]FieldAccessor = map[string]FieldAccessor{}
)

type FieldAccessor func(r *trail.Record) string

type RecordFieldResolver struct {
	record         *trail.Record
	reflectedValue reflect.Value
}

func NewRecordValueAccessor(r *trail.Record) RecordFieldResolver {
	return RecordFieldResolver{
		record:         r,
		reflectedValue: reflect.ValueOf(r),
	}
}

func (r RecordFieldResolver) Value(field string) string {
	if accessor, ok := FieldAccessors[field]; ok {
		return accessor(r.record)
	}

	if strings.Contains(field, ".") {
		parts := strings.Split(field, ".")
		value := r.reflectedValue
		for _, p := range parts {
			value = reflect.Indirect(value).FieldByName(p)
		}
		return fmt.Sprintf("%s", value.Interface())
	}

	value := reflect.Indirect(r.reflectedValue).FieldByName(field)
	return fmt.Sprintf("%s", value.Interface())
}

func init() {
	FieldAccessors[FieldAwsRegion] = func(r *trail.Record) string {
		return r.AwsRegion
	}
	FieldAccessors[FieldErrorCode] = func(r *trail.Record) string {
		return r.ErrorCode
	}
	FieldAccessors[FieldErrorMessage] = func(r *trail.Record) string {
		return r.ErrorMessage
	}
	FieldAccessors[FieldEventID] = func(r *trail.Record) string {
		return r.EventID
	}
	FieldAccessors[FieldEventName] = func(r *trail.Record) string {
		return r.EventName
	}
	FieldAccessors[FieldEventSource] = func(r *trail.Record) string {
		return r.EventSource
	}
	FieldAccessors[FieldEventTime] = func(r *trail.Record) string {
		return r.EventTime.String()
	}
	FieldAccessors[FieldEventType] = func(r *trail.Record) string {
		return r.EventType
	}
	FieldAccessors[FieldRecipientAccountID] = func(r *trail.Record) string {
		return r.RecipientAccountID
	}
	FieldAccessors[FieldUserIdentityArn] = func(r *trail.Record) string {
		return r.UserIdentity.Arn
	}
}
