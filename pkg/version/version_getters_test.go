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

import "testing"

var (
	version1 *Version = NewVersion("testtool", "something", "1.0.0", "Kendall Tauser", "kendall.tauser@securusmonitoring.com")
	version2 *Version = NewVersion("buildtool", "toodles", "1.0.1", "Billy Bob", "billy.bob@securusmonitoring.com")
)

func TestVersion_GetTool(t *testing.T) {
	tests := []struct {
		name string
		v    *Version
		want string
	}{
		{
			name: "1",
			v:    version1,
			want: "testtool",
		},
		{
			name: "2",
			v:    version2,
			want: "buildtool",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetTool(); got != tt.want {
				t.Errorf("Version.GetTool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_GetUse(t *testing.T) {
	tests := []struct {
		name string
		v    *Version
		want string
	}{
		{
			name: "1",
			v:    version1,
			want: "something",
		},
		{
			name: "2",
			v:    version2,
			want: "toodles",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetUse(); got != tt.want {
				t.Errorf("Version.GetUse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_GetVersion(t *testing.T) {
	tests := []struct {
		name string
		v    *Version
		want string
	}{
		{
			name: "1",
			v:    version1,
			want: "1.0.0",
		},
		{
			name: "2",
			v:    version2,
			want: "1.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetVersion(); got != tt.want {
				t.Errorf("Version.GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_GetAuthor(t *testing.T) {
	tests := []struct {
		name string
		v    *Version
		want string
	}{
		{
			name: "1",
			v:    version1,
			want: "Kendall Tauser",
		},
		{
			name: "2",
			v:    version2,
			want: "Billy Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetAuthor(); got != tt.want {
				t.Errorf("Version.GetAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_GetAuthorEmail(t *testing.T) {
	tests := []struct {
		name string
		v    *Version
		want string
	}{
		{
			name: "1",
			v:    version1,
			want: "kendall.tauser@securusmonitoring.com",
		},
		{
			name: "2",
			v:    version2,
			want: "billy.bob@securusmonitoring.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetAuthorEmail(); got != tt.want {
				t.Errorf("Version.GetAuthorEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
