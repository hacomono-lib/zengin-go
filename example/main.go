package main

import (
	"fmt"
	"log"

	"github.com/hacomono-lib/zengin-go"
)

func main() {
	fmt.Println("=== zengin-go Examples ===")
	fmt.Println()

	// Example 1: Get all banks
	fmt.Println("=== Example 1: Get all banks ===")
	allBanks := zengin.AllBanks()
	fmt.Printf("Total banks: %d\n", len(allBanks))
	fmt.Println()

	// Example 2: Get bank by code
	fmt.Println("=== Example 2: Get bank by code ===")
	bank, err := zengin.FindBank("0001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Code: %s\n", bank.Code)
	fmt.Printf("Name: %s\n", bank.Name)
	fmt.Printf("Kana: %s\n", bank.Kana)
	fmt.Printf("Hira: %s\n", bank.Hira)
	fmt.Printf("Roma: %s\n", bank.Roma)
	fmt.Printf("Total branches: %d\n", len(bank.Branches))
	fmt.Println()

	// Example 3: Get branch by bank code and branch code
	fmt.Println("=== Example 3: Get branch by bank code and branch code ===")
	branch, err := zengin.FindBranch("0001", "001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Code: %s\n", branch.Code)
	fmt.Printf("Name: %s\n", branch.Name)
	fmt.Printf("Kana: %s\n", branch.Kana)
	fmt.Printf("Hira: %s\n", branch.Hira)
	fmt.Printf("Roma: %s\n", branch.Roma)
	// Branch has reference to Bank (bidirectional relationship)
	fmt.Printf("Branch's bank: %s (%s)\n", branch.Bank.Name, branch.Bank.Code)
	fmt.Println()

	// Example 4: Get all branches for a bank
	fmt.Println("=== Example 4: Get all branches for a bank ===")
	allBranches, err := zengin.AllBranches("0001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total branches for bank 0001: %d\n", len(allBranches))
	
	// Show first 5 branches
	count := 0
	for _, br := range allBranches {
		if count >= 5 {
			fmt.Println("...")
			break
		}
		fmt.Printf("  - %s (%s)\n", br.Name, br.Code)
		count++
	}
	fmt.Println()

	// Example 5: Using instance methods (advanced)
	fmt.Println("=== Example 5: Using instance methods (advanced) ===")
	z, err := zengin.New()
	if err != nil {
		log.Fatal(err)
	}
	bank2, _ := z.FindBank("0005")
	fmt.Printf("Bank from instance: %s (%s)\n", bank2.Name, bank2.Code)
}

