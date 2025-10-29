package main

import (
	"fmt"
	"log"

	"github.com/hacomono-lib/zengin-go"
)

func main() {
	// Create a new Zengin instance
	z, err := zengin.New()
	if err != nil {
		log.Fatal(err)
	}

	// Example 1: Get bank by code
	fmt.Println("=== Example 1: Get bank by code ===")
	bank, err := z.GetBank("0001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Code: %s\n", bank.Code)
	fmt.Printf("Name: %s\n", bank.Name)
	fmt.Printf("Kana: %s\n", bank.Kana)
	fmt.Printf("Hira: %s\n", bank.Hira)
	fmt.Printf("Roma: %s\n", bank.Roma)
	fmt.Println()

	// Example 2: Find banks by name pattern
	fmt.Println("=== Example 2: Find banks by name pattern ===")
	banks, err := z.FindBanksByName(".*みずほ.*")
	if err != nil {
		log.Fatal(err)
	}
	for _, b := range banks {
		fmt.Printf("Found bank: %s (%s)\n", b.Name, b.Code)
	}
	fmt.Println()

	// Example 3: Get branch by bank code and branch code
	fmt.Println("=== Example 3: Get branch by bank code and branch code ===")
	branch, err := z.GetBranch("0001", "001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Code: %s\n", branch.Code)
	fmt.Printf("Name: %s\n", branch.Name)
	fmt.Printf("Kana: %s\n", branch.Kana)
	fmt.Printf("Hira: %s\n", branch.Hira)
	fmt.Printf("Roma: %s\n", branch.Roma)
	fmt.Println()

	// Example 4: Find branches by name pattern
	fmt.Println("=== Example 4: Find branches by name pattern ===")
	branches, err := z.FindBranchesByName("0001", ".*東京.*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d branches with '東京' in name:\n", len(branches))
	for i, br := range branches {
		if i >= 5 { // Show only first 5
			fmt.Println("...")
			break
		}
		fmt.Printf("  - %s (%s)\n", br.Name, br.Code)
	}
	fmt.Println()

	// Example 5: Get all banks
	fmt.Println("=== Example 5: Get all banks ===")
	allBanks := z.GetAllBanks()
	fmt.Printf("Total banks: %d\n", len(allBanks))
	fmt.Println()

	// Example 6: Get all branches for a bank
	fmt.Println("=== Example 6: Get all branches for a bank ===")
	allBranches, err := z.GetAllBranches("0001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total branches for bank 0001: %d\n", len(allBranches))
}
