package consensus

import (
	cmn "github.com/tendermint/tmlibs/common"
)

// kind of arbitrary
var Spec     = "1" // async
var Major    = "0" //
var Minor    = "2" // replay refactor
var Revision = "2" // validation -> commit

var Version  = cmn.Fmt("v%s/%s.%s.%s", Spec, Major, Minor, Revision)
