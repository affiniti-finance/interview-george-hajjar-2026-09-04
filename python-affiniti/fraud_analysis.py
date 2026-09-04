"""
Fraud Detection System

This module implements a fraud detection system for credit card transactions.
"""

import json
from dataclasses import dataclass
from typing import Literal


@dataclass
class Transaction:
    """A Transaction represents a card trying to spend money"""
    id: str
    cardNumber: str  # 16-digit PAN
    amount: int      # amount in cents
    country: str
    timestamp: str   # ISO timestamp, e.g., "2025-03-13T09:00:00.000Z"
    category: str


# These are the only fraud rules to support for today
FraudRule = Literal["Unusual Amount", "Unusual Location", "Suspicious Category Sequence"]


@dataclass
class FraudScore:
    """
    A FraudScore captures the transaction decision, the card status, a risk score,
    and the rules that contributed to the FraudScore
    """
    decision: Literal["ACCEPTED", "DECLINED"]
    cardStatus: Literal["MONITORING", "AT_RISK"]
    fraudRiskScore: int
    rulesTriggered: list[FraudRule]


def check_fraud_score(actual: FraudScore, expected: FraudScore) -> list[str]:
    """
    Determines if a FraudScore is correct

    Args:
        actual: the fraud score as calculated by the code
        expected: the expected fraud score from the test transaction data

    Returns:
        list of field names that are incorrect
    """
    incorrect: list[str] = []

    if actual.decision != expected.decision:
        incorrect.append("decision")
    if actual.cardStatus != expected.cardStatus:
        incorrect.append("cardStatus")
    if actual.fraudRiskScore != expected.fraudRiskScore:
        incorrect.append("fraudRiskScore")
    if sorted(actual.rulesTriggered) != sorted(expected.rulesTriggered):
        incorrect.append("rulesTriggered")

    return incorrect


def calculate_transaction_fraud_score(transaction: Transaction) -> FraudScore:
    """
    Calculate a FraudScore for a transaction.

    - Determine what the Fraud Risk Score of a card is based on the transactions for that card
    - If the Card's Fraud Risk Score > 80
        then the card is in `AT_RISK` status
        and **the current transaction and all further transactions will be declined**
    - If the card has a Fraud Risk Score <= 80 then the card is in `MONITORING` status

    Fraud Rules to implement:

    1. **Unusual Location**: If the transaction is from a country that is not one of the
       last 5 countries transacted on this card
        Then ==>  +51 Fraud Risk Score on the card
        And  ==>  Transaction is immediately DECLINED

    2. **Unusual Amount**: If the transaction amount is more than double the current
       transaction average for this card
        Then ==> +20 Fraud Risk Score on the card

    3. **Suspicious Category Sequence**: If there are back to back transactions at the
       following merchant categories:
        'atm' -> 'jewelry'
        'online_gaming' -> 'luxury_goods'
        'electronics' -> 'pawn_shop'
        Then ==> +32 Fraud Risk Score

    Args:
        transaction: a Transaction object

    Returns:
        a FraudScore for the transaction
    """
    # TODO: This placeholder return value allows every transaction
    # TODO: Replace with code to enforce the fraud rules described above
    return FraudScore(
        decision="ACCEPTED",
        cardStatus="MONITORING",
        fraudRiskScore=0,
        rulesTriggered=[],
    )


def main() -> None:
    """
    Evaluate transactions against fraud rules. Results are printed to console.
    """
    # Load test transactions
    with open("test-transactions.json") as f:
        test_data = json.load(f)

    # Process each transaction
    for test_transaction in test_data:
        # Extract transaction data
        transaction = Transaction(
            id=test_transaction["id"],
            cardNumber=test_transaction["cardNumber"],
            amount=test_transaction["amount"],
            country=test_transaction["country"],
            timestamp=test_transaction["timestamp"],
            category=test_transaction["category"],
        )

        # Calculate fraud score
        fraud_score = calculate_transaction_fraud_score(transaction)

        # Extract expected results
        expected_data = test_transaction["expected"]
        expected = FraudScore(
            decision=expected_data["decision"],
            cardStatus=expected_data["cardStatus"],
            fraudRiskScore=expected_data["fraudRiskScore"],
            rulesTriggered=expected_data["rulesTriggered"],
        )

        # Check if correct
        incorrect_fields = check_fraud_score(fraud_score, expected)
        if incorrect_fields:
            print(f"Transaction {transaction.id} has errors: {', '.join(incorrect_fields)}")
        else:
            print(f"Transaction {transaction.id} has correct fraud score")


if __name__ == "__main__":
    main()
