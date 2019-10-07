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

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Loadtestdata() dashboard {

	testdata := "testdash.json"

	jsonFile, err := os.Open(testdata)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	var dash = dashboard{}
	json.Unmarshal(b, &dash)

	return dash
}
func TestParsePanel(t *testing.T) {
	dash := Loadtestdata()
	p := panel{
		ID:           2,
		Title:        "Counter",
		Type:         "graph",
		Description:  "This is a counter Panel",
		Bars:         false,
		DashLength:   10,
		Dashes:       false,
		Fill:         1,
		FillGradient: 0,

		Legend: panelLegend{
			AlignAsTable: true,
			Avg:          false,
			Current:      false,
			HideEmpty:    true,
			HideZero:     true,
			Max:          false,
			Min:          false,
			RightSide:    true,
			Show:         true,
			SideWidth:    1,
			Total:        false,
			Values:       false,
		},

		Lines:         true,
		LinesWidth:    1,
		NullPointMode: "null",
		Percentage:    false,
		PointRadius:   2,
		Points:        false,
		Render:        "flot",
		SpaceLength:   10,
		Stack:         false,
		SteppedLine:   false,
		XAxis: panelXAxis{
			//Buckets: nil,
			Mode: "time",
			Name: "",
			Show: true,
			//Values: nil ,
		},
		YAxes: []panelYAxes{
			{
				Decimals: 6,
				Format:   "short",
				Label:    "LeftY Label",
				LogBase:  1,
				//Max:      nil,
				//Min:      nil,
				Show: true,
			},
			{
				Decimals: 6,
				Format:   "short",
				Label:    "Right Y Label",
				LogBase:  1,
				//Max:      nil,
				//Min:      nil,
				Show: true,
			},
		},

		YAxis: panelYAxis{
			Align:      true,
			AlignLevel: 1,
		},

		GridPos: panelGridPos{
			X: 0,
			Y: 1,
			H: 9,
			W: 12,
		}, //GridPos

		Targets: []panelTarget{
			{
				Expr:  "sum(rate(process_cpu_seconds_total [1m]))",
				RefID: "A",
			},
		}, //Targets

	}

	if p.GridPos != dash.Panels[0].GridPos {
		t.Errorf("Panel Grid Posistions do not match, expected %v got %v ", p.GridPos, dash.Panels[0].GridPos)
	}
	if p.Description != dash.Panels[0].Description {
		t.Errorf("Panel Descriptions do not match, expected %v got %v ", p.Description, dash.Panels[0].Description)
	}
	if p.ID != dash.Panels[0].ID {
		t.Errorf("Panel IDs do not match, expected %v got %v ", p.ID, dash.Panels[0].ID)
	}
	if p.Targets[0] != dash.Panels[0].Targets[0] {
		t.Errorf("Panel Targets do not match, expected %v got %v ", p.Targets[0], dash.Panels[0].Targets[0])
	}
	if p.Title != dash.Panels[0].Title {
		t.Errorf("Panel Titles do not match, expected %v got %v ", p.Title, dash.Panels[0].Title)
	}
	if p.Type != dash.Panels[0].Type {
		t.Errorf("Panel Types do not match, expected %v got %v ", p.Type, dash.Panels[0].Type)
	}
	if p.Legend != dash.Panels[0].Legend {
		t.Errorf("Panel Legends do not match, expected %v got %v ", p.Legend, dash.Panels[0].Legend)
	}
	if p.Bars != dash.Panels[0].Bars {
		t.Errorf("Panel Bars Option do not match, expected %v got %v ", p.Bars, dash.Panels[0].Bars)
	}
	if p.DashLength != dash.Panels[0].DashLength {
		t.Errorf("Panel DashLength Option do not match, expected %v got %v ", p.DashLength, dash.Panels[0].DashLength)
	}
	if p.Dashes != dash.Panels[0].Dashes {
		t.Errorf("Panel Dashes Option do not match, expected %v got %v ", p.Dashes, dash.Panels[0].Dashes)
	}
	if p.Fill != dash.Panels[0].Fill {
		t.Errorf("Panel Fill Option do not match, expected %v got %v ", p.Fill, dash.Panels[0].Fill)
	}
	if p.FillGradient != dash.Panels[0].FillGradient {
		t.Errorf("Panel FillGradient Option do not match, expected %v got %v ", p.FillGradient, dash.Panels[0].FillGradient)
	}
	if p.Lines != dash.Panels[0].Lines {
		t.Errorf("Panel Lines Option do not match, expected %v got %v ", p.Lines, dash.Panels[0].Lines)
	}
	if p.LinesWidth != dash.Panels[0].LinesWidth {
		t.Errorf("Panel LinesWidth Option do not match, expected %v got %v ", p.LinesWidth, dash.Panels[0].LinesWidth)
	}
	if p.NullPointMode != dash.Panels[0].NullPointMode {
		t.Errorf("Panel NullPointMode Option do not match, expected %v got %v ", p.NullPointMode, dash.Panels[0].NullPointMode)
	}
	if p.Percentage != dash.Panels[0].Percentage {
		t.Errorf("Panel Percentage Option do not match, expected %v got %v ", p.Percentage, dash.Panels[0].Percentage)
	}
	if p.PointRadius != dash.Panels[0].PointRadius {
		t.Errorf("Panel PointRadius Option do not match, expected %v got %v ", p.PointRadius, dash.Panels[0].PointRadius)
	}
	if p.Points != dash.Panels[0].Points {
		t.Errorf("Panel Points Option do not match, expected %v got %v ", p.Points, dash.Panels[0].Points)
	}
	if p.Render != dash.Panels[0].Render {
		t.Errorf("Panel Render Option do not match, expected %v got %v ", p.Render, dash.Panels[0].Render)
	}
	if p.SpaceLength != dash.Panels[0].SpaceLength {
		t.Errorf("Panel Render Option do not match, expected %v got %v ", p.SpaceLength, dash.Panels[0].SpaceLength)
	}
	if p.Stack != dash.Panels[0].Stack {
		t.Errorf("Panel Render Option do not match, expected %v got %v ", p.Stack, dash.Panels[0].Stack)
	}
	if p.SteppedLine != dash.Panels[0].SteppedLine {
		t.Errorf("Panel SteppedLine Option do not match, expected %v got %v ", p.SteppedLine, dash.Panels[0].SteppedLine)
	}

	if p.YAxes[0] != dash.Panels[0].YAxes[0] {
		t.Errorf("Panel Yaxes do not match, expected %v got %v ", p.YAxes, dash.Panels[0].YAxes)
	}
	if p.YAxes[1] != dash.Panels[0].YAxes[1] {
		t.Errorf("Panel Yaxes do not match, expected %v got %v ", p.YAxes, dash.Panels[0].YAxes)
	}
	if p.YAxis != dash.Panels[0].YAxis {
		t.Errorf("Panel YAxis do not match, expected %v got %v ", p.YAxis, dash.Panels[0].YAxis)
	}
}
