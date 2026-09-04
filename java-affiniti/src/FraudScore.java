import java.util.List;

public record FraudScore(
    Decision decision,
    CardStatus cardStatus,
    int fraudRiskScore,
    List<String> rulesTriggered
) {}
