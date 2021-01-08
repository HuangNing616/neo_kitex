package main

import "github.com/koding/kite"

/**
	Author: huang ning
	Date: 2021/01/06
	Func:
 */

func main() {

	k := kite.New("nep_server", "1.0.0")
	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		result := a * a
		return result, nil
	}).DisableAuthentication()
	k.Config.Port = 3636
	k.Run()

}
