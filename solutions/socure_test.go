package solutions

import (
	"testing"
)

func TestVisible_HappyPath_WithExtras(t *testing.T) {
	input := `{
		"tenantId": "cust_acme",
		"entityKeys": { "userId": "u1", "deviceId": "d1" },
		"signals": { "network": { "ip": "203.0.113.1" } },
		"enrichments": [
			{ "name": "device_risk", "status": "ok", "version": "2026-01-01", "latencyMs": 12, "weirdField": 99 }
		],
		"randomModelThoughts": "I think risk is low"
	}`

	got, errs := ValidateAndNormalizeLLMBundle(input)
	if len(errs) != 0 {
		t.Fatalf("expected no errors, got: %+v", errs)
	}
	if got.TenantID != "cust_acme" {
		t.Fatalf("tenantId mismatch: %q", got.TenantID)
	}
	if got.EntityKeys["userId"] != "u1" || got.EntityKeys["deviceId"] != "d1" {
		t.Fatalf("entityKeys mismatch: %+v", got.EntityKeys)
	}
	if got.Extras == nil || got.Extras["randomModelThoughts"] == nil {
		t.Fatalf("expected extras to include randomModelThoughts, got: %+v", got.Extras)
	}
	if got.Extras["enrichments"] == nil {
		t.Fatalf("expected enrichment extras stored under extras.enrichments, got: %+v", got.Extras)
	}
	if len(got.Enrichments) != 1 || got.Enrichments[0].Name != "device_risk" {
		t.Fatalf("enrichments mismatch: %+v", got.Enrichments)
	}
}

func TestVisible_InvalidJSON_ReturnsInvalidJsonOnly(t *testing.T) {
	input := `Sure! Here is the JSON:
{ "tenantId": "cust_acme", "entityKeys": { "userId": "u1" } }`

	_, errs := ValidateAndNormalizeLLMBundle(input)
	if len(errs) == 0 {
		t.Fatalf("expected errors")
	}
	if errs[0].Code != "invalid_json" {
		t.Fatalf("expected invalid_json, got: %+v", errs)
	}
}
