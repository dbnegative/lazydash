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
