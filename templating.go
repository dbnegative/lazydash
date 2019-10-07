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

//TemplatingVarState contains the current state of the dashboard variable
type templatingVarState struct {
	Tags  []string `json:"tags,omitempty"`
	Text  string   `json:"text,omitempty"`
	Value string   `json:"value,omitempty"`
}

//TemplatingOption contains all Templating option attributes
type templatingOption struct {
	Selected bool   `json:"selected,omitempty"`
	Text     string `json:"text,omitempty"`
	Value    string `json:"value,omitempty"`
}

//TemplatingVar defines a single Dashboard variable
type templatingVar struct {
	AllFormat      string             `json:"allFormat,omitempty"` //: "wildcard",
	Current        templatingVarState `json:"current,omitempty"`
	Datasource     string             `json:"datasource,omitempty"`
	Definition     string             `json:"definition,omitempty"`
	IncludeAll     bool               `json:"includeAll,omitempty"`
	Name           string             `json:"name,omitempty"`
	Options        []templatingOption `json:"options,omitempty"`
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
type templating struct {
	Enable bool            `json:"enable,omitempty"`
	List   []templatingVar `json:"list,omitempty"`
}

func NewTemplatingVar() *templatingVar {
	return &templatingVar{}
}
