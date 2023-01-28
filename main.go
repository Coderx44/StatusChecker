package main

import "context"

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}

type httpChecker struct {
}

func main() {
	checkHttp := httpChecker{}
	_ = checkHttp

}
