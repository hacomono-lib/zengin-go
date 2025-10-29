// Package zengin provides access to Japanese bank and branch code data (Zengin Code).
//
// ZenginCode is a comprehensive dataset of bank codes and branch codes for Japanese financial institutions.
// This library embeds the data directly using go:embed and automatically preloads it at startup.
//
// # Basic Usage
//
// Just import and use the package-level functions:
//
//	import "github.com/hacomono-lib/zengin-go"
//
//	// Get all banks
//	banks := zengin.AllBanks()
//	fmt.Printf("Total banks: %d\n", len(banks))
//
//	// Get bank by code
//	bank, err := zengin.FindBank("0001")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Bank: %s\n", bank.Name)
//
//	// Get branch by bank code and branch code
//	branch, err := zengin.FindBranch("0001", "001")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Branch: %s\n", branch.Name)
//
//	// Branch has reference to Bank (bidirectional relationship)
//	fmt.Printf("Branch's bank: %s\n", branch.Bank.Name)
//
// # Advanced Usage
//
// For advanced use cases (e.g., testing, custom data sources), you can create your own instance:
//
//	z, err := zengin.New()
//	if err != nil {
//		log.Fatal(err)
//	}
//	bank, _ := z.FindBank("0001")
//
// # Data Source
//
// This library uses data from the zengin-code project:
// https://github.com/zengin-code/source-data
package zengin
