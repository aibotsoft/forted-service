create or alter proc uspCreateMarket @MarketName varchar(60), @EventId int, @Url varchar(240), @Num tinyint as
begin
    set nocount on
    declare @MarketId int

    select @MarketId = MarketId from dbo.Market where MarketName = @MarketName and EventId = @EventId
    if @@rowcount = 0
        insert into Market (MarketName, EventId, Url, Num)
        output inserted.MarketId
        values (@MarketName, @EventId, @Url, @Num)
    else
        select @MarketId
end