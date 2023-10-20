package constants

const Schema string = `
PRAGMA FOREIGN_KEYS = ON;

CREATE TABLE http_mock (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    
    -- UUID
    mock_group TEXT UNIQUE NOT NULL,
    path TEXT DEFAULT '',
    method TEXT CHECK ( method IN ('GET', 'HEAD', 'POST', 'PUT', 'DELETE', 'CONNECT', 'OPTIONS', 'TRACE', 'PATCH') ) NOT NULL,

    UNIQUE (mock_group, path, method)
);

CREATE TABLE http_request_response (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    
    http_mock_id INTEGER,
    
	-- gob-encoded golang map
	request_header BLOB,
    request_body TEXT,
    request_body_mime_type TEXT,

	-- gob-encoded golang map
    additional_response_header BLOB,
    response_body TEXT,
    response_body_mime_type TEXT,
    response_code INT CHECK ( response_code >= 100 AND response_code < 600 ) DEFAULT 200,
    
    FOREIGN KEY (http_mock_id) REFERENCES http_mock(id),

    CHECK ( 
        ((request_body IS NULL) OR (request_body IS NOT NULL AND request_body_mime_type IS NOT NULL))
            AND
        ((response_body IS NULL) OR (response_body IS NOT NULL AND response_body_mime_type IS NOT NULL))
    )
);
CREATE INDEX http_request_response__mock_request ON http_request_response (http_mock_id, request_header, request_body, request_body_mime_type);
`
