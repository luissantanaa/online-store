package models

type order struct {
	order_id   int
	client_id  int
	items      map[string]int
	ordered_at string
	sent       bool
	sent_at    string
}
