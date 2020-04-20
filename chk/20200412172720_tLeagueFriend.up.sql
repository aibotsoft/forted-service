create table chk.LeagueFriend
(
    LeagueBaseId int                                        not null,
    LeagueId     int                                        not null,
    CreatedAt   datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_LeagueFriend primary key (LeagueBaseId, LeagueId)
)