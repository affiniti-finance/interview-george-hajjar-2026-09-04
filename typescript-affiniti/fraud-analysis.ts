
import testTransactionsData from "./test-transactions.json";

/**
 * A Transaction represents a card trying to spend money
 */
type Transaction = {
  id: string;
  cardNumber: string; // 16-digit PAN
  amount: number;     // amount in cents
  country: string;
  timestamp: string;  // ISO timestamp, for example "2025-03-13T09:00:00.000Z",
  category: string;
};

// These are the only fraud rules to support for today
type FraudRule = "Unusual Amount" | "Unusual Location" | "Suspicious Category Sequence"

/**
 * A FraudScore captures the transaction decision, the card status, a risk score,
 * and the rules that contributed to the FraudScore
 */
type FraudScore = {
  decision: "ACCEPTED" | "DECLINED";
  cardStatus: "MONITORING" | "AT_RISK";
  fraudRiskScore: number;
  rulesTriggered: FraudRule[];
};

/**
 * The TestTransaction data is simply a Transaction with the expected FraudScore in the "expected" property
 */
type TestTransaction = Transaction & {
  expected: FraudScore
}

/**
 * Determines if a FraudScore is correct
 * @param actual the fraud score as calculated by the code
 * @param expected the expected fraud score from the test transaction data
 */
function checkFraudScore (actual: FraudScore, expected: FraudScore): string[] {
  const incorrect: string[] = [];
  if (actual.decision !== expected.decision) {
    incorrect.push("decision");
  }
  if (actual.cardStatus !== expected.cardStatus) {
    incorrect.push("cardStatus");
  }
  if (actual.fraudRiskScore !== expected.fraudRiskScore) {
    incorrect.push("fraudRiskScore");
  }
  if (JSON.stringify([...actual.rulesTriggered].sort()) !== JSON.stringify([...expected.rulesTriggered].sort())) {
    incorrect.push("rulesTriggered");
  }
  return incorrect;
}

/**
 * Calculate a FraudScore for a transaction.
 *
 * - Determine what the Fraud Risk Score of a card is based on the transactions for that card
 * - If the Card's Fraud Risk Score > 80
 *     then the card is in `AT_RISK` status
 *     and **the current transaction and all further transactions will be declined**
 * - If the card has a Fraud Risk Score <= 80 then the card is in `MONITORING` status
 *
 * Fraud Rules to implement:
 *
 * 1. **Unusual Location**: If the transaction is from a country that is not one of the last 5 countries transacted on this card
 *     Then ==>  +51 Fraud Risk Score on the card
 *     And  ==>  Transaction is immediately DECLINED
 *
 * 2. **Unusual Amount**: If the transaction amount is more than double the current transaction average for this card
 *     Then ==> +20 Fraud Risk Score on the card
 *
 * 3. **Suspicious Category Sequence**: If there are back to back transactions at the following merchant categories
 *     'atm' -> 'jewelry'
 *     'online_gaming' -> 'luxury_goods'
 *     'electronics' -> 'pawn_shop'
 *     Then ==> +32 Fraud Risk Score
 *
 * @param transaction a Transaction object
 * @return a FraudScore for the transaction
 */
function calculateTransactionFraudScore(transaction: Transaction): FraudScore {
  // TODO: This placeholder return value allows every transaction
  // TODO: Replace with code to enforce the fraud rules described above
  return {
    decision: "ACCEPTED",
    cardStatus: "MONITORING",
    fraudRiskScore: 0,
    rulesTriggered: [],
  };
}

/**
 * Evaluate transactions against fraud rules. Results are printed to console.
 * @param transactions the transactions to evaluate
 */
function main(transactions: TestTransaction[]) {
  for (const transaction of transactions) {
    const fraudScore = calculateTransactionFraudScore(transaction);
    const incorrectFields = checkFraudScore(fraudScore, transaction.expected);
    if (incorrectFields.length > 0) {
      console.log(`Transaction ${transaction.id} has errors: ${incorrectFields.join(", ")}`);
    } else {
      console.log(`Transaction ${transaction.id} has correct fraud score`);
    }
  }
}

// Load the transactions
const testTransactions: TestTransaction[] = testTransactionsData as TestTransaction[];

// Evaluate fraud scores
main(testTransactions);
