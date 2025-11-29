DROP TABLE IF EXISTS PROJECT;
CREATE TABLE IF NOT EXISTS PROJECT
(
    id      integer primary key autoincrement,
    name    varchar not null
);


DROP TABLE IF EXISTS ITEM;
CREATE TABLE IF NOT EXISTS ITEM
(
	id          integer primary key autoincrement,
    id_project  integer not null,
    name        varchar not null,
    comment     varchar,
    FOREIGN KEY (id_project) REFERENCES project (id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS SLIDE;
CREATE TABLE IF NOT EXISTS SLIDE
(
	id      	        integer primary key autoincrement,
	id_item		        integer not null,
    num     	        integer,
    name    	        text,
	content		        text,
    content_type        text not null default 'code',
    direct		        text not null default 'column',
    content_proportion  integer not null default 1,
    page_proportion     integer not null default 2,
    comment 	        text,
    FOREIGN KEY (id_item) REFERENCES item (id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS TAB;
CREATE TABLE IF NOT EXISTS TAB
(
    id      	    integer primary key autoincrement,
    id_slide        integer not null,
    num             integer,
    name            text,
    content		    text,
    content_type    text not null default 'table',
    comment 	    text,
    FOREIGN KEY (id_slide) REFERENCES slide (id) ON DELETE CASCADE
);