//
// main.go
// Copyright (C) 2023 rmelo <Ricardo Melo <rmelo@ludia.com>>
//
// Distributed under terms of the MIT license.
//

package main

import (
    "github.com/cropalato/grpc-simple-task/core"
    "os"
)

func main() {
    nodeType := os.Args[1]
    switch nodeType {
    case "master":
        core.GetMasterNode().Start()
    case "worker":
        core.GetWorkerNode().Start()
    default:
      panic("invalid node type. " + os.Args[0])
    }
}
