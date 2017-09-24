goupnp is a UPnP client library for Go

Installation

Run go get -u github.com/huin/goupnp.

Documentation

Supported DCPs (you probably want to start with one of these):

GoDoc av1 - Client for UPnP Device Control Protocol MediaServer v1 and MediaRenderer v1.
GoDoc internetgateway1 - Client for UPnP Device Control Protocol Internet Gateway Device v1.
GoDoc internetgateway2 - Client for UPnP Device Control Protocol Internet Gateway Device v2.
Core components:

GoDoc (goupnp) core library - contains datastructures and utilities typically used by the implemented DCPs.
GoDoc httpu HTTPU implementation, underlies SSDP.
GoDoc ssdp SSDP client implementation (simple service discovery protocol) - used to discover UPnP services on a network.
GoDoc soap SOAP client implementation (simple object access protocol) - used to communicate with discovered services.
Regenerating dcps generated source code:

Install gotasks: go get -u github.com/jingweno/gotask
Change to the gotasks directory: cd gotasks
Run specgen task: gotask specgen
Supporting additional UPnP devices and services:

Supporting additional services is, in the trivial case, simply a matter of adding the service to the dcpMetadata whitelist in gotasks/specgen_task.go, regenerating the source code (see above), and committing that source code.

However, it would be helpful if anyone needing such a service could test the service against the service they have, and then reporting any trouble encountered as an issue on this project. If it just works, then please report at least minimal working functionality as an issue, and optionally contribute the metadata upstream.




[root@localhost upnp]# curl ifconfig.me
111.199.188.230
