package providermanager

// Method for getting the chan id based on the name
func getChainID(name string) int64 {
	// TODO(nb): hardcode all the chain id for the chains that'll be used
	networksMap := map[string]int64{
		"ethereum":  int64(1),
		"goerli":    int64(5),
		"localhost": int64(1337),
		"mumbai":    int64(80001),
	}

	return networksMap[name]
}
