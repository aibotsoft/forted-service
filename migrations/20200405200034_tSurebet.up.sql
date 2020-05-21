create table dbo.Surebet
(
    SurebetId     int identity                               not null,
    FortedEventId int                                        not null,
    AMarketId     int                                        not null,
    BMarketId     int                                        not null,
    CMarketId     int                                        not null default 0,

    CreatedAt     datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_SurebetId primary key (SurebetId),
    constraint UQ_Surebet unique (FortedEventId, AMarketId, BMarketId, CMarketId),
    constraint FK_Surebet_AMarket foreign key (AMarketId) references dbo.Market,
    constraint FK_Surebet_BMarket foreign key (BMarketId) references dbo.Market on update cascade on delete cascade,
)

