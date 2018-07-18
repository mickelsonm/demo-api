create database if not exists demodata;

use demodata;

drop table if exists Tickets;

create table Tickets (
  ID INT PRIMARY KEY AUTO_INCREMENT,
  ShortDesc VARCHAR(50),
  LongDesc VARCHAR(100),
  Created date,
  LastUpdated date
);

insert into Tickets(ShortDesc,LongDesc,Created,LastUpdated) values ('test ticket', 'this is a beginning test ticket', UTC_TIMESTAMP(), null);
