# lazydash [![CircleCI](https://circleci.com/gh/dbnegative/lazydash/tree/master.svg?style=svg&circle-token=166cef586b42bb07d2e81ffaffaac8bd371970d2)](https://circleci.com/gh/dbnegative/lazydash/tree/master)

Auto generate Grafana dashboards based on prometheus metrics endpoints

# Notes

* Only Supports Counter and Gauge Types at the moment
* Assumes metrics adhere to prometheus metrics naming conventions and standards

# Features
* Auto labelling
* Graphs use metrics HELP as description

# Usage

```
usage: lazydash [<flags>]

generate grafana dashboard json from prometheus metrics data via file or by | pipe

Flags:
  -h, --help          Show context-sensitive help (also try --help-long and --help-man).
  -f, --file=FILE     Parse metrics from file.
  -t, --title="Demo"  Dashboard title
      --stdin         Read from stdin
  -p, --pretty        Print pretty indented JSON
  -g, --gauges        Render gauge values as gauge panel types instead of graph
      --set-counter-expr="sum(rate(:METRIC: [1m]))"  
                      Set custom meterics query expression for counter type metric
      --set-gauge-expr=":METRIC:"  
                      Set custom meterics query expression for gauge type metric
      --set-delimiter=":METRIC:"  
                      Set custom meterics delimiter used to insert metric name into expression, only used if a custom expression is set
      --set-counter-legend="Job:[{{job}}]"  
                      Set the default counter panel legend format
      --set-gauge-legend="Job:[{{job}}]"  
                      Set the default counter panel legend format
      --version       Show application version.
```
# Examples

```
curl -s http://localhost:9323/metrics | ./lazydash -t "Demo" -p
```
```
./lazydash < promdata.txt
```
```
cat promdata | ./lazydash -t "Demo" -p
```
```
echo "# HELP builder_builds_triggered_total Number of triggered image builds \n# TYPE builder_builds_triggered_total counter\nbuilder_builds_triggered_total 0" |./lazydash -t "simple dashboard" -p
```
