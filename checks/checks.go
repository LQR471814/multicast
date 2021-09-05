package checks

type RuleContext struct {
	Interface int64
}

type Rule = func(RuleContext) (bool, error)

func Check_Win(ctx RuleContext) (bool, error) {
	return all(ctx, []Rule{ //? Returns true if all rules pass
		win_firewall_check,
		win_routing_check,
	})
}

func all(ctx RuleContext, rules []Rule) (bool, error) {
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
