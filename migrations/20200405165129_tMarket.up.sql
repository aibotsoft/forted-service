create table dbo.Market
(
    MarketId   int identity                               not null,
    MarketName varchar(60)                                not null,
    EventId    int                                        not null,
    Url        varchar(240)                               not null,
    Num        tinyint                                    not null,
    CreatedAt  datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_MarketId primary key (MarketId),
    constraint UQ_Market unique (MarketName, EventId),
    constraint FK_Market_Event foreign key (EventId) references Event on update cascade on delete cascade
)
