create table dbo.FortedEvent
(
    FortedEventId int identity                               not null,
    StartsAt      datetimeoffset default sysdatetimeoffset() not null,
    FortedHomeId  int                                        not null,
    FortedAwayId  int                                        not null,
    CreatedAt     datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_FortedEventId primary key nonclustered (FortedEventId),
    constraint UQ_FortedEvent unique (StartsAt, FortedHomeId, FortedAwayId),
    constraint FK_FortedHome_Team foreign key (FortedHomeId) references Team on delete NO ACTION,
    constraint FK_FortedAway_Team foreign key (FortedAwayId) references Team on delete NO ACTION
)