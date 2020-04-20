create or alter function chk.fnEventCountByLeagueId(@LeagueId int) returns int as
begin
    declare @cnt int;
    select @cnt = count(E.EventId)
    from Sport S
             join League L on S.SportId = L.SportId
             join Event E on L.LeagueId = E.LeagueId
    where L.LeagueId = @LeagueId

    return @cnt
end