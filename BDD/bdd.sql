CREATE TABLE category ( 
	ID                   integer NOT NULL  PRIMARY KEY autoincrement ,
	name                 text NOT NULL    
 );

CREATE TABLE user ( 
	pseudo               text NOT NULL  PRIMARY KEY  ,
	mail                 text NOT NULL    ,
	password             text NOT NULL    
 );

CREATE TABLE topic ( 
	ID                   integer NOT NULL  PRIMARY KEY autoincrement ,
	title                text NOT NULL    ,
	content              text     ,
	user_pseudo          text NOT NULL    ,
	category_id          integer NOT NULL    ,
	FOREIGN KEY ( category_id ) REFERENCES category( ID ) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY ( user_pseudo ) REFERENCES user( pseudo ) ON DELETE CASCADE ON UPDATE CASCADE
 );

CREATE TABLE post ( 
	ID                   integer NOT NULL  PRIMARY KEY autoincrement ,
	user_pseudo          text NOT NULL    ,
	content              text NOT NULL    ,
	date                 datetime NOT NULL DEFAULT CURRENT_TIMESTAMP   ,
	topic_id             integer NOT NULL    ,
	FOREIGN KEY ( topic_id ) REFERENCES topic( ID ) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY ( user_pseudo ) REFERENCES user( pseudo ) ON DELETE CASCADE ON UPDATE CASCADE
 );

CREATE TABLE comment ( 
	user_pseudo          text NOT NULL    ,
	content              text NOT NULL    ,
	post_id              integer NOT NULL    ,
	ID                   integer NOT NULL    ,
	CONSTRAINT sqlite_autoindex_comment_1 UNIQUE ( ID ) ,
	FOREIGN KEY ( post_id ) REFERENCES post( ID ) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY ( user_pseudo ) REFERENCES user( pseudo ) ON DELETE CASCADE ON UPDATE CASCADE
 );

CREATE TABLE like ( 
	user_pseudo          text NOT NULL    ,
	liked                boolean NOT NULL DEFAULT false   ,
	post_id              integer NOT NULL    ,
	ID                   integer NOT NULL    ,
	CONSTRAINT sqlite_autoindex_like_1 UNIQUE ( ID ) ,
	FOREIGN KEY ( post_id ) REFERENCES post( ID ) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY ( user_pseudo ) REFERENCES user( pseudo ) ON DELETE CASCADE ON UPDATE CASCADE
 );

INSERT INTO user( pseudo, mail, password ) VALUES ( 'liv44', 'liv44@gmail.com', 'test1234');
INSERT INTO user( pseudo, mail, password ) VALUES ( 'ShildowTV', 'shildowTV@gmail.com', 'azerty');
INSERT INTO user( pseudo, mail, password ) VALUES ( 'OlivMo', 'molive@gmail.com', 'qwerty');
INSERT INTO user( pseudo, mail, password ) VALUES ( 'RenJag', 'renjag@gmail.com', 'qwerty');
INSERT INTO user( pseudo, mail, password ) VALUES ( 'Mat', 'mat@gmail.com', 'qwerty');