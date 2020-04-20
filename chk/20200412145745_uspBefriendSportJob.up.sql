create or alter proc chk.uspBefriendSportJob @MinCount int as
begin
    DECLARE @ASportId int, @BSportId int

    DECLARE Friend_Cursor CURSOR FOR
        select ASportId, BSportId
        from chk.fnGroupedSportFriends()
        where TotalCount > @MinCount
          and not exists(select 1
                         from chk.SportFriend a
                                  join chk.SportFriend b
                                       on a.SportBaseId = b.SportBaseId
                         where a.SportId = ASportId
                           and b.SportId = BSportId)
        order by TotalCount desc

    OPEN Friend_Cursor;
    FETCH NEXT FROM Friend_Cursor INTO @ASportId, @BSportId;
    WHILE @@FETCH_STATUS = 0
        BEGIN
            exec chk.uspBefriendSports @ASportId, @BSportId
            FETCH NEXT FROM Friend_Cursor INTO @ASportId, @BSportId;
        END;

    CLOSE Friend_Cursor;
    DEALLOCATE Friend_Cursor;
end