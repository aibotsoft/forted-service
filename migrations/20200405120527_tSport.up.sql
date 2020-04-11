create table dbo.Sport
(
    SportId   int identity                               not null,
    SportName varchar(50)                                not null,
    ServiceId tinyint                                    not null,
    CreatedAt datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_SportId primary key (SportId),
    constraint UQ_Sport unique (SportName, ServiceId),
    constraint FK_Sport_Service foreign key (ServiceId) references Service on update cascade on delete cascade,
)

