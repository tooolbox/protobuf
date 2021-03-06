DROP TABLE IF EXISTS starfleet;

CREATE TABLE starfleet (
	id int NOT NULL AUTO_INCREMENT,
	name TEXT NOT NULL,
	no_of_passengers INT(6),
	mission_statement TEXT,
	we_are_leaving_at TIMESTAMP,
	PRIMARY KEY (id)
);

INSERT INTO starfleet (name, no_of_passengers, mission_statement, we_are_leaving_at)
VALUES ("USS Enterprise", 654, NULL, now());

INSERT INTO starfleet (name, no_of_passengers, mission_statement, we_are_leaving_at)
VALUES ("USS Discovery", NULL, "spore drive development", NULL);
