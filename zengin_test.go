package zengin

import (
	"testing"
)

func TestNew(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	if z == nil {
		t.Fatal("New() returned nil")
	}

	if len(z.banks) == 0 {
		t.Fatal("No banks loaded")
	}
}

func TestBank(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	tests := []struct {
		name     string
		code     string
		wantName string
		wantErr  bool
	}{
		{
			name:     "みずほ銀行",
			code:     "0001",
			wantName: "みずほ",
			wantErr:  false,
		},
		{
			name:     "三菱UFJ銀行",
			code:     "0005",
			wantName: "三菱ＵＦＪ",
			wantErr:  false,
		},
		{
			name:    "存在しない銀行",
			code:    "9999",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank, err := z.FindBank(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && bank.Name != tt.wantName {
				t.Errorf("Bank() bank.Name = %v, want %v", bank.Name, tt.wantName)
			}
		})
	}
}

func TestBranch(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	tests := []struct {
		name       string
		bankCode   string
		branchCode string
		wantName   string
		wantErr    bool
	}{
		{
			name:       "みずほ銀行 東京営業部",
			bankCode:   "0001",
			branchCode: "001",
			wantName:   "東京営業部",
			wantErr:    false,
		},
		{
			name:       "存在しない銀行",
			bankCode:   "9999",
			branchCode: "001",
			wantErr:    true,
		},
		{
			name:       "存在しない支店",
			bankCode:   "0001",
			branchCode: "999",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			branch, err := z.FindBranch(tt.bankCode, tt.branchCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Branch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && branch.Name != tt.wantName {
				t.Errorf("Branch() branch.Name = %v, want %v", branch.Name, tt.wantName)
			}
		})
	}
}

func TestAllBanks(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	banks := z.AllBanks()
	if len(banks) == 0 {
		t.Error("AllBanks() returned no banks")
	}
}

func TestAllBranches(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	tests := []struct {
		name       string
		bankCode   string
		wantMinLen int
		wantErr    bool
	}{
		{
			name:       "みずほ銀行の全支店",
			bankCode:   "0001",
			wantMinLen: 1,
			wantErr:    false,
		},
		{
			name:     "存在しない銀行",
			bankCode: "9999",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			branches, err := z.AllBranches(tt.bankCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllBranches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(branches) < tt.wantMinLen {
				t.Errorf("AllBranches() returned %d branches, want at least %d", len(branches), tt.wantMinLen)
			}
		})
	}
}

// TestBidirectionalReference tests the bidirectional relationship between Bank and Branch
func TestBidirectionalReference(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	// Get a bank
	bank, err := z.FindBank("0001")
	if err != nil {
		t.Fatalf("Bank() failed: %v", err)
	}

	// Check that bank has branches
	if len(bank.Branches) == 0 {
		t.Fatal("Bank has no branches")
	}

	// Get first branch
	var firstBranch *Branch
	for _, branch := range bank.Branches {
		firstBranch = branch
		break
	}

	// Test 1: Branch should have reference to Bank
	if firstBranch.Bank == nil {
		t.Fatal("Branch.Bank is nil")
	}

	// Test 2: Branch.Bank should point to the same bank
	if firstBranch.Bank.Code != bank.Code {
		t.Errorf("Branch.Bank.Code = %v, want %v", firstBranch.Bank.Code, bank.Code)
	}
	if firstBranch.Bank.Name != bank.Name {
		t.Errorf("Branch.Bank.Name = %v, want %v", firstBranch.Bank.Name, bank.Name)
	}

	// Test 3: Branch.Bank.Branches should contain the branch
	if _, exists := firstBranch.Bank.Branches[firstBranch.Code]; !exists {
		t.Errorf("Branch.Bank.Branches does not contain branch code %v", firstBranch.Code)
	}

	// Test 4: Get branch through Branch and verify its Bank reference
	branch, err := z.FindBranch("0001", "001")
	if err != nil {
		t.Fatalf("Branch() failed: %v", err)
	}

	if branch.Bank == nil {
		t.Fatal("Branch() returned branch with nil Bank")
	}

	if branch.Bank.Code != "0001" {
		t.Errorf("branch.Bank.Code = %v, want 0001", branch.Bank.Code)
	}
}

// Package-level function tests

func TestPackageAllBanks(t *testing.T) {
	banks := AllBanks()
	if len(banks) == 0 {
		t.Error("AllBanks() returned no banks")
	}

	// Verify it returns the same data as instance method
	z, _ := New()
	instanceBanks := z.AllBanks()
	if len(banks) != len(instanceBanks) {
		t.Errorf("AllBanks() returned %d banks, instance method returned %d", len(banks), len(instanceBanks))
	}
}

func TestPackageBank(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		wantName string
		wantErr  bool
	}{
		{
			name:     "みずほ銀行",
			code:     "0001",
			wantName: "みずほ",
			wantErr:  false,
		},
		{
			name:     "三菱UFJ銀行",
			code:     "0005",
			wantName: "三菱ＵＦＪ",
			wantErr:  false,
		},
		{
			name:    "存在しない銀行",
			code:    "9999",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank, err := FindBank(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && bank.Name != tt.wantName {
				t.Errorf("Bank() bank.Name = %v, want %v", bank.Name, tt.wantName)
			}
		})
	}
}

func TestPackageBranch(t *testing.T) {
	branch, err := FindBranch("0001", "001")
	if err != nil {
		t.Fatalf("Branch() failed: %v", err)
	}
	if branch.Name != "東京営業部" {
		t.Errorf("Branch() branch.Name = %v, want 東京営業部", branch.Name)
	}

	// Verify bidirectional reference
	if branch.Bank == nil {
		t.Fatal("Branch() returned branch with nil Bank")
	}
	if branch.Bank.Code != "0001" {
		t.Errorf("branch.Bank.Code = %v, want 0001", branch.Bank.Code)
	}
}

func TestPackageAllBranches(t *testing.T) {
	branches, err := AllBranches("0001")
	if err != nil {
		t.Fatalf("AllBranches() failed: %v", err)
	}
	if len(branches) == 0 {
		t.Error("AllBranches() returned no branches")
	}
}
