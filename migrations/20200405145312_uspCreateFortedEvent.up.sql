create or alter proc uspCreateFortedEvent @StartsAt datetimeoffset, @FortedHomeId int, @FortedAwayId int as
begin
    set nocount on
    declare @FortedEventId int

    select @FortedEventId = FortedEventId
    from FortedEvent
    where StartsAt = @StartsAt
      and FortedHomeId = @FortedHomeId
      and FortedAwayId = @FortedAwayId
    if @@rowcount = 0
        insert into FortedEvent (StartsAt, FortedHomeId, FortedAwayId)
        output inserted.FortedEventId
        values (@StartsAt, @FortedHomeId, @FortedAwayId)
    else
        select @FortedEventId
end