USE auth;

CREATE TABLE secret
(
    id          BINARY(16)                                                            NOT NULL PRIMARY KEY,
    active      BOOLEAN     DEFAULT TRUE NOT NULL,
    created     DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6)                              NOT NULL,
    valid_from  DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6)                              NOT NULL,
    valid_to    DATETIME(6) DEFAULT (DATE_ADD(CURRENT_TIMESTAMP(6), INTERVAL 1 YEAR)) NOT NULL,
    private_key TEXT                                                                  NOT NULL
);

