func TestCookbookDenom(t *testing.T) {
	valid1, _ := CookbookDenom("test", "pylons")
	valid2, _ := CookbookDenom("a12341234", "pylons")
	invalid1 := "1234/567"
	invalid2 := "pylons"


func TestIBCDenom(t *testing.T) {
	valid1Trace := ibctypes.ParseDenomTrace("transfer/channel-0/uatom")
	valid1 := valid1Trace.IBCDenom()
	valid2 := "ibc/529ba5e3e86ba7796d7caab4fc02728935fbc75c0f7b25a9e611c49dd7d68a35"

	invalid1 := "pylons"

		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			ok := IsIBCDenomRepresentation(tc.denom)
			fmt.Printf("ok = %v \n", ok)
			require.Equal(t, tc.is, ok)
		})}
