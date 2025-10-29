package zengin

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

//go:embed source-data/data/banks.json
//go:embed source-data/data/branches/*.json
var dataFS embed.FS

// Zengin represents the main zengin code library
type Zengin struct {
	banks map[string]*BankWithBranches
}

// New creates a new Zengin instance
func New() (*Zengin, error) {
	z := &Zengin{
		banks: make(map[string]*BankWithBranches),
	}

	if err := z.loadBanks(); err != nil {
		return nil, fmt.Errorf("failed to load banks: %w", err)
	}

	return z, nil
}

// loadBanks loads all bank data from embedded JSON
func (z *Zengin) loadBanks() error {
	// Load banks.json
	banksData, err := dataFS.ReadFile("source-data/data/banks.json")
	if err != nil {
		return fmt.Errorf("failed to read banks.json: %w", err)
	}

	var banksMap map[string]*Bank
	if err := json.Unmarshal(banksData, &banksMap); err != nil {
		return fmt.Errorf("failed to unmarshal banks.json: %w", err)
	}

	// Load each bank's branches
	for code, bank := range banksMap {
		bankWithBranches := &BankWithBranches{
			Bank:     *bank,
			Branches: make(map[string]*Branch),
		}

		// Load branch data
		branchFile := fmt.Sprintf("source-data/data/branches/%s.json", code)
		branchData, err := dataFS.ReadFile(branchFile)
		if err != nil {
			// Some banks might not have branch data, skip them
			z.banks[code] = bankWithBranches
			continue
		}

		var branchesMap map[string]*Branch
		if err := json.Unmarshal(branchData, &branchesMap); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", branchFile, err)
		}

		bankWithBranches.Branches = branchesMap
		z.banks[code] = bankWithBranches
	}

	return nil
}

// GetBank returns a bank by its code
func (z *Zengin) GetBank(code string) (*Bank, error) {
	bank, exists := z.banks[code]
	if !exists {
		return nil, errors.New("bank not found")
	}
	return &bank.Bank, nil
}

// FindBanksByName finds banks by name pattern (regex)
func (z *Zengin) FindBanksByName(pattern string) ([]*Bank, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	var results []*Bank
	for _, bank := range z.banks {
		if re.MatchString(bank.Name) {
			bankCopy := bank.Bank
			results = append(results, &bankCopy)
		}
	}

	return results, nil
}

// GetBranch returns a branch by bank code and branch code
func (z *Zengin) GetBranch(bankCode, branchCode string) (*Branch, error) {
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

// FindBranchesByName finds branches by bank code and name pattern (regex)
func (z *Zengin) FindBranchesByName(bankCode, pattern string) ([]*Branch, error) {
	bank, exists := z.banks[bankCode]
	if !exists {
		return nil, errors.New("bank not found")
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	var results []*Branch
	for _, branch := range bank.Branches {
		if re.MatchString(branch.Name) {
			results = append(results, branch)
		}
	}

	return results, nil
}

// GetAllBanks returns all banks
func (z *Zengin) GetAllBanks() map[string]*Bank {
	banks := make(map[string]*Bank)
	for code, bank := range z.banks {
		bankCopy := bank.Bank
		banks[code] = &bankCopy
	}
	return banks
}

// GetAllBranches returns all branches for a specific bank
func (z *Zengin) GetAllBranches(bankCode string) (map[string]*Branch, error) {
	bank, exists := z.banks[bankCode]
	if !exists {
		return nil, errors.New("bank not found")
	}
	return bank.Branches, nil
}
