create table dbo.Surebet
(
    SurebetId     bigint not null default DATEDIFF_BIG(millisecond, '1970-01-01 00:00:00', sysdatetimeoffset()),
    FortedEventId int    not null,
    AMarketId     int    not null,
    BMarketId     int    not null,
    CMarketId     int    not null default 0,

    CreatedAt     datetimeoffset(2)  default sysdatetimeoffset() not null,
    constraint PK_Surebet primary key (SurebetId),
    constraint UQ_Surebet unique (FortedEventId, AMarketId, BMarketId, CMarketId),
    constraint FK_Surebet_AMarket foreign key (AMarketId) references dbo.Market,
    constraint FK_Surebet_BMarket foreign key (BMarketId) references dbo.Market on update cascade on delete cascade,
)

