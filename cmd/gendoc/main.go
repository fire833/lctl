/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package main

import (
	_ "github.com/fire833/lctl/pkg/commands/account"
	_ "github.com/fire833/lctl/pkg/commands/card"
	_ "github.com/fire833/lctl/pkg/commands/funding"
	"github.com/fire833/lctl/pkg/commands/lctl"
	_ "github.com/fire833/lctl/pkg/commands/status"
	_ "github.com/fire833/lctl/pkg/commands/transaction"
	"github.com/spf13/cobra/doc"
)

func main() {
	doc.GenMarkdownTree(lctl.Cmd, "Documentation")
}
