// Code generated by ent, DO NOT EDIT.

package outbox

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// Topic applies equality check predicate on the "topic" field. It's identical to TopicEQ.
func Topic(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopic), v))
	})
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// Payload applies equality check predicate on the "payload" field. It's identical to PayloadEQ.
func Payload(v []byte) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPayload), v))
	})
}

// RetryCount applies equality check predicate on the "retry_count" field. It's identical to RetryCountEQ.
func RetryCount(v int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRetryCount), v))
	})
}

// LastRetry applies equality check predicate on the "last_retry" field. It's identical to LastRetryEQ.
func LastRetry(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastRetry), v))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// TopicEQ applies the EQ predicate on the "topic" field.
func TopicEQ(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopic), v))
	})
}

// TopicNEQ applies the NEQ predicate on the "topic" field.
func TopicNEQ(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTopic), v))
	})
}

// TopicIn applies the In predicate on the "topic" field.
func TopicIn(vs ...string) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTopic), v...))
	})
}

// TopicNotIn applies the NotIn predicate on the "topic" field.
func TopicNotIn(vs ...string) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTopic), v...))
	})
}

// TopicGT applies the GT predicate on the "topic" field.
func TopicGT(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTopic), v))
	})
}

// TopicGTE applies the GTE predicate on the "topic" field.
func TopicGTE(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTopic), v))
	})
}

// TopicLT applies the LT predicate on the "topic" field.
func TopicLT(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTopic), v))
	})
}

// TopicLTE applies the LTE predicate on the "topic" field.
func TopicLTE(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTopic), v))
	})
}

// TopicContains applies the Contains predicate on the "topic" field.
func TopicContains(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTopic), v))
	})
}

// TopicHasPrefix applies the HasPrefix predicate on the "topic" field.
func TopicHasPrefix(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTopic), v))
	})
}

// TopicHasSuffix applies the HasSuffix predicate on the "topic" field.
func TopicHasSuffix(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTopic), v))
	})
}

// TopicEqualFold applies the EqualFold predicate on the "topic" field.
func TopicEqualFold(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTopic), v))
	})
}

// TopicContainsFold applies the ContainsFold predicate on the "topic" field.
func TopicContainsFold(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTopic), v))
	})
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKey), v))
	})
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldKey), v...))
	})
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldKey), v...))
	})
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKey), v))
	})
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKey), v))
	})
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKey), v))
	})
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKey), v))
	})
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKey), v))
	})
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKey), v))
	})
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKey), v))
	})
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKey), v))
	})
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKey), v))
	})
}

// PayloadEQ applies the EQ predicate on the "payload" field.
func PayloadEQ(v []byte) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPayload), v))
	})
}

// PayloadNEQ applies the NEQ predicate on the "payload" field.
func PayloadNEQ(v []byte) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPayload), v))
	})
}

// PayloadIn applies the In predicate on the "payload" field.
func PayloadIn(vs ...[]byte) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPayload), v...))
	})
}

// PayloadNotIn applies the NotIn predicate on the "payload" field.
func PayloadNotIn(vs ...[]byte) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPayload), v...))
	})
}

// PayloadGT applies the GT predicate on the "payload" field.
func PayloadGT(v []byte) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPayload), v))
	})
}

// PayloadGTE applies the GTE predicate on the "payload" field.
func PayloadGTE(v []byte) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPayload), v))
	})
}

// PayloadLT applies the LT predicate on the "payload" field.
func PayloadLT(v []byte) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPayload), v))
	})
}

// PayloadLTE applies the LTE predicate on the "payload" field.
func PayloadLTE(v []byte) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPayload), v))
	})
}

// RetryCountEQ applies the EQ predicate on the "retry_count" field.
func RetryCountEQ(v int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRetryCount), v))
	})
}

// RetryCountNEQ applies the NEQ predicate on the "retry_count" field.
func RetryCountNEQ(v int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRetryCount), v))
	})
}

// RetryCountIn applies the In predicate on the "retry_count" field.
func RetryCountIn(vs ...int) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRetryCount), v...))
	})
}

// RetryCountNotIn applies the NotIn predicate on the "retry_count" field.
func RetryCountNotIn(vs ...int) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRetryCount), v...))
	})
}

// RetryCountGT applies the GT predicate on the "retry_count" field.
func RetryCountGT(v int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRetryCount), v))
	})
}

// RetryCountGTE applies the GTE predicate on the "retry_count" field.
func RetryCountGTE(v int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRetryCount), v))
	})
}

// RetryCountLT applies the LT predicate on the "retry_count" field.
func RetryCountLT(v int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRetryCount), v))
	})
}

// RetryCountLTE applies the LTE predicate on the "retry_count" field.
func RetryCountLTE(v int) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRetryCount), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// LastRetryEQ applies the EQ predicate on the "last_retry" field.
func LastRetryEQ(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastRetry), v))
	})
}

// LastRetryNEQ applies the NEQ predicate on the "last_retry" field.
func LastRetryNEQ(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastRetry), v))
	})
}

// LastRetryIn applies the In predicate on the "last_retry" field.
func LastRetryIn(vs ...time.Time) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastRetry), v...))
	})
}

// LastRetryNotIn applies the NotIn predicate on the "last_retry" field.
func LastRetryNotIn(vs ...time.Time) predicate.Outbox {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastRetry), v...))
	})
}

// LastRetryGT applies the GT predicate on the "last_retry" field.
func LastRetryGT(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastRetry), v))
	})
}

// LastRetryGTE applies the GTE predicate on the "last_retry" field.
func LastRetryGTE(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastRetry), v))
	})
}

// LastRetryLT applies the LT predicate on the "last_retry" field.
func LastRetryLT(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastRetry), v))
	})
}

// LastRetryLTE applies the LTE predicate on the "last_retry" field.
func LastRetryLTE(v time.Time) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastRetry), v))
	})
}

// LastRetryIsNil applies the IsNil predicate on the "last_retry" field.
func LastRetryIsNil() predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastRetry)))
	})
}

// LastRetryNotNil applies the NotNil predicate on the "last_retry" field.
func LastRetryNotNil() predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastRetry)))
	})
}

// ProcessingErrorsIsNil applies the IsNil predicate on the "processing_errors" field.
func ProcessingErrorsIsNil() predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldProcessingErrors)))
	})
}

// ProcessingErrorsNotNil applies the NotNil predicate on the "processing_errors" field.
func ProcessingErrorsNotNil() predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldProcessingErrors)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Outbox) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Outbox) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Outbox) predicate.Outbox {
	return predicate.Outbox(func(s *sql.Selector) {
		p(s.Not())
	})
}
