create or alter proc dbo.uspCheckSkynetId @SkynetId int as
begin
    set nocount on
    select LogId from Log where SkynetId = @SkynetId
end