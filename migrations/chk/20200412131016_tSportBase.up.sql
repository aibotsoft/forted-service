create table chk.SportBase
(
    SportBaseId   int identity                               not null,
    SportBaseName varchar(50)                                not null,
    CreatedAt     datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_SportBase primary key (SportBaseId)
)