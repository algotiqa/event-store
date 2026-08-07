-- ======================================================================
-- ===   Sql Script for Database : Event Store
-- ===
-- === Build : 317
-- ======================================================================

CREATE TABLE event
  (
    id          INT              auto_increment,
    username    VARCHAR(32),
    event_date  DATETIME         not null,
    level       TINYINT          not null,
    title       VARCHAR(64)      not null,
    message     VARCHAR(512),
    parameters  VARBINARY(500),

    primary key(id)
  )
 ENGINE = InnoDB ;

CREATE INDEX eventIDX1 ON event(username);
CREATE INDEX eventIDX2 ON event(event_date);

-- ======================================================================

