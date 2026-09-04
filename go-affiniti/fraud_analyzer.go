package main

import "fmt"

/**
 * Fraud detection system for credit card transactions.
 *
 * TODO: Implement the CalculateTransactionFraudScore method below.
 * You may need to add fields to the FraudAnalyzer struct to maintain state.
 */

// CardState is a suggested helper struct for tracking per-card state.
// You can modify this or create your own approach.
type CardState struct {
	// Example fields you might need:
	// TransactionCount int
	// TotalAmount      int
	// Countries        []string // Last 5 countries
	// LastCategory     string
	// FraudRiskScore   int
}

type FraudAnalyzer struct {
	// TODO: Add fields here to track state for each card
	// Example: cardStates map[string]*CardState
}

/**
 * NewFraudAnalyzer creates a new instance of the fraud analyzer.
 */
func NewFraudAnalyzer() *FraudAnalyzer {
	return &FraudAnalyzer{
		// TODO: Initialize your state here
		// Example: cardStates: make(map[string]*CardState),
	}
}

/**
 * CalculateTransactionFraudScore analyzes a transaction and calculates its fraud score.
 *
 * Requirements:
 *
 * 1. Maintain state for each credit card across all transactions:
 *    - Transaction history (for calculating averages)
 *    - Running fraud risk score (cumulative and persistent)
 *    - Last 5 countries used
 *    - Previous transaction category
 *
 * 2. Apply three fraud detection rules:
 *
 *    a) Unusual Location (+51 points):
 *       - Triggers when transaction is from a country NOT in the last 5 countries used
 *       - Transaction is immediately DECLINED
 *
 *    b) Unusual Amount (+20 points):
 *       - Triggers when amount > 2x the current average for this card
 *       - Transaction can still be ACCEPTED
 *
 *    c) Suspicious Category Sequence (+32 points):
 *       - Triggers on these back-to-back merchant category pairs:
 *         * atm → jewelry
 *         * online_gaming → luxury_goods
 *         * electronics → pawn_shop
 *       - Transaction can still be ACCEPTED
 *
 * 3. Determine card status:
 *    - MONITORING: fraud risk score <= 80
 *    - AT_RISK: fraud risk score > 80
 *      * When AT_RISK, the current transaction AND all future transactions are DECLINED
 *      * Cards remain AT_RISK permanently (score doesn't decrease)
 *
 * 4. Return a FraudScore struct with all required fields populated.
 */
func (fa *FraudAnalyzer) CalculateTransactionFraudScore(transaction Transaction) FraudScore {
	// TODO: Implement fraud detection logic here
	//
	// Implementation steps:
	// 1. Get or create card state for transaction.CardNumber
	// 2. Check if card is already AT_RISK (score > 80) - if so, decline immediately
	// 3. Initialize: rulesTriggered := []FraudRule{}
	// 4. Check each fraud rule:
	//    - Unusual Location: Is transaction.Country in last 5 countries?
	//    - Unusual Amount: Is transaction.Amount > 2x average?
	//    - Suspicious Category: Does lastCategory + transaction.Category match patterns?
	// 5. Add points to card's fraud score for each triggered rule
	// 6. Determine decision (DECLINED if Unusual Location OR score > 80)
	// 7. Update card state (add country, update category, increment counts)
	// 8. Return FraudScore with decision, status, score, and triggered rules

	return FraudScore{
		Decision:       DecisionAccepted,
		CardStatus:     CardStatusMonitoring,
		FraudRiskScore: 0,
		RulesTriggered: []FraudRule{},
	}
}

// CheckFraudScore validates that an actual FraudScore matches the expected result.
// This is a helper function for testing your implementation incrementally.
//
// Returns true if the actual result matches the expected result, false otherwise.
// Prints detailed error messages for each field that doesn't match.
//
// Example usage:
//
//	analyzer := NewFraudAnalyzer()
//	actual := analyzer.CalculateTransactionFraudScore(transaction)
//	if !CheckFraudScore(actual, transaction.Expected) {
//	    fmt.Printf("Test failed for transaction %s\n", transaction.ID)
//	}
func CheckFraudScore(actual FraudScore, expected *ExpectedResult) bool {
	if expected == nil {
		return true // No expected result to compare against
	}

	allMatch := true

	// Check Decision
	if string(actual.Decision) != expected.Decision {
		fmt.Printf("  ❌ Decision mismatch: got %s, expected %s\n", actual.Decision, expected.Decision)
		allMatch = false
	}

	// Check CardStatus
	if string(actual.CardStatus) != expected.CardStatus {
		fmt.Printf("  ❌ CardStatus mismatch: got %s, expected %s\n", actual.CardStatus, expected.CardStatus)
		allMatch = false
	}

	// Check FraudRiskScore
	if actual.FraudRiskScore != expected.FraudRiskScore {
		fmt.Printf("  ❌ FraudRiskScore mismatch: got %d, expected %d\n", actual.FraudRiskScore, expected.FraudRiskScore)
		allMatch = false
	}

	// Check RulesTriggered
	if len(actual.RulesTriggered) != len(expected.RulesTriggered) {
		fmt.Printf("  ❌ RulesTriggered count mismatch: got %d rules, expected %d rules\n",
			len(actual.RulesTriggered), len(expected.RulesTriggered))
		fmt.Printf("     Got: %v\n", actual.RulesTriggered)
		fmt.Printf("     Expected: %v\n", expected.RulesTriggered)
		allMatch = false
	} else {
		// Convert actual rules to strings for comparison
		actualRulesStr := make(map[string]bool)
		for _, rule := range actual.RulesTriggered {
			actualRulesStr[string(rule)] = true
		}

		expectedRulesStr := make(map[string]bool)
		for _, rule := range expected.RulesTriggered {
			expectedRulesStr[rule] = true
		}

		// Check if all expected rules are present
		for expectedRule := range expectedRulesStr {
			if !actualRulesStr[expectedRule] {
				fmt.Printf("  ❌ RulesTriggered missing: %s\n", expectedRule)
				allMatch = false
			}
		}

		// Check if there are any unexpected rules
		for actualRule := range actualRulesStr {
			if !expectedRulesStr[actualRule] {
				fmt.Printf("  ❌ RulesTriggered unexpected: %s\n", actualRule)
				allMatch = false
			}
		}
	}

	return allMatch
}
