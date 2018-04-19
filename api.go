package dotray

func QueryNodes(max int) []string {
	var addrs []string
	if node.seedAddr != "" {
		addrs = append(addrs, node.seedAddr)
	}

	for addr := range node.downstreams {
		if addr != "" && len(addrs) < max {
			addrs = append(addrs, addr)
		}
	}

	for _, seed := range node.seedBackup {
		if seed.addr != "" && len(addrs) < max {
			addrs = append(addrs, seed.addr)
		}
	}

	return addrs
}
