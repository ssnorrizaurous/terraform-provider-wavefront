---
layout: "wavefront"
page_title: "Wavefront: Dashboard JSON"
description: |-
  Provides a Wavefront Dashboard JSON resource.  This allows dashboards to be created, updated, and deleted.
---

# Resource : wavefront_dashboard_json

Provides a Wavefront Dashboard JSON resource.  This allows dashboards to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_dashboard_json" "test_dashboard_json" {
	dashboard_json = <<EOF
{
  "name": "Terraform Test Dashboard Json",
  "description": "a",
  "eventFilterType": "BYCHART",
  "eventQuery": "",
  "defaultTimeWindow": "",
  "url": "tftestimport",
  "displayDescription": false,
  "displaySectionTableOfContents": true,
  "displayQueryParameters": false,
  "sections": [
    {
      "name": "section 1",
      "rows": [
        {
          "charts": [
            {
              "name": "chart 1",
              "sources": [
                {
                  "name": "source 1",
                  "query": "ts()",
                  "scatterPlotSource": "Y",
                  "querybuilderEnabled": false,
                  "sourceDescription": ""
                }
              ],
              "units": "someunit",
              "base": 0,
              "noDefaultEvents": false,
              "interpolatePoints": false,
              "includeObsoleteMetrics": false,
              "description": "This is chart 1, showing something",
              "chartSettings": {
                "type": "markdown-widget",
                "max": 100,
                "expectedDataSpacing": 120,
                "windowing": "full",
                "windowSize": 10,
                "autoColumnTags": false,
                "columnTags": "deprecated",
                "tagMode": "all",
                "numTags": 2,
                "customTags": [
                  "tag1",
                  "tag2"
                ],
                "groupBySource": true,
                "y1Max": 100,
                "y1Units": "units",
                "y0ScaleSIBy1024": true,
                "y1ScaleSIBy1024": true,
                "y0UnitAutoscaling": true,
                "y1UnitAutoscaling": true,
                "fixedLegendEnabled": true,
                "fixedLegendUseRawStats": true,
                "fixedLegendPosition": "RIGHT",
                "fixedLegendDisplayStats": [
                  "stat1",
                  "stat2"
                ],
                "fixedLegendFilterSort": "TOP",
                "fixedLegendFilterLimit": 1,
                "fixedLegendFilterField": "CURRENT",
                "plainMarkdownContent": "markdown content"
              },
              "summarization": "MEAN"
            }
          ],
          "heightFactor": 50
        }
      ]
    }
  ],
  "parameterDetails": {
    "param": {
      "hideFromView": false,
      "description": null,
      "allowAll": null,
      "tagKey": null,
      "queryValue": null,
      "dynamicFieldType": null,
      "reverseDynSort": null,
      "parameterType": "SIMPLE",
      "label": "test",
      "defaultValue": "Label",
      "valuesToReadableStrings": {
        "Label": "test"
      },
      "selectedLabel": "Label",
      "value": "test"
    }
  },
  "tags" :{
    "customerTags":  ["terraform"]
  }
}
EOF
}
```

## Argument Reference

The following arguments are supported:

* `dashboard_json` - (Required) See the [Wavefront API Documentation](https://docs.wavefront.com/wavefront_api.html#api-documentation-wavefront-instance) 
for instructions on how to get to your API documentation for more details. 

## Import

Dashboard JSON can be imported by using the `id`, e.g.:

```
$ terraform import wavefront_dashboard_json.dashboard_json tftestimport
```