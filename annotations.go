package main

func NewAnnotation(name string, datasource string, enable bool, hide bool, atype string) *Annotation {
	return &Annotation{
		Datasource: datasource,
		Enable:     enable,
		Hide:       hide,
		IconColor:  "rgba(0, 211, 255, 1)",
		Name:       name,
		Type:       atype,
	}
}
