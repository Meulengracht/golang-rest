## Golang REST-API Server example

This repository contains an example of a very simple rest-api server that supports two endpoints:
/version - Returns a hardcoded string version in json
/duration - Returns the boot times for kernel and userspace (retrieved with systemd-analyze)

### Version (Get)

Implemented as an hardcoded string, this could be set in a proper fashion to include git SHA, actual build numbers
and also build dates during linking by providing linker-values for the flags. This is a nice extension for a future
verison of this repo.

Example output
``` {"version":"1.0.0"} ```

### Duration (Get)

Utilizes a combination of systemd-analyze (sorry windows) and regex to retrieve the bootup times for both the kernel
itself and the userspace, and returns this as an json object. This is deeply coupled to systemd-analyze output format
and would require some modifications and additional code-handling for windows platform, none of which is in the current
scope of this example.

Example output
``` {"kernel":"12.126s","userspace":"23.315s"} ```

### Unit tests

Unit test should be added for the time_controller as it contains actual logic that relies on proper input,
and would therefor be ideal to add some unit tests. This has not been the scope for this example either, but
would be a nice further enhancement.
