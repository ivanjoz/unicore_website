package core

// LimitAction names a family of rate limits. It is the first half of a limiter's key, the
// identifier being the second, which is what lets two features count on the same number without
// ever colliding: the contact form's IP 3405803821 and some future action's account 3405803821
// are different keys.
//
// Give a new feature its own action rather than reusing one, so the two never share a budget: a
// visitor sending a message while some other form is in flight from the same IP is not the abuse
// either limit is looking for.
//
// The values are written into DynamoDB partition keys, so they are storage, not just code. Never
// renumber one that is deployed — a rolling deploy would leave two versions disagreeing about
// what a number means, which is exactly the collision the namespace exists to prevent. Retire a
// value instead, and take the next free one.
type LimitAction uint16

const (
	// ActionContactByIP throttles the public contact form per client IP.
	ActionContactByIP LimitAction = 1
)
