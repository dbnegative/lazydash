package main

//RefreshIntervals for quick reference
var RefreshIntervals = []string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"}

//TimeOptions for quick referenece
var TimeOptions = []string{"5m", "15m", "1h", "3h", "6h", "12h", "24h", "2d", "3d", "4d", "7d", "30d"}

//TimeRange contains a range of time
type TimeRange struct {
	From string `json:"from"` //: "now-6h",
	To   string `json:"to"`   //: "now"
}

//Dashboard holds all other Grafana sub containers
type Dashboard struct {
	ID            int         `json:"id,omitempty"`
	UID           string      `json:"uid,ommitempty"`
	Title         string      `json:"title,omitempty"` // "New dashboard",
	Tags          []string    `json:"tags,omitempty"`
	TimeZone      string      `json:"timezone,omitempty"`
	Editable      bool        `json:"editable,omitempty"`
	HideControls  string      `json:"hideControls,omitempty"`
	GraphToolTip  int         `json:"graphTooltip,omitempty"`
	Panels        []Panel     `json:"panels,omitempty"`
	Time          TimeRange   `json:"time,omitempty"`
	TimePicker    TimePicker  `json:"timepicker,omitempty"`
	Templating    Templating  `json:"templating,omitempty"`
	Annotations   Annotations `json:"annotations,omitempty"`
	Refresh       string      `json:"refresh,omitempty"`
	SchemaVersion int         `json:"schemaVersion,omitempty"`
	Version       int         `json:"version,omitempty"`
	Links         []string    `json:"links,omitempty"`
	Style         string      `json:"style,omitempty"`
}

//Panel contains all Grafana panel attributes
type Panel struct {
	GridPos       GridPos     `json:"gridPos"`
	Type          string      `json:"type"`
	Title         string      `json:"title"`
	ID            int         `json:"id"`
	Mode          string      `json:"mode"`
	Content       string      `json:"content,omitempty"`
	Targets       []Target    `json:"targets,omitempty"`
	Description   string      `json:"description,omitempty"`
	Legend        PanelLegend `json:"legend,omitempty"`
	Bars          bool        `json:"bars,omitempty"`
	DashLength    int         `json:"dashLength"`
	Dashes        bool        `json:"dashes,omitempty"`
	Fill          int         `json:"fill,omitempty"`
	FillGradient  int         `json:"fillGradient,omitempty"`
	Lines         bool        `json:"lines,omitempty"`
	LinesWidth    int         `json:"linewidth,omitempty"`
	NullPointMode string      `json:"nullPointMode,omitempty"`
	Percentage    bool        `json:"percentage,omitempty"`
	PointRadius   int         `json:"pointradius,omitempty"`
	Points        bool        `json:"points,omitempty"`
	Render        string      `json:"renderer,omitempty"`
	//SeriesOverrides "seriesOverrides": [],
	SpaceLength int          `json:"spaceLength,omitempty"`
	Stack       bool         `json:"stack,omitempty"`
	SteppedLine bool         `json:"steppedLine,omitempty"`
	YAxes       []PanelYAxes `json:"yaxes,omitempty"`
	YAxis       PanelYAxis   `json:"yaxis,omitempty"`
	XAxis       PanelXAxis   `json:"xaxis,omitempty"`
	ToolTip     PanelToolTip `json:"tooltip,omitempty"`
	Options     PanelOptions `json:"options,omitempty"`
}

type PanelFieldOptionsThershold struct {
	Value int    `json:"value,omitempty"`
	Color string `json:"color,omitempty"`
}

type PanelFieldOptionsDefaults struct {
	Thresholds []PanelFieldOptionsThershold `json:"thresholds,omitempty"`
	//Mappings[]
	Unit string `json:"unit,omitempty"`
}

type PanelFieldOptions struct {
	Values   bool                      `json:"values,omitempty"`
	Calcs    []string                  `json:"calcs,omitempty"`
	Defaults PanelFieldOptionsDefaults `json:"defaults,omitempty"`
	//Override": {}
}

type PanelOptions struct {
	ShowThresholdMarkers bool              `json:"showThresholdMarkers,omitempty"`
	ShowThresholdLabels  bool              `json:"showThresholdLabels,omitempty"`
	FieldOptions         PanelFieldOptions `json:"fieldOptions,omitempty"`
	Orientation          string            `json:"orientation,omitempty"`
}

//PanelLegend contains all Legend options
type PanelLegend struct {
	Avg          bool `json:"avg,omitempty"`
	Current      bool `json:"current,omitempty"`
	Max          bool `json:"max,omitempty"`
	Min          bool `json:"min,omitempty"`
	Show         bool `json:"show,omitempty"`
	Total        bool `json:"total,omitempty"`
	Values       bool `json:"values,omitempty"`
	HideEmpty    bool `json:"hideEmpty,omitempty"`
	HideZero     bool `json:"hideZero,omitempty"`
	RightSide    bool `json:"rightSide,omitempty"`
	SideWidth    int  `json:"sideWidth,omitempty"`
	AlignAsTable bool `json:"alignAsTable,omitempty"`
}

