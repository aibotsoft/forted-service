create or alter proc chk.uspBefriendSports @ASportId int, @BSportId int as
begin
    declare @ASportBaseId int, @BSportBaseId int;
    declare @AEventCount int, @BEventCount int
    declare @BaseId int, @SportBaseId int

    select @ASportBaseId = SportBaseId from chk.SportFriend where SportId = @ASportId
    select @BSportBaseId = SportBaseId from chk.SportFriend where SportId = @BSportId

    if @ASportBaseId = @BSportBaseId
        return

    if @ASportBaseId is null and @BSportBaseId is null
        -- Создать запись в SportBase
        begin
            set @AEventCount = chk.fnEventCountBySportId(@ASportId)
            set @BEventCount = chk.fnEventCountBySportId(@BSportId)
            if @AEventCount >= @BEventCount
                set @BaseId = @ASportId
            else
                set @BaseId = @BSportId

            insert into chk.SportBase (SportBaseName)
            select SportName
            from Sport
            where SportId = @BaseId
            set @SportBaseId = SCOPE_IDENTITY()

            insert into chk.SportFriend (SportBaseId, SportId)
            values (@SportBaseId, @ASportId),
                   (@SportBaseId, @BSportId)
            return
        end

    if @ASportBaseId is null
        begin
            insert into chk.SportFriend (SportBaseId, SportId) values (@BSportBaseId, @ASportId)
            return
        end

    if @BSportBaseId is null
        begin
            insert into chk.SportFriend (SportBaseId, SportId) values (@ASportBaseId, @BSportId)
            return
        end

    if @ASportBaseId < @BSportBaseId
        begin
            update chk.SportFriend SET SportBaseId = @ASportBaseId WHERE SportBaseId = @BSportBaseId
            delete chk.SportBase where SportBaseId = @BSportBaseId
        end
    else
        begin
            update chk.SportFriend SET SportBaseId = @BSportBaseId WHERE SportBaseId = @ASportBaseId
            delete chk.SportBase where SportBaseId = @ASportBaseId
        end
end