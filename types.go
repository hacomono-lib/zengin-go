package zengin

// Bank represents a bank information
type Bank struct {
	Code     string             `json:"code"`
	Name     string             `json:"name"`
	Kana     string             `json:"kana"`
	Hira     string             `json:"hira"`
	Roma     string             `json:"roma"`
	Branches map[string]*Branch `json:"branches,omitempty"`
}

// Branch represents a branch information
type Branch struct {
	Bank *Bank  `json:"-"` // Reference to parent bank (not serialized to avoid circular reference)
	Code string `json:"code"`
	Name string `json:"name"`
	Kana string `json:"kana"`
	Hira string `json:"hira"`
	Roma string `json:"roma"`
}
