-- example HTTP POST script which demonstrates setting the
-- HTTP method, body, and adding a header

wrk.method = "POST"
wrk.body   = "{\"make\":\"Ford\", \"model\":\"Focus\", \"colour\":\"Green\", \"owner\":\"Tom\"}"
wrk.headers["Content-Type"] = "application/json"