create or alter proc uspCreateLeague @LeagueName varchar(120), @SportId int  as
begin
    set nocount on
    declare @LeagueId int

    select @LeagueId = LeagueId from League where LeagueName = @LeagueName and SportId = @SportId
    if @@rowcount = 0
        insert into League (LeagueName, SportId) output inserted.LeagueId values (@LeagueName, @SportId)
    else
        select @LeagueId
end