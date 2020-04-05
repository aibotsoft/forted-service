create table dbo.League
(
    LeagueId   int identity,
    LeagueName varchar(50)                                not null,
    SportId    int                                        not null,
    CreatedAt  datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_LeagueId primary key nonclustered (LeagueId),
    constraint UQ_League unique (LeagueName, SportId),
    constraint FK_League_Sport foreign key (SportId) references Sport on update cascade on delete cascade
)