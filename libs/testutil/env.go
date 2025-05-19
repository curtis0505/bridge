package testutil

type TestENV string

const (
	DEV  = TestENV("dev")
	DEV2 = TestENV("dev2")
	DQ   = TestENV("dq")
	LIVE = TestENV("prod")
)
