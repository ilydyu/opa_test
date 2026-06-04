package app.auth

default allow = false

allow if {
    "reader" in input.roles
    input.method == "GET"
}

allow if {
    "admin" in input.roles
}