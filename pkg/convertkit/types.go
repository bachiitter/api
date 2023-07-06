package convertkit

import "time"

type AddSubscriberResponse struct {
	Subscription Subscription `json:"subscription"`
}

type Subscriber struct {
	ID int `json:"id"`
}

type Subscription struct {
	ID               int        `json:"id"`
	State            string     `json:"state"`
	CreatedAt        time.Time  `json:"created_at"`
	Source           *string    `json:"source"`
	Referrer         *string    `json:"referrer"`
	SubscribableID   int        `json:"subscribable_id"`
	SubscribableType string     `json:"subscribable_type"`
	Subscriber       Subscriber `json:"subscriber"`
}

type TotalsSubsResponse struct {
	TotalSubscriptions int            `json:"total_subscriptions"`
	Page               int            `json:"page"`
	TotalPages         int            `json:"total_pages"`
	Subscriptions      []Subscription `json:"subscriptions"`
}
