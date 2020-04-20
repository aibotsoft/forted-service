create or alter proc dbo.uspCreatePrice @Price decimal(9, 6), @MarketId int, @ReceivedAt datetimeoffset as
begin
    set nocount on
    declare @LastPriceId int
    declare @LastPrice decimal(9, 6)

    select top 1 @LastPriceId = PriceId, @LastPrice = Price
    from dbo.Price
    where MarketId = @MarketId
    order by ReceivedAt desc

    if @@ROWCOUNT = 0 or @LastPrice != @Price
        begin
            insert into Price (Price, MarketId, ReceivedAt)
            output inserted.PriceId, inserted.CreatedAt
            values (@Price, @MarketId, @ReceivedAt)
        end
    else
        select @LastPriceId, null
end
