package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/**
 * Main program to validate the FraudAnalyzer implementation.
 * Loads test transactions and compares actual vs expected results.
 */
func main() {
	// Load test transactions
	data, err := os.ReadFile("test-transactions.json")
	if err != nil {
		fmt.Printf("Error reading test-transactions.json: %v\n", err)
		os.Exit(1)
	}

	var transactions []Transaction
	if err := json.Unmarshal(data, &transactions); err != nil {
		fmt.Printf("Error parsing test-transactions.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Running Fraud Detection Tests...\n")

	// Create analyzer
	analyzer := NewFraudAnalyzer()

	passed := 0
	failed := 0

	// Process each transaction
	for _, txn := range transactions {
		actual := analyzer.CalculateTransactionFraudScore(txn)
		expected := txn.Expected

		testPassed := validateResult(txn, actual, expected)

		if testPassed {
			passed++
			fmt.Printf("✓ %s PASSED\n", txn.ID)
		} else {
			failed++
			fmt.Printf("✗ %s FAILED\n", txn.ID)
			printMismatch(txn, actual, expected)
		}
	}

	// Print summary
	fmt.Println()
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Results: %d passed, %d failed\n", passed, failed)
	fmt.Println(strings.Repeat("=", 50))

	// Exit with appropriate status code
	if failed == 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func validateResult(txn Transaction, actual FraudScore, expected *ExpectedResult) bool {
	if expected == nil {
		return false
	}

	// Check decision
	if string(actual.Decision) != expected.Decision {
		return false
	}

	// Check card status
	if string(actual.CardStatus) != expected.CardStatus {
		return false
	}

	// Check fraud risk score
	if actual.FraudRiskScore != expected.FraudRiskScore {
		return false
	}

	// Check rules triggered
	if len(actual.RulesTriggered) != len(expected.RulesTriggered) {
		return false
	}

	// Convert actual rules to strings for comparison
	actualRulesMap := make(map[string]bool)
	for _, rule := range actual.RulesTriggered {
		actualRulesMap[string(rule)] = true
	}

	for _, expectedRule := range expected.RulesTriggered {
		if !actualRulesMap[expectedRule] {
			return false
		}
	}

	return true
}

func printMismatch(txn Transaction, actual FraudScore, expected *ExpectedResult) {
	if expected == nil {
		return
	}

	// Mask card number (show last 4 digits)
	maskedCard := "****" + txn.CardNumber[len(txn.CardNumber)-4:]
	fmt.Printf("  Card: %s\n", maskedCard)
	fmt.Printf("  Amount: %d, Country: %s, Category: %s\n", txn.Amount, txn.Country, txn.Category)

	if string(actual.Decision) != expected.Decision {
		fmt.Printf("  Decision: expected %s, got %s\n", expected.Decision, actual.Decision)
	}

	if string(actual.CardStatus) != expected.CardStatus {
		fmt.Printf("  Card Status: expected %s, got %s\n", expected.CardStatus, actual.CardStatus)
	}

	if actual.FraudRiskScore != expected.FraudRiskScore {
		fmt.Printf("  Fraud Score: expected %d, got %d\n", expected.FraudRiskScore, actual.FraudRiskScore)
	}

	// Check rules triggered mismatch
	actualRulesMap := make(map[string]bool)
	actualRulesList := []string{}
	for _, rule := range actual.RulesTriggered {
		ruleStr := string(rule)
		actualRulesMap[ruleStr] = true
		actualRulesList = append(actualRulesList, ruleStr)
	}

	rulesMatch := len(actual.RulesTriggered) == len(expected.RulesTriggered)
	if rulesMatch {
		for _, expectedRule := range expected.RulesTriggered {
			if !actualRulesMap[expectedRule] {
				rulesMatch = false
				break
			}
		}
	}

	if !rulesMatch {
		fmt.Printf("  Rules Triggered: expected %v, got %v\n", expected.RulesTriggered, actualRulesList)
	}

	fmt.Println()
}
