create table dbo.Event
(
    EventId   int identity                               not null,
    StartsAt  datetimeoffset default sysdatetimeoffset() not null,
    HomeId    int                                        not null,
    AwayId    int                                        not null,
    LeagueId  int                                        not null,
    CreatedAt datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_EventId primary key nonclustered (EventId),
    constraint UQ_Event unique (StartsAt, HomeId, AwayId, LeagueId),
    constraint FK_Home_Team foreign key (HomeId) references Team on delete NO ACTION,
    constraint FK_Away_Team foreign key (AwayId) references Team on delete NO ACTION,
    constraint FK_Event_League foreign key (LeagueId) references League on delete cascade ,

)