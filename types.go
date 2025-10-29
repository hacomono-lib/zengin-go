package zengin

// Bank represents a bank information
type Bank struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Kana string `json:"kana"`
	Hira string `json:"hira"`
	Roma string `json:"roma"`
}

// Branch represents a branch information
type Branch struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Kana string `json:"kana"`
	Hira string `json:"hira"`
	Roma string `json:"roma"`
}

// BankWithBranches represents a bank with its branches
type BankWithBranches struct {
	Bank
	Branches map[string]*Branch `json:"branches"`
}
