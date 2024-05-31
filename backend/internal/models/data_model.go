package models

// PingModel represents the ping model.
type PingModel struct{}

// Ping method returns "pong" message.
func (m *PingModel) Ping() string {
	return "pong"
}
