package app.auth

test_reader_get_allowed if {
    allow with input as {"method": "GET", "path": "/resource", "roles": ["reader"]}
}

test_reader_post_denied if {
    not allow with input as {"method": "POST", "path": "/resource", "roles": ["reader"]}
}

test_no_roles_denied if {
    not allow with input as {"method": "GET", "path": "/resource", "roles": []}
}

test_wrong_role_denied if {
    not allow with input as {"method": "GET", "path": "/resource", "roles": ["guest"]}
}