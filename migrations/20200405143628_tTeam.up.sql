create table dbo.Team
(
    TeamId    int identity                               not null,
    TeamName  varchar(120)                                not null,
    SportId   int                                        not null,
    CreatedAt datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_TeamId primary key (TeamId),
    constraint UQ_Team unique (TeamName, SportId),
    constraint FK_Team_Sport foreign key (SportId) references Sport on update cascade on delete cascade
)

