create or alter proc chk.uspBefriendLeagueJob @MinCount int as
begin
    DECLARE @ALeagueId int, @BLeagueId int

    DECLARE Friend_Cursor CURSOR FOR
        select ALeagueId, BLeagueId
        from chk.fnGroupedLeagueFriends()
        where not exists(select 1
                         from chk.LeagueFriend a
                                  join chk.LeagueFriend b
                                       on a.LeagueBaseId = b.LeagueBaseId
                         where a.LeagueId = ALeagueId
                           and b.LeagueId = BLeagueId)
          and TotalCount > @MinCount
        order by TotalCount desc

    OPEN Friend_Cursor;
    FETCH NEXT FROM Friend_Cursor INTO @ALeagueId, @BLeagueId;
    WHILE @@FETCH_STATUS = 0
        BEGIN
            exec chk.uspBefriendLeagues @ALeagueId, @BLeagueId
            FETCH NEXT FROM Friend_Cursor INTO @ALeagueId, @BLeagueId;
        END;

    CLOSE Friend_Cursor;
    DEALLOCATE Friend_Cursor;
end