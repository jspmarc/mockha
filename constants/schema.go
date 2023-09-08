package constants

const Schema string = `
PRAGMA FOREIGN_KEYS = ON;

CREATE TABLE http_mock (
    id INT PRIMARY KEY AUTOINCREMENT,
    
    mock_group TEXT,
    path TEXT DEFAULT '',
    method TEXT CHECK ( method IN ('GET', 'HEAD', 'POST', 'PUT', 'DELETE', 'CONNECT', 'OPTIONS', 'TRACE', 'PATCH') ) NOT NULL,

	-- gob-encoded golang map
    request_header BLOB,
    request_body TEXT,
    request_body_content_type TEXT,

	-- gob-encoded golang map
    additional_response_header BLOB,
    response_body TEXT,
    response_body_content_type TEXT,
    response_code INT CHECK ( response_code >= 100 AND response_code < 600 ) DEFAULT 200,

    UNIQUE (mock_group, path, method),

    CHECK ( 
        ((request_body IS NULL) OR (request_body IS NOT NULL AND request_body_content_type IS NOT NULL))
            AND
        ((response_body IS NULL) OR (response_body IS NOT NULL AND response_body_content_type IS NOT NULL))
     )
);
`
