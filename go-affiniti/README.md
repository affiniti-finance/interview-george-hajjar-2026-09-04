# Fraud Detection Coding Challenge (Go)

## Prerequisites

- **Go 1.18 or later** (required)

## Getting Started

This project uses Go modules, which are built into Go. No additional tools needed.

### Run the Program

```bash
go run .
```

Many tests will initially fail. Your task is to implement the fraud detection logic to make them pass.

## The Challenge

Implement the `CalculateTransactionFraudScore()` method in `fraud_analyzer.go`.

The program loads test transactions from `test-transactions.json`, processes each one through your implementation, and compares the actual results against expected results. You'll see which tests pass or fail with detailed feedback.

## Problem Statement

Implement the fraud detection logic to analyze credit card transactions and determine whether they should be ACCEPTED or DECLINED based on fraud detection rules.

Your implementation must:
- Analyze transactions in real-time
- Maintain state for each credit card
- Calculate cumulative fraud risk scores
- Apply fraud detection rules
- Return proper fraud analysis results

If a card's fraud risk score exceeds 80, the card is in `AT_RISK` status, and **the current transaction and all further transactions will be declined**. If the card has a fraud risk score of 80 or less, the card is in `MONITORING` status.

## Fraud Rules to Implement

1. **Unusual Location**: If the transaction is from a country that is not one of the last 5 countries transacted on this card.
   - +51 Fraud Risk Score on the card.
   - Transaction is immediately DECLINED.

2. **Unusual Amount**: If the transaction amount is more than double the current transaction average for this card
   - +20 Fraud Risk Score on the card.

3. **Suspicious Category Sequence**: If there are back-to-back transactions at the following merchant categories:
   ```
   'atm' -> 'jewelry'
   'online_gaming' -> 'luxury_goods'
   'electronics' -> 'pawn_shop'
   ```
   - +32 Fraud Risk Score

## State Management

Your implementation must track the following state for each card:

- **Transaction history**: All previous transactions for the card
- **Fraud risk score**: Cumulative score that persists across transactions
- **Last 5 countries**: A sliding window of up to 5 unique countries
- **Previous category**: The merchant category of the previous transaction

## Decision Logic

- If fraud risk score > 80, the card status is `AT_RISK`
  - All transactions (current and future) are DECLINED
  - The card remains AT_RISK permanently

- If fraud risk score ≤ 80, the card status is `MONITORING`
  - Transactions may be ACCEPTED or DECLINED based on rules
  - Only the "Unusual Location" rule immediately declines a transaction
  - Other rules add points but don't automatically decline

## Expected Return Values

Your implementation should return a `FraudScore` struct containing:

- **Decision**: `DecisionAccepted` or `DecisionDeclined`
- **CardStatus**: `CardStatusMonitoring` or `CardStatusAtRisk`
- **FraudRiskScore**: Current cumulative score for the card (integer)
- **RulesTriggered**: Slice of fraud rule constants that triggered on this transaction
  - Possible values: `FraudRuleUnusualLocation`, `FraudRuleUnusualAmount`, `FraudRuleSuspiciousCategorySequence`
  - Multiple rules can trigger on the same transaction

## Important Notes

- **Fraud scores are cumulative**: Once points are added, they persist across all future transactions
- **Process transactions sequentially**: The order matters for state management
- **Average calculation**: For "Unusual Amount", use the average of PREVIOUS transactions (not including the current one)
- **Last 5 countries**: This is a sliding window - when a 6th country appears, the oldest is removed from consideration
- **First transaction**: A card's first transaction has no history, so some rules won't trigger

## Helper Functions & Tools

### CheckFraudScore() Helper

A `CheckFraudScore()` function is available in `fraud_analyzer.go` to help validate your implementation incrementally:

```go
analyzer := NewFraudAnalyzer()
actual := analyzer.CalculateTransactionFraudScore(transaction)
if !CheckFraudScore(actual, transaction.Expected) {
    fmt.Printf("Test failed for transaction %s\n", transaction.ID)
}
```

This provides detailed output showing which fields don't match expected results.

### CardState Struct Template

A `CardState` struct template is provided in `fraud_analyzer.go` with example fields for tracking per-card state. You can use it as-is, modify it, or create your own structure.

## Project Structure

```
go-affiniti/
├── main.go                # Test runner (don't modify)
├── fraud_analyzer.go      # Implement this! (includes helper functions)
├── types.go               # Data structures and constants
├── test-transactions.json # Test data with expected results
├── go.mod                 # Go module definition
└── README.md
```

## Workflow

1. Run `go run .` - see failures
2. Edit `fraud_analyzer.go` and implement the logic
   - Add fields to the `FraudAnalyzer` struct to track state
   - Implement the `CalculateTransactionFraudScore()` method
3. Run `go run .` again - see tests pass
4. When all tests pass, you're done!

## Test Data

The [`test-transactions.json`](./test-transactions.json) file contains 22 test transactions with expected results for validation. Each transaction object includes an `expected` field that shows the correct results your application should produce after processing that transaction.

## Submission

When all tests pass:

1. Verify: `go run .`
2. Commit your changes
3. Push to your repository

Good luck!
