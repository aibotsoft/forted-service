create table dbo.Log
(
    LogId        int identity                               not null,
    SurebetId    int                                        not null,
    FilterName   varchar(50)                                not null,
    FortedProfit decimal(9, 6)                              not null,
    InitiatorId  int,
    SkynetId     int,
    ReceivedAt   datetimeoffset                             not null,
    CreatedAt    datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_LogId primary key (LogId),
)
create index Log_SkynetId_index on Log (SkynetId desc)
