// Copyright 2023 Zane(泽恩) <zanekwok73@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zanekwok/miniblog.

package main

import (
	"github.com/zanekwok/miniblog/internal/miniblog"
	_ "go.uber.org/automaxprocs/maxprocs"
	"os"
)

func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
