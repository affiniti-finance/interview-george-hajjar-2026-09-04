package main

// Decision represents whether a transaction should be accepted or declined based on fraud analysis.
//
// Example usage:
//
//	score := FraudScore{
//	    Decision: DecisionAccepted, // or DecisionDeclined
//	    // ... other fields
//	}
type Decision string

const (
	// DecisionAccepted indicates the transaction passed fraud checks and should be approved
	DecisionAccepted Decision = "ACCEPTED"
	// DecisionDeclined indicates the transaction failed fraud checks and should be rejected
	DecisionDeclined Decision = "DECLINED"
)

// CardStatus represents the risk status of a card based on its transaction history.
// Cards start as MONITORING and can be elevated to AT_RISK based on fraud patterns.
//
// Example usage:
//
//	score := FraudScore{
//	    CardStatus: CardStatusMonitoring, // or CardStatusAtRisk
//	    // ... other fields
//	}
type CardStatus string

const (
	// CardStatusMonitoring indicates the card is under normal monitoring (default state)
	CardStatusMonitoring CardStatus = "MONITORING"
	// CardStatusAtRisk indicates the card has triggered fraud patterns and requires heightened scrutiny
	CardStatusAtRisk CardStatus = "AT_RISK"
)

// FraudRule represents the different types of fraud detection rules that can be triggered.
// These constants should be used directly in the RulesTriggered slice of FraudScore.
//
// Example usage:
//
//	rules := []FraudRule{
//	    FraudRuleUnusualLocation,
//	    FraudRuleUnusualAmount,
//	}
//	score := FraudScore{
//	    RulesTriggered: rules,
//	    // ... other fields
//	}
type FraudRule string

const (
	// FraudRuleUnusualLocation is triggered when a transaction occurs in a different country
	// than the card's previous transactions
	FraudRuleUnusualLocation FraudRule = "Unusual Location"
	// FraudRuleUnusualAmount is triggered when a transaction amount is significantly higher
	// than the card's historical average
	FraudRuleUnusualAmount FraudRule = "Unusual Amount"
	// FraudRuleSuspiciousCategorySequence is triggered when a card shows a suspicious pattern
	// of merchant categories (e.g., gas station followed by jewelry store)
	FraudRuleSuspiciousCategorySequence FraudRule = "Suspicious Category Sequence"
)

// Transaction represents a credit card transaction to be analyzed for fraud.
// Transactions are processed sequentially in timestamp order for each card.
type Transaction struct {
	ID         string          `json:"id"`         // Unique transaction identifier
	CardNumber string          `json:"cardNumber"` // Credit card number (identifier)
	Amount     int             `json:"amount"`     // Transaction amount in cents
	Country    string          `json:"country"`    // Two-letter country code (e.g., "US", "MX")
	Timestamp  string          `json:"timestamp"`  // ISO 8601 timestamp (e.g., "2024-01-01T10:00:00Z")
	Category   string          `json:"category"`   // Merchant category (e.g., "Gas Station", "Jewelry Store")
	Expected   *ExpectedResult `json:"expected,omitempty"` // Expected result for testing (nil in production)
}

// ExpectedResult contains the expected fraud analysis result for testing purposes.
// This is only populated in test data files and should not be used in the fraud analysis logic.
type ExpectedResult struct {
	Decision       string   `json:"decision"`       // Expected decision: "ACCEPTED" or "DECLINED"
	CardStatus     string   `json:"cardStatus"`     // Expected card status: "MONITORING" or "AT_RISK"
	FraudRiskScore int      `json:"fraudRiskScore"` // Expected fraud risk score
	RulesTriggered []string `json:"rulesTriggered"` // Expected fraud rules triggered (as strings)
}

// FraudScore represents the fraud analysis result for a transaction.
// This is the primary return type for the CalculateTransactionFraudScore function.
//
// Example:
//
//	score := FraudScore{
//	    Decision:       DecisionAccepted,
//	    CardStatus:     CardStatusMonitoring,
//	    FraudRiskScore: 10,
//	    RulesTriggered: []FraudRule{FraudRuleUnusualAmount},
//	}
type FraudScore struct {
	Decision       Decision    `json:"decision"`       // Whether to accept or decline the transaction
	CardStatus     CardStatus  `json:"cardStatus"`     // Current risk status of the card
	FraudRiskScore int         `json:"fraudRiskScore"` // Cumulative fraud risk score for the card
	RulesTriggered []FraudRule `json:"rulesTriggered"` // List of fraud rules triggered by this transaction
}
