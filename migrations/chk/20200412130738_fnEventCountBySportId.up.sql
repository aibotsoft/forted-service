create or alter function chk.fnEventCountBySportId(@SportId tinyint) returns int as
begin
    declare @cnt int;
    select @cnt = count(E.EventId)
    from Sport S
             join League L on S.SportId = L.SportId
             join Event E on L.LeagueId = E.LeagueId
    where S.SportId = @SportId

    return @cnt
end