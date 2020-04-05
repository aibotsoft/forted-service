create table dbo.Team
(
    TeamId    int identity,
    TeamName  varchar(50)                                not null,
    SportId   int                                        not null,
    CreatedAt datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_TeamId primary key nonclustered (TeamId),
    constraint UQ_Team unique (TeamName, SportId),
    constraint FK_Team_Sport foreign key (SportId) references Sport on update cascade on delete cascade
)