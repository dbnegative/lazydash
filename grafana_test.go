package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var dash Dashboard
var testdata = "testdash.json"

func init() {
	jsonFile, err := os.Open(testdata)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(b, &dash)
	defer jsonFile.Close()
}

func TestParseDashboard(t *testing.T) {

	d := &Dashboard{
		Title:    "demo",
		TimeZone: "",
		ID:       1,
		Links:    []string{""},
		UID:      "QIbx6hhZz",
		Version:  4,
		Time: TimeRange{
			From: "now-5m",
			To:   "now",
		},
		TimePicker: TimePicker{
			RefreshIntervals: RefreshIntervals,
		},
		SchemaVersion: 20,
		Style:         "dark",
		Tags:          []string{""},
		Templating: Templating{
			List: []TemplatingVar{},
		},
	}

	//fmt.Printf("%v", d)

	if d.ID != dash.ID {
		t.Errorf("ID does not match expected %v got %v ", d.ID, dash.ID)
	}
	if d.Title != dash.Title {
		t.Errorf("Title does not match expected %v got %v ", d.Title, dash.Title)
	}

	if d.SchemaVersion != dash.SchemaVersion {
		t.Errorf("Schema Versions do not match, expected %v got %v ", d.SchemaVersion, d.SchemaVersion)
	}
	if d.Style != dash.Style {
		t.Errorf("Styles do not match, expected %v got %v ", d.Panels[0].Type, dash.Panels[0].Type)
	}

}

//Check that Annotations conform
func TestParseAnnotation(t *testing.T) {

	a := Annotation{
		Datasource: "-- Grafana --",
		Enable:     true,
		Hide:       true,
		IconColor:  "rgba(0, 211, 255, 1)",
		Name:       "Annotations & Alerts",
		Type:       "dashboard",
	}

	if a.Datasource != dash.Annotations.List[0].Datasource {
		t.Errorf("Datasource do not match, expected %v got %v ", a.Datasource, dash.Annotations.List[0].Datasource)
	}
	if a.Enable != dash.Annotations.List[0].Enable {
		t.Errorf("Enabled does not match, expected %v got %v ", a.Enable, dash.Annotations.List[0].Enable)
	}
	if a.Hide != dash.Annotations.List[0].Hide {
		t.Errorf("Hide does not match, expected %v got %v ", a.Hide, dash.Annotations.List[0].Hide)
	}
	if a.IconColor != dash.Annotations.List[0].IconColor {
		t.Errorf("IconColor does not match, expected %v got %v ", a.IconColor, dash.Annotations.List[0].IconColor)
	}
	if a.Name != dash.Annotations.List[0].Name {
		t.Errorf("Name does not match, expected %v got %v ", a.Name, dash.Annotations.List[0].Name)
	}
	if a.Type != dash.Annotations.List[0].Type {
		t.Errorf("Type does not match, expected %v got %v ", a.Type, dash.Annotations.List[0].Type)
	}

}

//Check that Panels conform
func TestParsePanel(t *testing.T) {
	p := Panel{
		ID:           2,
		Title:        "Counter",
		Type:         "graph",
		Description:  "This is a counter Panel",
		Bars:         false,
		DashLength:   10,
		Dashes:       false,
		Fill:         1,
		FillGradient: 0,

		Legend: PanelLegend{
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
		XAxis: PanelXAxis{
			//Buckets: nil,
			Mode: "time",
			Name: "",
			Show: true,
			//Values: nil ,
		},
		YAxes: []PanelYAxes{
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

		YAxis: PanelYAxis{
			Align:      true,
			AlignLevel: 1,
		},

		GridPos: GridPos{
			X: 0,
			Y: 1,
			H: 9,
			W: 12,
		}, //GridPos

		Targets: []Target{
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
