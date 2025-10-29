// Package zengin provides access to Japanese bank and branch code data (Zengin Code).
//
// ZenginCode is a comprehensive dataset of bank codes and branch codes for Japanese financial institutions.
// This library embeds the data directly using go:embed, making it easy to use without external dependencies.
//
// # Basic Usage
//
//	z, err := zengin.New()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get bank by code
//	bank, err := z.GetBank("0001")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Bank: %s\n", bank.Name)
//
//	// Search banks by name pattern (regex)
//	banks, err := z.FindBanksByName(".*みずほ.*")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, bank := range banks {
//		fmt.Printf("Found: %s\n", bank.Name)
//	}
//
//	// Get branch by bank code and branch code
//	branch, err := z.GetBranch("0001", "001")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Branch: %s\n", branch.Name)
//
// # Data Source
//
// This library uses data from the zengin-code project:
// https://github.com/zengin-code/source-data
package zengin
