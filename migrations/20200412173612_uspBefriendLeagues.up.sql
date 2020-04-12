create or alter proc chk.uspBefriendLeagues @ALeagueId int, @BLeagueId int as
begin
    declare @ALeagueBaseId int, @BLeagueBaseId int;
    declare @AEventCount int, @BEventCount int
    declare @BaseId int, @LeagueBaseId int

    select @ALeagueBaseId = LeagueBaseId from chk.LeagueFriend where LeagueId = @ALeagueId
    select @BLeagueBaseId = LeagueBaseId from chk.LeagueFriend where LeagueId = @BLeagueId

    if @ALeagueBaseId = @BLeagueBaseId
        return

    if @ALeagueBaseId is null and @BLeagueBaseId is null
        -- Создать запись в LeagueBase
        begin
            set @AEventCount = chk.fnEventCountByLeagueId(@ALeagueId)
            set @BEventCount = chk.fnEventCountByLeagueId(@BLeagueId)
            if @AEventCount >= @BEventCount
                set @BaseId = @ALeagueId
            else
                set @BaseId = @BLeagueId

            insert into chk.LeagueBase (LeagueBaseName)
            select LeagueName
            from League
            where LeagueId = @BaseId
            set @LeagueBaseId = SCOPE_IDENTITY()

            insert into chk.LeagueFriend (LeagueBaseId, LeagueId)
            values (@LeagueBaseId, @ALeagueId),
                   (@LeagueBaseId, @BLeagueId)
            return
        end

    if @ALeagueBaseId is null
        begin
            insert into chk.LeagueFriend (LeagueBaseId, LeagueId) values (@BLeagueBaseId, @ALeagueId)
            return
        end

    if @BLeagueBaseId is null
        begin
            insert into chk.LeagueFriend (LeagueBaseId, LeagueId) values (@ALeagueBaseId, @BLeagueId)
            return
        end

    if @ALeagueBaseId < @BLeagueBaseId
        begin
            update chk.LeagueFriend SET LeagueBaseId = @ALeagueBaseId WHERE LeagueBaseId = @BLeagueBaseId
            delete chk.LeagueBase where LeagueBaseId = @BLeagueBaseId
        end
    else
        begin
            update chk.LeagueFriend SET LeagueBaseId = @BLeagueBaseId WHERE LeagueBaseId = @ALeagueBaseId
            delete chk.LeagueBase where LeagueBaseId = @ALeagueBaseId
        end
end