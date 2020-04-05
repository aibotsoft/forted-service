create table dbo.Service
(
    ServiceId   tinyint identity,
    ServiceName varchar(50) unique                         not null,
    CreatedAt   datetimeoffset default sysdatetimeoffset() not null,
    constraint PK_ServiceId primary key nonclustered (ServiceId),
)