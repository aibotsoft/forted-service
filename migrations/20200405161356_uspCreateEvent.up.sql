create or alter proc uspCreateEvent @StartsAt datetimeoffset, @HomeId int, @AwayId int , @LeagueId int as
begin
    set nocount on
    declare @EventId int

    select @EventId = EventId
    from dbo.Event
    where StartsAt = @StartsAt
      and HomeId = @HomeId
      and AwayId = @AwayId
      and LeagueId = @LeagueId
    if @@rowcount = 0
        insert into Event (StartsAt, HomeId, AwayId, LeagueId)
        output inserted.EventId
        values (@StartsAt, @HomeId, @AwayId, @LeagueId)
    else
        select @EventId
end