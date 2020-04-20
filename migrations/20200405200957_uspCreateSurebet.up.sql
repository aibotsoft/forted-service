create or alter proc dbo.uspCreateSurebet @FortedEventId int, @AMarketId int, @BMarketId int as
begin
    set nocount on
    declare @SurebetId int

    select @SurebetId = SurebetId
    from dbo.Surebet
    where FortedEventId = @FortedEventId
      and AMarketId = @AMarketId
      and BMarketId = @BMarketId

    if @@rowcount = 0
        insert into dbo.Surebet (FortedEventId, AMarketId, BMarketId)
        output inserted.SurebetId
        values (@FortedEventId, @AMarketId, @BMarketId)
    else
        select @SurebetId
end

