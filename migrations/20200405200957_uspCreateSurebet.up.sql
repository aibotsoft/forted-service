create or alter proc dbo.uspCreateSurebet @FortedEventId int,
                                          @AMarketId int,
                                          @BMarketId int,
                                          @CMarketId int
as
begin
    set nocount on
    declare @SurebetId int

    select @SurebetId = SurebetId
    from dbo.Surebet
    where FortedEventId = @FortedEventId
      and AMarketId = @AMarketId
      and BMarketId = @BMarketId
      and CMarketId = @CMarketId

    if @@rowcount = 0
        insert into dbo.Surebet (FortedEventId, AMarketId, BMarketId, CMarketId)
        output inserted.SurebetId
        values (@FortedEventId, @AMarketId, @BMarketId, @CMarketId)
    else
        select @SurebetId
end

