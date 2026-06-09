package main

import "context"

const msg = "context root"

func ctx() context.Context {
	ct := context.Background()

	ct = context.WithValue(ct, "msg", msg)

	return ct
}
