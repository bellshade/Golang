package main

import (
	"context"
	"fmt"
)

func doSomethingWithContext(ctx context.Context){
	fmt.Printf("This is : %v \n",ctx)
}

func doAnotherFromSomething(ctx context.Context){
	// menghasilkan anotherKey from doAnotherCtx anotherKey : anotherVal
	fmt.Printf("key from doAnotherCtx anotherKey : %s \n", ctx.Value("key"))
}

func doSomethingContextWithValue(ctx context.Context){
	fmt.Printf("key from ctxWithVal : %s \n", ctx.Value("key"))
	doAnotherCtx := context.WithValue(ctx, "key", "anotherVal")
	doAnotherFromSomething(doAnotherCtx)
	// menghasilkan key from ctxWithVal : val
	fmt.Printf("key from ctxWithVal : %s \n", ctx.Value("key"))
}

func main() {
	ctxTODO := context.TODO()
	ctxBackground := context.Background()
	// menghasilkan This is : context.TODO
	doSomethingWithContext(ctxTODO)
	// menghasilkan This is : context.Background
	doSomethingWithContext(ctxBackground)
	ctxWithVal := context.WithValue(ctxBackground, "key", "val")
	// menghasilkan key from ctx_with_val : val
	doSomethingContextWithValue(ctxWithVal)
}