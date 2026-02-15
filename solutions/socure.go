package solutions

import (
	"encoding/json"
	"fmt"
)

const (
	MaxInputBytes  = 100 * 1024
	MaxEnrichments = 200
	MaxSignalsKeys = 200
	MaxStringLen   = 256
)

type ValidationError struct {
	Path    string
	Code    string
	Message string
}

type RiskContextBundle struct {
	TenantID    string            `json:"tenantId"`
	EntityKeys  map[string]string `json:"entityKeys"`
	Signals     map[string]any    `json:"signals"`
	Enrichments []Enrichment      `json:"enrichments"`
	Extras      map[string]any    `json:"extras"`
}

type Enrichment struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Version   string `json:"version"`
	LatencyMs int    `json:"latencyMs"`
}

// ValidateAndNormalizeLLMBundle parses and validates an LLM provided JSON bundle.
// It returns a normalized bundle plus validation errors.
// If JSON is invalid, return an empty bundle and one error with Code "invalid_json".
func ValidateAndNormalizeLLMBundle(input string) (RiskContextBundle, []ValidationError) {
	
	var output map[string]any
	
	err := json.Unmarshal([]byte(input), &output)
	
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(output)
	}

	return RiskContextBundle{}, nil
}

func addErr(errs *[]ValidationError, path, code, msg string) {
	*errs = append(*errs, ValidationError{Path: path, Code: code, Message: msg})
}

func isJSONObject(v any) (map[string]any, bool) {
	m, ok := v.(map[string]any)
	return m, ok
}

func asString(v any) (string, bool) {
	s, ok := v.(string)
	return s, ok
}

func asStringMap(v any) (map[string]string, bool) {
	raw, ok := v.(map[string]any)
	if !ok {
		return nil, false
	}
	out := make(map[string]string, len(raw))
	for k, vv := range raw {
		s, ok := vv.(string)
		if !ok {
			return nil, false
		}
		out[k] = s
	}
	return out, true
}

func marshalUnmarshal[T any](v any) (T, error) {
	var out T
	b, err := json.Marshal(v)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(b, &out)
	return out, err
}
