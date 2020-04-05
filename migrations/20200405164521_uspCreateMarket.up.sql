create or alter proc uspCreateMarket @MarketName varchar(60), @EventId int as
begin
    set nocount on
    declare @MarketId int

    select @MarketId = MarketId from dbo.Market where MarketName = @MarketName and EventId = @EventId
    if @@rowcount = 0
        insert into Market (MarketName, EventId) output inserted.MarketId values (@MarketName, @EventId)
    else
        select @MarketId
end