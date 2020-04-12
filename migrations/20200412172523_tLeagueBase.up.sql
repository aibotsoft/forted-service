create table chk.LeagueBase
(
    LeagueBaseId   int identity                               not null,
    LeagueBaseName varchar(120)                               not null,
    CreatedAt      datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_LeagueBase primary key (LeagueBaseId)
)