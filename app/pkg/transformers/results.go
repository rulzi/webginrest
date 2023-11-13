package transformers

// Result User Model
type Result struct {
	Status      int          `json:"status,omitempty"`
	Data        *interface{} `json:"data,omitempty"`
	UserMassage *UserMassage `json:"user_message,omitempty"`
	Error       *Error       `json:"error,omitempty"`
}

// UserMassage User Model
type UserMassage struct {
	Title   *string      `json:"title,omitempty"`
	Message *interface{} `json:"message,omitempty"`
}

// UserMassage User Model
type Error struct {
	Message *string      `json:"message,omitempty"`
	Reason  *interface{} `json:"reason,omitempty"`
}
