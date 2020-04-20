create schema chk

create table chk.SportFriend
(
    SportBaseId int                                        not null,
    SportId     int                                        not null,
    CreatedAt   datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_SportFriend primary key (SportBaseId, SportId)
)