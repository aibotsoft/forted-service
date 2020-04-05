create table dbo.Surebet
(
    SurebetId     int identity                               not null,
    FortedEventId int                                        not null,
    AMarketId     int                                        not null,
    BMarketId     int                                        not null,
    CreatedAt     datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_SurebetId primary key nonclustered (SurebetId),
    constraint UQ_Surebet unique (FortedEventId, AMarketId, BMarketId),
    constraint FK_Surebet_AMarket foreign key (AMarketId) references Market,
    constraint FK_Surebet_BMarket foreign key (BMarketId) references Market on update cascade on delete cascade,
)