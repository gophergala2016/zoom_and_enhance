## HTTP2fer
HTTP/2 proxy for HTTP1.x servers.

### Overview
Bring your app into the next generation (without having to change a thing). HTTP2fer allows you to set up a lightweight transforming proxy between your backend and the outside world, so HTTP2-enabled clients can send HTTP2 requests, while your HTTP1 backend can continue to work as it always has.

### Features
- [x] Proxy
- [x] SSL (required by most browsers)
- [ ] Load Balancing (multiple origin servers)
- [ ] Pipelining
- [x] Stats
- [ ] Throttling
- [ ] Health Checking
- [x] Web UI

# Frontend
A lightweight UI will provide users access to statistics and metrics about each endpoint (traffic, health, etc), as well as the ability to create, modify, and delete proxied endpoints.

### The Dashboard
The dashboard will provide a list of monitored and proxied endpoints, including their current status, whether or not they are secured, and the total number of requests made to date.

The dashboard will also provide access to statistics and metrics about the proxy as a whole, including the total number of requests handled, their sources and destinations, and an estimate of reduced latency; this will all be in the form of graphs displaying live data sent over a websocket.

### Viewing Endpoints
Each endpoint will have its own detail page, showing the name, a health check, the proxy settings, and statistics about the endpoints performance over the last (by default) week.

### Creating Endpoints
Endpoints will be easily created through the UI and will configure the proxy to route requests to and from various clients and endpoints over various HTTP methods.

Endpoint settings include:
* Source and destination
* HTTP1 or HTTP2
* Whether or not to use server push
* Transparent passthrough for HTTP2-enabled clients to access HTTP2 services

### Deleting Endpoints
Endpoints can be deleted by clicking the delete icon on the main endpoint list, as well as the `delete endpoint` button on the endpoint view page.
