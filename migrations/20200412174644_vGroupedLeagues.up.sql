create view chk.vGroupedLeagues as
select AServiceId,
       ASportName,
       ASportId,
       ALeagueId,
       ALeagueName,
       BServiceId,
       BSportName,
       BSportId,
       BLeagueId,
       BLeagueName,
       count(ASportId) TotalCount
from (select a_sport.ServiceId   AServiceId,
             a_sport.SportName   ASportName,
             a_sport.SportId     ASportId,
             a_league.LeagueId   ALeagueId,
             a_league.LeagueName ALeagueName,
             b_sport.ServiceId   BServiceId,
             b_sport.SportName   BSportName,
             b_sport.SportId     BSportId,
             b_league.LeagueId   BLeagueId,
             b_league.LeagueName BLeagueName
      from vSurebetMirrored sb
               join Market aM on sb.AMarketId = am.MarketId
               join Event a_event on aM.EventId = a_event.EventId
               join League a_league on a_event.LeagueId = a_league.LeagueId
               join Sport a_sport on a_league.SportId = a_sport.SportId
               join chk.SportFriend a_sport_friend on a_sport_friend.SportId = a_sport.SportId

               join Market b_market on sb.BMarketId = b_market.MarketId
               join Event b_event on b_market.EventId = b_event.EventId
               join League b_league on b_league.LeagueId = b_event.LeagueId
               join Sport b_sport on b_sport.SportId = b_league.SportId
               join chk.SportFriend b_sport_friend on b_sport_friend.SportId = b_sport.SportId
      where a_sport.ServiceId < b_sport.ServiceId
        and a_sport_friend.SportBaseId = b_sport_friend.SportBaseId) t
group by AServiceId, ASportName, ASportId, ALeagueId, ALeagueName, BServiceId, BSportName, BSportId, BLeagueId,
         BLeagueName
having count(ASportId) > 4