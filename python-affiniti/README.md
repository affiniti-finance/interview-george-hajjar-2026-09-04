# Fraud Detection Coding Challenge (Python)

A Python coding challenge that implements a fraud detection system for credit card transactions.

## Setup

**Requirements:** Python 3.12 or higher

## Running the Challenge

Run the test suite:
```bash
python fraud_analysis.py
# or
python3 fraud_analysis.py
```

The output will show which transactions pass or fail validation against expected fraud scores.

## Your Task

Implement the `calculate_transaction_fraud_score()` function in `fraud_analysis.py` to make all 22 test transactions pass.

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

Your implementation should return a `FraudScore` object containing:

- **decision**: `"ACCEPTED"` or `"DECLINED"`
- **cardStatus**: `"MONITORING"` or `"AT_RISK"`
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

The [`test-transactions.json`](./test-transactions.json) file contains 22 test transactions with expected results for validation. Each transaction object includes an `expected` field that shows the correct results your application should produce after processing that transaction.
