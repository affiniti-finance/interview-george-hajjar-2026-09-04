public enum FraudRule {
    UNUSUAL_LOCATION("Unusual Location"),
    UNUSUAL_AMOUNT("Unusual Amount"),
    SUSPICIOUS_CATEGORY_SEQUENCE("Suspicious Category Sequence");

    private final String displayName;

    FraudRule(String displayName) {
        this.displayName = displayName;
    }

    public String getDisplayName() {
        return displayName;
    }
}
