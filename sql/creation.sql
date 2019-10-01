DROP DATABASE opendownloader;
create database opendownloader;

USE opendownloader;

create table links (
    id     int(10) unsigned NOT NULL AUTO_INCREMENT,
    link    varchar(512) UNIQUE,
    downloadCount  int(8) unsigned,
    lastDownloaded  DATETIME,

    PRIMARY KEY (id)
);

create table downloadedAudio (
    id     int(10) unsigned NOT NULL AUTO_INCREMENT,
    originalLinkId    int(10) unsigned NOT NULL,
    downloadedLink  varchar(512) NOT NULL UNIQUE,
    downloadedFormatCode    int(5) unsigned NOT NULL,
    downloadedFile  MEDIUMBLOB NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (originalLinkId) REFERENCES links(id)
);

create table downloadedVideo (
    id     int(10) unsigned NOT NULL AUTO_INCREMENT,
    originalLinkId    int(10) unsigned NOT NULL,
    downloadedLink  varchar(512) NOT NULL UNIQUE,
    downloadedFormatCode    int(5) unsigned NOT NULL,
    downloadedFile  LONGBLOB NOT NULL,
    downloadedTime  DATETIME NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (originalLinkId) REFERENCES links(id)
);