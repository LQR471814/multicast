package reset

func Win(intf int) error {
	_, err := Win_routing_reset(intf)
	if err != nil {
		return err
	}

	_, err = Win_firewall_reset()

	return err
}
