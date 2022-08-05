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

package version

import "github.com/fire833/lctl/pkg"

type Version struct {
	v      string
	author string
	tool   string
	use    string
	email  string
	commit string
	date   string
}

func NewVersion(tool, use, version, author, email string) *Version {
	return &Version{
		v:      version,
		author: author,
		tool:   tool,
		use:    use,
		email:  email,
		commit: pkg.Commit,
		date:   pkg.BuildTime,
	}
}

// Parses contents of version object into complete string that can be printed
// to stdout when the command is run with the -v flag or version subcommand.
func (v *Version) Version() string {
	var versionstr string

	if v.tool != "" && v.use != "" && v.v != "" {
		versionstr += (v.tool + " (" + v.use + ") \n" + "Version: " + v.v + " \n")
	} else if v.tool == "" && v.use != "" && v.v != "" {
		versionstr += (v.use + " \n" + "Version: " + v.v + " \n")
	} else if v.tool == "" && v.use == "" && v.v != "" {
		versionstr += (v.v + " \n")
	} else if v.tool != "" && v.use == "" && v.v != "" {
		versionstr += (v.tool + " - " + v.v + " \n")
	} else if v.tool != "" && v.use == "" && v.v == "" {
		versionstr += (v.tool + " \n")
	}

	if v.author != "" && v.email != "" {
		versionstr += ("Author: " + v.author + " <" + v.email + "> \n")
	} else if v.author != "" && v.email == "" {
		versionstr += ("Author: " + v.author + " \n")
	}

	return versionstr

}

func (v *Version) GetTool() string        { return v.tool }
func (v *Version) GetUse() string         { return v.use }
func (v *Version) GetVersion() string     { return v.v }
func (v *Version) GetAuthor() string      { return v.author }
func (v *Version) GetAuthorEmail() string { return v.email }
func (v *Version) GetCommit() string      { return v.commit }
func (v *Version) GetDate() string        { return v.date }
