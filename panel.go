//MIT License

//Copyright (c) 2019 Jason Witting

//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:

//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.

//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package main

import "strings"

func NewPanel(title string, gtype string, description string, metric string, format string, x int, y int, expr string, deliminator string, label string) *Panel {
	return &Panel{

		Title:       title,
		Type:        gtype,
		Description: description,

		Targets: []Target{
			{
				Expr:         strings.Replace(expr, deliminator, metric, -1),
				RefID:        "A",
				LegendFormat: label,
			},
		}, //Targets

		GridPos: GridPos{
			X: x,
			Y: y,
			H: 9,
			W: 12,
		}, //GridPos

		YAxes: []PanelYAxes{
			{
				//Decimals: 6,
				Format: format,
				//Label:    "LeftY Label",
				LogBase: 1,
				//Max:      nil,
				//Min:      nil,
				Show: true,
			},
			{
				Show: true,
			},
		},

		XAxis: PanelXAxis{
			//Buckets: nil,
			Mode: "time",
			Name: "",
			Show: true,
			//Values: nil ,
		},

		Options: PanelOptions{
			FieldOptions: PanelFieldOptions{
				Defaults: PanelFieldOptionsDefaults{
					Unit: format,
				},
			},
		},
	}

}
