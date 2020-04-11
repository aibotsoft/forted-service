create table dbo.Price
(
    PriceId    int identity                               not null,
    Price      decimal(9, 6)                              not null,
    MarketId   int                                        not null,
    CreatedAt  datetimeoffset default sysdatetimeoffset() not null,
    ReceivedAt datetimeoffset                             not null,

    constraint PK_PriceId primary key (PriceId),
    constraint FK_Price_Market foreign key (MarketId) references Market on update cascade on delete cascade,
    index IX_Price_MarketId (MarketId)
)


