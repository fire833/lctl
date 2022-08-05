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

package funding

import (
	"github.com/fire833/lctl/pkg/commands/lctl"
	"github.com/fire833/lctl/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type cmdOpts struct{}

var (
	cmdVersion *version.Version = version.NewVersion("funding", "Interact with funding sources in the Lithic API.",
		"0.0.1", "Kendall Tauser", "kttpsy@gmail.com")

	cmd *cobra.Command = &cobra.Command{
		Use:        cmdVersion.GetTool(),
		Short:      cmdVersion.GetUse(),
		Version:    cmdVersion.GetVersion(),
		Aliases:    []string{"f", "fund", "funds"},
		SuggestFor: []string{},
		Long: `
   `,
		RunE: exec,
	}

	opts *cmdOpts
)

func exec(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	set := pflag.NewFlagSet(cmdVersion.GetTool(), pflag.ExitOnError)

	o := &cmdOpts{}

	cmd.Flags().AddFlagSet(set)
	lctl.Cmd.AddCommand(cmd)
	opts = o
}
