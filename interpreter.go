package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func helloWorld() string {
	vm := otto.New()
	vm.Run(`

	async function asyncAwait() {
		console.log("Knock, knock!");
	  
		await delay(1000);
		console.log("Who's there?");
	  
		await delay(1000);
		console.log("async/await!");
	}

	asyncAwait();

		function wow() {
			console.log("wow")
		}

		wow()

		console.log("hello world")
		hello = 'hello world!'

	`)
	if value, err := vm.Get("hello"); err == nil {
		if value_str, err := value.ToString(); err == nil {
			fmt.Println("testing testing", value_str, err)
			return value_str
		}
	}
	return ""
}
