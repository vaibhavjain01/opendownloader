create database tradescribe;

create table links (
    id     int(5) unsigned NOT NULL AUTO_INCREMENT,
    link    varchar(2048),
    downloaded  int(8),
    
    PRIMARY KEY ('id'),
    UNIQUE KEY 'link' ('link')
) DEFAULT CHARSET=utf8;