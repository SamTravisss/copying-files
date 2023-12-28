func TestValidateGenesis(t *testing.T) {
	for _, tc := range []struct {
		desc string
		req  *GenesisState
		err  error
	}{

			desc: "Duplicate id for trade",
			req: &GenesisState{
				TradeList: []Trade{
					{Id: uint64(1)},
					{Id: uint64(1)},
				},
			},
			err: fmt.Errorf("duplicated id for trade"),
		},
		{
			desc: "Duplicate purchaseToken for googleInAppPurchaseOrder",
			req: &GenesisState{
				GoogleInAppPurchaseOrderList: []GoogleInAppPurchaseOrder{
					{PurchaseToken: "test"},
					{PurchaseToken: "test"},
				},
			},
