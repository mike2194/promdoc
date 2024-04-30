package generate

import "fmt"

// CSV finds all rules at the given path and its
// subdirectories and generates a CSV file of the found rules.
func CSV(path string, input string) (string, error) {
	ruleGroups, err := getRuleGroups(path, input)
	if err != nil {
		return "", fmt.Errorf("get rule groups: %w", err)
	}

	document := "Name,RuleGroup,Summary,Description,Severity,Expr,For,Runbook,Dashboard\n"
	for _, ruleGroup := range ruleGroups {
		for _, rule := range ruleGroup.Rules {
			var description string
			if val, ok := rule.Annotations["description"]; ok {
				description = trimText(val)
			} else if val, ok := rule.Annotations["message"]; ok {
				description = trimText(val)
			}

			expr := trimText(rule.Expr)
			summary := rule.Annotations["summary"]
			severity := rule.Labels["severity"]
			runbookURL := rule.Annotations["runbook_url"]
			dashboardURL := rule.Annotations["dashboard_url"]

			document += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s", rule.Alert, ruleGroup.Name, summary, description, severity, expr, rule.For, runbookURL, dashboardURL)
			document += "\n"
		}
	}

	return document, nil
}
