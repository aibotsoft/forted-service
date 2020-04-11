create table dbo.Service
(
    ServiceId   tinyint identity                           not null,
    ServiceName varchar(50)                                not null,
    CreatedAt   datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_ServiceId primary key (ServiceId),
    constraint UQ_Service unique (ServiceName) ,

)
