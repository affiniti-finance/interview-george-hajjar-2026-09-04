# Fraud Detection Coding Challenge (Java)

## Prerequisites

- **Java 17 or later** (required)

## Getting Started

This project supports both **Maven** and **Gradle**. Choose your preferred build tool:

### Option 1: Gradle

```bash
./gradlew run
```

On Windows:
```cmd
gradlew.bat run
```

### Option 2: Maven

```bash
./mvnw compile exec:java
```

On Windows:
```cmd
mvnw.cmd compile exec:java
```

**Note:** This project uses build tool wrappers (Gradle Wrapper and Maven Wrapper), so you don't need to have Gradle or Maven installed.

Many tests will initially fail. Your task is to implement the fraud detection logic to make them pass.

## The Challenge

Implement the `calculateTransactionFraudScore()` method in `src/FraudAnalyzer.java`.

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

Your implementation should return a `FraudScore` object containing:

- **decision**: `Decision.ACCEPTED` or `Decision.DECLINED`
- **cardStatus**: `CardStatus.MONITORING` or `CardStatus.AT_RISK`
- **fraudRiskScore**: Current cumulative score for the card (integer)
- **rulesTriggered**: List of rule display names that triggered on this transaction
  - Use the display names: `"Unusual Location"`, `"Unusual Amount"`, `"Suspicious Category Sequence"`
  - Multiple rules can trigger on the same transaction

## Important Notes

- **Fraud scores are cumulative**: Once points are added, they persist across all future transactions
- **Process transactions sequentially**: The order matters for state management
- **Average calculation**: For "Unusual Amount", use the average of PREVIOUS transactions (not including the current one)
- **Last 5 countries**: This is a sliding window - when a 6th country appears, the oldest is removed from consideration
- **First transaction**: A card's first transaction has no history, so some rules won't trigger

## Project Structure

```
java-affiniti/
├── src/
│   ├── Main.java              # Test runner (don't modify)
│   ├── FraudAnalyzer.java     # Implement this!
│   ├── Transaction.java       # Data models
│   ├── FraudScore.java
│   ├── Decision.java
│   ├── CardStatus.java
│   └── FraudRule.java
├── test-transactions.json     # Test data with expected results
├── build.gradle               # Gradle configuration
├── pom.xml                    # Maven configuration
├── gradle/                    # Gradle wrapper files
├── .mvn/                      # Maven wrapper files
├── gradlew, gradlew.bat       # Gradle wrapper scripts
├── mvnw, mvnw.cmd             # Maven wrapper scripts
└── README.md
```

## Workflow

### Using Gradle:
1. Run `./gradlew run` - see failures
2. Edit `src/FraudAnalyzer.java` and implement the logic
3. Run `./gradlew run` again - see tests pass
4. When all tests pass, you're done!

### Using Maven:
1. Run `./mvnw compile exec:java` - see failures
2. Edit `src/FraudAnalyzer.java` and implement the logic
3. Run `./mvnw compile exec:java` again - see tests pass
4. When all tests pass, you're done!

## Test Data

The [`test-transactions.json`](./test-transactions.json) file contains 22 test transactions with expected results for validation. Each transaction object includes an `expected` field that shows the correct results your application should produce after processing that transaction.

## Submission

When all tests pass:

1. Verify with your chosen build tool:
   - Gradle: `./gradlew run`
   - Maven: `./mvnw compile exec:java`
2. Commit your changes
3. Push to your repository

Good luck!
