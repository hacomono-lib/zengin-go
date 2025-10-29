package zengin

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
)

//go:embed data/banks.json
//go:embed data/branches/*.json
var dataFS embed.FS

// defaultZengin is the global instance, preloaded on init
var defaultZengin *Zengin

func init() {
	var err error
	defaultZengin, err = New()
	if err != nil {
		panic(fmt.Sprintf("failed to preload zengin data: %v", err))
	}
}

// Zengin represents the main zengin code library
type Zengin struct {
	banks map[string]*Bank
}

// New creates a new Zengin instance
func New() (*Zengin, error) {
	z := &Zengin{
		banks: make(map[string]*Bank),
	}

	if err := z.loadBanks(); err != nil {
		return nil, fmt.Errorf("failed to load banks: %w", err)
	}

	return z, nil
}

// loadBanks loads all bank data from embedded JSON
func (z *Zengin) loadBanks() error {
	// Load banks.json
	banksData, err := dataFS.ReadFile("data/banks.json")
	if err != nil {
		return fmt.Errorf("failed to read banks.json: %w", err)
	}

	var banksMap map[string]*Bank
	if err := json.Unmarshal(banksData, &banksMap); err != nil {
		return fmt.Errorf("failed to unmarshal banks.json: %w", err)
	}

	// Load each bank's branches
	for code, bank := range banksMap {
		bank.Branches = make(map[string]*Branch)

		// Load branch data
		branchFile := fmt.Sprintf("data/branches/%s.json", code)
		branchData, err := dataFS.ReadFile(branchFile)
		if err != nil {
			// Some banks might not have branch data, skip them
			z.banks[code] = bank
			continue
		}

		var branchesMap map[string]*Branch
		if err := json.Unmarshal(branchData, &branchesMap); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", branchFile, err)
		}

		// Set up bidirectional relationship
		for _, branch := range branchesMap {
			branch.Bank = bank
		}

		bank.Branches = branchesMap
		z.banks[code] = bank
	}

	return nil
}

// FindBank returns a bank by its code
func (z *Zengin) FindBank(code string) (*Bank, error) {
	bank, exists := z.banks[code]
	if !exists {
		return nil, errors.New("bank not found")
	}
	return bank, nil
}

// AllBanks returns all banks
func (z *Zengin) AllBanks() map[string]*Bank {
	return z.banks
}

// FindBranch returns a branch by bank code and branch code
func (z *Zengin) FindBranch(bankCode, branchCode string) (*Branch, error) {
	bank, exists := z.banks[bankCode]
	if !exists {
		return nil, errors.New("bank not found")
	}

	branch, exists := bank.Branches[branchCode]
	if !exists {
		return nil, errors.New("branch not found")
	}

	return branch, nil
}

// AllBranches returns all branches for a specific bank
func (z *Zengin) AllBranches(bankCode string) (map[string]*Branch, error) {
	bank, exists := z.banks[bankCode]
	if !exists {
		return nil, errors.New("bank not found")
	}
	return bank.Branches, nil
}

// Package-level functions using the global instance

// AllBanks returns all banks from the global instance
func AllBanks() map[string]*Bank {
	return defaultZengin.AllBanks()
}

// FindBank returns a bank by its code from the global instance
func FindBank(code string) (*Bank, error) {
	return defaultZengin.FindBank(code)
}

// FindBranch returns a branch by bank code and branch code from the global instance
func FindBranch(bankCode, branchCode string) (*Branch, error) {
	return defaultZengin.FindBranch(bankCode, branchCode)
}

// AllBranches returns all branches for a specific bank from the global instance
func AllBranches(bankCode string) (map[string]*Branch, error) {
	return defaultZengin.AllBranches(bankCode)
}
