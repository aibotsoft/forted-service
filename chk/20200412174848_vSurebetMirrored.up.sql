create or alter view chk.vSurebetMirrored as
select AMarketId, BMarketId
from (
         select AMarketId, BMarketId
         from Surebet
         union all
         select BMarketId, AMarketId
         from Surebet) t