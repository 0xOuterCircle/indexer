create table environment(
    id serial unique,
    network char(10),
    chain_id int
);
insert into environment (network, chain_id) values ('ethereum', 5);
select * from environment;

create table dao(
    id serial unique,
    environment int references environment(id),
    creator text,
    parent int references dao(id),
    registry text unique,
    governance text unique
);
drop table dao;
insert into dao (environment, creator, parent, registry, governance)
values (1, '0x0000000000000000000000000000000000000000', null, '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000');