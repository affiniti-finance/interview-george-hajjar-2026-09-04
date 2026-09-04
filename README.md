# Fraud Detection Coding Challenge

This is a coding interview challenge focused on implementing fraud detection logic for credit card transactions.

## Overview

You will implement a fraud detection system that:
- Processes a series of credit card transactions
- Tracks card state across multiple transactions
- Applies fraud detection rules to calculate risk scores
- Makes accept/decline decisions based on fraud analysis

The challenge includes a set of test transactions with expected outcomes. Your goal is to implement the fraud detection logic to make all tests pass.

## Problem Statement

Implement the fraud detection logic to analyze credit card transactions and determine whether they should be ACCEPTED or DECLINED based on fraud detection rules.

Your implementation must:
- Analyze transactions
- Maintain state for each credit card
- Calculate cumulative fraud risk scores
- Apply fraud detection rules
- Return proper fraud analysis results

If a card's fraud risk score exceeds 80, the card is in `AT_RISK` status, and **the current transaction and all further transactions will be declined**. If the card has a fraud risk score of 80 or less, the card is in `MONITORING` status.

## Requirements

### Fraud Rules to Implement

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

### State Management

Your implementation must track the following state for each card:

- **Transaction history**: All previous transactions for the card
- **Fraud risk score**: Cumulative score that persists across transactions
- **Last 5 countries**: A sliding window of up to 5 unique countries
- **Previous category**: The merchant category of the previous transaction

### Decision Logic

- If fraud risk score > 80, the card status is `AT_RISK`
  - All transactions (current and future) are DECLINED
  - The card remains AT_RISK permanently

- If fraud risk score ≤ 80, the card status is `MONITORING`
  - Transactions may be ACCEPTED or DECLINED based on rules
  - Only the "Unusual Location" rule immediately declines a transaction
  - Other rules add points but don't automatically decline

### Expected Return Values

Your implementation should return fraud analysis results containing:

- **decision**: ACCEPTED or DECLINED
- **cardStatus**: MONITORING or AT_RISK
- **fraudRiskScore**: Current cumulative score for the card (integer)
- **rulesTriggered**: List of rule names that triggered on this transaction
  - Possible values: `"Unusual Location"`, `"Unusual Amount"`, `"Suspicious Category Sequence"`
  - Multiple rules can trigger on the same transaction

## Important Notes

- **Fraud scores are cumulative**: Once points are added, they persist across all future transactions
- **Process transactions sequentially**: The order matters for state management
- **Average calculation**: For "Unusual Amount", use the average of PREVIOUS transactions (not including the current one)
- **Last 5 countries**: This is a sliding window - when a 6th country appears, the oldest is removed from consideration
- **First transaction**: A card's first transaction has no history, so some rules won't trigger

## Test Data

Each language implementation includes a `test-transactions.json` file with 22 test transactions. Each transaction object includes an `expected` field that shows the correct results your application should produce after processing that transaction.

## Getting Started

Choose a language-specific directory and follow the README instructions within that directory to begin the challenge:

- **[typescript-affiniti/](./typescript-affiniti/)** - TypeScript implementation
- **[java-affiniti/](./java-affiniti/)** - Java (supports both Gradle and Maven)
- **[python-affiniti/](./python-affiniti/)** - Python implementation
- **[go-affiniti/](./go-affiniti/)** - Go implementation

Each directory contains:
- Setup instructions and prerequisites
- Scaffolded code with the function/method to implement
- Test runner to validate your implementation
- The same test data with expected results
