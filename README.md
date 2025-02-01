### GTFS Tracker

This tool will provide:
1. module for ingesting data from GTFS transit sources in-memory
2. general API for providing API access for transit requests, which should be portable from different GTFS-based systems.
3. functionality to serve this data in a template-based web format
4. a webserver for serving the above template-generated timetables
5. a mapping tool to visualize where transit vehicles are at any given point in the transit system

#### Architecture diagram
Please see GTFSTool.drawio for details

#### Data collection
To workaround API request limits, we should have a worker to re-request feed updates from realtime/static sources every N minutes. This will also enable data collection for statistical analysis.

### Implementation details

Use Go + realtime protobuf bindings

