create table dbo.League
(
    LeagueId   int identity                               not null,
    LeagueName varchar(120)                               not null,
    SportId    int                                        not null,
    CreatedAt  datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_LeagueId primary key (LeagueId),
    constraint UQ_League unique (LeagueName, SportId),
    index IX_LeagueSport (SportId),
    constraint FK_League_Sport foreign key (SportId) references Sport on update cascade on delete cascade
)
