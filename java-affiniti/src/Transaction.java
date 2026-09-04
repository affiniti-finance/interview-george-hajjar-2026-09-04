import com.fasterxml.jackson.annotation.JsonProperty;

public record Transaction(
    String id,
    String cardNumber,
    int amount,
    String country,
    String timestamp,
    String category,
    @JsonProperty("expected") ExpectedResult expected
) {
    public record ExpectedResult(
        String decision,
        String cardStatus,
        int fraudRiskScore,
        String[] rulesTriggered
    ) {}
}
