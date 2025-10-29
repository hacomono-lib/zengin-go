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

func TestGetBank(t *testing.T) {
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
			bank, err := z.GetBank(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && bank.Name != tt.wantName {
				t.Errorf("GetBank() bank.Name = %v, want %v", bank.Name, tt.wantName)
			}
		})
	}
}

func TestFindBanksByName(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	tests := []struct {
		name       string
		pattern    string
		wantMinLen int
		wantErr    bool
	}{
		{
			name:       "みずほを含む銀行",
			pattern:    ".*みずほ.*",
			wantMinLen: 1,
			wantErr:    false,
		},
		{
			name:       "三井を含む銀行",
			pattern:    ".*三井.*",
			wantMinLen: 1,
			wantErr:    false,
		},
		{
			name:    "無効な正規表現",
			pattern: "[invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			banks, err := z.FindBanksByName(tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBanksByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(banks) < tt.wantMinLen {
				t.Errorf("FindBanksByName() returned %d banks, want at least %d", len(banks), tt.wantMinLen)
			}
		})
	}
}

func TestGetBranch(t *testing.T) {
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
			branch, err := z.GetBranch(tt.bankCode, tt.branchCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && branch.Name != tt.wantName {
				t.Errorf("GetBranch() branch.Name = %v, want %v", branch.Name, tt.wantName)
			}
		})
	}
}

func TestFindBranchesByName(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	tests := []struct {
		name       string
		bankCode   string
		pattern    string
		wantMinLen int
		wantErr    bool
	}{
		{
			name:       "みずほ銀行の本店を含む支店",
			bankCode:   "0001",
			pattern:    ".*本店.*",
			wantMinLen: 0, // 本店という名前がない可能性があるので0
			wantErr:    false,
		},
		{
			name:       "みずほ銀行の東京を含む支店",
			bankCode:   "0001",
			pattern:    ".*東京.*",
			wantMinLen: 1,
			wantErr:    false,
		},
		{
			name:     "存在しない銀行",
			bankCode: "9999",
			pattern:  ".*",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			branches, err := z.FindBranchesByName(tt.bankCode, tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBranchesByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(branches) < tt.wantMinLen {
				t.Errorf("FindBranchesByName() returned %d branches, want at least %d", len(branches), tt.wantMinLen)
			}
		})
	}
}

func TestGetAllBanks(t *testing.T) {
	z, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	banks := z.GetAllBanks()
	if len(banks) == 0 {
		t.Error("GetAllBanks() returned no banks")
	}
}

func TestGetAllBranches(t *testing.T) {
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
			branches, err := z.GetAllBranches(tt.bankCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllBranches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(branches) < tt.wantMinLen {
				t.Errorf("GetAllBranches() returned %d branches, want at least %d", len(branches), tt.wantMinLen)
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
	bank, err := z.GetBank("0001")
	if err != nil {
		t.Fatalf("GetBank() failed: %v", err)
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

	// Test 4: Get branch through GetBranch and verify its Bank reference
	branch, err := z.GetBranch("0001", "001")
	if err != nil {
		t.Fatalf("GetBranch() failed: %v", err)
	}

	if branch.Bank == nil {
		t.Fatal("GetBranch() returned branch with nil Bank")
	}

	if branch.Bank.Code != "0001" {
		t.Errorf("branch.Bank.Code = %v, want 0001", branch.Bank.Code)
	}
}

// Package-level function tests (similar to ZenginCode::Bank.all in Ruby)

func TestAllBanks(t *testing.T) {
	banks := AllBanks()
	if len(banks) == 0 {
		t.Error("AllBanks() returned no banks")
	}

	// Verify it returns the same data as instance method
	z, _ := New()
	instanceBanks := z.GetAllBanks()
	if len(banks) != len(instanceBanks) {
		t.Errorf("AllBanks() returned %d banks, instance method returned %d", len(banks), len(instanceBanks))
	}
}

func TestFindBank(t *testing.T) {
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
				t.Errorf("FindBank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && bank.Name != tt.wantName {
				t.Errorf("FindBank() bank.Name = %v, want %v", bank.Name, tt.wantName)
			}
		})
	}
}

func TestPackageFindBanksByName(t *testing.T) {
	banks, err := FindBanksByName(".*みずほ.*")
	if err != nil {
		t.Fatalf("FindBanksByName() failed: %v", err)
	}
	if len(banks) == 0 {
		t.Error("FindBanksByName() returned no banks")
	}
}

func TestFindBranch(t *testing.T) {
	branch, err := FindBranch("0001", "001")
	if err != nil {
		t.Fatalf("FindBranch() failed: %v", err)
	}
	if branch.Name != "東京営業部" {
		t.Errorf("FindBranch() branch.Name = %v, want 東京営業部", branch.Name)
	}

	// Verify bidirectional reference
	if branch.Bank == nil {
		t.Fatal("FindBranch() returned branch with nil Bank")
	}
	if branch.Bank.Code != "0001" {
		t.Errorf("branch.Bank.Code = %v, want 0001", branch.Bank.Code)
	}
}

func TestPackageFindBranchesByName(t *testing.T) {
	branches, err := FindBranchesByName("0001", ".*東京.*")
	if err != nil {
		t.Fatalf("FindBranchesByName() failed: %v", err)
	}
	if len(branches) == 0 {
		t.Error("FindBranchesByName() returned no branches")
	}
}

func TestAllBranches(t *testing.T) {
	branches, err := AllBranches("0001")
	if err != nil {
		t.Fatalf("AllBranches() failed: %v", err)
	}
	if len(branches) == 0 {
		t.Error("AllBranches() returned no branches")
	}
}
