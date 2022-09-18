// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/exchange/shared/repositories"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, repositories.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q repositories.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, repositories.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m repositories.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, repositories.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, repositories.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ repositories.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ repositories.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op repositories.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m repositories.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op repositories.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m repositories.Mutation) error {
		return Denyf("repositories/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The CoinQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CoinQueryRuleFunc func(context.Context, *repositories.CoinQuery) error

// EvalQuery return f(ctx, q).
func (f CoinQueryRuleFunc) EvalQuery(ctx context.Context, q repositories.Query) error {
	if q, ok := q.(*repositories.CoinQuery); ok {
		return f(ctx, q)
	}
	return Denyf("repositories/privacy: unexpected query type %T, expect *repositories.CoinQuery", q)
}

// The CoinMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CoinMutationRuleFunc func(context.Context, *repositories.CoinMutation) error

// EvalMutation calls f(ctx, m).
func (f CoinMutationRuleFunc) EvalMutation(ctx context.Context, m repositories.Mutation) error {
	if m, ok := m.(*repositories.CoinMutation); ok {
		return f(ctx, m)
	}
	return Denyf("repositories/privacy: unexpected mutation type %T, expect *repositories.CoinMutation", m)
}

// The ExchangeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ExchangeQueryRuleFunc func(context.Context, *repositories.ExchangeQuery) error

// EvalQuery return f(ctx, q).
func (f ExchangeQueryRuleFunc) EvalQuery(ctx context.Context, q repositories.Query) error {
	if q, ok := q.(*repositories.ExchangeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("repositories/privacy: unexpected query type %T, expect *repositories.ExchangeQuery", q)
}

// The ExchangeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ExchangeMutationRuleFunc func(context.Context, *repositories.ExchangeMutation) error

// EvalMutation calls f(ctx, m).
func (f ExchangeMutationRuleFunc) EvalMutation(ctx context.Context, m repositories.Mutation) error {
	if m, ok := m.(*repositories.ExchangeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("repositories/privacy: unexpected mutation type %T, expect *repositories.ExchangeMutation", m)
}

// The OutboxQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OutboxQueryRuleFunc func(context.Context, *repositories.OutboxQuery) error

// EvalQuery return f(ctx, q).
func (f OutboxQueryRuleFunc) EvalQuery(ctx context.Context, q repositories.Query) error {
	if q, ok := q.(*repositories.OutboxQuery); ok {
		return f(ctx, q)
	}
	return Denyf("repositories/privacy: unexpected query type %T, expect *repositories.OutboxQuery", q)
}

// The OutboxMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OutboxMutationRuleFunc func(context.Context, *repositories.OutboxMutation) error

// EvalMutation calls f(ctx, m).
func (f OutboxMutationRuleFunc) EvalMutation(ctx context.Context, m repositories.Mutation) error {
	if m, ok := m.(*repositories.OutboxMutation); ok {
		return f(ctx, m)
	}
	return Denyf("repositories/privacy: unexpected mutation type %T, expect *repositories.OutboxMutation", m)
}

// The TickerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TickerQueryRuleFunc func(context.Context, *repositories.TickerQuery) error

// EvalQuery return f(ctx, q).
func (f TickerQueryRuleFunc) EvalQuery(ctx context.Context, q repositories.Query) error {
	if q, ok := q.(*repositories.TickerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("repositories/privacy: unexpected query type %T, expect *repositories.TickerQuery", q)
}

// The TickerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TickerMutationRuleFunc func(context.Context, *repositories.TickerMutation) error

// EvalMutation calls f(ctx, m).
func (f TickerMutationRuleFunc) EvalMutation(ctx context.Context, m repositories.Mutation) error {
	if m, ok := m.(*repositories.TickerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("repositories/privacy: unexpected mutation type %T, expect *repositories.TickerMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q repositories.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m repositories.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q repositories.Query) (Filter, error) {
	switch q := q.(type) {
	case *repositories.CoinQuery:
		return q.Filter(), nil
	case *repositories.ExchangeQuery:
		return q.Filter(), nil
	case *repositories.OutboxQuery:
		return q.Filter(), nil
	case *repositories.TickerQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("repositories/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m repositories.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *repositories.CoinMutation:
		return m.Filter(), nil
	case *repositories.ExchangeMutation:
		return m.Filter(), nil
	case *repositories.OutboxMutation:
		return m.Filter(), nil
	case *repositories.TickerMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("repositories/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
