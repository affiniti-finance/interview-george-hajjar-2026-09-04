import java.util.ArrayList;

/**
 * Fraud detection system for credit card transactions.
 *
 * TODO: Implement the calculateTransactionFraudScore method below.
 */
public class FraudAnalyzer {

    /**
     * Analyzes a transaction and calculates its fraud score.
     *
     * Requirements:
     *
     * 1. Maintain state for each credit card across all transactions:
     *    - Transaction history
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
     * 4. Return a FraudScore object with:
     *    - decision: ACCEPTED or DECLINED
     *    - cardStatus: MONITORING or AT_RISK
     *    - fraudRiskScore: current cumulative score for the card
     *    - rulesTriggered: list of rule display names that triggered
     *
     * @param transaction The transaction to analyze
     * @return FraudScore containing the decision and fraud analysis results
     */
    public FraudScore calculateTransactionFraudScore(Transaction transaction) {
        // TODO: Implement fraud detection logic here

        return new FraudScore(
            Decision.ACCEPTED,
            CardStatus.MONITORING,
            0,
            new ArrayList<>()
        );
    }
}
