# Meet the Virtual orb

This orb is a part of the testing infrastructure to validate the main system
* When it runs in a CI pipeline simulate sign-up and submit images to the API with an associated id.
* When it is deployed as a workload it periodically reports status by calling an API and submitting battery, cpu usage, cpu temp, disk space.

# How to use

To build the orb: `make build`

To run orb in workflow: `./bin/virtual_orb`  

To run orb in ci: `./bin/virtual_orb -ci`  
