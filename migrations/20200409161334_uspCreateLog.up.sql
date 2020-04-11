create or alter proc uspCreateLog @SurebetId int,
                                  @FilterName varchar(50),
                                  @FortedProfit decimal(9, 6),
                                  @InitiatorId int,
                                  @SkynetId int,
                                  @ReceivedAt datetimeoffset as
begin
    set nocount on
    insert into dbo.Log (SurebetId, FilterName, FortedProfit, InitiatorId, SkynetId, ReceivedAt)
    output inserted.LogId
    values (@SurebetId, @FilterName, @FortedProfit, @InitiatorId,@SkynetId, @ReceivedAt)
end