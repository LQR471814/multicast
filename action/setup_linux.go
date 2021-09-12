package action

func Setup(exec string, intf int) error {
	store.UpdateStore(store.Store{
		Interface: int64(intf),
	})

	return nil
}
