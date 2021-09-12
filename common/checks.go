package common

type RuleContext struct {
	Interface int64
}

type Rule = func(RuleContext) (bool, error)

func All(ctx RuleContext, rules []Rule) (bool, error) {
	var err error
	result := true

	for _, rule := range rules {
		result, err = rule(ctx)
		if !result {
			return false, err
		}
	}

	return result, err
}
