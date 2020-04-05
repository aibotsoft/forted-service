create or alter proc uspCreateService @ServiceName varchar(50) as
begin
    set nocount on
    declare @ServiceId tinyint

    select @ServiceId = ServiceId from Service where ServiceName = @ServiceName
    if @@rowcount = 0
        insert into Service (ServiceName) output inserted.ServiceId values (@ServiceName)
    else
        select @ServiceId
end