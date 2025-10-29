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
