create or alter proc uspCreateSport @SportName varchar(50), @ServiceId tinyint  as
begin
    set nocount on
    declare @SportId int

    select @SportId = SportId from Sport where SportName = @SportName and ServiceId = @ServiceId
    if @@rowcount = 0
        insert into Sport (SportName, ServiceId) output inserted.SportId values (@SportName, @ServiceId)
    else
        select @SportId
end