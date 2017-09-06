
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE WORDS (
    ID                INTEGER       PRIMARY KEY AUTOINCREMENT
,   LOGICAL_NAME      VARCHAR(64)   NOT NULL
,   PHYSICAL_NAME     VARCHAR(64)   NOT NULL
,   PHYSICAL_NAME_ABB VARCHAR(64)   NOT NULL
,   STATUS            CHAR(1)       NOT NULL
,   CREATED_AT        TIMESTAMP
,   UPDATED_AT        TIMESTAMP
,   DELETED_AT        TIMESTAMP
,   CREATOR_ID        INTEGER
,   UPDATER_ID        INTEGER
,   DELETER_ID        INTEGER
,   LOCK_VERSION      INTEGER       NOT NULL  DEFAULT 0
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE WORDS;
