create or alter proc uspCreateTeam @TeamName varchar(120), @SportId int  as
begin
    set nocount on
    declare @TeamId int

    select @TeamId = TeamId from Team where TeamName = @TeamName and SportId = @SportId
    if @@rowcount = 0
        insert into Team (TeamName, SportId) output inserted.TeamId values (@TeamName, @SportId)
    else
        select @TeamId
end