package main

import "github.com/ericcornelissen/gofunchecks/internal/utils"

const legalMessage = `

The source code for this software is created by Eric Cornelissen and it is
distributed under an MIT license. See the following for details:

* https://github.com/ericcornelissen/gofunchecks/blob/main/LICENSE
* https://tldrlegal.com/license/mit-license`

func printLegal(p utils.Printer) {
	p.Print(legalMessage)
}
