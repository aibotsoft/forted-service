create or alter proc uspCreateMarket @eventId int, @Market varchar(50) as
begin
    set nocount on
    declare @MarketId int

    select @MarketId = MarketId from Market where TeamName = @TeamName and SportId = @SportId
    if @@rowcount = 0
        insert into Team (TeamName, SportId) output inserted.TeamId values (@TeamName, @SportId)
    else
        select @TeamId
end