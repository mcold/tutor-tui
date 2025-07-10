CREATE TABLE IF NOT EXISTS ITEM 
(
	id      integer primary key autoincrement,
    name    varchar not null,
    comment varchar
);

DROP TABLE IF EXISTS CONTENT;
CREATE TABLE IF NOT EXISTS CONTENT
(
	id      	integer primary key autoincrement,
	id_item		integer,
	id_parent	integer,
    num     	integer,
    name    	text,
	code		text,
	output		text,
	comment 	text,
	direct		text not null default 'row',
	out_type 	text not null default 'table',
    code_proportion integer not null default 1,
    out_proportion integer not null default 2,
    FOREIGN KEY (id_item) REFERENCES item (id) ON DELETE CASCADE,
    FOREIGN KEY (id_parent) REFERENCES content (id) ON DELETE CASCADE
);


INSERT INTO item (name)
  VALUES ('TRAIT');


INSERT INTO content(id_item, num, name, code)
  values(1, 1, 'first', 'some code')
  		, (1, 2, 'second', 'another one')
  		, (1, 3, 'third', 'final');

INSERT INTO content(id_item, num, name, code)
  values(1, 4, 'fourth', 'final');


INSERT INTO content(id_item, num, name, code)
  values(1, 5, 'fifth', 'final');