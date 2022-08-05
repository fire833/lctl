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

func TestVersion_Version(t *testing.T) {
	type fields struct {
		v      string
		author string
		tool   string
		email  string
		use    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				v:      "0.0.1",
				tool:   "tooling",
				use:    "a tool to get stuff done",
				author: "Kendall Tauser",
				email:  "kendall.tauser@securusmonitoring.com",
			},
			want: "tooling (a tool to get stuff done) \nVersion: 0.0.1 \nAuthor: Kendall Tauser <kendall.tauser@securusmonitoring.com> \n",
		},
		{
			name: "2",
			fields: fields{
				v:      "1.13.4",
				author: "Forrest Gump",
				email:  "",
				tool:   "",
				use:    "A testing tool",
			},
			want: "A testing tool \nVersion: 1.13.4 \nAuthor: Forrest Gump \n",
		},
		{
			name: "3",
			fields: fields{
				v:     "",
				tool:  "tool",
				email: "admin@admin.org",
			},
			want: "tool \n",
		},
		{
			name: "4",
			fields: fields{
				v: "2.3.4",
			},
			want: "2.3.4 \n",
		},
		{
			name: "5",
			fields: fields{
				v:    "4.5.6",
				tool: "something",
			},
			want: "something - 4.5.6 \n",
		},
		{
			name: "6",
			fields: fields{
				v:      "1.0pre1",
				author: "Somebody",
			},
			want: "1.0pre1 \nAuthor: Somebody \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Version{
				v:      tt.fields.v,
				author: tt.fields.author,
				tool:   tt.fields.tool,
				email:  tt.fields.email,
				use:    tt.fields.use,
			}
			if got := v.Version(); got != tt.want {
				t.Errorf("Version.Version() = %v, want %v", got, tt.want)
			}
		})
	}
}
