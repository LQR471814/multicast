package checks

func interface_store_check(ctx RuleContext) (bool, error) {
	if ctx.Interface < 0 {
		return false, nil
	}

	return true, nil
}