//PanelToolTip contains panel tooltip options
type PanelToolTip struct {
	Shared    bool   `json:"shared,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	ValueType string `json:"value_type,omitempty"`
}

//PanelXAxis contains panel xaxis options
type PanelXAxis struct {
	Buckets int    `json:"buckets,omitempty"`
	Mode    string `json:"mode,omitempty"`
	Name    string `json:"name,omitempty"`
	Show    bool   `json:"show,omitempty"`
}

//PanelYAxes contains Yaxes options
type PanelYAxes struct {
	Decimals int    `json:"decimals,omitempty"`
	Format   string `json:"format,omitempty"`
	Label    string `json:"label,omitempty"`
	LogBase  int    `json:"logBase,omitempty"`
	Max      int    `json:"max,omitempty"`
	Min      int    `json:"min,omitempty"`
	Show     bool   `json:"show,omitempty"`
}

//PanelYAxis contains panel yaxis options
type PanelYAxis struct {
	Align      bool `json:"align,omitempty"`
	AlignLevel int  `json:"alignLevel,omitempty"`
}

//GridPos holds a panels grid posistion
type GridPos struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	H int `json:"h,omitempty"`
	W int `json:"w,omitempty"`
}

//Target contains a Panel metrics experssion and order
type Target struct {
	Expr         string `json:"expr,ommitempty"`        // "sum(rate(process_cpu_seconds_total [1m]))",
	RefID        string `json:"refId,omitempty"`        //"A"
	LegendFormat string `json:"legendFormat,omitempty"` //: "{{state}}"
}

//TimePicker comtains all attribuyes used to set dashboard time options
type TimePicker struct {
	Collapse         bool     `json:"collapse,omitempty"`
	Enable           bool     `json:"enable,omitempty"`
	Notice           bool     `json:"notice,omitempty"`
	Now              bool     `json:"now,omitempty"`
	RefreshIntervals []string `json:"refresh_intervals,omitempty"`
	Status           string   `json:"status,omitempty"`
	TimeOptions      []string `json:"time_options"`
	Type             string   `json:"type,omitempty"`
	NowDelay         string   `json:"nowDelay,omitempty"`
}

//TemplatingVarState contains the current state of the dashboard variable
type TemplatingVarState struct {
	Tags  []string `json:"tags,omitempty"`
	Text  string   `json:"text,omitempty"`
	Value string   `json:"value,omitempty"`
}

//TemplatingOption contains all Templating option attributes
type TemplatingOption struct {
	Selected bool   `json:"selected,omitempty"`
	Text     string `json:"text,omitempty"`
	Value    string `json:"value,omitempty"`
}

//TemplatingVar defines a single Dashboard variable
type TemplatingVar struct {
	AllFormat      string             `json:"allFormat,omitempty"` //: "wildcard",
	Current        TemplatingVarState `json:"current,omitempty"`
	Datasource     string             `json:"datasource,omitempty"`
	Definition     string             `json:"definition,omitempty"`
	IncludeAll     bool               `json:"includeAll,omitempty"`
	Name           string             `json:"name,omitempty"`
	Options        []TemplatingOption `json:"options,omitempty"`
	Query          string             `json:"query,omitempty"`
	Regex          string             `json:"regex,omitempty"`
	Refresh        int                `json:"refresh,omitempty"`
	Type           string             `json:"type,omitempty"`
	Multi          bool               `json:"multi,omitempty"`
	MultiFormat    string             `json:"multiFormat,omitempty"`
	SkipURLSync    bool               `json:"skipUrlSync,omitempty"`
	Sort           int                `json:"sort,omitempty"`
	Tags           []string           `json:"tags,omitempty"`
	TagValuesQuery string             `json:"tagValuesQuery,omitempty"`
	TagsQuery      string             `json:"tagsQuery,omitempty"`
	UseTags        bool               `json:"useTags,omitempty"`
}

//Templating defines all templated dashbaord variables
type Templating struct {
	Enable bool            `json:"enable,omitempty"`
	List   []TemplatingVar `json:"list,omitempty"`
}

//Annotation defines a single Dashboard annotation
type Annotation struct {
	Datasource string `json:"datasource,omitempty"`
	Enable     bool   `json:"enable,omitempty"`
	Hide       bool   `json:"hide,omitempty"`
	IconColor  string `json:"iconColor,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
}

//Annotations contains a list of annotations
type Annotations struct {
	List []Annotation `json:"list,omitempty"`
}
