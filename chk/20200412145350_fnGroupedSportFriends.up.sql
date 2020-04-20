create or alter function chk.fnGroupedSportFriends()
    returns table return
        select ASportName, AServiceName, ASportId, BSportName, BServiceName, BSportId, count(aSportId) TotalCount
        from (
                 select a.SportName   aSportName,
                        a.ServiceName aServiceName,
                        a.SportId     aSportId,
                        b.SportName   BSportName,
                        b.ServiceName BServiceName,
                        b.SportId     BSportId
                 from Surebet sur
                          cross apply fnFullByMarketId(sur.AMarketId) a
                          cross apply fnFullByMarketId(sur.BMarketId) b
                 where a.ServiceId < b.ServiceId
                 union all
                 select b.SportName, b.ServiceName, b.SportId, a.SportName, a.ServiceName, a.SportId
                 from Surebet sur
                          cross apply fnFullByMarketId(sur.AMarketId) a
                          cross apply fnFullByMarketId(sur.BMarketId) b
                 where a.ServiceId > b.ServiceId
             ) t
        group by t.aSportName, t.aServiceName, aSportId, t.BSportName, t.BServiceName, t.BSportId