// Code generated by ent, DO NOT EDIT.

package repositories

import (
	"github.com/omiga-group/omiga/src/order/shared/repositories/order"
	"github.com/omiga-group/omiga/src/order/shared/repositories/outbox"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
		Type: "Order",
		Fields: map[string]*sqlgraph.FieldSpec{
			order.FieldOrderID: {Type: field.TypeUUID, Column: order.FieldOrderID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   outbox.Table,
			Columns: outbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outbox.FieldID,
			},
		},
		Type: "Outbox",
		Fields: map[string]*sqlgraph.FieldSpec{
			outbox.FieldTimestamp:        {Type: field.TypeTime, Column: outbox.FieldTimestamp},
			outbox.FieldTopic:            {Type: field.TypeString, Column: outbox.FieldTopic},
			outbox.FieldKey:              {Type: field.TypeString, Column: outbox.FieldKey},
			outbox.FieldPayload:          {Type: field.TypeBytes, Column: outbox.FieldPayload},
			outbox.FieldHeaders:          {Type: field.TypeJSON, Column: outbox.FieldHeaders},
			outbox.FieldRetryCount:       {Type: field.TypeInt, Column: outbox.FieldRetryCount},
			outbox.FieldStatus:           {Type: field.TypeEnum, Column: outbox.FieldStatus},
			outbox.FieldLastRetry:        {Type: field.TypeTime, Column: outbox.FieldLastRetry},
			outbox.FieldProcessingErrors: {Type: field.TypeJSON, Column: outbox.FieldProcessingErrors},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (oq *OrderQuery) addPredicate(pred func(s *sql.Selector)) {
	oq.predicates = append(oq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OrderQuery builder.
func (oq *OrderQuery) Filter() *OrderFilter {
	return &OrderFilter{config: oq.config, predicateAdder: oq}
}

// addPredicate implements the predicateAdder interface.
func (m *OrderMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OrderMutation builder.
func (m *OrderMutation) Filter() *OrderFilter {
	return &OrderFilter{config: m.config, predicateAdder: m}
}

// OrderFilter provides a generic filtering capability at runtime for OrderQuery.
type OrderFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OrderFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *OrderFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(order.FieldID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *OrderFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(order.FieldOrderID))
}

// addPredicate implements the predicateAdder interface.
func (oq *OutboxQuery) addPredicate(pred func(s *sql.Selector)) {
	oq.predicates = append(oq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OutboxQuery builder.
func (oq *OutboxQuery) Filter() *OutboxFilter {
	return &OutboxFilter{config: oq.config, predicateAdder: oq}
}

// addPredicate implements the predicateAdder interface.
func (m *OutboxMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OutboxMutation builder.
func (m *OutboxMutation) Filter() *OutboxFilter {
	return &OutboxFilter{config: m.config, predicateAdder: m}
}

// OutboxFilter provides a generic filtering capability at runtime for OutboxQuery.
type OutboxFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OutboxFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *OutboxFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(outbox.FieldID))
}

// WhereTimestamp applies the entql time.Time predicate on the timestamp field.
func (f *OutboxFilter) WhereTimestamp(p entql.TimeP) {
	f.Where(p.Field(outbox.FieldTimestamp))
}

// WhereTopic applies the entql string predicate on the topic field.
func (f *OutboxFilter) WhereTopic(p entql.StringP) {
	f.Where(p.Field(outbox.FieldTopic))
}

// WhereKey applies the entql string predicate on the key field.
func (f *OutboxFilter) WhereKey(p entql.StringP) {
	f.Where(p.Field(outbox.FieldKey))
}

// WherePayload applies the entql []byte predicate on the payload field.
func (f *OutboxFilter) WherePayload(p entql.BytesP) {
	f.Where(p.Field(outbox.FieldPayload))
}

// WhereHeaders applies the entql json.RawMessage predicate on the headers field.
func (f *OutboxFilter) WhereHeaders(p entql.BytesP) {
	f.Where(p.Field(outbox.FieldHeaders))
}

// WhereRetryCount applies the entql int predicate on the retry_count field.
func (f *OutboxFilter) WhereRetryCount(p entql.IntP) {
	f.Where(p.Field(outbox.FieldRetryCount))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *OutboxFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(outbox.FieldStatus))
}

// WhereLastRetry applies the entql time.Time predicate on the last_retry field.
func (f *OutboxFilter) WhereLastRetry(p entql.TimeP) {
	f.Where(p.Field(outbox.FieldLastRetry))
}

// WhereProcessingErrors applies the entql json.RawMessage predicate on the processing_errors field.
func (f *OutboxFilter) WhereProcessingErrors(p entql.BytesP) {
	f.Where(p.Field(outbox.FieldProcessingErrors))
}
