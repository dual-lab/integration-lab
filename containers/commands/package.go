// Copyright 2022 dmike16. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package commads provides cli supported command
package commands

import "github.com/urfave/cli/v2"


var Commands []*cli.Command
func init() {
  Commands = append(Commands, box)
}
