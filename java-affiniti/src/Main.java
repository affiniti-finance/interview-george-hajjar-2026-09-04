import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.File;
import java.util.Arrays;
import java.util.List;

/**
 * Main program to validate the FraudAnalyzer implementation.
 * Loads test transactions and compares actual vs expected results.
 */
public class Main {
    public static void main(String[] args) throws Exception {
        // Load test transactions
        ObjectMapper mapper = new ObjectMapper();
        Transaction[] transactionsArray = mapper.readValue(
            new File("test-transactions.json"),
            Transaction[].class
        );
        List<Transaction> transactions = Arrays.asList(transactionsArray);

        System.out.println("Running Fraud Detection Tests...\n");

        // Create analyzer
        FraudAnalyzer analyzer = new FraudAnalyzer();

        int passed = 0;
        int failed = 0;

        // Process each transaction
        for (Transaction txn : transactions) {
            FraudScore actual = analyzer.calculateTransactionFraudScore(txn);
            Transaction.ExpectedResult expected = txn.expected();

            boolean testPassed = validateResult(txn, actual, expected);

            if (testPassed) {
                passed++;
                System.out.println("✓ " + txn.id() + " PASSED");
            } else {
                failed++;
                System.out.println("✗ " + txn.id() + " FAILED");
                printMismatch(txn, actual, expected);
            }
        }

        // Print summary
        System.out.println("\n" + "=".repeat(50));
        System.out.println("Results: " + passed + " passed, " + failed + " failed");
        System.out.println("=".repeat(50));

        // Exit with appropriate status code
        System.exit(failed == 0 ? 0 : 1);
    }

    private static boolean validateResult(Transaction txn, FraudScore actual, Transaction.ExpectedResult expected) {
        // Check decision
        if (!actual.decision().name().equals(expected.decision())) {
            return false;
        }

        // Check card status
        if (!actual.cardStatus().name().equals(expected.cardStatus())) {
            return false;
        }

        // Check fraud risk score
        if (actual.fraudRiskScore() != expected.fraudRiskScore()) {
            return false;
        }

        // Check rules triggered
        List<String> expectedRules = Arrays.asList(expected.rulesTriggered());
        if (actual.rulesTriggered().size() != expectedRules.size()) {
            return false;
        }

        if (!actual.rulesTriggered().containsAll(expectedRules)) {
            return false;
        }

        return true;
    }

    private static void printMismatch(Transaction txn, FraudScore actual, Transaction.ExpectedResult expected) {
        System.out.println("  Card: ****" + txn.cardNumber().substring(12));
        System.out.println("  Amount: " + txn.amount() + ", Country: " + txn.country() + ", Category: " + txn.category());

        if (!actual.decision().name().equals(expected.decision())) {
            System.out.println("  Decision: expected " + expected.decision() + ", got " + actual.decision());
        }

        if (!actual.cardStatus().name().equals(expected.cardStatus())) {
            System.out.println("  Card Status: expected " + expected.cardStatus() + ", got " + actual.cardStatus());
        }

        if (actual.fraudRiskScore() != expected.fraudRiskScore()) {
            System.out.println("  Fraud Score: expected " + expected.fraudRiskScore() + ", got " + actual.fraudRiskScore());
        }

        List<String> expectedRules = Arrays.asList(expected.rulesTriggered());
        if (actual.rulesTriggered().size() != expectedRules.size() || !actual.rulesTriggered().containsAll(expectedRules)) {
            System.out.println("  Rules Triggered: expected " + expectedRules + ", got " + actual.rulesTriggered());
        }

        System.out.println();
    }
}
