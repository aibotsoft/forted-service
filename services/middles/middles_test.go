package middles

import (
	"encoding/json"
	"github.com/aibotsoft/micro/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var log *zap.SugaredLogger

type DemoSurebet struct {
	Stavka1 string
	Stavka2 string
	Tip     string
}

func Test_calcMiddle(t *testing.T) {
	log = logger.New()
	var list []DemoSurebet
	err := json.Unmarshal([]byte(s), &list)
	assert.NoError(t, err)
	for _, surebet := range list {
		got, err := CalcMiddle(surebet.Stavka1, surebet.Stavka2)
		if assert.NoError(t, err) {
			//if surebet.Tip == "VILKA" {
			//	assert.Equal(t, float64(0), got, surebet)
			//}
			if surebet.Tip == "KARIDOR" {
				assert.Greater(t, got, float64(0), surebet)
			}
		}
		log.Info(got, surebet)
	}
}

var s = `
[
  {
    "stavka1": "1",
    "stavka2": "Ф2(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2",
    "stavka2": "Ф1(-0,75)",
    "tip": "VILKA"
  }
]
`

var s2 = `
[
  {
    "stavka1": "(карты) ТМ(2,5)",
    "stavka2": "(карты) ТБ(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) ТМ(3,5)",
    "stavka2": "(карты) ТБ(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) ТМ(4,5)",
    "stavka2": "(карты) ТБ(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) ТМ(5,5)",
    "stavka2": "(карты) ТБ(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(-0,5)",
    "stavka2": "(карты) Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(-1,5)",
    "stavka2": "(карты) Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(-1,5)",
    "stavka2": "(карты) Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(-2,5)",
    "stavka2": "(карты) Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(-2,5)",
    "stavka2": "(карты) Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(0,5)",
    "stavka2": "(карты) Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(1,5)",
    "stavka2": "(карты) Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(1,5)",
    "stavka2": "(карты) Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф1(2,5)",
    "stavka2": "(карты) Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф2(-1,5)",
    "stavka2": "(карты) Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф2(-1,5)",
    "stavka2": "(карты) Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф2(-2,5)",
    "stavka2": "(карты) Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф2(-2,5)",
    "stavka2": "(карты) Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф2(1,5)",
    "stavka2": "(карты) Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф2(1,5)",
    "stavka2": "(карты) Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(карты) Ф2(2,5)",
    "stavka2": "(карты) Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) П1",
    "stavka2": "(сеты) Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) П1",
    "stavka2": "(сеты) Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) П1",
    "stavka2": "(сеты) Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) П1",
    "stavka2": "(сеты) Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) П2",
    "stavka2": "(сеты) Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) П2",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) П2",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(-0,5)",
    "stavka2": "(сеты) Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(-0,5)",
    "stavka2": "(сеты) Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(-1)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(-1,5)",
    "stavka2": "(сеты) Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(-1,5)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(-1,5)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(-1,5)",
    "stavka2": "(сеты) Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(-1,5)",
    "stavka2": "(сеты) Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(-2,5)",
    "stavka2": "(сеты) Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(-2,5)",
    "stavka2": "(сеты) Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(-3,5)",
    "stavka2": "(сеты) Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(0)",
    "stavka2": "(сеты) Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(0,5)",
    "stavka2": "(сеты) П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(0,5)",
    "stavka2": "(сеты) П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(0,5)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(1)",
    "stavka2": "(сеты) П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(1)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(1,5)",
    "stavka2": "(сеты) Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(1,5)",
    "stavka2": "(сеты) Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(1,5)",
    "stavka2": "(сеты) Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(1,5)",
    "stavka2": "(сеты) Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(1,5)",
    "stavka2": "(сеты) Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(1,5)",
    "stavka2": "(сеты) Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(1,5)",
    "stavka2": "(сеты) Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(1,5)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(2)",
    "stavka2": "(сеты) П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(2)",
    "stavka2": "(сеты) Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(2)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф1(2,5)",
    "stavka2": "(сеты) Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(2,5)",
    "stavka2": "(сеты) Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф1(2,5)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф2(-1,5)",
    "stavka2": "(сеты) Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф2(-1,5)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф2(-2,5)",
    "stavka2": "(сеты) Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф2(-2,5)",
    "stavka2": "(сеты) Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф2(0)",
    "stavka2": "(сеты) Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф2(1,5)",
    "stavka2": "(сеты) П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф2(1,5)",
    "stavka2": "(сеты) Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф2(1,5)",
    "stavka2": "(сеты) Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф2(2,5)",
    "stavka2": "(сеты) Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "(сеты) Ф2(2,5)",
    "stavka2": "(сеты) Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "(сеты) Ф2(2,5)",
    "stavka2": "(сеты) Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1",
    "stavka2": "1Х",
    "tip": "VILKA"
  },

  {
    "stavka1": "1",
    "stavka2": "Х2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1",
    "stavka2": "Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1",
    "stavka2": "Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1",
    "stavka2": "Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1",
    "stavka2": "Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1",
    "stavka2": "Ф2(0,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 1",
    "stavka2": "1/2 Х2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 1",
    "stavka2": "1/2 Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 1",
    "stavka2": "1/2 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 1Х",
    "stavka2": "1/2 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 2",
    "stavka2": "1/2 Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 2",
    "stavka2": "1/2 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(0,5)",
    "stavka2": "1/2 ИТ1Б(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(0,75)",
    "stavka2": "1/2 ИТ1Б(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(1)",
    "stavka2": "1/2 ИТ1Б(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(1,5)",
    "stavka2": "1/2 ИТ1Б(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(53)",
    "stavka2": "1/2 ИТ2Б(52,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(53,5)",
    "stavka2": "1/2 ИТ2Б(53,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(54)",
    "stavka2": "1/2 ИТ2Б(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(54,5)",
    "stavka2": "1/2 ИТ2Б(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(55)",
    "stavka2": "1/2 ИТ1Б(54,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(56)",
    "stavka2": "1/2 ИТ2Б(55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(56,5)",
    "stavka2": "1/2 ИТ2Б(56)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(56,5)",
    "stavka2": "1/2 ИТ2Б(56,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(57)",
    "stavka2": "1/2 ИТ2Б(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(57,5)",
    "stavka2": "1/2 ИТ2Б(57)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(58)",
    "stavka2": "1/2 ИТ2Б(57,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(58)",
    "stavka2": "1/2 ИТ2Б(57,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(59)",
    "stavka2": "1/2 ИТ2Б(59)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ1М(59,5)",
    "stavka2": "1/2 ИТ2Б(58)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(59,5)",
    "stavka2": "1/2 ИТ2Б(59)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ1М(60,5)",
    "stavka2": "1/2 ИТ1Б(60,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(0,5)",
    "stavka2": "1/2 ИТ2Б(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(2,5)",
    "stavka2": "1/2 ИТ1Б(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(3,5)",
    "stavka2": "1/2 ИТ1Б(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(45,5)",
    "stavka2": "1/2 ИТ2Б(45,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(51)",
    "stavka2": "1/2 ИТ1Б(51)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(51,5)",
    "stavka2": "1/2 ИТ2Б(51,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(52,5)",
    "stavka2": "1/2 ИТ1Б(51,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(52,5)",
    "stavka2": "1/2 ИТ2Б(51,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(53)",
    "stavka2": "1/2 ИТ1Б(52,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(53)",
    "stavka2": "1/2 ИТ1Б(53)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(53,5)",
    "stavka2": "1/2 ИТ1Б(52)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(53,5)",
    "stavka2": "1/2 ИТ1Б(53)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(53,5)",
    "stavka2": "1/2 ИТ1Б(53,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(53,5)",
    "stavka2": "1/2 ИТ2Б(53,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(54)",
    "stavka2": "1/2 ИТ1Б(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(54,5)",
    "stavka2": "1/2 ИТ1Б(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(54,5)",
    "stavka2": "1/2 ИТ1Б(54)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(55)",
    "stavka2": "1/2 ИТ1Б(54,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(55)",
    "stavka2": "1/2 ИТ1Б(55)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(55,5)",
    "stavka2": "1/2 ИТ1Б(55,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(56)",
    "stavka2": "1/2 ИТ1Б(55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(56,5)",
    "stavka2": "1/2 ИТ1Б(56)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(56,5)",
    "stavka2": "1/2 ИТ1Б(56,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(56,5)",
    "stavka2": "1/2 ИТ2Б(56,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(57,5)",
    "stavka2": "1/2 ИТ2Б(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(58)",
    "stavka2": "1/2 ИТ1Б(57,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(60)",
    "stavka2": "1/2 ИТ1Б(59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ИТ2М(61,5)",
    "stavka2": "1/2 ИТ1Б(61,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ИТ2М(62)",
    "stavka2": "1/2 ИТ1Б(61,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 П1",
    "stavka2": "1/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 П1",
    "stavka2": "1/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 П1",
    "stavka2": "1/2 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 П1",
    "stavka2": "1/2 Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 П1",
    "stavka2": "1/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 П1",
    "stavka2": "1/2 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 П1",
    "stavka2": "1/2 Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 П2",
    "stavka2": "1/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 П2",
    "stavka2": "1/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 П2",
    "stavka2": "1/2 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 П2",
    "stavka2": "1/2 Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 П2",
    "stavka2": "1/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 П2",
    "stavka2": "1/2 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 П2",
    "stavka2": "1/2 Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 П2",
    "stavka2": "1/2 Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(0,5)",
    "stavka2": "1/2 ТБ(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(0,75)",
    "stavka2": "1/2 ТБ(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(1)",
    "stavka2": "1/2 ТБ(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(1,25)",
    "stavka2": "1/2 ТБ(1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(1,5)",
    "stavka2": "1/2 ТБ(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(1,75)",
    "stavka2": "1/2 ТБ(1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(100)",
    "stavka2": "1/2 ТБ(100)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(100)",
    "stavka2": "1/2 ТБ(99,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(100,5)",
    "stavka2": "1/2 ТБ(100)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(100,5)",
    "stavka2": "1/2 ТБ(100,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(100,5)",
    "stavka2": "1/2 ТБ(99,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(101)",
    "stavka2": "1/2 ТБ(100,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(101,5)",
    "stavka2": "1/2 ТБ(101)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(101,5)",
    "stavka2": "1/2 ТБ(101,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(102)",
    "stavka2": "1/2 ТБ(101,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(102)",
    "stavka2": "1/2 ТБ(102)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(102,5)",
    "stavka2": "1/2 ТБ(100,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(102,5)",
    "stavka2": "1/2 ТБ(101,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(102,5)",
    "stavka2": "1/2 ТБ(102)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(102,5)",
    "stavka2": "1/2 ТБ(102,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(103)",
    "stavka2": "1/2 ТБ(102,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(103)",
    "stavka2": "1/2 ТБ(103)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(103,5)",
    "stavka2": "1/2 ТБ(102)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(103,5)",
    "stavka2": "1/2 ТБ(103)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(103,5)",
    "stavka2": "1/2 ТБ(103)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(103,5)",
    "stavka2": "1/2 ТБ(103,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(104)",
    "stavka2": "1/2 ТБ(104)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(104,5)",
    "stavka2": "1/2 ТБ(103,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(104,5)",
    "stavka2": "1/2 ТБ(104)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(104,5)",
    "stavka2": "1/2 ТБ(104,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(105)",
    "stavka2": "1/2 ТБ(103,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(105)",
    "stavka2": "1/2 ТБ(104)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(105)",
    "stavka2": "1/2 ТБ(105)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(105,5)",
    "stavka2": "1/2 ТБ(105)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(105,5)",
    "stavka2": "1/2 ТБ(105,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(106)",
    "stavka2": "1/2 ТБ(105,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(106)",
    "stavka2": "1/2 ТБ(106)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(106,5)",
    "stavka2": "1/2 ТБ(104,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(106,5)",
    "stavka2": "1/2 ТБ(105)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(106,5)",
    "stavka2": "1/2 ТБ(105,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(106,5)",
    "stavka2": "1/2 ТБ(106)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(106,5)",
    "stavka2": "1/2 ТБ(106,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(107)",
    "stavka2": "1/2 ТБ(106)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(107)",
    "stavka2": "1/2 ТБ(106,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(107)",
    "stavka2": "1/2 ТБ(107)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(107,5)",
    "stavka2": "1/2 ТБ(105,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(107,5)",
    "stavka2": "1/2 ТБ(106,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(107,5)",
    "stavka2": "1/2 ТБ(107)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(107,5)",
    "stavka2": "1/2 ТБ(107,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(108)",
    "stavka2": "1/2 ТБ(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(108)",
    "stavka2": "1/2 ТБ(108)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(108,5)",
    "stavka2": "1/2 ТБ(106,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(108,5)",
    "stavka2": "1/2 ТБ(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(108,5)",
    "stavka2": "1/2 ТБ(108)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(108,5)",
    "stavka2": "1/2 ТБ(108,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(109)",
    "stavka2": "1/2 ТБ(108,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(109)",
    "stavka2": "1/2 ТБ(108,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(109)",
    "stavka2": "1/2 ТБ(109)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(109,5)",
    "stavka2": "1/2 ТБ(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(109,5)",
    "stavka2": "1/2 ТБ(108)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(109,5)",
    "stavka2": "1/2 ТБ(108,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(109,5)",
    "stavka2": "1/2 ТБ(109)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(109,5)",
    "stavka2": "1/2 ТБ(109,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(110)",
    "stavka2": "1/2 ТБ(108,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(110)",
    "stavka2": "1/2 ТБ(109,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(110)",
    "stavka2": "1/2 ТБ(110)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(110,5)",
    "stavka2": "1/2 ТБ(109)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(110,5)",
    "stavka2": "1/2 ТБ(110)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(110,5)",
    "stavka2": "1/2 ТБ(110,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(111)",
    "stavka2": "1/2 ТБ(110)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(111)",
    "stavka2": "1/2 ТБ(110,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(111)",
    "stavka2": "1/2 ТБ(111)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(111,5)",
    "stavka2": "1/2 ТБ(110)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(111,5)",
    "stavka2": "1/2 ТБ(111)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(111,5)",
    "stavka2": "1/2 ТБ(111,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(112)",
    "stavka2": "1/2 ТБ(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(112)",
    "stavka2": "1/2 ТБ(112)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(112,5)",
    "stavka2": "1/2 ТБ(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(112,5)",
    "stavka2": "1/2 ТБ(112)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(112,5)",
    "stavka2": "1/2 ТБ(112)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(112,5)",
    "stavka2": "1/2 ТБ(112,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(113)",
    "stavka2": "1/2 ТБ(112,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(113)",
    "stavka2": "1/2 ТБ(113)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(113,5)",
    "stavka2": "1/2 ТБ(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(113,5)",
    "stavka2": "1/2 ТБ(112,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(113,5)",
    "stavka2": "1/2 ТБ(113)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(113,5)",
    "stavka2": "1/2 ТБ(113,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(114)",
    "stavka2": "1/2 ТБ(113)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(114)",
    "stavka2": "1/2 ТБ(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(114)",
    "stavka2": "1/2 ТБ(114)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(114,5)",
    "stavka2": "1/2 ТБ(114)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(114,5)",
    "stavka2": "1/2 ТБ(114)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(114,5)",
    "stavka2": "1/2 ТБ(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(114,5)",
    "stavka2": "1/2 ТБ(114,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(115)",
    "stavka2": "1/2 ТБ(111)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(115)",
    "stavka2": "1/2 ТБ(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(115)",
    "stavka2": "1/2 ТБ(115)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(115,5)",
    "stavka2": "1/2 ТБ(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(115,5)",
    "stavka2": "1/2 ТБ(115)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(115,5)",
    "stavka2": "1/2 ТБ(115,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(116)",
    "stavka2": "1/2 ТБ(115,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(116)",
    "stavka2": "1/2 ТБ(116)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(116,5)",
    "stavka2": "1/2 ТБ(115,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(116,5)",
    "stavka2": "1/2 ТБ(116)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(116,5)",
    "stavka2": "1/2 ТБ(116,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(117)",
    "stavka2": "1/2 ТБ(116,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(117)",
    "stavka2": "1/2 ТБ(117)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(117,5)",
    "stavka2": "1/2 ТБ(116,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(117,5)",
    "stavka2": "1/2 ТБ(117)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(117,5)",
    "stavka2": "1/2 ТБ(117,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(118)",
    "stavka2": "1/2 ТБ(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(118)",
    "stavka2": "1/2 ТБ(117,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(118)",
    "stavka2": "1/2 ТБ(118)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(118,5)",
    "stavka2": "1/2 ТБ(117)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(118,5)",
    "stavka2": "1/2 ТБ(118)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(118,5)",
    "stavka2": "1/2 ТБ(118,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(119)",
    "stavka2": "1/2 ТБ(117)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(119)",
    "stavka2": "1/2 ТБ(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(119)",
    "stavka2": "1/2 ТБ(118)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(119)",
    "stavka2": "1/2 ТБ(118,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(119)",
    "stavka2": "1/2 ТБ(119)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(119,5)",
    "stavka2": "1/2 ТБ(119)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(119,5)",
    "stavka2": "1/2 ТБ(119,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(120)",
    "stavka2": "1/2 ТБ(119,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(120)",
    "stavka2": "1/2 ТБ(120)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(118,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(118,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(119)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(119)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(119,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(119,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(120)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(120,5)",
    "stavka2": "1/2 ТБ(120,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(121)",
    "stavka2": "1/2 ТБ(121)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(121,5)",
    "stavka2": "1/2 ТБ(121)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(121,5)",
    "stavka2": "1/2 ТБ(121,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(122)",
    "stavka2": "1/2 ТБ(121,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(122)",
    "stavka2": "1/2 ТБ(122)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(122,5)",
    "stavka2": "1/2 ТБ(122,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(123)",
    "stavka2": "1/2 ТБ(122,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(123,5)",
    "stavka2": "1/2 ТБ(123)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(124)",
    "stavka2": "1/2 ТБ(123,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(124,5)",
    "stavka2": "1/2 ТБ(124)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(124,5)",
    "stavka2": "1/2 ТБ(124,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(125)",
    "stavka2": "1/2 ТБ(124,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(125)",
    "stavka2": "1/2 ТБ(125)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(125,5)",
    "stavka2": "1/2 ТБ(124,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(125,5)",
    "stavka2": "1/2 ТБ(125,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(126,5)",
    "stavka2": "1/2 ТБ(126,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(127)",
    "stavka2": "1/2 ТБ(126,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(156,5)",
    "stavka2": "1/2 ТБ(156,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(18,5)",
    "stavka2": "1/2 ТБ(18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(19)",
    "stavka2": "1/2 ТБ(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(2)",
    "stavka2": "1/2 ТБ(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(2,25)",
    "stavka2": "1/2 ТБ(2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(21)",
    "stavka2": "1/2 ТБ(21)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(21,5)",
    "stavka2": "1/2 ТБ(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(22)",
    "stavka2": "1/2 ТБ(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(22,5)",
    "stavka2": "1/2 ТБ(22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(23)",
    "stavka2": "1/2 ТБ(23)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(23,5)",
    "stavka2": "1/2 ТБ(23)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(23,5)",
    "stavka2": "1/2 ТБ(23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(24)",
    "stavka2": "1/2 ТБ(24)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(24,5)",
    "stavka2": "1/2 ТБ(24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(25,5)",
    "stavka2": "1/2 ТБ(24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(25,5)",
    "stavka2": "1/2 ТБ(25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(26,5)",
    "stavka2": "1/2 ТБ(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(27,5)",
    "stavka2": "1/2 ТБ(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(28,5)",
    "stavka2": "1/2 ТБ(28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(29,5)",
    "stavka2": "1/2 ТБ(29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(3)",
    "stavka2": "1/2 ТБ(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(3,5)",
    "stavka2": "1/2 ТБ(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(30,5)",
    "stavka2": "1/2 ТБ(30,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(31)",
    "stavka2": "1/2 ТБ(31)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(31,5)",
    "stavka2": "1/2 ТБ(31,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(32)",
    "stavka2": "1/2 ТБ(31,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(32)",
    "stavka2": "1/2 ТБ(31,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(32,5)",
    "stavka2": "1/2 ТБ(32,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(33)",
    "stavka2": "1/2 ТБ(32,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(33,5)",
    "stavka2": "1/2 ТБ(33,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(34,5)",
    "stavka2": "1/2 ТБ(34,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(35,5)",
    "stavka2": "1/2 ТБ(35,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(36,5)",
    "stavka2": "1/2 ТБ(36,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(38,5)",
    "stavka2": "1/2 ТБ(38,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(4)",
    "stavka2": "1/2 ТБ(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(4,5)",
    "stavka2": "1/2 ТБ(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(41,5)",
    "stavka2": "1/2 ТБ(41,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(48,5)",
    "stavka2": "1/2 ТБ(48)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(49,5)",
    "stavka2": "1/2 ТБ(48,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(49,5)",
    "stavka2": "1/2 ТБ(49,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(5)",
    "stavka2": "1/2 ТБ(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(5,5)",
    "stavka2": "1/2 ТБ(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(50,5)",
    "stavka2": "1/2 ТБ(50,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(52)",
    "stavka2": "1/2 ТБ(52)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(53)",
    "stavka2": "1/2 ТБ(53)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(53,5)",
    "stavka2": "1/2 ТБ(53)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(53,5)",
    "stavka2": "1/2 ТБ(53,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(54)",
    "stavka2": "1/2 ТБ(54)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(54,5)",
    "stavka2": "1/2 ТБ(54,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(55)",
    "stavka2": "1/2 ТБ(54,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(55)",
    "stavka2": "1/2 ТБ(55)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(55,5)",
    "stavka2": "1/2 ТБ(55)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(55,5)",
    "stavka2": "1/2 ТБ(55,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(56)",
    "stavka2": "1/2 ТБ(55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(56)",
    "stavka2": "1/2 ТБ(56)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(56,5)",
    "stavka2": "1/2 ТБ(55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(56,5)",
    "stavka2": "1/2 ТБ(56,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(57)",
    "stavka2": "1/2 ТБ(55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(57)",
    "stavka2": "1/2 ТБ(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(57)",
    "stavka2": "1/2 ТБ(57)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(57,5)",
    "stavka2": "1/2 ТБ(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(57,5)",
    "stavka2": "1/2 ТБ(57)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(57,5)",
    "stavka2": "1/2 ТБ(57)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(57,5)",
    "stavka2": "1/2 ТБ(57,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(58)",
    "stavka2": "1/2 ТБ(57,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(58)",
    "stavka2": "1/2 ТБ(57,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(58)",
    "stavka2": "1/2 ТБ(58)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(58,5)",
    "stavka2": "1/2 ТБ(57,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(58,5)",
    "stavka2": "1/2 ТБ(58)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(58,5)",
    "stavka2": "1/2 ТБ(58)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(58,5)",
    "stavka2": "1/2 ТБ(58,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(59)",
    "stavka2": "1/2 ТБ(58,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(59)",
    "stavka2": "1/2 ТБ(59)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(59,5)",
    "stavka2": "1/2 ТБ(59)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(59,5)",
    "stavka2": "1/2 ТБ(59,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(6)",
    "stavka2": "1/2 ТБ(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(6,5)",
    "stavka2": "1/2 ТБ(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(60)",
    "stavka2": "1/2 ТБ(59)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(60)",
    "stavka2": "1/2 ТБ(59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(60)",
    "stavka2": "1/2 ТБ(60)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(60,5)",
    "stavka2": "1/2 ТБ(57)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(60,5)",
    "stavka2": "1/2 ТБ(58)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(60,5)",
    "stavka2": "1/2 ТБ(58,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(60,5)",
    "stavka2": "1/2 ТБ(59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(60,5)",
    "stavka2": "1/2 ТБ(60)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(60,5)",
    "stavka2": "1/2 ТБ(60,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(61)",
    "stavka2": "1/2 ТБ(59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61)",
    "stavka2": "1/2 ТБ(60)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61)",
    "stavka2": "1/2 ТБ(60,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61)",
    "stavka2": "1/2 ТБ(61)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(61,5)",
    "stavka2": "1/2 ТБ(58,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61,5)",
    "stavka2": "1/2 ТБ(59)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61,5)",
    "stavka2": "1/2 ТБ(59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61,5)",
    "stavka2": "1/2 ТБ(60)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61,5)",
    "stavka2": "1/2 ТБ(60,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61,5)",
    "stavka2": "1/2 ТБ(61)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(61,5)",
    "stavka2": "1/2 ТБ(61)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(61,5)",
    "stavka2": "1/2 ТБ(61,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(62)",
    "stavka2": "1/2 ТБ(60,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(62)",
    "stavka2": "1/2 ТБ(61)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(62)",
    "stavka2": "1/2 ТБ(61)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(62)",
    "stavka2": "1/2 ТБ(61,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(62)",
    "stavka2": "1/2 ТБ(61,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(62)",
    "stavka2": "1/2 ТБ(62)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(62,5)",
    "stavka2": "1/2 ТБ(59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(62,5)",
    "stavka2": "1/2 ТБ(60)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(62,5)",
    "stavka2": "1/2 ТБ(60,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(62,5)",
    "stavka2": "1/2 ТБ(61,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(62,5)",
    "stavka2": "1/2 ТБ(62)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(62,5)",
    "stavka2": "1/2 ТБ(62)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(62,5)",
    "stavka2": "1/2 ТБ(62,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(63)",
    "stavka2": "1/2 ТБ(62)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(63)",
    "stavka2": "1/2 ТБ(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(63)",
    "stavka2": "1/2 ТБ(62,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(63)",
    "stavka2": "1/2 ТБ(63)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(63,5)",
    "stavka2": "1/2 ТБ(60)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(63,5)",
    "stavka2": "1/2 ТБ(61)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(63,5)",
    "stavka2": "1/2 ТБ(61)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(63,5)",
    "stavka2": "1/2 ТБ(61,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(63,5)",
    "stavka2": "1/2 ТБ(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(63,5)",
    "stavka2": "1/2 ТБ(63)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(63,5)",
    "stavka2": "1/2 ТБ(63)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(63,5)",
    "stavka2": "1/2 ТБ(63,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(64)",
    "stavka2": "1/2 ТБ(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(64)",
    "stavka2": "1/2 ТБ(63)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(64)",
    "stavka2": "1/2 ТБ(63,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(64)",
    "stavka2": "1/2 ТБ(63,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(64)",
    "stavka2": "1/2 ТБ(64)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(64,5)",
    "stavka2": "1/2 ТБ(61,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(64,5)",
    "stavka2": "1/2 ТБ(62)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(64,5)",
    "stavka2": "1/2 ТБ(63)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(64,5)",
    "stavka2": "1/2 ТБ(63,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(64,5)",
    "stavka2": "1/2 ТБ(64)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(64,5)",
    "stavka2": "1/2 ТБ(64)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(64,5)",
    "stavka2": "1/2 ТБ(64,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(65)",
    "stavka2": "1/2 ТБ(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65)",
    "stavka2": "1/2 ТБ(63)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65)",
    "stavka2": "1/2 ТБ(63,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65)",
    "stavka2": "1/2 ТБ(64)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65)",
    "stavka2": "1/2 ТБ(64)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(65)",
    "stavka2": "1/2 ТБ(64,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65)",
    "stavka2": "1/2 ТБ(65)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(65,5)",
    "stavka2": "1/2 ТБ(61,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65,5)",
    "stavka2": "1/2 ТБ(63,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65,5)",
    "stavka2": "1/2 ТБ(64)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65,5)",
    "stavka2": "1/2 ТБ(64,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65,5)",
    "stavka2": "1/2 ТБ(65)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(65,5)",
    "stavka2": "1/2 ТБ(65)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(65,5)",
    "stavka2": "1/2 ТБ(65,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(63)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(63)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(63,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(64)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(64,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(65)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(65,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(65,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(66)",
    "stavka2": "1/2 ТБ(66)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(66,5)",
    "stavka2": "1/2 ТБ(65)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66,5)",
    "stavka2": "1/2 ТБ(65,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66,5)",
    "stavka2": "1/2 ТБ(66)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(66,5)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(67)",
    "stavka2": "1/2 ТБ(64)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(67)",
    "stavka2": "1/2 ТБ(65,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(67)",
    "stavka2": "1/2 ТБ(66)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(67)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(67)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(67)",
    "stavka2": "1/2 ТБ(67)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(67,5)",
    "stavka2": "1/2 ТБ(66)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(67,5)",
    "stavka2": "1/2 ТБ(66)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(67,5)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(67,5)",
    "stavka2": "1/2 ТБ(67)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(67,5)",
    "stavka2": "1/2 ТБ(67)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(67,5)",
    "stavka2": "1/2 ТБ(67,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(68)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(68)",
    "stavka2": "1/2 ТБ(67)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(68)",
    "stavka2": "1/2 ТБ(67,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(68)",
    "stavka2": "1/2 ТБ(67,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(68)",
    "stavka2": "1/2 ТБ(68)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(68,5)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(68,5)",
    "stavka2": "1/2 ТБ(67)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(68,5)",
    "stavka2": "1/2 ТБ(67,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(68,5)",
    "stavka2": "1/2 ТБ(68)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(68,5)",
    "stavka2": "1/2 ТБ(68)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(68,5)",
    "stavka2": "1/2 ТБ(68,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(69)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(69)",
    "stavka2": "1/2 ТБ(67,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(69)",
    "stavka2": "1/2 ТБ(68)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(69)",
    "stavka2": "1/2 ТБ(68,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(69)",
    "stavka2": "1/2 ТБ(69)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(69,5)",
    "stavka2": "1/2 ТБ(68)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(69,5)",
    "stavka2": "1/2 ТБ(68,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(69,5)",
    "stavka2": "1/2 ТБ(69)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(69,5)",
    "stavka2": "1/2 ТБ(69)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(69,5)",
    "stavka2": "1/2 ТБ(69,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(7)",
    "stavka2": "1/2 ТБ(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(70)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70)",
    "stavka2": "1/2 ТБ(68,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70)",
    "stavka2": "1/2 ТБ(69)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70)",
    "stavka2": "1/2 ТБ(69,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70)",
    "stavka2": "1/2 ТБ(69,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(70)",
    "stavka2": "1/2 ТБ(70)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(66,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(67)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(67)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(68,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(69)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(69,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(70)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(70,5)",
    "stavka2": "1/2 ТБ(70,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(71)",
    "stavka2": "1/2 ТБ(67,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(71)",
    "stavka2": "1/2 ТБ(67,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(71)",
    "stavka2": "1/2 ТБ(69,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(71)",
    "stavka2": "1/2 ТБ(70)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(71)",
    "stavka2": "1/2 ТБ(70,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(71)",
    "stavka2": "1/2 ТБ(71)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(71,5)",
    "stavka2": "1/2 ТБ(70,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(71,5)",
    "stavka2": "1/2 ТБ(71)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(71,5)",
    "stavka2": "1/2 ТБ(71)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(71,5)",
    "stavka2": "1/2 ТБ(71,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(72)",
    "stavka2": "1/2 ТБ(71)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(72)",
    "stavka2": "1/2 ТБ(71,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(72)",
    "stavka2": "1/2 ТБ(72)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(72,5)",
    "stavka2": "1/2 ТБ(71)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(72,5)",
    "stavka2": "1/2 ТБ(71,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(72,5)",
    "stavka2": "1/2 ТБ(72)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(72,5)",
    "stavka2": "1/2 ТБ(72)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(72,5)",
    "stavka2": "1/2 ТБ(72,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(73)",
    "stavka2": "1/2 ТБ(72)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(73)",
    "stavka2": "1/2 ТБ(72,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(73)",
    "stavka2": "1/2 ТБ(73)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(73,5)",
    "stavka2": "1/2 ТБ(72,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(73,5)",
    "stavka2": "1/2 ТБ(73)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(73,5)",
    "stavka2": "1/2 ТБ(73)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(73,5)",
    "stavka2": "1/2 ТБ(73,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(74)",
    "stavka2": "1/2 ТБ(72,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(74)",
    "stavka2": "1/2 ТБ(73)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(74)",
    "stavka2": "1/2 ТБ(73)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(74)",
    "stavka2": "1/2 ТБ(73,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(74)",
    "stavka2": "1/2 ТБ(74)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(74,5)",
    "stavka2": "1/2 ТБ(73)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(74,5)",
    "stavka2": "1/2 ТБ(73,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(74,5)",
    "stavka2": "1/2 ТБ(74)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(74,5)",
    "stavka2": "1/2 ТБ(74,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(75)",
    "stavka2": "1/2 ТБ(74)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(75)",
    "stavka2": "1/2 ТБ(74,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(75)",
    "stavka2": "1/2 ТБ(75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(75,5)",
    "stavka2": "1/2 ТБ(74)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(75,5)",
    "stavka2": "1/2 ТБ(74,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(75,5)",
    "stavka2": "1/2 ТБ(75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(75,5)",
    "stavka2": "1/2 ТБ(75,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(76)",
    "stavka2": "1/2 ТБ(73,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(76)",
    "stavka2": "1/2 ТБ(75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(76)",
    "stavka2": "1/2 ТБ(75,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(76)",
    "stavka2": "1/2 ТБ(75,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(76)",
    "stavka2": "1/2 ТБ(76)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(76,5)",
    "stavka2": "1/2 ТБ(74,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(76,5)",
    "stavka2": "1/2 ТБ(75,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(76,5)",
    "stavka2": "1/2 ТБ(76)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(76,5)",
    "stavka2": "1/2 ТБ(76,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(77)",
    "stavka2": "1/2 ТБ(76)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(77)",
    "stavka2": "1/2 ТБ(76,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(77)",
    "stavka2": "1/2 ТБ(77)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(77,5)",
    "stavka2": "1/2 ТБ(75,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(77,5)",
    "stavka2": "1/2 ТБ(76)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(77,5)",
    "stavka2": "1/2 ТБ(76,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(77,5)",
    "stavka2": "1/2 ТБ(77)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(77,5)",
    "stavka2": "1/2 ТБ(77,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(78)",
    "stavka2": "1/2 ТБ(76,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(78)",
    "stavka2": "1/2 ТБ(77)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(78)",
    "stavka2": "1/2 ТБ(77,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(78)",
    "stavka2": "1/2 ТБ(78)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(78,5)",
    "stavka2": "1/2 ТБ(77,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(78,5)",
    "stavka2": "1/2 ТБ(78)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(78,5)",
    "stavka2": "1/2 ТБ(78,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(79)",
    "stavka2": "1/2 ТБ(78,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(79)",
    "stavka2": "1/2 ТБ(78,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(79)",
    "stavka2": "1/2 ТБ(79)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(79,5)",
    "stavka2": "1/2 ТБ(78,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(79,5)",
    "stavka2": "1/2 ТБ(78,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(79,5)",
    "stavka2": "1/2 ТБ(79)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(79,5)",
    "stavka2": "1/2 ТБ(79,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(80)",
    "stavka2": "1/2 ТБ(78,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(80)",
    "stavka2": "1/2 ТБ(79)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(80)",
    "stavka2": "1/2 ТБ(79,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(80)",
    "stavka2": "1/2 ТБ(79,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(80)",
    "stavka2": "1/2 ТБ(80)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(80,5)",
    "stavka2": "1/2 ТБ(78,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(80,5)",
    "stavka2": "1/2 ТБ(79)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(80,5)",
    "stavka2": "1/2 ТБ(79,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(80,5)",
    "stavka2": "1/2 ТБ(80)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(80,5)",
    "stavka2": "1/2 ТБ(80)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(80,5)",
    "stavka2": "1/2 ТБ(80,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(81)",
    "stavka2": "1/2 ТБ(80)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(81)",
    "stavka2": "1/2 ТБ(80,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(81)",
    "stavka2": "1/2 ТБ(81)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(81,5)",
    "stavka2": "1/2 ТБ(79,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(81,5)",
    "stavka2": "1/2 ТБ(80)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(81,5)",
    "stavka2": "1/2 ТБ(80,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(81,5)",
    "stavka2": "1/2 ТБ(80,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(81,5)",
    "stavka2": "1/2 ТБ(81)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(81,5)",
    "stavka2": "1/2 ТБ(81,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(82)",
    "stavka2": "1/2 ТБ(81,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(82)",
    "stavka2": "1/2 ТБ(82)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(82,5)",
    "stavka2": "1/2 ТБ(81)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(82,5)",
    "stavka2": "1/2 ТБ(81,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(82,5)",
    "stavka2": "1/2 ТБ(82)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(82,5)",
    "stavka2": "1/2 ТБ(82,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(83)",
    "stavka2": "1/2 ТБ(82)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(83)",
    "stavka2": "1/2 ТБ(82,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(83)",
    "stavka2": "1/2 ТБ(83)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(83,5)",
    "stavka2": "1/2 ТБ(82,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(83,5)",
    "stavka2": "1/2 ТБ(83)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(83,5)",
    "stavka2": "1/2 ТБ(83,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(84)",
    "stavka2": "1/2 ТБ(82)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(84)",
    "stavka2": "1/2 ТБ(83)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(84)",
    "stavka2": "1/2 ТБ(83,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(84)",
    "stavka2": "1/2 ТБ(83,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(84)",
    "stavka2": "1/2 ТБ(84)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(84,5)",
    "stavka2": "1/2 ТБ(83,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(84,5)",
    "stavka2": "1/2 ТБ(84)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(84,5)",
    "stavka2": "1/2 ТБ(84)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(84,5)",
    "stavka2": "1/2 ТБ(84,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(85)",
    "stavka2": "1/2 ТБ(82,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(85)",
    "stavka2": "1/2 ТБ(84)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(85)",
    "stavka2": "1/2 ТБ(84,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(85)",
    "stavka2": "1/2 ТБ(84,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(85)",
    "stavka2": "1/2 ТБ(85)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(85,5)",
    "stavka2": "1/2 ТБ(85)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(85,5)",
    "stavka2": "1/2 ТБ(85,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(86)",
    "stavka2": "1/2 ТБ(85)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(86)",
    "stavka2": "1/2 ТБ(85,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(86)",
    "stavka2": "1/2 ТБ(86)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(86,5)",
    "stavka2": "1/2 ТБ(83)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(86,5)",
    "stavka2": "1/2 ТБ(85)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(86,5)",
    "stavka2": "1/2 ТБ(86)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(86,5)",
    "stavka2": "1/2 ТБ(86)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(86,5)",
    "stavka2": "1/2 ТБ(86,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(87)",
    "stavka2": "1/2 ТБ(86)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(87)",
    "stavka2": "1/2 ТБ(86,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(87)",
    "stavka2": "1/2 ТБ(87)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(87,5)",
    "stavka2": "1/2 ТБ(86)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(87,5)",
    "stavka2": "1/2 ТБ(86,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(87,5)",
    "stavka2": "1/2 ТБ(87)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(87,5)",
    "stavka2": "1/2 ТБ(87,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(88)",
    "stavka2": "1/2 ТБ(87,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(88)",
    "stavka2": "1/2 ТБ(88)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(88,5)",
    "stavka2": "1/2 ТБ(87,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(88,5)",
    "stavka2": "1/2 ТБ(88)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(88,5)",
    "stavka2": "1/2 ТБ(88,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(89)",
    "stavka2": "1/2 ТБ(87)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(89)",
    "stavka2": "1/2 ТБ(87,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(89)",
    "stavka2": "1/2 ТБ(88,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(89)",
    "stavka2": "1/2 ТБ(89)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(89,5)",
    "stavka2": "1/2 ТБ(88,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(89,5)",
    "stavka2": "1/2 ТБ(89)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(89,5)",
    "stavka2": "1/2 ТБ(89,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(90)",
    "stavka2": "1/2 ТБ(89,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(90)",
    "stavka2": "1/2 ТБ(90)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(90,5)",
    "stavka2": "1/2 ТБ(89,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(90,5)",
    "stavka2": "1/2 ТБ(90)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(90,5)",
    "stavka2": "1/2 ТБ(90,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(91)",
    "stavka2": "1/2 ТБ(90,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(91)",
    "stavka2": "1/2 ТБ(90,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(91,5)",
    "stavka2": "1/2 ТБ(90)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(91,5)",
    "stavka2": "1/2 ТБ(90,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(91,5)",
    "stavka2": "1/2 ТБ(91)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(91,5)",
    "stavka2": "1/2 ТБ(91,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(92)",
    "stavka2": "1/2 ТБ(92)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(92,5)",
    "stavka2": "1/2 ТБ(92)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(92,5)",
    "stavka2": "1/2 ТБ(92,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(93)",
    "stavka2": "1/2 ТБ(92,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(93,5)",
    "stavka2": "1/2 ТБ(93,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(94)",
    "stavka2": "1/2 ТБ(92,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(94)",
    "stavka2": "1/2 ТБ(93)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(94)",
    "stavka2": "1/2 ТБ(93,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(94)",
    "stavka2": "1/2 ТБ(93,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(94)",
    "stavka2": "1/2 ТБ(94)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(94,5)",
    "stavka2": "1/2 ТБ(94)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(94,5)",
    "stavka2": "1/2 ТБ(94,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(95)",
    "stavka2": "1/2 ТБ(94,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(95,5)",
    "stavka2": "1/2 ТБ(95,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(96)",
    "stavka2": "1/2 ТБ(95,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(96,5)",
    "stavka2": "1/2 ТБ(95,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(96,5)",
    "stavka2": "1/2 ТБ(96,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(97,5)",
    "stavka2": "1/2 ТБ(97,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(98,5)",
    "stavka2": "1/2 ТБ(97,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(98,5)",
    "stavka2": "1/2 ТБ(98)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(98,5)",
    "stavka2": "1/2 ТБ(98,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 ТМ(99)",
    "stavka2": "1/2 ТБ(98,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(99,5)",
    "stavka2": "1/2 ТБ(98,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(99,5)",
    "stavka2": "1/2 ТБ(99)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 ТМ(99,5)",
    "stavka2": "1/2 ТБ(99,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-0,25)",
    "stavka2": "1/2 Ф1(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-0,25)",
    "stavka2": "1/2 Ф2(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-0,5)",
    "stavka2": "1/2 Х2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-0,5)",
    "stavka2": "1/2 Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-0,5)",
    "stavka2": "1/2 Ф1(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-0,5)",
    "stavka2": "1/2 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-0,5)",
    "stavka2": "1/2 Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-0,75)",
    "stavka2": "1/2 Ф2(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-1)",
    "stavka2": "1/2 Ф1(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-1)",
    "stavka2": "1/2 Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-1)",
    "stavka2": "1/2 Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-1)",
    "stavka2": "1/2 Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-1)",
    "stavka2": "1/2 Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-1,25)",
    "stavka2": "1/2 Ф2(1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-1,5)",
    "stavka2": "1/2 Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-1,5)",
    "stavka2": "1/2 Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-1,5)",
    "stavka2": "1/2 Ф1(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-1,5)",
    "stavka2": "1/2 Ф1(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-1,5)",
    "stavka2": "1/2 Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-1,5)",
    "stavka2": "1/2 Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-1,5)",
    "stavka2": "1/2 Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-1,5)",
    "stavka2": "1/2 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-1,75)",
    "stavka2": "1/2 Ф2(1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-10)",
    "stavka2": "1/2 Ф1(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-10)",
    "stavka2": "1/2 Ф2(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-10,5)",
    "stavka2": "1/2 Ф1(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-10,5)",
    "stavka2": "1/2 Ф1(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-10,5)",
    "stavka2": "1/2 Ф1(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-10,5)",
    "stavka2": "1/2 Ф2(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-10,5)",
    "stavka2": "1/2 Ф2(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-11)",
    "stavka2": "1/2 Ф1(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-11)",
    "stavka2": "1/2 Ф1(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-11)",
    "stavka2": "1/2 Ф2(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-11,5)",
    "stavka2": "1/2 Ф1(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-11,5)",
    "stavka2": "1/2 Ф1(12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-11,5)",
    "stavka2": "1/2 Ф2(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-12)",
    "stavka2": "1/2 Ф1(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-12)",
    "stavka2": "1/2 Ф1(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-12)",
    "stavka2": "1/2 Ф2(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-12,5)",
    "stavka2": "1/2 Ф1(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-12,5)",
    "stavka2": "1/2 Ф1(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-12,5)",
    "stavka2": "1/2 Ф2(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-13)",
    "stavka2": "1/2 Ф1(13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-13)",
    "stavka2": "1/2 Ф1(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-13,5)",
    "stavka2": "1/2 Ф1(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-13,5)",
    "stavka2": "1/2 Ф1(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-13,5)",
    "stavka2": "1/2 Ф2(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-13,5)",
    "stavka2": "1/2 Ф2(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-14)",
    "stavka2": "1/2 Ф1(14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-14)",
    "stavka2": "1/2 Ф1(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-14)",
    "stavka2": "1/2 Ф1(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-14)",
    "stavka2": "1/2 Ф2(14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-14,5)",
    "stavka2": "1/2 Ф1(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-14,5)",
    "stavka2": "1/2 Ф1(15)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-14,5)",
    "stavka2": "1/2 Ф2(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-15)",
    "stavka2": "1/2 Ф1(15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-15)",
    "stavka2": "1/2 Ф2(15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-15,5)",
    "stavka2": "1/2 Ф1(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-15,5)",
    "stavka2": "1/2 Ф1(16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-15,5)",
    "stavka2": "1/2 Ф2(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-16)",
    "stavka2": "1/2 Ф1(16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-16,5)",
    "stavka2": "1/2 Ф1(16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-16,5)",
    "stavka2": "1/2 Ф1(17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-16,5)",
    "stavka2": "1/2 Ф2(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-17)",
    "stavka2": "1/2 Ф1(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-17)",
    "stavka2": "1/2 Ф2(17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-17)",
    "stavka2": "1/2 Ф2(18)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-17,5)",
    "stavka2": "1/2 Ф1(17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-18)",
    "stavka2": "1/2 Ф1(18)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-18)",
    "stavka2": "1/2 Ф2(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-18,5)",
    "stavka2": "1/2 Ф1(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-19)",
    "stavka2": "1/2 Ф2(21)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-19,5)",
    "stavka2": "1/2 Ф1(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-2)",
    "stavka2": "1/2 Ф1(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-2)",
    "stavka2": "1/2 Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-2)",
    "stavka2": "1/2 Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-2)",
    "stavka2": "1/2 Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-2,25)",
    "stavka2": "1/2 Ф2(2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-2,5)",
    "stavka2": "1/2 Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-2,5)",
    "stavka2": "1/2 Ф1(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-2,5)",
    "stavka2": "1/2 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-2,5)",
    "stavka2": "1/2 Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-2,5)",
    "stavka2": "1/2 Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-2,5)",
    "stavka2": "1/2 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-2,5)",
    "stavka2": "1/2 Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-20)",
    "stavka2": "1/2 Ф2(20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-21,5)",
    "stavka2": "1/2 Ф2(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-22)",
    "stavka2": "1/2 Ф2(22)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-22,5)",
    "stavka2": "1/2 Ф2(22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-27,5)",
    "stavka2": "1/2 Ф1(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-3)",
    "stavka2": "1/2 Ф1(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-3)",
    "stavka2": "1/2 Ф1(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-3)",
    "stavka2": "1/2 Ф1(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-3)",
    "stavka2": "1/2 Ф1(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-3)",
    "stavka2": "1/2 Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-3)",
    "stavka2": "1/2 Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-3)",
    "stavka2": "1/2 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-3,5)",
    "stavka2": "1/2 Ф1(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-3,5)",
    "stavka2": "1/2 Ф1(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-3,5)",
    "stavka2": "1/2 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-3,5)",
    "stavka2": "1/2 Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-4)",
    "stavka2": "1/2 Ф1(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-4)",
    "stavka2": "1/2 Ф1(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-4)",
    "stavka2": "1/2 Ф2(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-4)",
    "stavka2": "1/2 Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-4,5)",
    "stavka2": "1/2 Ф1(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-4,5)",
    "stavka2": "1/2 Ф1(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-4,5)",
    "stavka2": "1/2 Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-4,5)",
    "stavka2": "1/2 Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-4,5)",
    "stavka2": "1/2 Ф2(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-4,5)",
    "stavka2": "1/2 Ф2(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-5)",
    "stavka2": "1/2 Ф1(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-5)",
    "stavka2": "1/2 Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-5)",
    "stavka2": "1/2 Ф2(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-5)",
    "stavka2": "1/2 Ф2(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-5,5)",
    "stavka2": "1/2 Ф1(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-5,5)",
    "stavka2": "1/2 Ф1(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-5,5)",
    "stavka2": "1/2 Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-5,5)",
    "stavka2": "1/2 Ф2(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-6)",
    "stavka2": "1/2 Ф1(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-6)",
    "stavka2": "1/2 Ф1(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-6)",
    "stavka2": "1/2 Ф2(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-6,5)",
    "stavka2": "1/2 Ф1(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-6,5)",
    "stavka2": "1/2 Ф1(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-6,5)",
    "stavka2": "1/2 Ф1(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-6,5)",
    "stavka2": "1/2 Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-6,5)",
    "stavka2": "1/2 Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-7)",
    "stavka2": "1/2 Ф1(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-7)",
    "stavka2": "1/2 Ф1(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-7)",
    "stavka2": "1/2 Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-7)",
    "stavka2": "1/2 Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-7,5)",
    "stavka2": "1/2 Ф1(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-7,5)",
    "stavka2": "1/2 Ф1(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-7,5)",
    "stavka2": "1/2 Ф2(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-8)",
    "stavka2": "1/2 Ф1(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-8)",
    "stavka2": "1/2 Ф1(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-8)",
    "stavka2": "1/2 Ф2(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-8)",
    "stavka2": "1/2 Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-8)",
    "stavka2": "1/2 Ф2(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-8,5)",
    "stavka2": "1/2 Ф1(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-8,5)",
    "stavka2": "1/2 Ф2(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-8,5)",
    "stavka2": "1/2 Ф2(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-9)",
    "stavka2": "1/2 Ф1(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-9)",
    "stavka2": "1/2 Ф1(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-9)",
    "stavka2": "1/2 Ф1(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-9)",
    "stavka2": "1/2 Ф2(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-9,5)",
    "stavka2": "1/2 Ф1(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(-9,5)",
    "stavka2": "1/2 Ф1(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(-9,5)",
    "stavka2": "1/2 Ф2(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0)",
    "stavka2": "1/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0)",
    "stavka2": "1/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0)",
    "stavka2": "1/2 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0)",
    "stavka2": "1/2 Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(0)",
    "stavka2": "1/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,25)",
    "stavka2": "1/2 Ф1(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,25)",
    "stavka2": "1/2 Ф2(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,5)",
    "stavka2": "1/2 Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(0,75)",
    "stavka2": "1/2 Ф1(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(0,75)",
    "stavka2": "1/2 Ф2(-0,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(0,75)",
    "stavka2": "1/2 Ф2(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1)",
    "stavka2": "1/2 Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1,25)",
    "stavka2": "1/2 Ф2(-1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф2(0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(1,5)",
    "stavka2": "1/2 Ф2(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(10)",
    "stavka2": "1/2 Ф1(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(10)",
    "stavka2": "1/2 Ф2(-10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(10,5)",
    "stavka2": "1/2 Ф1(-10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(10,5)",
    "stavka2": "1/2 Ф2(-10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(11)",
    "stavka2": "1/2 Ф1(-11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(11)",
    "stavka2": "1/2 Ф2(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(11,5)",
    "stavka2": "1/2 Ф1(-11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(11,5)",
    "stavka2": "1/2 Ф2(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(12)",
    "stavka2": "1/2 Ф2(-12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(12,5)",
    "stavka2": "1/2 Ф1(-12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(12,5)",
    "stavka2": "1/2 Ф2(-12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(13)",
    "stavka2": "1/2 Ф1(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(13)",
    "stavka2": "1/2 Ф1(-13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(13)",
    "stavka2": "1/2 Ф2(-12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(13)",
    "stavka2": "1/2 Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(13)",
    "stavka2": "1/2 Ф2(-13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(14)",
    "stavka2": "1/2 Ф1(-14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(14,5)",
    "stavka2": "1/2 Ф1(-14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(14,5)",
    "stavka2": "1/2 Ф2(-14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(15)",
    "stavka2": "1/2 Ф1(-15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(15,5)",
    "stavka2": "1/2 Ф2(-14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(16)",
    "stavka2": "1/2 Ф2(-15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(16,5)",
    "stavka2": "1/2 Ф1(-16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(16,5)",
    "stavka2": "1/2 Ф2(-16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(17,5)",
    "stavka2": "1/2 Ф2(-17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(18,5)",
    "stavka2": "1/2 Ф2(-17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(2)",
    "stavka2": "1/2 Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(2)",
    "stavka2": "1/2 Ф1(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(2)",
    "stavka2": "1/2 Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(2)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(2)",
    "stavka2": "1/2 Ф2(0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(2,5)",
    "stavka2": "1/2 Ф1(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(2,5)",
    "stavka2": "1/2 Ф1(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(2,5)",
    "stavka2": "1/2 Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(2,5)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(2,5)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(2,5)",
    "stavka2": "1/2 Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(21,5)",
    "stavka2": "1/2 Ф2(-21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(25,5)",
    "stavka2": "1/2 Ф2(-25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(26)",
    "stavka2": "1/2 Ф2(-23)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3)",
    "stavka2": "1/2 2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3)",
    "stavka2": "1/2 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3)",
    "stavka2": "1/2 Ф1(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3)",
    "stavka2": "1/2 Ф1(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(3)",
    "stavka2": "1/2 Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3)",
    "stavka2": "1/2 Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(3,5)",
    "stavka2": "1/2 Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3,5)",
    "stavka2": "1/2 Ф1(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3,5)",
    "stavka2": "1/2 Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(3,5)",
    "stavka2": "1/2 Ф1(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3,5)",
    "stavka2": "1/2 Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(3,5)",
    "stavka2": "1/2 Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3,5)",
    "stavka2": "1/2 Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(3,5)",
    "stavka2": "1/2 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(4)",
    "stavka2": "1/2 Ф1(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(4)",
    "stavka2": "1/2 Ф1(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(4)",
    "stavka2": "1/2 Ф1(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(4)",
    "stavka2": "1/2 Ф2(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(4)",
    "stavka2": "1/2 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(4)",
    "stavka2": "1/2 Ф2(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(4,5)",
    "stavka2": "1/2 Ф1(-4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(4,5)",
    "stavka2": "1/2 Ф1(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(4,5)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(4,5)",
    "stavka2": "1/2 Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(5)",
    "stavka2": "1/2 Ф1(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(5)",
    "stavka2": "1/2 Ф1(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(5)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(5)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(5)",
    "stavka2": "1/2 Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(5)",
    "stavka2": "1/2 Ф2(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(5,5)",
    "stavka2": "1/2 Ф1(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(5,5)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(5,5)",
    "stavka2": "1/2 Ф2(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(6,5)",
    "stavka2": "1/2 Ф1(-6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(6,5)",
    "stavka2": "1/2 Ф1(-6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(6,5)",
    "stavka2": "1/2 Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(6,5)",
    "stavka2": "1/2 Ф2(-6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(6,5)",
    "stavka2": "1/2 Ф2(-6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(7)",
    "stavka2": "1/2 Ф1(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(7)",
    "stavka2": "1/2 Ф1(-7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(7)",
    "stavka2": "1/2 Ф2(-7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(7,5)",
    "stavka2": "1/2 Ф1(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(7,5)",
    "stavka2": "1/2 Ф2(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(8)",
    "stavka2": "1/2 Ф1(-8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(8)",
    "stavka2": "1/2 Ф2(-7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(8)",
    "stavka2": "1/2 Ф2(-8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(8,5)",
    "stavka2": "1/2 Ф1(-8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(8,5)",
    "stavka2": "1/2 Ф1(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(8,5)",
    "stavka2": "1/2 Ф2(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(9)",
    "stavka2": "1/2 Ф1(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(9)",
    "stavka2": "1/2 Ф1(-9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(9)",
    "stavka2": "1/2 Ф2(-9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(9,5)",
    "stavka2": "1/2 Ф1(-9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф1(9,5)",
    "stavka2": "1/2 Ф1(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф1(9,5)",
    "stavka2": "1/2 Ф2(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-0,25)",
    "stavka2": "1/2 Ф1(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-0,25)",
    "stavka2": "1/2 Ф2(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-0,5)",
    "stavka2": "1/2 Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-0,5)",
    "stavka2": "1/2 Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-0,5)",
    "stavka2": "1/2 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-0,5)",
    "stavka2": "1/2 Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-0,75)",
    "stavka2": "1/2 Ф2(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-1)",
    "stavka2": "1/2 Ф1(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-1)",
    "stavka2": "1/2 Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-1)",
    "stavka2": "1/2 Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-1)",
    "stavka2": "1/2 Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-1,5)",
    "stavka2": "1/2 Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-1,5)",
    "stavka2": "1/2 Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-1,5)",
    "stavka2": "1/2 Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-1,5)",
    "stavka2": "1/2 Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-1,5)",
    "stavka2": "1/2 Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-10)",
    "stavka2": "1/2 Ф2(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-10,5)",
    "stavka2": "1/2 Ф2(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-11)",
    "stavka2": "1/2 Ф1(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-11)",
    "stavka2": "1/2 Ф2(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-11,5)",
    "stavka2": "1/2 Ф2(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-12)",
    "stavka2": "1/2 Ф2(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-13)",
    "stavka2": "1/2 Ф1(13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-13)",
    "stavka2": "1/2 Ф2(13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-13,5)",
    "stavka2": "1/2 Ф2(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-15,5)",
    "stavka2": "1/2 Ф2(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-16)",
    "stavka2": "1/2 Ф2(16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-16,5)",
    "stavka2": "1/2 Ф2(16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-2)",
    "stavka2": "1/2 Ф1(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-2)",
    "stavka2": "1/2 Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-2)",
    "stavka2": "1/2 Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-2,5)",
    "stavka2": "1/2 Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-2,5)",
    "stavka2": "1/2 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-2,5)",
    "stavka2": "1/2 Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-20)",
    "stavka2": "1/2 Ф2(20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-20,5)",
    "stavka2": "1/2 Ф2(20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-21,5)",
    "stavka2": "1/2 Ф2(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-25,5)",
    "stavka2": "1/2 Ф2(25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-3)",
    "stavka2": "1/2 Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-3)",
    "stavka2": "1/2 Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-3)",
    "stavka2": "1/2 Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-3,5)",
    "stavka2": "1/2 Ф1(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-3,5)",
    "stavka2": "1/2 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-3,5)",
    "stavka2": "1/2 Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-4)",
    "stavka2": "1/2 Ф1(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-4)",
    "stavka2": "1/2 Ф2(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-4)",
    "stavka2": "1/2 Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-4,5)",
    "stavka2": "1/2 Ф1(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-4,5)",
    "stavka2": "1/2 Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-5)",
    "stavka2": "1/2 Ф1(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-5)",
    "stavka2": "1/2 Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-5)",
    "stavka2": "1/2 Ф2(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-5)",
    "stavka2": "1/2 Ф2(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-5)",
    "stavka2": "1/2 Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-5)",
    "stavka2": "1/2 Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-5,5)",
    "stavka2": "1/2 Ф1(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-5,5)",
    "stavka2": "1/2 Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-5,5)",
    "stavka2": "1/2 Ф2(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-6)",
    "stavka2": "1/2 Ф1(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-6)",
    "stavka2": "1/2 Ф2(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-6)",
    "stavka2": "1/2 Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-6,5)",
    "stavka2": "1/2 Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-6,5)",
    "stavka2": "1/2 Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-6,5)",
    "stavka2": "1/2 Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-7)",
    "stavka2": "1/2 Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-7)",
    "stavka2": "1/2 Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-7,5)",
    "stavka2": "1/2 Ф2(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-7,5)",
    "stavka2": "1/2 Ф2(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-7,5)",
    "stavka2": "1/2 Ф2(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-8)",
    "stavka2": "1/2 Ф2(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-8)",
    "stavka2": "1/2 Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-8,5)",
    "stavka2": "1/2 Ф2(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-8,5)",
    "stavka2": "1/2 Ф2(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-8,5)",
    "stavka2": "1/2 Ф2(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-9)",
    "stavka2": "1/2 Ф2(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(-9)",
    "stavka2": "1/2 Ф2(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(-9,5)",
    "stavka2": "1/2 Ф2(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0)",
    "stavka2": "1/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0)",
    "stavka2": "1/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0)",
    "stavka2": "1/2 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0)",
    "stavka2": "1/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0)",
    "stavka2": "1/2 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(0,25)",
    "stavka2": "1/2 Ф1(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0,25)",
    "stavka2": "1/2 Ф2(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0,5)",
    "stavka2": "1/2 1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0,5)",
    "stavka2": "1/2 2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0,5)",
    "stavka2": "1/2 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(0,5)",
    "stavka2": "1/2 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0,5)",
    "stavka2": "1/2 Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(0,5)",
    "stavka2": "1/2 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(0,5)",
    "stavka2": "1/2 Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(0,75)",
    "stavka2": "1/2 Ф1(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(1)",
    "stavka2": "1/2 Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(1)",
    "stavka2": "1/2 Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(1)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(1,5)",
    "stavka2": "1/2 Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(1,5)",
    "stavka2": "1/2 Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(1,5)",
    "stavka2": "1/2 Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(1,5)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(1,5)",
    "stavka2": "1/2 Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(12)",
    "stavka2": "1/2 Ф2(-12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(12,5)",
    "stavka2": "1/2 Ф2(-12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(12,5)",
    "stavka2": "1/2 Ф2(-12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(13)",
    "stavka2": "1/2 Ф1(-13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(13)",
    "stavka2": "1/2 Ф2(-12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(13,5)",
    "stavka2": "1/2 Ф2(-13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(15,5)",
    "stavka2": "1/2 Ф1(-15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(15,5)",
    "stavka2": "1/2 Ф2(-15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(16,5)",
    "stavka2": "1/2 Ф1(-16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(17)",
    "stavka2": "1/2 Ф2(-16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(17,5)",
    "stavka2": "1/2 Ф2(-17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(2)",
    "stavka2": "1/2 П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(2)",
    "stavka2": "1/2 Ф1(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(2)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(2,5)",
    "stavka2": "1/2 Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(2,5)",
    "stavka2": "1/2 Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(2,5)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(2,5)",
    "stavka2": "1/2 Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(3)",
    "stavka2": "1/2 Ф1(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(3)",
    "stavka2": "1/2 Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(3)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(3)",
    "stavka2": "1/2 Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(3)",
    "stavka2": "1/2 Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(3)",
    "stavka2": "1/2 Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(3,5)",
    "stavka2": "1/2 П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(3,5)",
    "stavka2": "1/2 Ф1(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(3,5)",
    "stavka2": "1/2 Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(3,5)",
    "stavka2": "1/2 Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(3,5)",
    "stavka2": "1/2 Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(3,5)",
    "stavka2": "1/2 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(4)",
    "stavka2": "1/2 Ф1(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(4)",
    "stavka2": "1/2 Ф2(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(4)",
    "stavka2": "1/2 Ф2(-4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(4)",
    "stavka2": "1/2 Ф2(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(4,5)",
    "stavka2": "1/2 Ф1(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(4,5)",
    "stavka2": "1/2 Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(4,5)",
    "stavka2": "1/2 Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(5)",
    "stavka2": "1/2 Ф1(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(5)",
    "stavka2": "1/2 Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(5)",
    "stavka2": "1/2 Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(5)",
    "stavka2": "1/2 Ф2(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(5,5)",
    "stavka2": "1/2 Ф1(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(5,5)",
    "stavka2": "1/2 Ф1(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(5,5)",
    "stavka2": "1/2 Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(5,5)",
    "stavka2": "1/2 Ф1(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(5,5)",
    "stavka2": "1/2 Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(5,5)",
    "stavka2": "1/2 Ф2(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(6)",
    "stavka2": "1/2 Ф2(-6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(6,5)",
    "stavka2": "1/2 Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(6,5)",
    "stavka2": "1/2 Ф2(-6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(7)",
    "stavka2": "1/2 Ф2(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(7)",
    "stavka2": "1/2 Ф2(-7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(7,5)",
    "stavka2": "1/2 Ф1(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(7,5)",
    "stavka2": "1/2 Ф2(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(8)",
    "stavka2": "1/2 Ф2(-8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(8,5)",
    "stavka2": "1/2 Ф1(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(8,5)",
    "stavka2": "1/2 Ф2(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(9)",
    "stavka2": "1/2 Ф1(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(9)",
    "stavka2": "1/2 Ф1(-9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(9)",
    "stavka2": "1/2 Ф2(-9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Ф2(9,5)",
    "stavka2": "1/2 Ф2(-9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/2 Ф2(9,5)",
    "stavka2": "1/2 Ф2(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/2 Чёт",
    "stavka2": "1/2 Нечёт",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 П1",
    "stavka2": "1/3 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 П1",
    "stavka2": "1/3 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/3 П1",
    "stavka2": "1/3 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 П2",
    "stavka2": "1/3 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 П2",
    "stavka2": "1/3 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 ТМ(0,5)",
    "stavka2": "1/3 ТБ(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 ТМ(1,5)",
    "stavka2": "1/3 ТБ(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 ТМ(2)",
    "stavka2": "1/3 ТБ(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(-0,5)",
    "stavka2": "1/3 Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(-0,5)",
    "stavka2": "1/3 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(-1)",
    "stavka2": "1/3 Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(-2)",
    "stavka2": "1/3 Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(-2,5)",
    "stavka2": "1/3 Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/3 Ф1(-2,5)",
    "stavka2": "1/3 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(-3,5)",
    "stavka2": "1/3 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(-4,5)",
    "stavka2": "1/3 Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(-5,5)",
    "stavka2": "1/3 Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(0)",
    "stavka2": "1/3 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(0,5)",
    "stavka2": "1/3 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(1,5)",
    "stavka2": "1/3 Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(2,5)",
    "stavka2": "1/3 Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(3,5)",
    "stavka2": "1/3 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(4,5)",
    "stavka2": "1/3 Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф1(5,5)",
    "stavka2": "1/3 Ф2(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф2(-0,5)",
    "stavka2": "1/3 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф2(-1,5)",
    "stavka2": "1/3 Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф2(-2,5)",
    "stavka2": "1/3 Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф2(-3,5)",
    "stavka2": "1/3 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф2(0,5)",
    "stavka2": "1/3 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/3 Ф2(2,5)",
    "stavka2": "1/3 Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П1",
    "stavka2": "1/4 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П1",
    "stavka2": "1/4 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П1",
    "stavka2": "1/4 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П1",
    "stavka2": "1/4 Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 П1",
    "stavka2": "1/4 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П1",
    "stavka2": "1/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 П2",
    "stavka2": "1/4 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П2",
    "stavka2": "1/4 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П2",
    "stavka2": "1/4 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П2",
    "stavka2": "1/4 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 П2",
    "stavka2": "1/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(10)",
    "stavka2": "1/4 ТБ(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(10,5)",
    "stavka2": "1/4 ТБ(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(11,5)",
    "stavka2": "1/4 ТБ(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(12,5)",
    "stavka2": "1/4 ТБ(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(13)",
    "stavka2": "1/4 ТБ(13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(13,5)",
    "stavka2": "1/4 ТБ(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(14)",
    "stavka2": "1/4 ТБ(14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(14,5)",
    "stavka2": "1/4 ТБ(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(15,5)",
    "stavka2": "1/4 ТБ(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(17)",
    "stavka2": "1/4 ТБ(17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(30)",
    "stavka2": "1/4 ТБ(30)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(30,5)",
    "stavka2": "1/4 ТБ(30,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(31,5)",
    "stavka2": "1/4 ТБ(31,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(32)",
    "stavka2": "1/4 ТБ(32)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(32,5)",
    "stavka2": "1/4 ТБ(32)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(32,5)",
    "stavka2": "1/4 ТБ(32,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(33)",
    "stavka2": "1/4 ТБ(33)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(33,5)",
    "stavka2": "1/4 ТБ(33,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(34)",
    "stavka2": "1/4 ТБ(34)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(34,5)",
    "stavka2": "1/4 ТБ(34)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(34,5)",
    "stavka2": "1/4 ТБ(34,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(35)",
    "stavka2": "1/4 ТБ(34,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(35)",
    "stavka2": "1/4 ТБ(35)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(35,5)",
    "stavka2": "1/4 ТБ(35,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(36)",
    "stavka2": "1/4 ТБ(35,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(36)",
    "stavka2": "1/4 ТБ(36)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(36,5)",
    "stavka2": "1/4 ТБ(36)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(36,5)",
    "stavka2": "1/4 ТБ(36,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(37)",
    "stavka2": "1/4 ТБ(36,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(37)",
    "stavka2": "1/4 ТБ(37)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(37,5)",
    "stavka2": "1/4 ТБ(36,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(37,5)",
    "stavka2": "1/4 ТБ(37)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(37,5)",
    "stavka2": "1/4 ТБ(37,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(38)",
    "stavka2": "1/4 ТБ(37,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(38)",
    "stavka2": "1/4 ТБ(38)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(38,5)",
    "stavka2": "1/4 ТБ(38)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(38,5)",
    "stavka2": "1/4 ТБ(38,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(39)",
    "stavka2": "1/4 ТБ(38)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(39)",
    "stavka2": "1/4 ТБ(38,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(39)",
    "stavka2": "1/4 ТБ(39)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(39,5)",
    "stavka2": "1/4 ТБ(39)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(39,5)",
    "stavka2": "1/4 ТБ(39,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(40)",
    "stavka2": "1/4 ТБ(39,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(40)",
    "stavka2": "1/4 ТБ(39,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(40)",
    "stavka2": "1/4 ТБ(40)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(40)",
    "stavka2": "1/4 ТБ(40)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(40,5)",
    "stavka2": "1/4 ТБ(39,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(40,5)",
    "stavka2": "1/4 ТБ(40)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(40,5)",
    "stavka2": "1/4 ТБ(40,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(41)",
    "stavka2": "1/4 ТБ(40,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(41)",
    "stavka2": "1/4 ТБ(40,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(41)",
    "stavka2": "1/4 ТБ(41)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(41,5)",
    "stavka2": "1/4 ТБ(40,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(41,5)",
    "stavka2": "1/4 ТБ(41)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(41,5)",
    "stavka2": "1/4 ТБ(41)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(41,5)",
    "stavka2": "1/4 ТБ(41,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(42)",
    "stavka2": "1/4 ТБ(41,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(42)",
    "stavka2": "1/4 ТБ(42)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(42,5)",
    "stavka2": "1/4 ТБ(41,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(42,5)",
    "stavka2": "1/4 ТБ(42)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(42,5)",
    "stavka2": "1/4 ТБ(42)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(42,5)",
    "stavka2": "1/4 ТБ(42,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(43)",
    "stavka2": "1/4 ТБ(42)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(43)",
    "stavka2": "1/4 ТБ(42,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(43)",
    "stavka2": "1/4 ТБ(43)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(43,5)",
    "stavka2": "1/4 ТБ(42,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(43,5)",
    "stavka2": "1/4 ТБ(43)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(43,5)",
    "stavka2": "1/4 ТБ(43,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(44)",
    "stavka2": "1/4 ТБ(43)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(44)",
    "stavka2": "1/4 ТБ(43,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(44)",
    "stavka2": "1/4 ТБ(44)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(44,5)",
    "stavka2": "1/4 ТБ(43,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(44,5)",
    "stavka2": "1/4 ТБ(44)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(44,5)",
    "stavka2": "1/4 ТБ(44)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(44,5)",
    "stavka2": "1/4 ТБ(44,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(45)",
    "stavka2": "1/4 ТБ(44,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(45)",
    "stavka2": "1/4 ТБ(45)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(45,5)",
    "stavka2": "1/4 ТБ(44)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(45,5)",
    "stavka2": "1/4 ТБ(45)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(45,5)",
    "stavka2": "1/4 ТБ(45,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(46)",
    "stavka2": "1/4 ТБ(45)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(46)",
    "stavka2": "1/4 ТБ(45,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(46)",
    "stavka2": "1/4 ТБ(46)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(46,5)",
    "stavka2": "1/4 ТБ(46)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(46,5)",
    "stavka2": "1/4 ТБ(46,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(47)",
    "stavka2": "1/4 ТБ(46,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(47)",
    "stavka2": "1/4 ТБ(47)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(47,5)",
    "stavka2": "1/4 ТБ(46,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(47,5)",
    "stavka2": "1/4 ТБ(47)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(47,5)",
    "stavka2": "1/4 ТБ(47,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(48)",
    "stavka2": "1/4 ТБ(47,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(48)",
    "stavka2": "1/4 ТБ(48)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(48,5)",
    "stavka2": "1/4 ТБ(47,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(48,5)",
    "stavka2": "1/4 ТБ(48)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(48,5)",
    "stavka2": "1/4 ТБ(48,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(49)",
    "stavka2": "1/4 ТБ(47,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(49)",
    "stavka2": "1/4 ТБ(48)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(49)",
    "stavka2": "1/4 ТБ(48,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(49)",
    "stavka2": "1/4 ТБ(49)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(49,5)",
    "stavka2": "1/4 ТБ(48)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(49,5)",
    "stavka2": "1/4 ТБ(48,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(49,5)",
    "stavka2": "1/4 ТБ(49,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(50)",
    "stavka2": "1/4 ТБ(49)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(50)",
    "stavka2": "1/4 ТБ(50)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(50,5)",
    "stavka2": "1/4 ТБ(49,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(50,5)",
    "stavka2": "1/4 ТБ(50)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(50,5)",
    "stavka2": "1/4 ТБ(50,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(51)",
    "stavka2": "1/4 ТБ(51)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(51,5)",
    "stavka2": "1/4 ТБ(51,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(52)",
    "stavka2": "1/4 ТБ(51,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(52)",
    "stavka2": "1/4 ТБ(52)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(52,5)",
    "stavka2": "1/4 ТБ(51)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(52,5)",
    "stavka2": "1/4 ТБ(51,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(52,5)",
    "stavka2": "1/4 ТБ(52)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(52,5)",
    "stavka2": "1/4 ТБ(52,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(53)",
    "stavka2": "1/4 ТБ(52,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(53)",
    "stavka2": "1/4 ТБ(53)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(53,5)",
    "stavka2": "1/4 ТБ(53)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(53,5)",
    "stavka2": "1/4 ТБ(53,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(54)",
    "stavka2": "1/4 ТБ(52,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(54)",
    "stavka2": "1/4 ТБ(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(54)",
    "stavka2": "1/4 ТБ(54)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(54,5)",
    "stavka2": "1/4 ТБ(54)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(54,5)",
    "stavka2": "1/4 ТБ(54,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(55)",
    "stavka2": "1/4 ТБ(54,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(55)",
    "stavka2": "1/4 ТБ(55)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(55,5)",
    "stavka2": "1/4 ТБ(55)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(55,5)",
    "stavka2": "1/4 ТБ(55,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(56)",
    "stavka2": "1/4 ТБ(55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(56)",
    "stavka2": "1/4 ТБ(56)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(56,5)",
    "stavka2": "1/4 ТБ(56)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(56,5)",
    "stavka2": "1/4 ТБ(56,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(57)",
    "stavka2": "1/4 ТБ(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(57)",
    "stavka2": "1/4 ТБ(57)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(57,5)",
    "stavka2": "1/4 ТБ(55)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(57,5)",
    "stavka2": "1/4 ТБ(56)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(57,5)",
    "stavka2": "1/4 ТБ(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(57,5)",
    "stavka2": "1/4 ТБ(57)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(57,5)",
    "stavka2": "1/4 ТБ(57,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(58)",
    "stavka2": "1/4 ТБ(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(58)",
    "stavka2": "1/4 ТБ(56,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(58)",
    "stavka2": "1/4 ТБ(57,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(58)",
    "stavka2": "1/4 ТБ(58)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(58,5)",
    "stavka2": "1/4 ТБ(58)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(58,5)",
    "stavka2": "1/4 ТБ(58,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(59)",
    "stavka2": "1/4 ТБ(58)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(59)",
    "stavka2": "1/4 ТБ(58,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(59)",
    "stavka2": "1/4 ТБ(59)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(59,5)",
    "stavka2": "1/4 ТБ(58)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(59,5)",
    "stavka2": "1/4 ТБ(59)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(59,5)",
    "stavka2": "1/4 ТБ(59,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(60)",
    "stavka2": "1/4 ТБ(59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(60)",
    "stavka2": "1/4 ТБ(60)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(60,5)",
    "stavka2": "1/4 ТБ(60)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(60,5)",
    "stavka2": "1/4 ТБ(60,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(61)",
    "stavka2": "1/4 ТБ(61)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(61,5)",
    "stavka2": "1/4 ТБ(61,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(62)",
    "stavka2": "1/4 ТБ(61,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(62,5)",
    "stavka2": "1/4 ТБ(62,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(63)",
    "stavka2": "1/4 ТБ(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(63)",
    "stavka2": "1/4 ТБ(63)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(7)",
    "stavka2": "1/4 ТБ(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(7,5)",
    "stavka2": "1/4 ТБ(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(80,5)",
    "stavka2": "1/4 ТБ(80,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(86,5)",
    "stavka2": "1/4 ТБ(86,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 ТМ(9,5)",
    "stavka2": "1/4 ТБ(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(9,5)",
    "stavka2": "1/4 ТБ(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 ТМ(9,5)",
    "stavka2": "1/4 ТБ(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-0,5)",
    "stavka2": "1/4 Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-0,5)",
    "stavka2": "1/4 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-0,5)",
    "stavka2": "1/4 Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(-1)",
    "stavka2": "1/4 Ф1(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-1)",
    "stavka2": "1/4 Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-1)",
    "stavka2": "1/4 Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(-1,5)",
    "stavka2": "1/4 Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-1,5)",
    "stavka2": "1/4 Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-1,5)",
    "stavka2": "1/4 Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(-10)",
    "stavka2": "1/4 Ф2(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-10,5)",
    "stavka2": "1/4 Ф2(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-11,5)",
    "stavka2": "1/4 Ф2(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-2)",
    "stavka2": "1/4 Ф1(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-2)",
    "stavka2": "1/4 Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-2)",
    "stavka2": "1/4 Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(-2,5)",
    "stavka2": "1/4 Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-2,5)",
    "stavka2": "1/4 Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(-2,5)",
    "stavka2": "1/4 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-3)",
    "stavka2": "1/4 Ф1(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-3)",
    "stavka2": "1/4 Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-3,5)",
    "stavka2": "1/4 Ф1(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-3,5)",
    "stavka2": "1/4 Ф1(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(-3,5)",
    "stavka2": "1/4 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-4)",
    "stavka2": "1/4 Ф1(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-4)",
    "stavka2": "1/4 Ф2(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-4,5)",
    "stavka2": "1/4 Ф1(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-4,5)",
    "stavka2": "1/4 Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-5)",
    "stavka2": "1/4 Ф2(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-5,5)",
    "stavka2": "1/4 Ф1(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-5,5)",
    "stavka2": "1/4 Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-6)",
    "stavka2": "1/4 Ф2(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-6,5)",
    "stavka2": "1/4 Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-7)",
    "stavka2": "1/4 Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(-7,5)",
    "stavka2": "1/4 Ф2(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(0)",
    "stavka2": "1/4 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(0)",
    "stavka2": "1/4 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(0)",
    "stavka2": "1/4 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(0)",
    "stavka2": "1/4 Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(0)",
    "stavka2": "1/4 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(0)",
    "stavka2": "1/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(0)",
    "stavka2": "1/4 Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(0,5)",
    "stavka2": "1/4 П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(0,5)",
    "stavka2": "1/4 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(0,5)",
    "stavka2": "1/4 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(0,5)",
    "stavka2": "1/4 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(0,5)",
    "stavka2": "1/4 Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(1)",
    "stavka2": "1/4 Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(1)",
    "stavka2": "1/4 Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(1)",
    "stavka2": "1/4 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(1)",
    "stavka2": "1/4 Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(1,5)",
    "stavka2": "1/4 Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(1,5)",
    "stavka2": "1/4 Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(1,5)",
    "stavka2": "1/4 Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(1,5)",
    "stavka2": "1/4 Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(10,5)",
    "stavka2": "1/4 Ф2(-10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(2)",
    "stavka2": "1/4 Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(2)",
    "stavka2": "1/4 Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(2)",
    "stavka2": "1/4 Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(2)",
    "stavka2": "1/4 Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(2,5)",
    "stavka2": "1/4 Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(2,5)",
    "stavka2": "1/4 Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(2,5)",
    "stavka2": "1/4 Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(3)",
    "stavka2": "1/4 Ф1(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(3)",
    "stavka2": "1/4 Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(3,5)",
    "stavka2": "1/4 Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(3,5)",
    "stavka2": "1/4 Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(3,5)",
    "stavka2": "1/4 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(4)",
    "stavka2": "1/4 Ф2(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(4,5)",
    "stavka2": "1/4 Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(5)",
    "stavka2": "1/4 Ф2(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(5,5)",
    "stavka2": "1/4 Ф1(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(5,5)",
    "stavka2": "1/4 Ф2(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(6)",
    "stavka2": "1/4 Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф1(6,5)",
    "stavka2": "1/4 Ф2(-6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(7,5)",
    "stavka2": "1/4 Ф2(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(8)",
    "stavka2": "1/4 Ф2(-8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(8,5)",
    "stavka2": "1/4 Ф2(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф1(9,5)",
    "stavka2": "1/4 Ф2(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-0,5)",
    "stavka2": "1/4 Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-0,5)",
    "stavka2": "1/4 Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(-0,5)",
    "stavka2": "1/4 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-0,5)",
    "stavka2": "1/4 Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(-0,5)",
    "stavka2": "1/4 Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(-1)",
    "stavka2": "1/4 Ф1(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-1)",
    "stavka2": "1/4 Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-1,5)",
    "stavka2": "1/4 Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-1,5)",
    "stavka2": "1/4 Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(-2)",
    "stavka2": "1/4 Ф1(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-2,5)",
    "stavka2": "1/4 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-3)",
    "stavka2": "1/4 Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-3,5)",
    "stavka2": "1/4 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-4)",
    "stavka2": "1/4 Ф2(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-4,5)",
    "stavka2": "1/4 Ф1(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-4,5)",
    "stavka2": "1/4 Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-5,5)",
    "stavka2": "1/4 Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-6)",
    "stavka2": "1/4 Ф2(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-6,5)",
    "stavka2": "1/4 Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(-7)",
    "stavka2": "1/4 Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(0)",
    "stavka2": "1/4 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(0)",
    "stavka2": "1/4 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(0)",
    "stavka2": "1/4 Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(0)",
    "stavka2": "1/4 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(0)",
    "stavka2": "1/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(0,5)",
    "stavka2": "1/4 П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(0,5)",
    "stavka2": "1/4 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(0,5)",
    "stavka2": "1/4 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(1)",
    "stavka2": "1/4 Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(1)",
    "stavka2": "1/4 Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(1,5)",
    "stavka2": "1/4 Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(1,5)",
    "stavka2": "1/4 Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(1,5)",
    "stavka2": "1/4 Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(1,5)",
    "stavka2": "1/4 Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(2)",
    "stavka2": "1/4 Ф1(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(2)",
    "stavka2": "1/4 Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(2,5)",
    "stavka2": "1/4 Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(2,5)",
    "stavka2": "1/4 Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "1/4 Ф2(2,5)",
    "stavka2": "1/4 Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(3)",
    "stavka2": "1/4 Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(3,5)",
    "stavka2": "1/4 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "1/4 Ф2(4,5)",
    "stavka2": "1/4 Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "12",
    "stavka2": "Х",
    "tip": "VILKA"
  },
  {
    "stavka1": "1Х",
    "stavka2": "1",
    "tip": "VILKA"
  },
  {
    "stavka1": "1Х",
    "stavka2": "2",
    "tip": "VILKA"
  },
  {
    "stavka1": "1Х",
    "stavka2": "Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2",
    "stavka2": "Х2",
    "tip": "VILKA"
  },
  {
    "stavka1": "2",
    "stavka2": "Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2",
    "stavka2": "Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2",
    "stavka2": "Ф2(0,75)",
    "tip": "KARIDOR"
  },

  {
    "stavka1": "2/2 П1",
    "stavka2": "2/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 П1",
    "stavka2": "2/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 П1",
    "stavka2": "2/2 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 П1",
    "stavka2": "2/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 П2",
    "stavka2": "2/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 П2",
    "stavka2": "2/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 П2",
    "stavka2": "2/2 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 П2",
    "stavka2": "2/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 ТМ(25,5)",
    "stavka2": "2/2 ТБ(25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 ТМ(26,5)",
    "stavka2": "2/2 ТБ(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 ТМ(27,5)",
    "stavka2": "2/2 ТБ(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(-2,5)",
    "stavka2": "2/2 Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(-2,5)",
    "stavka2": "2/2 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(-3,5)",
    "stavka2": "2/2 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(-4,5)",
    "stavka2": "2/2 Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(-5,5)",
    "stavka2": "2/2 Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(-6,5)",
    "stavka2": "2/2 Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(0)",
    "stavka2": "2/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(0)",
    "stavka2": "2/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(2,5)",
    "stavka2": "2/2 Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(2,5)",
    "stavka2": "2/2 Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(3,5)",
    "stavka2": "2/2 Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(3,5)",
    "stavka2": "2/2 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф1(4,5)",
    "stavka2": "2/2 Ф1(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф2(-2,5)",
    "stavka2": "2/2 Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф2(-3,5)",
    "stavka2": "2/2 Ф1(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф2(-4,5)",
    "stavka2": "2/2 Ф1(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф2(0)",
    "stavka2": "2/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф2(2,5)",
    "stavka2": "2/2 Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф2(3,5)",
    "stavka2": "2/2 Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/2 Ф2(3,5)",
    "stavka2": "2/2 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/4 П1",
    "stavka2": "2/4 Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 П1",
    "stavka2": "2/4 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/4 П1",
    "stavka2": "2/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 П2",
    "stavka2": "2/4 Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/4 ТМ(54,5)",
    "stavka2": "2/4 ТБ(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 Ф1(-0,5)",
    "stavka2": "2/4 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/4 Ф1(-2,5)",
    "stavka2": "2/4 Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 Ф1(-2,5)",
    "stavka2": "2/4 Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/4 Ф1(0)",
    "stavka2": "2/4 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 Ф1(0)",
    "stavka2": "2/4 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/4 Ф1(0)",
    "stavka2": "2/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 Ф1(0,5)",
    "stavka2": "2/4 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 Ф1(0,5)",
    "stavka2": "2/4 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/4 Ф1(0,5)",
    "stavka2": "2/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 Ф2(0)",
    "stavka2": "2/4 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "2/4 Ф2(0)",
    "stavka2": "2/4 Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "2/4 Ф2(0,5)",
    "stavka2": "2/4 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 П1",
    "stavka2": "3/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 П1",
    "stavka2": "3/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 П1",
    "stavka2": "3/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 П2",
    "stavka2": "3/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 П2",
    "stavka2": "3/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 П2",
    "stavka2": "3/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 ТМ(25,5)",
    "stavka2": "3/2 ТБ(25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 ТМ(26,5)",
    "stavka2": "3/2 ТБ(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 Ф1(-2,5)",
    "stavka2": "3/2 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 Ф1(0)",
    "stavka2": "3/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 Ф1(0)",
    "stavka2": "3/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 Ф1(2,5)",
    "stavka2": "3/2 Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 Ф1(3,5)",
    "stavka2": "3/2 Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/2 Ф2(-2,5)",
    "stavka2": "3/2 Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/4 П1",
    "stavka2": "3/4 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/4 П1",
    "stavka2": "3/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "3/4 П2",
    "stavka2": "3/4 Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "3/4 ТМ(51,5)",
    "stavka2": "3/4 ТБ(51,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "3/4 ТМ(51,5)",
    "stavka2": "3/4 ТБ(51,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/4 Ф1(-0,5)",
    "stavka2": "3/4 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/4 Ф1(-1)",
    "stavka2": "3/4 Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "3/4 Ф1(-1,5)",
    "stavka2": "3/4 Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/4 Ф1(0)",
    "stavka2": "3/4 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/4 Ф1(0)",
    "stavka2": "3/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "3/4 Ф1(0,5)",
    "stavka2": "3/4 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "3/4 Ф1(0,5)",
    "stavka2": "3/4 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "3/4 Ф1(0,5)",
    "stavka2": "3/4 Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "3/4 Ф2(0,5)",
    "stavka2": "3/4 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/2 П1",
    "stavka2": "4/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/2 П1",
    "stavka2": "4/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/2 П1",
    "stavka2": "4/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/2 П2",
    "stavka2": "4/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/2 П2",
    "stavka2": "4/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/2 Ф1(0)",
    "stavka2": "4/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/2 Ф1(0)",
    "stavka2": "4/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 П1",
    "stavka2": "4/4 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 П1",
    "stavka2": "4/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "4/4 ТМ(54,5)",
    "stavka2": "4/4 ТБ(51)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "4/4 ТМ(54,5)",
    "stavka2": "4/4 ТБ(54,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 ТМ(55)",
    "stavka2": "4/4 ТБ(51)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "4/4 ТМ(56)",
    "stavka2": "4/4 ТБ(56)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 ТМ(57)",
    "stavka2": "4/4 ТБ(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "4/4 ТМ(57)",
    "stavka2": "4/4 ТБ(56,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 ТМ(57,5)",
    "stavka2": "4/4 ТБ(57)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "4/4 ТМ(59)",
    "stavka2": "4/4 ТБ(59)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 Ф1(-0,5)",
    "stavka2": "4/4 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 Ф1(0)",
    "stavka2": "4/4 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "4/4 Ф1(0)",
    "stavka2": "4/4 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 Ф1(0)",
    "stavka2": "4/4 Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "4/4 Ф1(0,5)",
    "stavka2": "4/4 П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "4/4 Ф2(0)",
    "stavka2": "4/4 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "4/4 Ф2(0,5)",
    "stavka2": "4/4 Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "5/2 П1",
    "stavka2": "5/2 П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "5/2 П1",
    "stavka2": "5/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "5/2 П2",
    "stavka2": "5/2 П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Х2",
    "stavka2": "1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Х2",
    "stavka2": "2",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК 1",
    "stavka2": "ЖК Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК 1/2 ТМ(0,5)",
    "stavka2": "ЖК 1/2 ТБ(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК 1/2 ТМ(1)",
    "stavka2": "ЖК 1/2 ТБ(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК 1/2 ТМ(1,5)",
    "stavka2": "ЖК 1/2 ТБ(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК 1/2 ТМ(2)",
    "stavka2": "ЖК 1/2 ТБ(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК 1/2 Ф1(0)",
    "stavka2": "ЖК 1/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК 1/2 Ф1(0,5)",
    "stavka2": "ЖК 1/2 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК 1/2 Ф2(-0,25)",
    "stavka2": "ЖК 1/2 Ф1(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК ТМ(2)",
    "stavka2": "ЖК ТБ(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК ТМ(2,5)",
    "stavka2": "ЖК ТБ(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК ТМ(3)",
    "stavka2": "ЖК ТБ(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК ТМ(3,5)",
    "stavka2": "ЖК ТБ(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК ТМ(4)",
    "stavka2": "ЖК ТБ(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК ТМ(4,5)",
    "stavka2": "ЖК ТБ(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК ТМ(5,5)",
    "stavka2": "ЖК ТБ(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК ТМ(6,5)",
    "stavka2": "ЖК ТБ(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф1(-0,25)",
    "stavka2": "ЖК Ф2(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф1(-0,5)",
    "stavka2": "ЖК Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф1(-1)",
    "stavka2": "ЖК Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф1(0)",
    "stavka2": "ЖК Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф1(0,5)",
    "stavka2": "ЖК 2",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф1(0,5)",
    "stavka2": "ЖК Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф1(1)",
    "stavka2": "ЖК Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф1(1)",
    "stavka2": "ЖК Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф2(0)",
    "stavka2": "ЖК Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ЖК Ф2(1)",
    "stavka2": "ЖК Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(0,5)",
    "stavka2": "ИТ1Б(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(0,5)",
    "stavka2": "ИТ2Б(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(1)",
    "stavka2": "ИТ1Б(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(1,5)",
    "stavka2": "ИТ1Б(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(100)",
    "stavka2": "ИТ1Б(100)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(100)",
    "stavka2": "ИТ1Б(99,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(100)",
    "stavka2": "ИТ2Б(99,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(100,5)",
    "stavka2": "ИТ1Б(100)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(100,5)",
    "stavka2": "ИТ1Б(100,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(100,5)",
    "stavka2": "ИТ2Б(100)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(101,5)",
    "stavka2": "ИТ1Б(100)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(101,5)",
    "stavka2": "ИТ1Б(101,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(102)",
    "stavka2": "ИТ1Б(102)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(102)",
    "stavka2": "ИТ2Б(102)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(102,5)",
    "stavka2": "ИТ1Б(101,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(102,5)",
    "stavka2": "ИТ1Б(102)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(103)",
    "stavka2": "ИТ2Б(103)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(104)",
    "stavka2": "ИТ2Б(104)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(104,5)",
    "stavka2": "ИТ1Б(104)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(104,5)",
    "stavka2": "ИТ1Б(104,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(104,5)",
    "stavka2": "ИТ2Б(104,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(105)",
    "stavka2": "ИТ1Б(104,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(105,5)",
    "stavka2": "ИТ1Б(105,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(105,5)",
    "stavka2": "ИТ2Б(105)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(105,5)",
    "stavka2": "ИТ2Б(105,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(106)",
    "stavka2": "ИТ1Б(106)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(106,5)",
    "stavka2": "ИТ1Б(106,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(106,5)",
    "stavka2": "ИТ2Б(106)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(107)",
    "stavka2": "ИТ1Б(106,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(107)",
    "stavka2": "ИТ1Б(106,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(107)",
    "stavka2": "ИТ2Б(106,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(107)",
    "stavka2": "ИТ2Б(106,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(107,5)",
    "stavka2": "ИТ1Б(107,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(108)",
    "stavka2": "ИТ1Б(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(108)",
    "stavka2": "ИТ2Б(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(108)",
    "stavka2": "ИТ2Б(107,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(108)",
    "stavka2": "ИТ2Б(108)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(108,5)",
    "stavka2": "ИТ1Б(108)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(108,5)",
    "stavka2": "ИТ2Б(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(108,5)",
    "stavka2": "ИТ2Б(108)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(109,5)",
    "stavka2": "ИТ1Б(109)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(109,5)",
    "stavka2": "ИТ1Б(109,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(109,5)",
    "stavka2": "ИТ2Б(108,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(109,5)",
    "stavka2": "ИТ2Б(109)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(109,5)",
    "stavka2": "ИТ2Б(109)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(110)",
    "stavka2": "ИТ1Б(109)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(110)",
    "stavka2": "ИТ1Б(109,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(110)",
    "stavka2": "ИТ2Б(110)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(110,5)",
    "stavka2": "ИТ1Б(110)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(110,5)",
    "stavka2": "ИТ2Б(110)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(110,5)",
    "stavka2": "ИТ2Б(110,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(111)",
    "stavka2": "ИТ1Б(110,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(111)",
    "stavka2": "ИТ1Б(111)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(111)",
    "stavka2": "ИТ2Б(109,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(111)",
    "stavka2": "ИТ2Б(109,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(111)",
    "stavka2": "ИТ2Б(110,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(111,5)",
    "stavka2": "ИТ1Б(111)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(111,5)",
    "stavka2": "ИТ1Б(111)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(111,5)",
    "stavka2": "ИТ1Б(111,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(111,5)",
    "stavka2": "ИТ2Б(110,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(111,5)",
    "stavka2": "ИТ2Б(111,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(112)",
    "stavka2": "ИТ2Б(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(112)",
    "stavka2": "ИТ2Б(111,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(112,5)",
    "stavka2": "ИТ1Б(111)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(112,5)",
    "stavka2": "ИТ1Б(112)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(112,5)",
    "stavka2": "ИТ1Б(112,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(112,5)",
    "stavka2": "ИТ2Б(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(112,5)",
    "stavka2": "ИТ2Б(112)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(113)",
    "stavka2": "ИТ1Б(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(113)",
    "stavka2": "ИТ1Б(112,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(113)",
    "stavka2": "ИТ2Б(111)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(113)",
    "stavka2": "ИТ2Б(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(113)",
    "stavka2": "ИТ2Б(112,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(113)",
    "stavka2": "ИТ2Б(113)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(113,5)",
    "stavka2": "ИТ2Б(113)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(113,5)",
    "stavka2": "ИТ2Б(113,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(114)",
    "stavka2": "ИТ1Б(113)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(114)",
    "stavka2": "ИТ1Б(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(114)",
    "stavka2": "ИТ2Б(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(114)",
    "stavka2": "ИТ2Б(114)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(114,5)",
    "stavka2": "ИТ1Б(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(114,5)",
    "stavka2": "ИТ1Б(114)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(114,5)",
    "stavka2": "ИТ1Б(114)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(114,5)",
    "stavka2": "ИТ1Б(114,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(114,5)",
    "stavka2": "ИТ2Б(113)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(114,5)",
    "stavka2": "ИТ2Б(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(114,5)",
    "stavka2": "ИТ2Б(114,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(115)",
    "stavka2": "ИТ1Б(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(115)",
    "stavka2": "ИТ2Б(115)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(115,5)",
    "stavka2": "ИТ2Б(115,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(116)",
    "stavka2": "ИТ1Б(115,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(116)",
    "stavka2": "ИТ2Б(115)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(116)",
    "stavka2": "ИТ2Б(115,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(116)",
    "stavka2": "ИТ2Б(116)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(116,5)",
    "stavka2": "ИТ1Б(116)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(116,5)",
    "stavka2": "ИТ1Б(116,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(116,5)",
    "stavka2": "ИТ2Б(115,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(117)",
    "stavka2": "ИТ2Б(116,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(117)",
    "stavka2": "ИТ2Б(116,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(117)",
    "stavka2": "ИТ2Б(117)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(117,5)",
    "stavka2": "ИТ1Б(116,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(117,5)",
    "stavka2": "ИТ1Б(117)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(117,5)",
    "stavka2": "ИТ1Б(117,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(117,5)",
    "stavka2": "ИТ2Б(117)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(117,5)",
    "stavka2": "ИТ2Б(117,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(118)",
    "stavka2": "ИТ1Б(116)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(118)",
    "stavka2": "ИТ1Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(118)",
    "stavka2": "ИТ1Б(117,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(118)",
    "stavka2": "ИТ2Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(118)",
    "stavka2": "ИТ2Б(117,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(118,5)",
    "stavka2": "ИТ1Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(118,5)",
    "stavka2": "ИТ2Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(118,5)",
    "stavka2": "ИТ2Б(118,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(119)",
    "stavka2": "ИТ2Б(118,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(119,5)",
    "stavka2": "ИТ1Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(119,5)",
    "stavka2": "ИТ1Б(119)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(119,5)",
    "stavka2": "ИТ2Б(118,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(119,5)",
    "stavka2": "ИТ2Б(119)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(119,5)",
    "stavka2": "ИТ2Б(119)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(119,5)",
    "stavka2": "ИТ2Б(119,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(120)",
    "stavka2": "ИТ2Б(119,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(120)",
    "stavka2": "ИТ2Б(119,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(120,5)",
    "stavka2": "ИТ1Б(119,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(120,5)",
    "stavka2": "ИТ1Б(120,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(121)",
    "stavka2": "ИТ1Б(120,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(121)",
    "stavka2": "ИТ2Б(120,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(121,5)",
    "stavka2": "ИТ1Б(121)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(121,5)",
    "stavka2": "ИТ1Б(121,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(122)",
    "stavka2": "ИТ1Б(121)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(122)",
    "stavka2": "ИТ1Б(121,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(122,5)",
    "stavka2": "ИТ2Б(121,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(122,5)",
    "stavka2": "ИТ2Б(122,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(123,5)",
    "stavka2": "ИТ1Б(123)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(124,5)",
    "stavka2": "ИТ2Б(121,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(126)",
    "stavka2": "ИТ1Б(125,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(126,5)",
    "stavka2": "ИТ2Б(126,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(128)",
    "stavka2": "ИТ2Б(126,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(128,5)",
    "stavka2": "ИТ2Б(126,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(129)",
    "stavka2": "ИТ2Б(126,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(13,5)",
    "stavka2": "ИТ2Б(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(16,5)",
    "stavka2": "ИТ2Б(16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(17,5)",
    "stavka2": "ИТ2Б(17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(18,5)",
    "stavka2": "ИТ2Б(18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(19,5)",
    "stavka2": "ИТ2Б(19,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(2)",
    "stavka2": "ИТ1Б(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(20,5)",
    "stavka2": "ИТ2Б(20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(22,5)",
    "stavka2": "ТБ(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(22,5)",
    "stavka2": "ТБ(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(23,5)",
    "stavka2": "ИТ2Б(23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(23,5)",
    "stavka2": "ТБ(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(23,5)",
    "stavka2": "ТБ(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(23,5)",
    "stavka2": "ТБ(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(23,5)",
    "stavka2": "ТБ(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(24,5)",
    "stavka2": "ИТ2Б(24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(24,5)",
    "stavka2": "ТБ(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(24,5)",
    "stavka2": "ТБ(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(26,5)",
    "stavka2": "ИТ2Б(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(27,5)",
    "stavka2": "ИТ2Б(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(29,5)",
    "stavka2": "ИТ2Б(29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(3,5)",
    "stavka2": "ИТ1Б(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(318,5)",
    "stavka2": "ТБ(38,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(52)",
    "stavka2": "ИТ1Б(51,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(59,5)",
    "stavka2": "ИТ1Б(59,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(61,5)",
    "stavka2": "ИТ1Б(61,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(62)",
    "stavka2": "ИТ2Б(62)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(63)",
    "stavka2": "ИТ1Б(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(63)",
    "stavka2": "ИТ1Б(63)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(64)",
    "stavka2": "ИТ1Б(64)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(64,5)",
    "stavka2": "ИТ1Б(64)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(64,5)",
    "stavka2": "ИТ1Б(64,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(65)",
    "stavka2": "ИТ1Б(64)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(65,5)",
    "stavka2": "ИТ2Б(65,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(67,5)",
    "stavka2": "ИТ1Б(67)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(67,5)",
    "stavka2": "ИТ1Б(67,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(68,5)",
    "stavka2": "ИТ1Б(67,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(70,5)",
    "stavka2": "ИТ1Б(69,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(71,5)",
    "stavka2": "ИТ1Б(70,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(72)",
    "stavka2": "ИТ1Б(71,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(72,5)",
    "stavka2": "ИТ1Б(71,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(72,5)",
    "stavka2": "ИТ1Б(72,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(73,5)",
    "stavka2": "ИТ1Б(73)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(73,5)",
    "stavka2": "ТБ(17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(74)",
    "stavka2": "ИТ1Б(73)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(74,5)",
    "stavka2": "ИТ1Б(73,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(74,5)",
    "stavka2": "ИТ1Б(74)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(74,5)",
    "stavka2": "ИТ1Б(74,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(75)",
    "stavka2": "ИТ1Б(74,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(75)",
    "stavka2": "ИТ1Б(75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(75,5)",
    "stavka2": "ИТ1Б(75,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(75,5)",
    "stavka2": "ИТ2Б(75,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(76)",
    "stavka2": "ИТ1Б(76)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(76,5)",
    "stavka2": "ИТ1Б(76)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(76,5)",
    "stavka2": "ИТ1Б(76,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(77)",
    "stavka2": "ИТ1Б(76)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(77)",
    "stavka2": "ИТ1Б(76,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(77)",
    "stavka2": "ИТ1Б(76,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(77)",
    "stavka2": "ИТ1Б(77)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(77)",
    "stavka2": "ИТ2Б(77)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(77,5)",
    "stavka2": "ИТ1Б(76,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(77,5)",
    "stavka2": "ИТ1Б(77)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(77,5)",
    "stavka2": "ИТ1Б(77)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(77,5)",
    "stavka2": "ИТ1Б(77,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(78)",
    "stavka2": "ИТ1Б(77,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(78)",
    "stavka2": "ИТ1Б(78)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(78,5)",
    "stavka2": "ИТ1Б(77,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(78,5)",
    "stavka2": "ИТ1Б(78)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(78,5)",
    "stavka2": "ИТ1Б(78,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(79)",
    "stavka2": "ИТ1Б(78)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(79)",
    "stavka2": "ИТ1Б(78)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(79)",
    "stavka2": "ИТ1Б(79)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(79,5)",
    "stavka2": "ИТ1Б(78,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(79,5)",
    "stavka2": "ИТ1Б(79)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(79,5)",
    "stavka2": "ИТ1Б(79,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(80)",
    "stavka2": "ИТ1Б(79)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(80)",
    "stavka2": "ИТ1Б(79)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(80)",
    "stavka2": "ИТ1Б(79,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(80)",
    "stavka2": "ИТ1Б(80)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(80)",
    "stavka2": "ИТ2Б(80)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(80,5)",
    "stavka2": "ИТ1Б(80)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(80,5)",
    "stavka2": "ИТ1Б(80)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(80,5)",
    "stavka2": "ИТ1Б(80,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(81)",
    "stavka2": "ИТ1Б(80)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(81)",
    "stavka2": "ИТ1Б(80,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(81)",
    "stavka2": "ИТ1Б(80,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(81)",
    "stavka2": "ИТ1Б(81)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(81,5)",
    "stavka2": "ИТ1Б(80,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(81,5)",
    "stavka2": "ИТ1Б(81)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(81,5)",
    "stavka2": "ИТ1Б(81)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(81,5)",
    "stavka2": "ИТ1Б(81,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(82)",
    "stavka2": "ИТ1Б(80,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(82)",
    "stavka2": "ИТ1Б(82)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(82,5)",
    "stavka2": "ИТ1Б(80,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(82,5)",
    "stavka2": "ИТ1Б(82)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(82,5)",
    "stavka2": "ИТ1Б(82)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(82,5)",
    "stavka2": "ИТ1Б(82,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(83)",
    "stavka2": "ИТ1Б(82)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(83)",
    "stavka2": "ИТ1Б(82,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(83)",
    "stavka2": "ИТ1Б(83)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(83)",
    "stavka2": "ИТ2Б(83)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(83,5)",
    "stavka2": "ИТ1Б(82,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(83,5)",
    "stavka2": "ИТ1Б(83)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(83,5)",
    "stavka2": "ИТ1Б(83)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(83,5)",
    "stavka2": "ИТ1Б(83,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(84)",
    "stavka2": "ИТ1Б(83)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(84)",
    "stavka2": "ИТ1Б(83,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(84)",
    "stavka2": "ИТ1Б(84)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(84,5)",
    "stavka2": "ИТ1Б(84)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(84,5)",
    "stavka2": "ИТ1Б(84)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(84,5)",
    "stavka2": "ИТ1Б(84,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(85)",
    "stavka2": "ИТ1Б(84)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(85)",
    "stavka2": "ИТ1Б(84,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(85,5)",
    "stavka2": "ИТ1Б(84,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(85,5)",
    "stavka2": "ИТ1Б(85)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(85,5)",
    "stavka2": "ИТ1Б(85)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(85,5)",
    "stavka2": "ИТ1Б(85,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(85,5)",
    "stavka2": "ИТ2Б(85,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(86)",
    "stavka2": "ИТ1Б(84)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(86)",
    "stavka2": "ИТ1Б(84,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(86)",
    "stavka2": "ИТ1Б(85,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(86)",
    "stavka2": "ИТ1Б(85,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(86)",
    "stavka2": "ИТ1Б(86)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(86,5)",
    "stavka2": "ИТ1Б(86)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(86,5)",
    "stavka2": "ИТ1Б(86,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(87)",
    "stavka2": "ИТ1Б(86)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(87)",
    "stavka2": "ИТ1Б(86,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(87)",
    "stavka2": "ИТ1Б(86,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(87)",
    "stavka2": "ИТ1Б(87)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(87,5)",
    "stavka2": "ИТ1Б(86,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(87,5)",
    "stavka2": "ИТ1Б(87)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(87,5)",
    "stavka2": "ИТ1Б(87,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(88)",
    "stavka2": "ИТ1Б(87)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(88)",
    "stavka2": "ИТ1Б(87,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(88)",
    "stavka2": "ИТ1Б(88)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(88,5)",
    "stavka2": "ИТ1Б(87,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(88,5)",
    "stavka2": "ИТ1Б(88)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(88,5)",
    "stavka2": "ИТ1Б(88,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(89)",
    "stavka2": "ИТ1Б(87,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(89)",
    "stavka2": "ИТ1Б(88,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(89)",
    "stavka2": "ИТ1Б(89)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(89,5)",
    "stavka2": "ИТ1Б(88,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(89,5)",
    "stavka2": "ИТ1Б(89)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(89,5)",
    "stavka2": "ИТ1Б(89,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(90)",
    "stavka2": "ИТ1Б(89,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(90,5)",
    "stavka2": "ИТ1Б(90)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(90,5)",
    "stavka2": "ИТ1Б(90,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(91)",
    "stavka2": "ИТ1Б(90,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(91)",
    "stavka2": "ИТ1Б(91)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(91,5)",
    "stavka2": "ИТ1Б(91,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(92)",
    "stavka2": "ИТ1Б(91)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(92)",
    "stavka2": "ИТ1Б(91,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(92)",
    "stavka2": "ИТ1Б(92)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(92,5)",
    "stavka2": "ИТ1Б(91)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(92,5)",
    "stavka2": "ИТ1Б(92)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(92,5)",
    "stavka2": "ИТ1Б(92,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(93)",
    "stavka2": "ИТ1Б(92,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(93)",
    "stavka2": "ИТ1Б(93)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(93,5)",
    "stavka2": "ИТ1Б(92)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(93,5)",
    "stavka2": "ИТ1Б(92,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(93,5)",
    "stavka2": "ИТ1Б(93)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(93,5)",
    "stavka2": "ИТ1Б(93,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(94)",
    "stavka2": "ИТ1Б(93,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(94,5)",
    "stavka2": "ИТ1Б(94)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(94,5)",
    "stavka2": "ИТ1Б(94,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(95)",
    "stavka2": "ИТ1Б(94)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(95)",
    "stavka2": "ИТ1Б(95)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(95,5)",
    "stavka2": "ИТ1Б(95)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(95,5)",
    "stavka2": "ИТ1Б(95,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(96)",
    "stavka2": "ИТ1Б(95,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(96)",
    "stavka2": "ИТ1Б(96)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(96,5)",
    "stavka2": "ИТ1Б(96)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(96,5)",
    "stavka2": "ИТ1Б(96,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(97)",
    "stavka2": "ИТ1Б(96,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(97)",
    "stavka2": "ИТ1Б(97)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(97,5)",
    "stavka2": "ИТ1Б(96,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(97,5)",
    "stavka2": "ИТ1Б(96,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(97,5)",
    "stavka2": "ИТ1Б(97)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(98)",
    "stavka2": "ИТ1Б(97,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(98)",
    "stavka2": "ИТ1Б(98)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(98,5)",
    "stavka2": "ИТ1Б(97,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(98,5)",
    "stavka2": "ИТ1Б(98)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(98,5)",
    "stavka2": "ИТ1Б(98)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(99)",
    "stavka2": "ИТ1Б(97,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(99)",
    "stavka2": "ИТ1Б(98,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(99)",
    "stavka2": "ИТ1Б(99)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(99,5)",
    "stavka2": "ИТ1Б(98)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(99,5)",
    "stavka2": "ИТ1Б(99)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ1М(99,5)",
    "stavka2": "ИТ1Б(99)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(99,5)",
    "stavka2": "ИТ1Б(99,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ1М(99,5)",
    "stavka2": "ИТ2Б(99,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(0,5)",
    "stavka2": "ИТ1Б(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(0,5)",
    "stavka2": "ИТ2Б(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(1)",
    "stavka2": "ИТ1Б(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(1)",
    "stavka2": "ИТ2Б(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(1,5)",
    "stavka2": "ИТ2Б(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(1,5)",
    "stavka2": "ТБ(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(100)",
    "stavka2": "ИТ1Б(99,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(100,5)",
    "stavka2": "ИТ1Б(100)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(100,5)",
    "stavka2": "ИТ1Б(99,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(101)",
    "stavka2": "ИТ1Б(100,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(101)",
    "stavka2": "ИТ2Б(101)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(101,5)",
    "stavka2": "ИТ2Б(101)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(102)",
    "stavka2": "ИТ1Б(101,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(102)",
    "stavka2": "ИТ2Б(101,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(102)",
    "stavka2": "ТБ(22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(102,5)",
    "stavka2": "ИТ1Б(102)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(102,5)",
    "stavka2": "ИТ2Б(102)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(102,5)",
    "stavka2": "ИТ2Б(102,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(103)",
    "stavka2": "ИТ1Б(102,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(103)",
    "stavka2": "ИТ2Б(102)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(103)",
    "stavka2": "ИТ2Б(102,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(104)",
    "stavka2": "ИТ1Б(103,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(104)",
    "stavka2": "ИТ2Б(103,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(104,5)",
    "stavka2": "ИТ1Б(104,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(105)",
    "stavka2": "ИТ1Б(104,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(105)",
    "stavka2": "ИТ2Б(104,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(105)",
    "stavka2": "ИТ2Б(105)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(105,5)",
    "stavka2": "ИТ1Б(104,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(105,5)",
    "stavka2": "ИТ1Б(105)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(105,5)",
    "stavka2": "ИТ1Б(105,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(105,5)",
    "stavka2": "ИТ2Б(104,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(105,5)",
    "stavka2": "ИТ2Б(104,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(106)",
    "stavka2": "ИТ1Б(105,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(106)",
    "stavka2": "ИТ1Б(106)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(106)",
    "stavka2": "ИТ2Б(105,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(106,5)",
    "stavka2": "ИТ1Б(106,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(107)",
    "stavka2": "ИТ1Б(106)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(107)",
    "stavka2": "ИТ1Б(106,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(107)",
    "stavka2": "ИТ2Б(106)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(107)",
    "stavka2": "ИТ2Б(106,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(107,5)",
    "stavka2": "ИТ1Б(106,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(107,5)",
    "stavka2": "ИТ1Б(107)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(107,5)",
    "stavka2": "ИТ1Б(107,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(107,5)",
    "stavka2": "ИТ2Б(107,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(108)",
    "stavka2": "ИТ1Б(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(108)",
    "stavka2": "ИТ1Б(108)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(108)",
    "stavka2": "ИТ2Б(108)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(108,5)",
    "stavka2": "ИТ1Б(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(108,5)",
    "stavka2": "ИТ1Б(108)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(108,5)",
    "stavka2": "ИТ2Б(108)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(108,5)",
    "stavka2": "ИТ2Б(108,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(109)",
    "stavka2": "ИТ1Б(108)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(109)",
    "stavka2": "ИТ2Б(108,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(109,5)",
    "stavka2": "ИТ1Б(108,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(109,5)",
    "stavka2": "ИТ1Б(109)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(109,5)",
    "stavka2": "ИТ1Б(109,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(109,5)",
    "stavka2": "ИТ2Б(108,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(109,5)",
    "stavka2": "ИТ2Б(109)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(110)",
    "stavka2": "ИТ1Б(107,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(110)",
    "stavka2": "ИТ1Б(109,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(110)",
    "stavka2": "ИТ1Б(109,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(110)",
    "stavka2": "ИТ2Б(109,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(110,5)",
    "stavka2": "ИТ1Б(108)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(110,5)",
    "stavka2": "ИТ1Б(109,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(110,5)",
    "stavka2": "ИТ1Б(110)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(110,5)",
    "stavka2": "ИТ1Б(110,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(110,5)",
    "stavka2": "ИТ2Б(110)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(110,5)",
    "stavka2": "ИТ2Б(110,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(111)",
    "stavka2": "ИТ1Б(109,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(111)",
    "stavka2": "ИТ1Б(111)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(111)",
    "stavka2": "ИТ2Б(110,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(111)",
    "stavka2": "ИТ2Б(111)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(111,5)",
    "stavka2": "ИТ1Б(110,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(111,5)",
    "stavka2": "ИТ1Б(111)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(111,5)",
    "stavka2": "ИТ2Б(110,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(111,5)",
    "stavka2": "ИТ2Б(111,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(112)",
    "stavka2": "ИТ1Б(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(112)",
    "stavka2": "ИТ2Б(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(112,5)",
    "stavka2": "ИТ1Б(111,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(112,5)",
    "stavka2": "ИТ1Б(111,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(112,5)",
    "stavka2": "ИТ1Б(112)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(112,5)",
    "stavka2": "ИТ1Б(112,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(112,5)",
    "stavka2": "ИТ2Б(112)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(113,5)",
    "stavka2": "ИТ1Б(113)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(113,5)",
    "stavka2": "ИТ1Б(113,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(113,5)",
    "stavka2": "ИТ2Б(112,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(113,5)",
    "stavka2": "ИТ2Б(113)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(114)",
    "stavka2": "ИТ1Б(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(114)",
    "stavka2": "ИТ2Б(112,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(114,5)",
    "stavka2": "ИТ1Б(113)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(114,5)",
    "stavka2": "ИТ1Б(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(114,5)",
    "stavka2": "ИТ1Б(114)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(114,5)",
    "stavka2": "ИТ2Б(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(115)",
    "stavka2": "ИТ1Б(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(115)",
    "stavka2": "ИТ2Б(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(115,5)",
    "stavka2": "ИТ1Б(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(115,5)",
    "stavka2": "ИТ1Б(115)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(115,5)",
    "stavka2": "ИТ1Б(115)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(115,5)",
    "stavka2": "ИТ2Б(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(116)",
    "stavka2": "ИТ1Б(115,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(116,5)",
    "stavka2": "ИТ1Б(116)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(116,5)",
    "stavka2": "ИТ1Б(116,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(116,5)",
    "stavka2": "ИТ2Б(116)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(117)",
    "stavka2": "ИТ2Б(116,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(117,5)",
    "stavka2": "ИТ1Б(117)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(117,5)",
    "stavka2": "ИТ2Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(117,5)",
    "stavka2": "ИТ2Б(117,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(118)",
    "stavka2": "ИТ2Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(118,5)",
    "stavka2": "ИТ1Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(118,5)",
    "stavka2": "ИТ1Б(118,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(118,5)",
    "stavka2": "ИТ2Б(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(118,5)",
    "stavka2": "ИТ2Б(118,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(119)",
    "stavka2": "ИТ1Б(118,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(119)",
    "stavka2": "ИТ2Б(118,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(119,5)",
    "stavka2": "ИТ1Б(119,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(12)",
    "stavka2": "ТБ(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(120)",
    "stavka2": "ИТ1Б(119,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(120,5)",
    "stavka2": "ИТ1Б(119,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(120,5)",
    "stavka2": "ИТ1Б(120)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(120,5)",
    "stavka2": "ИТ1Б(120,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(120,5)",
    "stavka2": "ИТ2Б(120,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(121,5)",
    "stavka2": "ИТ1Б(120,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(121,5)",
    "stavka2": "ИТ1Б(121)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(121,5)",
    "stavka2": "ИТ1Б(121)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(122)",
    "stavka2": "ИТ1Б(121,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(124,5)",
    "stavka2": "ИТ1Б(124)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(126,5)",
    "stavka2": "ИТ1Б(126)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(126,5)",
    "stavka2": "ИТ2Б(125,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(126,5)",
    "stavka2": "ИТ2Б(126,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(127,5)",
    "stavka2": "ИТ1Б(126,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(128)",
    "stavka2": "ИТ2Б(127,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(128,5)",
    "stavka2": "ИТ2Б(127,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(129)",
    "stavka2": "ИТ2Б(127,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(129)",
    "stavka2": "ИТ2Б(127,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(14,5)",
    "stavka2": "ИТ1Б(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(16,5)",
    "stavka2": "ИТ1Б(16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(19,5)",
    "stavka2": "ИТ1Б(19,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(2)",
    "stavka2": "ИТ1Б(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(2)",
    "stavka2": "ИТ2Б(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(2,5)",
    "stavka2": "ИТ2Б(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(20,5)",
    "stavka2": "ИТ1Б(20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(20,5)",
    "stavka2": "ТБ(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(21,5)",
    "stavka2": "ИТ1Б(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(21,5)",
    "stavka2": "ТБ(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(22,5)",
    "stavka2": "ИТ1Б(22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(23,5)",
    "stavka2": "ИТ1Б(23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(23,5)",
    "stavka2": "ТБ(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(23,5)",
    "stavka2": "ТБ(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(24,5)",
    "stavka2": "ИТ1Б(24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(26,5)",
    "stavka2": "ИТ1Б(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(27,5)",
    "stavka2": "ИТ1Б(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(28,5)",
    "stavka2": "ИТ1Б(28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(3)",
    "stavka2": "ИТ2Б(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(30,5)",
    "stavka2": "ТБ(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(30,5)",
    "stavka2": "ТБ(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(56,5)",
    "stavka2": "ИТ2Б(56)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(59)",
    "stavka2": "ИТ2Б(59)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(59,5)",
    "stavka2": "ИТ2Б(59)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(59,5)",
    "stavka2": "ИТ2Б(59,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(60,5)",
    "stavka2": "ИТ2Б(60)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(61,5)",
    "stavka2": "ИТ2Б(61,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(62,5)",
    "stavka2": "ИТ2Б(62,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(63)",
    "stavka2": "ИТ2Б(63)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(65)",
    "stavka2": "ИТ1Б(65)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(65,5)",
    "stavka2": "ИТ2Б(65,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(66)",
    "stavka2": "ИТ2Б(66)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(66,5)",
    "stavka2": "ИТ2Б(65,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(66,5)",
    "stavka2": "ИТ2Б(66)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(67,5)",
    "stavka2": "ИТ2Б(67,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(69)",
    "stavka2": "ИТ2Б(68,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(69)",
    "stavka2": "ИТ2Б(68,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(69,5)",
    "stavka2": "ИТ1Б(69,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(70)",
    "stavka2": "ИТ2Б(70)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(70,5)",
    "stavka2": "ИТ2Б(70,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(71)",
    "stavka2": "ИТ2Б(70,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(71)",
    "stavka2": "ИТ2Б(71)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(71,5)",
    "stavka2": "ИТ2Б(71)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(71,5)",
    "stavka2": "ИТ2Б(71,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(72)",
    "stavka2": "ИТ2Б(72)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(72,5)",
    "stavka2": "ИТ2Б(71)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(72,5)",
    "stavka2": "ИТ2Б(71,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(72,5)",
    "stavka2": "ИТ2Б(72)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(72,5)",
    "stavka2": "ИТ2Б(72,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(73)",
    "stavka2": "ИТ2Б(72)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(73)",
    "stavka2": "ИТ2Б(72,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(73,5)",
    "stavka2": "ИТ2Б(73)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(73,5)",
    "stavka2": "ИТ2Б(73,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(74)",
    "stavka2": "ИТ2Б(73,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(74)",
    "stavka2": "ИТ2Б(73,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(74,5)",
    "stavka2": "ИТ2Б(73,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(74,5)",
    "stavka2": "ИТ2Б(74)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(74,5)",
    "stavka2": "ИТ2Б(74)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(74,5)",
    "stavka2": "ИТ2Б(74,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(75)",
    "stavka2": "ИТ2Б(74)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(75)",
    "stavka2": "ИТ2Б(74,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(75)",
    "stavka2": "ИТ2Б(75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(75,5)",
    "stavka2": "ИТ2Б(74,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(75,5)",
    "stavka2": "ИТ2Б(75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(75,5)",
    "stavka2": "ИТ2Б(75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(75,5)",
    "stavka2": "ИТ2Б(75,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(76)",
    "stavka2": "ИТ2Б(75,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(76)",
    "stavka2": "ИТ2Б(75,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(76)",
    "stavka2": "ИТ2Б(76)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(76,5)",
    "stavka2": "ИТ2Б(75,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(76,5)",
    "stavka2": "ИТ2Б(76)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(76,5)",
    "stavka2": "ИТ2Б(76,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(77)",
    "stavka2": "ИТ2Б(75,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(77)",
    "stavka2": "ИТ2Б(75,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(77)",
    "stavka2": "ИТ2Б(76,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(77)",
    "stavka2": "ИТ2Б(77)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(77,5)",
    "stavka2": "ИТ2Б(77)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(77,5)",
    "stavka2": "ИТ2Б(77,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(78)",
    "stavka2": "ИТ2Б(77,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(78)",
    "stavka2": "ИТ2Б(78)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(78,5)",
    "stavka2": "ИТ2Б(78)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(78,5)",
    "stavka2": "ИТ2Б(78,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(78,5)",
    "stavka2": "ТБ(17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(79)",
    "stavka2": "ИТ2Б(78)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(79)",
    "stavka2": "ИТ2Б(78,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(79)",
    "stavka2": "ИТ2Б(79)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(79,5)",
    "stavka2": "ИТ2Б(78,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(79,5)",
    "stavka2": "ИТ2Б(79)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(79,5)",
    "stavka2": "ИТ2Б(79)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(79,5)",
    "stavka2": "ИТ2Б(79,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(80)",
    "stavka2": "ИТ2Б(78,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(80)",
    "stavka2": "ИТ2Б(79)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(80)",
    "stavka2": "ИТ2Б(79,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(80)",
    "stavka2": "ИТ2Б(80)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(80,5)",
    "stavka2": "ИТ2Б(80)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(80,5)",
    "stavka2": "ИТ2Б(80,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(81)",
    "stavka2": "ИТ2Б(80,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(81)",
    "stavka2": "ИТ2Б(81)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(81,5)",
    "stavka2": "ИТ2Б(80)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(81,5)",
    "stavka2": "ИТ2Б(80,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(81,5)",
    "stavka2": "ИТ2Б(81)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(81,5)",
    "stavka2": "ИТ2Б(81,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(82)",
    "stavka2": "ИТ2Б(81,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(82)",
    "stavka2": "ИТ2Б(81,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(82)",
    "stavka2": "ИТ2Б(82)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(82,5)",
    "stavka2": "ИТ2Б(81,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(82,5)",
    "stavka2": "ИТ2Б(82)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(82,5)",
    "stavka2": "ИТ2Б(82,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(83)",
    "stavka2": "ИТ2Б(82)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(83)",
    "stavka2": "ИТ2Б(82,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(83)",
    "stavka2": "ИТ2Б(82,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(83)",
    "stavka2": "ИТ2Б(83)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(83,5)",
    "stavka2": "ИТ1Б(83,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(83,5)",
    "stavka2": "ИТ2Б(83)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(83,5)",
    "stavka2": "ИТ2Б(83,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(84)",
    "stavka2": "ИТ2Б(83,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(84)",
    "stavka2": "ИТ2Б(84)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(84,5)",
    "stavka2": "ИТ1Б(84,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(84,5)",
    "stavka2": "ИТ2Б(83)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(84,5)",
    "stavka2": "ИТ2Б(83,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(84,5)",
    "stavka2": "ИТ2Б(84)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(85)",
    "stavka2": "ИТ2Б(84)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(85)",
    "stavka2": "ИТ2Б(84,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(85,5)",
    "stavka2": "ИТ2Б(84,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(85,5)",
    "stavka2": "ИТ2Б(85)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(85,5)",
    "stavka2": "ИТ2Б(85,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(86)",
    "stavka2": "ИТ1Б(86)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(86)",
    "stavka2": "ИТ2Б(85,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(86)",
    "stavka2": "ИТ2Б(86)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(86,5)",
    "stavka2": "ИТ2Б(85,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(86,5)",
    "stavka2": "ИТ2Б(86)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(87)",
    "stavka2": "ИТ2Б(86,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(87)",
    "stavka2": "ИТ2Б(86,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(87)",
    "stavka2": "ИТ2Б(87)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(87,5)",
    "stavka2": "ИТ2Б(86,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(87,5)",
    "stavka2": "ИТ2Б(87)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(87,5)",
    "stavka2": "ИТ2Б(87)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(88)",
    "stavka2": "ИТ2Б(87,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(88)",
    "stavka2": "ИТ2Б(87,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(88,5)",
    "stavka2": "ИТ2Б(87)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(88,5)",
    "stavka2": "ИТ2Б(87,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(88,5)",
    "stavka2": "ИТ2Б(88)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(88,5)",
    "stavka2": "ИТ2Б(88,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(89,5)",
    "stavka2": "ИТ2Б(89)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(90)",
    "stavka2": "ИТ2Б(89,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(90)",
    "stavka2": "ИТ2Б(89,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(90,5)",
    "stavka2": "ИТ2Б(90,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(91)",
    "stavka2": "ИТ2Б(90,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(91)",
    "stavka2": "ИТ2Б(91)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(91,5)",
    "stavka2": "ИТ2Б(91,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(92)",
    "stavka2": "ИТ2Б(92)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(92,5)",
    "stavka2": "ИТ2Б(91,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(92,5)",
    "stavka2": "ИТ2Б(92)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(93)",
    "stavka2": "ИТ1Б(93)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(94,5)",
    "stavka2": "ИТ2Б(94)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(94,5)",
    "stavka2": "ИТ2Б(94,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(95)",
    "stavka2": "ИТ2Б(94,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(95,5)",
    "stavka2": "ИТ2Б(94,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(95,5)",
    "stavka2": "ИТ2Б(95)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(95,5)",
    "stavka2": "ИТ2Б(95)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(96)",
    "stavka2": "ИТ2Б(96)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(96,5)",
    "stavka2": "ИТ2Б(95,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(96,5)",
    "stavka2": "ИТ2Б(96)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(96,5)",
    "stavka2": "ИТ2Б(96,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(97)",
    "stavka2": "ИТ2Б(96,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(98,5)",
    "stavka2": "ИТ2Б(98,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ИТ2М(99,5)",
    "stavka2": "ИТ1Б(99)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ИТ2М(99,5)",
    "stavka2": "ИТ2Б(98,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф1(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П1",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П2",
    "stavka2": "П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф1(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф1(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "П2",
    "stavka2": "Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(0,5)",
    "stavka2": "ТБ(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(0,5)",
    "stavka2": "ТБ(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(1)",
    "stavka2": "ТБ(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(1,25)",
    "stavka2": "ТБ(1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(1,5)",
    "stavka2": "ТБ(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(1,75)",
    "stavka2": "ТБ(1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(10)",
    "stavka2": "ТБ(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(10,5)",
    "stavka2": "ТБ(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(10,5)",
    "stavka2": "ТБ(7,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(10,5)",
    "stavka2": "ТБ(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(102,5)",
    "stavka2": "ТБ(102,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(105,5)",
    "stavka2": "ТБ(105,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(109,5)",
    "stavka2": "ТБ(108,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(11)",
    "stavka2": "ТБ(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(11,5)",
    "stavka2": "ТБ(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(112,5)",
    "stavka2": "ТБ(112,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(113,5)",
    "stavka2": "ТБ(112)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(113,5)",
    "stavka2": "ТБ(112,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(114)",
    "stavka2": "ТБ(113,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(114)",
    "stavka2": "ТБ(114)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(114,5)",
    "stavka2": "ТБ(114)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(115)",
    "stavka2": "ТБ(114,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(115)",
    "stavka2": "ТБ(115)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(115,5)",
    "stavka2": "ТБ(115,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(116)",
    "stavka2": "ТБ(115,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(116,5)",
    "stavka2": "ТБ(115,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(116,5)",
    "stavka2": "ТБ(116)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(117)",
    "stavka2": "ТБ(116,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(117)",
    "stavka2": "ТБ(117)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(117,5)",
    "stavka2": "ТБ(117,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(118)",
    "stavka2": "ТБ(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(118)",
    "stavka2": "ТБ(118)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(118,5)",
    "stavka2": "ТБ(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(118,5)",
    "stavka2": "ТБ(118)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(118,5)",
    "stavka2": "ТБ(118,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(119)",
    "stavka2": "ТБ(117,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(119)",
    "stavka2": "ТБ(118)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(119)",
    "stavka2": "ТБ(118,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(119,5)",
    "stavka2": "ТБ(119)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(119,5)",
    "stavka2": "ТБ(119,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(12)",
    "stavka2": "ТБ(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(12,5)",
    "stavka2": "ТБ(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(120)",
    "stavka2": "ТБ(119,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(120,5)",
    "stavka2": "ТБ(119,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(120,5)",
    "stavka2": "ТБ(120,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(121)",
    "stavka2": "ТБ(120)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(121)",
    "stavka2": "ТБ(120,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(121)",
    "stavka2": "ТБ(120,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(121)",
    "stavka2": "ТБ(121)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(121,5)",
    "stavka2": "ТБ(120,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(121,5)",
    "stavka2": "ТБ(121)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(121,5)",
    "stavka2": "ТБ(121,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(122)",
    "stavka2": "ТБ(121,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(122,5)",
    "stavka2": "ТБ(122)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(122,5)",
    "stavka2": "ТБ(122,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(123)",
    "stavka2": "ТБ(121,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(123)",
    "stavka2": "ТБ(122,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(123)",
    "stavka2": "ТБ(123)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(123,5)",
    "stavka2": "ТБ(120,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(123,5)",
    "stavka2": "ТБ(121)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(123,5)",
    "stavka2": "ТБ(121,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(123,5)",
    "stavka2": "ТБ(122,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(123,5)",
    "stavka2": "ТБ(123)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(123,5)",
    "stavka2": "ТБ(123,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(124)",
    "stavka2": "ТБ(122,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(124)",
    "stavka2": "ТБ(123)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(124)",
    "stavka2": "ТБ(123,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(124)",
    "stavka2": "ТБ(124)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(124,5)",
    "stavka2": "ТБ(120,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(124,5)",
    "stavka2": "ТБ(123,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(124,5)",
    "stavka2": "ТБ(124)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(124,5)",
    "stavka2": "ТБ(124,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(125)",
    "stavka2": "ТБ(123,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(125)",
    "stavka2": "ТБ(124)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(125)",
    "stavka2": "ТБ(124,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(125,5)",
    "stavka2": "ТБ(124)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(125,5)",
    "stavka2": "ТБ(124,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(125,5)",
    "stavka2": "ТБ(125)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(125,5)",
    "stavka2": "ТБ(125,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(126)",
    "stavka2": "ТБ(124,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(126)",
    "stavka2": "ТБ(125,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(126)",
    "stavka2": "ТБ(126)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(126,5)",
    "stavka2": "ТБ(124,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(126,5)",
    "stavka2": "ТБ(125)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(126,5)",
    "stavka2": "ТБ(125,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(126,5)",
    "stavka2": "ТБ(126)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(126,5)",
    "stavka2": "ТБ(126,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(127)",
    "stavka2": "ТБ(125,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(127)",
    "stavka2": "ТБ(126)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(127)",
    "stavka2": "ТБ(126,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(127)",
    "stavka2": "ТБ(126,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(127)",
    "stavka2": "ТБ(127)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(127,5)",
    "stavka2": "ТБ(125)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(127,5)",
    "stavka2": "ТБ(125,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(127,5)",
    "stavka2": "ТБ(126,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(127,5)",
    "stavka2": "ТБ(126,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(127,5)",
    "stavka2": "ТБ(127)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(127,5)",
    "stavka2": "ТБ(127)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(127,5)",
    "stavka2": "ТБ(127,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(128)",
    "stavka2": "ТБ(125,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(128)",
    "stavka2": "ТБ(126)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(128)",
    "stavka2": "ТБ(127)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(128)",
    "stavka2": "ТБ(127,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(128)",
    "stavka2": "ТБ(127,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(128)",
    "stavka2": "ТБ(128)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(128,5)",
    "stavka2": "ТБ(126,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(128,5)",
    "stavka2": "ТБ(127)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(128,5)",
    "stavka2": "ТБ(127,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(128,5)",
    "stavka2": "ТБ(128)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(128,5)",
    "stavka2": "ТБ(128,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(129)",
    "stavka2": "ТБ(127)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(129)",
    "stavka2": "ТБ(127,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(129)",
    "stavka2": "ТБ(128,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(129)",
    "stavka2": "ТБ(129)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(129,5)",
    "stavka2": "ТБ(127)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(129,5)",
    "stavka2": "ТБ(128)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(129,5)",
    "stavka2": "ТБ(128,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(129,5)",
    "stavka2": "ТБ(128,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(129,5)",
    "stavka2": "ТБ(129)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(129,5)",
    "stavka2": "ТБ(129,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(13,5)",
    "stavka2": "ТБ(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(130)",
    "stavka2": "ТБ(127,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(130)",
    "stavka2": "ТБ(128,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(130)",
    "stavka2": "ТБ(129)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(130)",
    "stavka2": "ТБ(129)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(130)",
    "stavka2": "ТБ(129,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(130)",
    "stavka2": "ТБ(130)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(130,5)",
    "stavka2": "ТБ(128,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(130,5)",
    "stavka2": "ТБ(129)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(130,5)",
    "stavka2": "ТБ(129,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(130,5)",
    "stavka2": "ТБ(129,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(130,5)",
    "stavka2": "ТБ(130)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(130,5)",
    "stavka2": "ТБ(130)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(130,5)",
    "stavka2": "ТБ(130,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(131)",
    "stavka2": "ТБ(127,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(131)",
    "stavka2": "ТБ(130)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(131)",
    "stavka2": "ТБ(130,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(131)",
    "stavka2": "ТБ(130,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(131)",
    "stavka2": "ТБ(131)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(131,5)",
    "stavka2": "ТБ(129,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(131,5)",
    "stavka2": "ТБ(130)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(131,5)",
    "stavka2": "ТБ(130,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(131,5)",
    "stavka2": "ТБ(131)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(131,5)",
    "stavka2": "ТБ(131)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(131,5)",
    "stavka2": "ТБ(131,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(132)",
    "stavka2": "ТБ(130,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(132)",
    "stavka2": "ТБ(131)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(132)",
    "stavka2": "ТБ(131,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(132)",
    "stavka2": "ТБ(131,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(132)",
    "stavka2": "ТБ(132)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(132,5)",
    "stavka2": "ТБ(129,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(132,5)",
    "stavka2": "ТБ(131)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(132,5)",
    "stavka2": "ТБ(131,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(132,5)",
    "stavka2": "ТБ(132)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(132,5)",
    "stavka2": "ТБ(132)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(132,5)",
    "stavka2": "ТБ(132,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(133)",
    "stavka2": "ТБ(130,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(133)",
    "stavka2": "ТБ(131,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(133)",
    "stavka2": "ТБ(132)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(133)",
    "stavka2": "ТБ(132,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(133)",
    "stavka2": "ТБ(132,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(133)",
    "stavka2": "ТБ(133)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(133,5)",
    "stavka2": "ТБ(131,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(133,5)",
    "stavka2": "ТБ(132,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(133,5)",
    "stavka2": "ТБ(133)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(133,5)",
    "stavka2": "ТБ(133)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(133,5)",
    "stavka2": "ТБ(133,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(134)",
    "stavka2": "ТБ(131,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134)",
    "stavka2": "ТБ(132,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134)",
    "stavka2": "ТБ(133)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134)",
    "stavka2": "ТБ(133,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134)",
    "stavka2": "ТБ(134)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(134,5)",
    "stavka2": "ТБ(132,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134,5)",
    "stavka2": "ТБ(133)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134,5)",
    "stavka2": "ТБ(133,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134,5)",
    "stavka2": "ТБ(134)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134,5)",
    "stavka2": "ТБ(134,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(134,5)",
    "stavka2": "ТБ(134,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(135)",
    "stavka2": "ТБ(133)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135)",
    "stavka2": "ТБ(133,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135)",
    "stavka2": "ТБ(134)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135)",
    "stavka2": "ТБ(134)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(135)",
    "stavka2": "ТБ(134,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135)",
    "stavka2": "ТБ(134,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(135)",
    "stavka2": "ТБ(135)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(135,5)",
    "stavka2": "ТБ(132,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135,5)",
    "stavka2": "ТБ(133,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135,5)",
    "stavka2": "ТБ(134)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135,5)",
    "stavka2": "ТБ(134,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135,5)",
    "stavka2": "ТБ(135)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(135,5)",
    "stavka2": "ТБ(135)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(135,5)",
    "stavka2": "ТБ(135,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(136)",
    "stavka2": "ТБ(133,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(136)",
    "stavka2": "ТБ(134,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(136)",
    "stavka2": "ТБ(135)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(136)",
    "stavka2": "ТБ(135,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(136)",
    "stavka2": "ТБ(135,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(136)",
    "stavka2": "ТБ(136)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(136,5)",
    "stavka2": "ТБ(133,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(136,5)",
    "stavka2": "ТБ(134,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(136,5)",
    "stavka2": "ТБ(135,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(136,5)",
    "stavka2": "ТБ(135,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(136,5)",
    "stavka2": "ТБ(136)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(136,5)",
    "stavka2": "ТБ(136)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(136,5)",
    "stavka2": "ТБ(136,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(137)",
    "stavka2": "ТБ(135,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(137)",
    "stavka2": "ТБ(136)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(137)",
    "stavka2": "ТБ(136,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(137)",
    "stavka2": "ТБ(137)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(137,5)",
    "stavka2": "ТБ(135,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(137,5)",
    "stavka2": "ТБ(136)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(137,5)",
    "stavka2": "ТБ(136,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(137,5)",
    "stavka2": "ТБ(137)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(137,5)",
    "stavka2": "ТБ(137,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(138)",
    "stavka2": "ТБ(136,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(138)",
    "stavka2": "ТБ(137)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(138)",
    "stavka2": "ТБ(137)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(138)",
    "stavka2": "ТБ(137,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(138)",
    "stavka2": "ТБ(137,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(138)",
    "stavka2": "ТБ(138)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(138,5)",
    "stavka2": "ТБ(136)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(138,5)",
    "stavka2": "ТБ(136,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(138,5)",
    "stavka2": "ТБ(136,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(138,5)",
    "stavka2": "ТБ(137)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(138,5)",
    "stavka2": "ТБ(137,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(138,5)",
    "stavka2": "ТБ(138)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(138,5)",
    "stavka2": "ТБ(138,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(139)",
    "stavka2": "ТБ(136,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139)",
    "stavka2": "ТБ(137)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139)",
    "stavka2": "ТБ(137,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139)",
    "stavka2": "ТБ(138)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139)",
    "stavka2": "ТБ(138)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(139)",
    "stavka2": "ТБ(138,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139)",
    "stavka2": "ТБ(138,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(139)",
    "stavka2": "ТБ(139)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(139,5)",
    "stavka2": "ТБ(136,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139,5)",
    "stavka2": "ТБ(137)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139,5)",
    "stavka2": "ТБ(137)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(139,5)",
    "stavka2": "ТБ(137,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139,5)",
    "stavka2": "ТБ(138)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139,5)",
    "stavka2": "ТБ(138,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139,5)",
    "stavka2": "ТБ(139)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(139,5)",
    "stavka2": "ТБ(139,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(14)",
    "stavka2": "ТБ(14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(14,5)",
    "stavka2": "ТБ(13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140)",
    "stavka2": "ТБ(137,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140)",
    "stavka2": "ТБ(138)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140)",
    "stavka2": "ТБ(138,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140)",
    "stavka2": "ТБ(139)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140)",
    "stavka2": "ТБ(139,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140)",
    "stavka2": "ТБ(139,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(140)",
    "stavka2": "ТБ(140)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(140,5)",
    "stavka2": "ТБ(138,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140,5)",
    "stavka2": "ТБ(139)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140,5)",
    "stavka2": "ТБ(139,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140,5)",
    "stavka2": "ТБ(139,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(140,5)",
    "stavka2": "ТБ(140)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(140,5)",
    "stavka2": "ТБ(140,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(141)",
    "stavka2": "ТБ(138,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141)",
    "stavka2": "ТБ(139)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141)",
    "stavka2": "ТБ(139,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141)",
    "stavka2": "ТБ(140)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141)",
    "stavka2": "ТБ(140,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141)",
    "stavka2": "ТБ(140,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(141)",
    "stavka2": "ТБ(141)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(141,5)",
    "stavka2": "ТБ(139,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141,5)",
    "stavka2": "ТБ(140)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141,5)",
    "stavka2": "ТБ(140,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141,5)",
    "stavka2": "ТБ(141)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(141,5)",
    "stavka2": "ТБ(141,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(142)",
    "stavka2": "ТБ(139,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142)",
    "stavka2": "ТБ(140,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142)",
    "stavka2": "ТБ(141)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142)",
    "stavka2": "ТБ(141,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142)",
    "stavka2": "ТБ(142)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(142,5)",
    "stavka2": "ТБ(139)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142,5)",
    "stavka2": "ТБ(140)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142,5)",
    "stavka2": "ТБ(141)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142,5)",
    "stavka2": "ТБ(141,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142,5)",
    "stavka2": "ТБ(142)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(142,5)",
    "stavka2": "ТБ(142)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(142,5)",
    "stavka2": "ТБ(142,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(143)",
    "stavka2": "ТБ(139,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143)",
    "stavka2": "ТБ(140)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143)",
    "stavka2": "ТБ(141)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143)",
    "stavka2": "ТБ(141,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143)",
    "stavka2": "ТБ(142)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143)",
    "stavka2": "ТБ(142,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143)",
    "stavka2": "ТБ(142,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(143)",
    "stavka2": "ТБ(143)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(143,5)",
    "stavka2": "ТБ(140,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143,5)",
    "stavka2": "ТБ(141,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143,5)",
    "stavka2": "ТБ(142)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143,5)",
    "stavka2": "ТБ(142,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143,5)",
    "stavka2": "ТБ(142,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(143,5)",
    "stavka2": "ТБ(143)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(143,5)",
    "stavka2": "ТБ(143,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(144)",
    "stavka2": "ТБ(142)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144)",
    "stavka2": "ТБ(142,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144)",
    "stavka2": "ТБ(143)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144)",
    "stavka2": "ТБ(143,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144)",
    "stavka2": "ТБ(143,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(144)",
    "stavka2": "ТБ(144)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(144,5)",
    "stavka2": "ТБ(142)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144,5)",
    "stavka2": "ТБ(142)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(144,5)",
    "stavka2": "ТБ(142,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144,5)",
    "stavka2": "ТБ(143)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144,5)",
    "stavka2": "ТБ(143,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144,5)",
    "stavka2": "ТБ(143,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(144,5)",
    "stavka2": "ТБ(144)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(144,5)",
    "stavka2": "ТБ(144,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(145)",
    "stavka2": "ТБ(142,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145)",
    "stavka2": "ТБ(143)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145)",
    "stavka2": "ТБ(143,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145)",
    "stavka2": "ТБ(144)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145)",
    "stavka2": "ТБ(144,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145)",
    "stavka2": "ТБ(145)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(145,5)",
    "stavka2": "ТБ(143)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145,5)",
    "stavka2": "ТБ(143,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145,5)",
    "stavka2": "ТБ(144)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145,5)",
    "stavka2": "ТБ(144,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145,5)",
    "stavka2": "ТБ(145)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(145,5)",
    "stavka2": "ТБ(145,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(146)",
    "stavka2": "ТБ(144,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146)",
    "stavka2": "ТБ(145)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146)",
    "stavka2": "ТБ(145)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(146)",
    "stavka2": "ТБ(145,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146)",
    "stavka2": "ТБ(145,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(146)",
    "stavka2": "ТБ(146)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(143,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(144)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(144,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(145)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(145,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(145,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(146)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(146)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(146,5)",
    "stavka2": "ТБ(146,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(147)",
    "stavka2": "ТБ(145,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147)",
    "stavka2": "ТБ(146)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147)",
    "stavka2": "ТБ(146,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147)",
    "stavka2": "ТБ(146,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(147)",
    "stavka2": "ТБ(147)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(147,5)",
    "stavka2": "ТБ(143,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147,5)",
    "stavka2": "ТБ(144)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147,5)",
    "stavka2": "ТБ(145,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147,5)",
    "stavka2": "ТБ(146)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147,5)",
    "stavka2": "ТБ(146,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147,5)",
    "stavka2": "ТБ(146,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(147,5)",
    "stavka2": "ТБ(147)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(147,5)",
    "stavka2": "ТБ(147,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(148)",
    "stavka2": "ТБ(145,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148)",
    "stavka2": "ТБ(146,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148)",
    "stavka2": "ТБ(147)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148)",
    "stavka2": "ТБ(147,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148)",
    "stavka2": "ТБ(147,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(148)",
    "stavka2": "ТБ(148)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(145)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(145,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(146)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(146,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(147)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(147,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(148)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(148)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(148,5)",
    "stavka2": "ТБ(148,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(149)",
    "stavka2": "ТБ(147)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149)",
    "stavka2": "ТБ(147,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149)",
    "stavka2": "ТБ(148)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149)",
    "stavka2": "ТБ(148,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149)",
    "stavka2": "ТБ(148,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(149)",
    "stavka2": "ТБ(149)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(146,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(146,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(147)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(147,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(148)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(148,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(149)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(149)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(149,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(149,5)",
    "stavka2": "ТБ(149,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(15,5)",
    "stavka2": "ТБ(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(150)",
    "stavka2": "ТБ(146,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150)",
    "stavka2": "ТБ(147,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150)",
    "stavka2": "ТБ(148)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150)",
    "stavka2": "ТБ(148,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150)",
    "stavka2": "ТБ(149)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150)",
    "stavka2": "ТБ(149,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150)",
    "stavka2": "ТБ(149,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(150)",
    "stavka2": "ТБ(150)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(150,5)",
    "stavka2": "ТБ(146,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150,5)",
    "stavka2": "ТБ(147)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150,5)",
    "stavka2": "ТБ(148,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150,5)",
    "stavka2": "ТБ(149)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150,5)",
    "stavka2": "ТБ(149,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150,5)",
    "stavka2": "ТБ(150)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(150,5)",
    "stavka2": "ТБ(150)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(150,5)",
    "stavka2": "ТБ(150,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(151)",
    "stavka2": "ТБ(148,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151)",
    "stavka2": "ТБ(149)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151)",
    "stavka2": "ТБ(149,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151)",
    "stavka2": "ТБ(150)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151)",
    "stavka2": "ТБ(150,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151)",
    "stavka2": "ТБ(150,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(151)",
    "stavka2": "ТБ(151)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(151,5)",
    "stavka2": "ТБ(149,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151,5)",
    "stavka2": "ТБ(150)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151,5)",
    "stavka2": "ТБ(150,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151,5)",
    "stavka2": "ТБ(150,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(151,5)",
    "stavka2": "ТБ(151)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(151,5)",
    "stavka2": "ТБ(151,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(152)",
    "stavka2": "ТБ(150)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152)",
    "stavka2": "ТБ(150)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(152)",
    "stavka2": "ТБ(150,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152)",
    "stavka2": "ТБ(151)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152)",
    "stavka2": "ТБ(151,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152)",
    "stavka2": "ТБ(151,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(152)",
    "stavka2": "ТБ(152)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(149,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(150)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(150,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(151)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(151,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(151,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(152)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(152)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(152,5)",
    "stavka2": "ТБ(152,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(153)",
    "stavka2": "ТБ(149,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153)",
    "stavka2": "ТБ(150,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153)",
    "stavka2": "ТБ(151)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153)",
    "stavka2": "ТБ(151,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153)",
    "stavka2": "ТБ(152)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153)",
    "stavka2": "ТБ(152,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153)",
    "stavka2": "ТБ(153)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(153,5)",
    "stavka2": "ТБ(149,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153,5)",
    "stavka2": "ТБ(151,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153,5)",
    "stavka2": "ТБ(152)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153,5)",
    "stavka2": "ТБ(152,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153,5)",
    "stavka2": "ТБ(153)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(153,5)",
    "stavka2": "ТБ(153)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(153,5)",
    "stavka2": "ТБ(153,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(154)",
    "stavka2": "ТБ(152,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(154)",
    "stavka2": "ТБ(153)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(154)",
    "stavka2": "ТБ(153,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(154)",
    "stavka2": "ТБ(153,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(154)",
    "stavka2": "ТБ(154)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(154,5)",
    "stavka2": "ТБ(152,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(154,5)",
    "stavka2": "ТБ(153)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(154,5)",
    "stavka2": "ТБ(153,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(154,5)",
    "stavka2": "ТБ(154)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(154,5)",
    "stavka2": "ТБ(154)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(154,5)",
    "stavka2": "ТБ(154,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(155)",
    "stavka2": "ТБ(152,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(155)",
    "stavka2": "ТБ(153,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(155)",
    "stavka2": "ТБ(154,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(155)",
    "stavka2": "ТБ(154,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(155)",
    "stavka2": "ТБ(155)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(153)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(153,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(154)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(154)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(154,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(154,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(155)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(155)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(155,5)",
    "stavka2": "ТБ(155,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(156)",
    "stavka2": "ТБ(153,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156)",
    "stavka2": "ТБ(154)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156)",
    "stavka2": "ТБ(154,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156)",
    "stavka2": "ТБ(155)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156)",
    "stavka2": "ТБ(155)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(156)",
    "stavka2": "ТБ(155,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156)",
    "stavka2": "ТБ(156)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(153,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(154)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(154,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(155)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(155,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(155,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(156)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(156)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(156,5)",
    "stavka2": "ТБ(156,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(157)",
    "stavka2": "ТБ(155,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(157)",
    "stavka2": "ТБ(156)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(157)",
    "stavka2": "ТБ(156,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(157)",
    "stavka2": "ТБ(156,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(157)",
    "stavka2": "ТБ(157)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(157,5)",
    "stavka2": "ТБ(155,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(157,5)",
    "stavka2": "ТБ(155,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(157,5)",
    "stavka2": "ТБ(156)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(157,5)",
    "stavka2": "ТБ(156,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(157,5)",
    "stavka2": "ТБ(157)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(157,5)",
    "stavka2": "ТБ(157)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(157,5)",
    "stavka2": "ТБ(157,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(157,5)",
    "stavka2": "ТБ(157,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(158)",
    "stavka2": "ТБ(155,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158)",
    "stavka2": "ТБ(156,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158)",
    "stavka2": "ТБ(157)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158)",
    "stavka2": "ТБ(157)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(158)",
    "stavka2": "ТБ(157,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158)",
    "stavka2": "ТБ(157,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(158)",
    "stavka2": "ТБ(158)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(158,5)",
    "stavka2": "ТБ(156,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158,5)",
    "stavka2": "ТБ(157)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158,5)",
    "stavka2": "ТБ(157,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158,5)",
    "stavka2": "ТБ(158)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158,5)",
    "stavka2": "ТБ(158)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(158,5)",
    "stavka2": "ТБ(158,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(158,5)",
    "stavka2": "ТБ(158,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(159)",
    "stavka2": "ТБ(156)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159)",
    "stavka2": "ТБ(157,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159)",
    "stavka2": "ТБ(158)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159)",
    "stavka2": "ТБ(158,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159)",
    "stavka2": "ТБ(159)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(159,5)",
    "stavka2": "ТБ(157)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159,5)",
    "stavka2": "ТБ(157,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159,5)",
    "stavka2": "ТБ(158)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159,5)",
    "stavka2": "ТБ(158,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159,5)",
    "stavka2": "ТБ(159)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(159,5)",
    "stavka2": "ТБ(159)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(159,5)",
    "stavka2": "ТБ(159,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(160)",
    "stavka2": "ТБ(158,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(160)",
    "stavka2": "ТБ(159)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(160)",
    "stavka2": "ТБ(159,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(160)",
    "stavka2": "ТБ(159,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(160)",
    "stavka2": "ТБ(160)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(160,5)",
    "stavka2": "ТБ(158,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(160,5)",
    "stavka2": "ТБ(159)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(160,5)",
    "stavka2": "ТБ(159,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(160,5)",
    "stavka2": "ТБ(159,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(160,5)",
    "stavka2": "ТБ(160)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(160,5)",
    "stavka2": "ТБ(160,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(161)",
    "stavka2": "ТБ(157,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161)",
    "stavka2": "ТБ(158,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161)",
    "stavka2": "ТБ(160)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161)",
    "stavka2": "ТБ(160,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161)",
    "stavka2": "ТБ(160,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(161)",
    "stavka2": "ТБ(161)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(161,5)",
    "stavka2": "ТБ(157,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161,5)",
    "stavka2": "ТБ(159,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161,5)",
    "stavka2": "ТБ(160)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161,5)",
    "stavka2": "ТБ(160,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161,5)",
    "stavka2": "ТБ(161)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(161,5)",
    "stavka2": "ТБ(161)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(161,5)",
    "stavka2": "ТБ(161,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(162)",
    "stavka2": "ТБ(159,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(162)",
    "stavka2": "ТБ(160,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(162)",
    "stavka2": "ТБ(161)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(162)",
    "stavka2": "ТБ(161,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(162)",
    "stavka2": "ТБ(161,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(162)",
    "stavka2": "ТБ(162)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(162,5)",
    "stavka2": "ТБ(160,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(162,5)",
    "stavka2": "ТБ(161)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(162,5)",
    "stavka2": "ТБ(161,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(162,5)",
    "stavka2": "ТБ(162)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(162,5)",
    "stavka2": "ТБ(162)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(162,5)",
    "stavka2": "ТБ(162,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(163)",
    "stavka2": "ТБ(160,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163)",
    "stavka2": "ТБ(162)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163)",
    "stavka2": "ТБ(162)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(163)",
    "stavka2": "ТБ(162,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163)",
    "stavka2": "ТБ(162,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(163)",
    "stavka2": "ТБ(163)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(159,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(160,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(161,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(162)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(162,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(162,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(163)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(163)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(163,5)",
    "stavka2": "ТБ(163,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(164)",
    "stavka2": "ТБ(161)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164)",
    "stavka2": "ТБ(162,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164)",
    "stavka2": "ТБ(163)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164)",
    "stavka2": "ТБ(163,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164)",
    "stavka2": "ТБ(163,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(164)",
    "stavka2": "ТБ(164)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(164,5)",
    "stavka2": "ТБ(162)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164,5)",
    "stavka2": "ТБ(162,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164,5)",
    "stavka2": "ТБ(163)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164,5)",
    "stavka2": "ТБ(163,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164,5)",
    "stavka2": "ТБ(163,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(164,5)",
    "stavka2": "ТБ(164)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(164,5)",
    "stavka2": "ТБ(164)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(164,5)",
    "stavka2": "ТБ(164,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(165)",
    "stavka2": "ТБ(162,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165)",
    "stavka2": "ТБ(163,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165)",
    "stavka2": "ТБ(164)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165)",
    "stavka2": "ТБ(164,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165)",
    "stavka2": "ТБ(164,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(165)",
    "stavka2": "ТБ(165)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(165,5)",
    "stavka2": "ТБ(162,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165,5)",
    "stavka2": "ТБ(163)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165,5)",
    "stavka2": "ТБ(163,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165,5)",
    "stavka2": "ТБ(164,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165,5)",
    "stavka2": "ТБ(165)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(165,5)",
    "stavka2": "ТБ(165)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(165,5)",
    "stavka2": "ТБ(165,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(166)",
    "stavka2": "ТБ(163,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166)",
    "stavka2": "ТБ(163,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(166)",
    "stavka2": "ТБ(164,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166)",
    "stavka2": "ТБ(165)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166)",
    "stavka2": "ТБ(165)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(166)",
    "stavka2": "ТБ(165,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166)",
    "stavka2": "ТБ(165,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(166)",
    "stavka2": "ТБ(166)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(166,5)",
    "stavka2": "ТБ(162,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166,5)",
    "stavka2": "ТБ(164)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166,5)",
    "stavka2": "ТБ(164,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166,5)",
    "stavka2": "ТБ(165)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166,5)",
    "stavka2": "ТБ(165,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166,5)",
    "stavka2": "ТБ(165,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(166,5)",
    "stavka2": "ТБ(166)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(166,5)",
    "stavka2": "ТБ(166,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(167)",
    "stavka2": "ТБ(165)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(167)",
    "stavka2": "ТБ(165,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(167)",
    "stavka2": "ТБ(166,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(167)",
    "stavka2": "ТБ(166,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(167)",
    "stavka2": "ТБ(167)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(167,5)",
    "stavka2": "ТБ(166)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(167,5)",
    "stavka2": "ТБ(166,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(167,5)",
    "stavka2": "ТБ(166,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(167,5)",
    "stavka2": "ТБ(167)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(167,5)",
    "stavka2": "ТБ(167)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(167,5)",
    "stavka2": "ТБ(167,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(168)",
    "stavka2": "ТБ(164,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(168)",
    "stavka2": "ТБ(166,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(168)",
    "stavka2": "ТБ(167,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(168)",
    "stavka2": "ТБ(167,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(168)",
    "stavka2": "ТБ(168)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(168,5)",
    "stavka2": "ТБ(166,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(168,5)",
    "stavka2": "ТБ(167)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(168,5)",
    "stavka2": "ТБ(167,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(168,5)",
    "stavka2": "ТБ(168)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(168,5)",
    "stavka2": "ТБ(168)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(168,5)",
    "stavka2": "ТБ(168,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(168,5)",
    "stavka2": "ТБ(168,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(169)",
    "stavka2": "ТБ(167,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169)",
    "stavka2": "ТБ(168)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169)",
    "stavka2": "ТБ(168,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169)",
    "stavka2": "ТБ(168,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(169)",
    "stavka2": "ТБ(169)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(166)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(166,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(167)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(167,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(168)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(168,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(169)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(169)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(169,5)",
    "stavka2": "ТБ(169,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(17)",
    "stavka2": "ТБ(17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(17,5)",
    "stavka2": "ТБ(17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(170)",
    "stavka2": "ТБ(168,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170)",
    "stavka2": "ТБ(169)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170)",
    "stavka2": "ТБ(169,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170)",
    "stavka2": "ТБ(170)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(170,5)",
    "stavka2": "ТБ(167)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170,5)",
    "stavka2": "ТБ(167,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170,5)",
    "stavka2": "ТБ(168)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170,5)",
    "stavka2": "ТБ(168,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170,5)",
    "stavka2": "ТБ(169)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170,5)",
    "stavka2": "ТБ(169,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170,5)",
    "stavka2": "ТБ(170)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(170,5)",
    "stavka2": "ТБ(170,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(171)",
    "stavka2": "ТБ(168)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171)",
    "stavka2": "ТБ(169,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171)",
    "stavka2": "ТБ(170)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171)",
    "stavka2": "ТБ(170,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171)",
    "stavka2": "ТБ(171)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(171,5)",
    "stavka2": "ТБ(169)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171,5)",
    "stavka2": "ТБ(169,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171,5)",
    "stavka2": "ТБ(170)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171,5)",
    "stavka2": "ТБ(170,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171,5)",
    "stavka2": "ТБ(171)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(171,5)",
    "stavka2": "ТБ(171)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(171,5)",
    "stavka2": "ТБ(171,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(172)",
    "stavka2": "ТБ(170,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(172)",
    "stavka2": "ТБ(171,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(172)",
    "stavka2": "ТБ(171,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(172)",
    "stavka2": "ТБ(172)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(172,5)",
    "stavka2": "ТБ(170,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(172,5)",
    "stavka2": "ТБ(171,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(172,5)",
    "stavka2": "ТБ(172)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(172,5)",
    "stavka2": "ТБ(172)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(172,5)",
    "stavka2": "ТБ(172,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(173)",
    "stavka2": "ТБ(171,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173)",
    "stavka2": "ТБ(171,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(173)",
    "stavka2": "ТБ(172)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173)",
    "stavka2": "ТБ(172,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173)",
    "stavka2": "ТБ(173)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(169,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(170)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(170,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(171)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(172)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(172,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(173)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(173)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(173,5)",
    "stavka2": "ТБ(173,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(174)",
    "stavka2": "ТБ(171,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(174)",
    "stavka2": "ТБ(172,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(174)",
    "stavka2": "ТБ(173)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(174)",
    "stavka2": "ТБ(173,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(174)",
    "stavka2": "ТБ(173,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(174)",
    "stavka2": "ТБ(174)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(174,5)",
    "stavka2": "ТБ(173)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(174,5)",
    "stavka2": "ТБ(173)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(174,5)",
    "stavka2": "ТБ(173,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(174,5)",
    "stavka2": "ТБ(173,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(174,5)",
    "stavka2": "ТБ(174)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(174,5)",
    "stavka2": "ТБ(174)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(174,5)",
    "stavka2": "ТБ(174,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(175)",
    "stavka2": "ТБ(173,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(175)",
    "stavka2": "ТБ(174)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(175)",
    "stavka2": "ТБ(174,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(175)",
    "stavka2": "ТБ(174,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(175)",
    "stavka2": "ТБ(175)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(175,5)",
    "stavka2": "ТБ(173,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(175,5)",
    "stavka2": "ТБ(174)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(175,5)",
    "stavka2": "ТБ(174,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(175,5)",
    "stavka2": "ТБ(174,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(175,5)",
    "stavka2": "ТБ(175)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(175,5)",
    "stavka2": "ТБ(175,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(176)",
    "stavka2": "ТБ(174,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(176)",
    "stavka2": "ТБ(175)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(176)",
    "stavka2": "ТБ(175,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(176)",
    "stavka2": "ТБ(176)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(176,5)",
    "stavka2": "ТБ(174,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(176,5)",
    "stavka2": "ТБ(175)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(176,5)",
    "stavka2": "ТБ(175,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(176,5)",
    "stavka2": "ТБ(176)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(176,5)",
    "stavka2": "ТБ(176)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(176,5)",
    "stavka2": "ТБ(176,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(177)",
    "stavka2": "ТБ(175)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(177)",
    "stavka2": "ТБ(175,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(177)",
    "stavka2": "ТБ(176,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(177)",
    "stavka2": "ТБ(177)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(177,5)",
    "stavka2": "ТБ(176,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(177,5)",
    "stavka2": "ТБ(177)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(177,5)",
    "stavka2": "ТБ(177,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(178)",
    "stavka2": "ТБ(176,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(178)",
    "stavka2": "ТБ(177,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(178)",
    "stavka2": "ТБ(178)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(178,5)",
    "stavka2": "ТБ(177)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(178,5)",
    "stavka2": "ТБ(177,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(178,5)",
    "stavka2": "ТБ(178)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(178,5)",
    "stavka2": "ТБ(178,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(179)",
    "stavka2": "ТБ(177,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(179)",
    "stavka2": "ТБ(178,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(179)",
    "stavka2": "ТБ(179)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(179,5)",
    "stavka2": "ТБ(176,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(179,5)",
    "stavka2": "ТБ(177,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(179,5)",
    "stavka2": "ТБ(178)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(179,5)",
    "stavka2": "ТБ(178,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(179,5)",
    "stavka2": "ТБ(178,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(179,5)",
    "stavka2": "ТБ(179)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(179,5)",
    "stavka2": "ТБ(179)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(179,5)",
    "stavka2": "ТБ(179,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(18)",
    "stavka2": "ТБ(18)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(18,5)",
    "stavka2": "ТБ(18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(180)",
    "stavka2": "ТБ(176,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(180)",
    "stavka2": "ТБ(178,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(180)",
    "stavka2": "ТБ(179,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(180)",
    "stavka2": "ТБ(180)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(180,5)",
    "stavka2": "ТБ(179,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(180,5)",
    "stavka2": "ТБ(180)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(180,5)",
    "stavka2": "ТБ(180,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(181)",
    "stavka2": "ТБ(179,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(181)",
    "stavka2": "ТБ(180,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(181)",
    "stavka2": "ТБ(181)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(181,5)",
    "stavka2": "ТБ(179,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(181,5)",
    "stavka2": "ТБ(180,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(181,5)",
    "stavka2": "ТБ(181)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(181,5)",
    "stavka2": "ТБ(181,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(182)",
    "stavka2": "ТБ(181,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(182)",
    "stavka2": "ТБ(181,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(182)",
    "stavka2": "ТБ(182)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(182,5)",
    "stavka2": "ТБ(180,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(182,5)",
    "stavka2": "ТБ(182)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(182,5)",
    "stavka2": "ТБ(182,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(183)",
    "stavka2": "ТБ(181,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(183)",
    "stavka2": "ТБ(182,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(183)",
    "stavka2": "ТБ(183)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(183,5)",
    "stavka2": "ТБ(182,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(183,5)",
    "stavka2": "ТБ(183)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(183,5)",
    "stavka2": "ТБ(183,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(184)",
    "stavka2": "ТБ(183)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(184)",
    "stavka2": "ТБ(183,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(184)",
    "stavka2": "ТБ(184)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(184,5)",
    "stavka2": "ТБ(182,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(184,5)",
    "stavka2": "ТБ(183,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(184,5)",
    "stavka2": "ТБ(184)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(184,5)",
    "stavka2": "ТБ(184,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(185)",
    "stavka2": "ТБ(184,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(185)",
    "stavka2": "ТБ(185)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(185,5)",
    "stavka2": "ТБ(185)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(185,5)",
    "stavka2": "ТБ(185,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(186)",
    "stavka2": "ТБ(185,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(186)",
    "stavka2": "ТБ(186)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(186,5)",
    "stavka2": "ТБ(185,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(186,5)",
    "stavka2": "ТБ(186)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(186,5)",
    "stavka2": "ТБ(186,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(187)",
    "stavka2": "ТБ(185,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(187)",
    "stavka2": "ТБ(186,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(187,5)",
    "stavka2": "ТБ(186,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(187,5)",
    "stavka2": "ТБ(186,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(187,5)",
    "stavka2": "ТБ(187)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(187,5)",
    "stavka2": "ТБ(187,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(188)",
    "stavka2": "ТБ(184,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(188)",
    "stavka2": "ТБ(187,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(188)",
    "stavka2": "ТБ(188)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(188,5)",
    "stavka2": "ТБ(187,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(188,5)",
    "stavka2": "ТБ(188)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(188,5)",
    "stavka2": "ТБ(188,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(189)",
    "stavka2": "ТБ(189)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(189,5)",
    "stavka2": "ТБ(189)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(189,5)",
    "stavka2": "ТБ(189,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(19)",
    "stavka2": "ТБ(19)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(19,5)",
    "stavka2": "ТБ(19,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(190)",
    "stavka2": "ТБ(190)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(190,5)",
    "stavka2": "ТБ(189,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(190,5)",
    "stavka2": "ТБ(190)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(190,5)",
    "stavka2": "ТБ(190,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(191)",
    "stavka2": "ТБ(190)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(191)",
    "stavka2": "ТБ(190,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(191,5)",
    "stavka2": "ТБ(189,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(191,5)",
    "stavka2": "ТБ(190)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(191,5)",
    "stavka2": "ТБ(191)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(191,5)",
    "stavka2": "ТБ(191,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(192)",
    "stavka2": "ТБ(192)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(192,5)",
    "stavka2": "ТБ(192)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(192,5)",
    "stavka2": "ТБ(192,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(193)",
    "stavka2": "ТБ(192,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(193,5)",
    "stavka2": "ТБ(193)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(193,5)",
    "stavka2": "ТБ(193,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(194)",
    "stavka2": "ТБ(193,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(194)",
    "stavka2": "ТБ(194)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(194,5)",
    "stavka2": "ТБ(194,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(195)",
    "stavka2": "ТБ(194,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(195)",
    "stavka2": "ТБ(195)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(195,5)",
    "stavka2": "ТБ(194,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(195,5)",
    "stavka2": "ТБ(195)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(195,5)",
    "stavka2": "ТБ(195,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(196)",
    "stavka2": "ТБ(195,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(196)",
    "stavka2": "ТБ(196)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(196,5)",
    "stavka2": "ТБ(192,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(196,5)",
    "stavka2": "ТБ(193)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(196,5)",
    "stavka2": "ТБ(194,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(196,5)",
    "stavka2": "ТБ(195,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(196,5)",
    "stavka2": "ТБ(196)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(196,5)",
    "stavka2": "ТБ(196,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(197)",
    "stavka2": "ТБ(196,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(197,5)",
    "stavka2": "ТБ(194,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(197,5)",
    "stavka2": "ТБ(197)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(197,5)",
    "stavka2": "ТБ(197,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(198,5)",
    "stavka2": "ТБ(196,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(198,5)",
    "stavka2": "ТБ(198,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(199,5)",
    "stavka2": "ТБ(199)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(199,5)",
    "stavka2": "ТБ(199,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(2)",
    "stavka2": "ТБ(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(2)",
    "stavka2": "ТБ(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(2,25)",
    "stavka2": "ТБ(2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(2,5)",
    "stavka2": "ТБ(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(2,75)",
    "stavka2": "ТБ(2,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(20)",
    "stavka2": "ТБ(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(20)",
    "stavka2": "ТБ(20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(20,5)",
    "stavka2": "ТБ(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(20,5)",
    "stavka2": "ТБ(20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(200)",
    "stavka2": "ТБ(199,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(200,5)",
    "stavka2": "ТБ(197,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(200,5)",
    "stavka2": "ТБ(199)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(200,5)",
    "stavka2": "ТБ(200)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(200,5)",
    "stavka2": "ТБ(200,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(201)",
    "stavka2": "ТБ(200,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(201,5)",
    "stavka2": "ТБ(200,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(201,5)",
    "stavka2": "ТБ(201)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(201,5)",
    "stavka2": "ТБ(201,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(202)",
    "stavka2": "ТБ(198,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(202)",
    "stavka2": "ТБ(200,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(202)",
    "stavka2": "ТБ(201,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(202)",
    "stavka2": "ТБ(202)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(202,5)",
    "stavka2": "ТБ(201,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(202,5)",
    "stavka2": "ТБ(202)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(202,5)",
    "stavka2": "ТБ(202,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(203)",
    "stavka2": "ТБ(202,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(203)",
    "stavka2": "ТБ(203)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(203,5)",
    "stavka2": "ТБ(202,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(203,5)",
    "stavka2": "ТБ(203,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(204)",
    "stavka2": "ТБ(203,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(204)",
    "stavka2": "ТБ(204)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(204,5)",
    "stavka2": "ТБ(204)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(204,5)",
    "stavka2": "ТБ(204,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(205)",
    "stavka2": "ТБ(204,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(205)",
    "stavka2": "ТБ(204,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(205)",
    "stavka2": "ТБ(205)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(205,5)",
    "stavka2": "ТБ(201,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(205,5)",
    "stavka2": "ТБ(204,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(205,5)",
    "stavka2": "ТБ(205)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(205,5)",
    "stavka2": "ТБ(205,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(206)",
    "stavka2": "ТБ(205,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(206)",
    "stavka2": "ТБ(206)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(206,5)",
    "stavka2": "ТБ(204,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(206,5)",
    "stavka2": "ТБ(205)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(206,5)",
    "stavka2": "ТБ(205,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(206,5)",
    "stavka2": "ТБ(206)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(206,5)",
    "stavka2": "ТБ(206,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(207)",
    "stavka2": "ТБ(205,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(207)",
    "stavka2": "ТБ(206,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(207)",
    "stavka2": "ТБ(207)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(207,5)",
    "stavka2": "ТБ(206)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(207,5)",
    "stavka2": "ТБ(206,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(207,5)",
    "stavka2": "ТБ(207)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(207,5)",
    "stavka2": "ТБ(207)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(207,5)",
    "stavka2": "ТБ(207,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(208)",
    "stavka2": "ТБ(207,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(208)",
    "stavka2": "ТБ(208)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(208,5)",
    "stavka2": "ТБ(207)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(208,5)",
    "stavka2": "ТБ(207,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(208,5)",
    "stavka2": "ТБ(208)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(208,5)",
    "stavka2": "ТБ(208,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(209)",
    "stavka2": "ТБ(206,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(209)",
    "stavka2": "ТБ(207,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(209)",
    "stavka2": "ТБ(208)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(209)",
    "stavka2": "ТБ(208,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(209)",
    "stavka2": "ТБ(209)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(209,5)",
    "stavka2": "ТБ(208,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(209,5)",
    "stavka2": "ТБ(209)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(209,5)",
    "stavka2": "ТБ(209,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(21)",
    "stavka2": "ТБ(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(21)",
    "stavka2": "ТБ(20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(21)",
    "stavka2": "ТБ(21)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(21,5)",
    "stavka2": "ТБ(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(21,5)",
    "stavka2": "ТБ(20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(21,5)",
    "stavka2": "ТБ(21)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(21,5)",
    "stavka2": "ТБ(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(210)",
    "stavka2": "ТБ(209)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(210)",
    "stavka2": "ТБ(209,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(210)",
    "stavka2": "ТБ(209,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(210)",
    "stavka2": "ТБ(210)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(210,5)",
    "stavka2": "ТБ(209)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(210,5)",
    "stavka2": "ТБ(209)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(210,5)",
    "stavka2": "ТБ(209,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(210,5)",
    "stavka2": "ТБ(210)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(210,5)",
    "stavka2": "ТБ(210,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(211)",
    "stavka2": "ТБ(210)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(211)",
    "stavka2": "ТБ(210,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(211)",
    "stavka2": "ТБ(210,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(211)",
    "stavka2": "ТБ(211)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(211,5)",
    "stavka2": "ТБ(210,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(211,5)",
    "stavka2": "ТБ(211)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(211,5)",
    "stavka2": "ТБ(211,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(212)",
    "stavka2": "ТБ(210,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(212)",
    "stavka2": "ТБ(211)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(212)",
    "stavka2": "ТБ(211,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(212)",
    "stavka2": "ТБ(212)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(212,5)",
    "stavka2": "ТБ(209,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(212,5)",
    "stavka2": "ТБ(210,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(212,5)",
    "stavka2": "ТБ(211)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(212,5)",
    "stavka2": "ТБ(211,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(212,5)",
    "stavka2": "ТБ(212)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(212,5)",
    "stavka2": "ТБ(212)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(212,5)",
    "stavka2": "ТБ(212,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(213)",
    "stavka2": "ТБ(211,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(213)",
    "stavka2": "ТБ(212)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(213)",
    "stavka2": "ТБ(212,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(213)",
    "stavka2": "ТБ(213)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(213,5)",
    "stavka2": "ТБ(212)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(213,5)",
    "stavka2": "ТБ(212,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(213,5)",
    "stavka2": "ТБ(213)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(213,5)",
    "stavka2": "ТБ(213)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(213,5)",
    "stavka2": "ТБ(213,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(214)",
    "stavka2": "ТБ(212,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(214)",
    "stavka2": "ТБ(213)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(214)",
    "stavka2": "ТБ(213,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(214)",
    "stavka2": "ТБ(213,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(214)",
    "stavka2": "ТБ(214)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(214,5)",
    "stavka2": "ТБ(213)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(214,5)",
    "stavka2": "ТБ(213,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(214,5)",
    "stavka2": "ТБ(214)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(214,5)",
    "stavka2": "ТБ(214,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(215)",
    "stavka2": "ТБ(214,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(215)",
    "stavka2": "ТБ(215)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(215,5)",
    "stavka2": "ТБ(214)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(215,5)",
    "stavka2": "ТБ(214,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(215,5)",
    "stavka2": "ТБ(215)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(215,5)",
    "stavka2": "ТБ(215)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(215,5)",
    "stavka2": "ТБ(215,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(216)",
    "stavka2": "ТБ(215)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(216)",
    "stavka2": "ТБ(215,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(216)",
    "stavka2": "ТБ(216)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(216,5)",
    "stavka2": "ТБ(214,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(216,5)",
    "stavka2": "ТБ(215)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(216,5)",
    "stavka2": "ТБ(215,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(216,5)",
    "stavka2": "ТБ(216)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(216,5)",
    "stavka2": "ТБ(216)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(216,5)",
    "stavka2": "ТБ(216,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(217)",
    "stavka2": "ТБ(216)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(217)",
    "stavka2": "ТБ(216)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(217)",
    "stavka2": "ТБ(216,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(217)",
    "stavka2": "ТБ(216,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(217)",
    "stavka2": "ТБ(217)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(217,5)",
    "stavka2": "ТБ(216)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(217,5)",
    "stavka2": "ТБ(216,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(217,5)",
    "stavka2": "ТБ(217)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(217,5)",
    "stavka2": "ТБ(217,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(218)",
    "stavka2": "ТБ(217)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(218)",
    "stavka2": "ТБ(217,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(218)",
    "stavka2": "ТБ(217,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(218)",
    "stavka2": "ТБ(218)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(218,5)",
    "stavka2": "ТБ(217,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(218,5)",
    "stavka2": "ТБ(218)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(218,5)",
    "stavka2": "ТБ(218)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(218,5)",
    "stavka2": "ТБ(218,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(219)",
    "stavka2": "ТБ(218)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(219)",
    "stavka2": "ТБ(218,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(219)",
    "stavka2": "ТБ(218,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(219)",
    "stavka2": "ТБ(219)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(219)",
    "stavka2": "ТБ(219)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(219,5)",
    "stavka2": "ТБ(218)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(219,5)",
    "stavka2": "ТБ(218,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(219,5)",
    "stavka2": "ТБ(218,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(219,5)",
    "stavka2": "ТБ(219)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(219,5)",
    "stavka2": "ТБ(219)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(219,5)",
    "stavka2": "ТБ(219,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(22)",
    "stavka2": "ТБ(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(22)",
    "stavka2": "ТБ(22)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(22,5)",
    "stavka2": "ТБ(22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(22,5)",
    "stavka2": "ТБ(22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(220)",
    "stavka2": "ТБ(219)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(220)",
    "stavka2": "ТБ(219,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(220)",
    "stavka2": "ТБ(219,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(220)",
    "stavka2": "ТБ(220)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(220,5)",
    "stavka2": "ТБ(219,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(220,5)",
    "stavka2": "ТБ(220)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(220,5)",
    "stavka2": "ТБ(220,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(221)",
    "stavka2": "ТБ(220)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(221)",
    "stavka2": "ТБ(220,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(221)",
    "stavka2": "ТБ(221)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(221,5)",
    "stavka2": "ТБ(220)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(221,5)",
    "stavka2": "ТБ(220,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(221,5)",
    "stavka2": "ТБ(221)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(221,5)",
    "stavka2": "ТБ(221)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(221,5)",
    "stavka2": "ТБ(221,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(222)",
    "stavka2": "ТБ(220,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(222)",
    "stavka2": "ТБ(220,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(222)",
    "stavka2": "ТБ(221)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(222)",
    "stavka2": "ТБ(221,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(222)",
    "stavka2": "ТБ(221,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(222)",
    "stavka2": "ТБ(222)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(222,5)",
    "stavka2": "ТБ(221,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(222,5)",
    "stavka2": "ТБ(222)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(222,5)",
    "stavka2": "ТБ(222)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(222,5)",
    "stavka2": "ТБ(222,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(223)",
    "stavka2": "ТБ(221,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(223)",
    "stavka2": "ТБ(222)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(223)",
    "stavka2": "ТБ(222,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(223)",
    "stavka2": "ТБ(222,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(223)",
    "stavka2": "ТБ(223)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(223)",
    "stavka2": "ТБ(223)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(223,5)",
    "stavka2": "ТБ(221,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(223,5)",
    "stavka2": "ТБ(222,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(223,5)",
    "stavka2": "ТБ(222,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(223,5)",
    "stavka2": "ТБ(223)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(223,5)",
    "stavka2": "ТБ(223)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(223,5)",
    "stavka2": "ТБ(223,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(224)",
    "stavka2": "ТБ(222,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(224)",
    "stavka2": "ТБ(223)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(224)",
    "stavka2": "ТБ(223,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(224)",
    "stavka2": "ТБ(224)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(224,5)",
    "stavka2": "ТБ(223,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(224,5)",
    "stavka2": "ТБ(224)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(224,5)",
    "stavka2": "ТБ(224)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(224,5)",
    "stavka2": "ТБ(224,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(225)",
    "stavka2": "ТБ(222,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225)",
    "stavka2": "ТБ(223,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225)",
    "stavka2": "ТБ(224)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225)",
    "stavka2": "ТБ(224)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(225)",
    "stavka2": "ТБ(224,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225)",
    "stavka2": "ТБ(224,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(225)",
    "stavka2": "ТБ(225)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(225,5)",
    "stavka2": "ТБ(222,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225,5)",
    "stavka2": "ТБ(223,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225,5)",
    "stavka2": "ТБ(224)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225,5)",
    "stavka2": "ТБ(224,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225,5)",
    "stavka2": "ТБ(225)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(225,5)",
    "stavka2": "ТБ(225)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(225,5)",
    "stavka2": "ТБ(225,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(226)",
    "stavka2": "ТБ(223,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(226)",
    "stavka2": "ТБ(224,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(226)",
    "stavka2": "ТБ(225)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(226)",
    "stavka2": "ТБ(225,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(226)",
    "stavka2": "ТБ(226)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(226,5)",
    "stavka2": "ТБ(223,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(226,5)",
    "stavka2": "ТБ(225)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(226,5)",
    "stavka2": "ТБ(225,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(226,5)",
    "stavka2": "ТБ(226)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(226,5)",
    "stavka2": "ТБ(226,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(227)",
    "stavka2": "ТБ(225,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(227)",
    "stavka2": "ТБ(226)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(227)",
    "stavka2": "ТБ(226)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(227)",
    "stavka2": "ТБ(226,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(227)",
    "stavka2": "ТБ(226,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(227)",
    "stavka2": "ТБ(227)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(227,5)",
    "stavka2": "ТБ(224)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(227,5)",
    "stavka2": "ТБ(225,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(227,5)",
    "stavka2": "ТБ(226)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(227,5)",
    "stavka2": "ТБ(226,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(227,5)",
    "stavka2": "ТБ(226,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(227,5)",
    "stavka2": "ТБ(227)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(227,5)",
    "stavka2": "ТБ(227)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(227,5)",
    "stavka2": "ТБ(227,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(228)",
    "stavka2": "ТБ(226,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(228)",
    "stavka2": "ТБ(227)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(228)",
    "stavka2": "ТБ(227)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(228)",
    "stavka2": "ТБ(227,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(228)",
    "stavka2": "ТБ(227,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(228)",
    "stavka2": "ТБ(228)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(228,5)",
    "stavka2": "ТБ(225,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(228,5)",
    "stavka2": "ТБ(227)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(228,5)",
    "stavka2": "ТБ(227,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(228,5)",
    "stavka2": "ТБ(227,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(228,5)",
    "stavka2": "ТБ(228)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(228,5)",
    "stavka2": "ТБ(228)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(228,5)",
    "stavka2": "ТБ(228,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(229)",
    "stavka2": "ТБ(227,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(229)",
    "stavka2": "ТБ(228)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(229)",
    "stavka2": "ТБ(228,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(229)",
    "stavka2": "ТБ(228,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(229)",
    "stavka2": "ТБ(229)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(229)",
    "stavka2": "ТБ(229)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(229,5)",
    "stavka2": "ТБ(225,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(229,5)",
    "stavka2": "ТБ(228)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(229,5)",
    "stavka2": "ТБ(228,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(229,5)",
    "stavka2": "ТБ(229)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(229,5)",
    "stavka2": "ТБ(229)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(229,5)",
    "stavka2": "ТБ(229,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(23)",
    "stavka2": "ТБ(22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(23)",
    "stavka2": "ТБ(23)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(23,5)",
    "stavka2": "ТБ(22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(23,5)",
    "stavka2": "ТБ(23)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(23,5)",
    "stavka2": "ТБ(23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(230)",
    "stavka2": "ТБ(228)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(230)",
    "stavka2": "ТБ(228,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(230)",
    "stavka2": "ТБ(229)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(230)",
    "stavka2": "ТБ(229,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(230)",
    "stavka2": "ТБ(230)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(230,5)",
    "stavka2": "ТБ(227)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(230,5)",
    "stavka2": "ТБ(228,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(230,5)",
    "stavka2": "ТБ(229,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(230,5)",
    "stavka2": "ТБ(230)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(230,5)",
    "stavka2": "ТБ(230)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(230,5)",
    "stavka2": "ТБ(230,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(231)",
    "stavka2": "ТБ(230)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(231)",
    "stavka2": "ТБ(230,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(231)",
    "stavka2": "ТБ(230,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(231)",
    "stavka2": "ТБ(231)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(231,5)",
    "stavka2": "ТБ(227,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(231,5)",
    "stavka2": "ТБ(228,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(231,5)",
    "stavka2": "ТБ(229,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(231,5)",
    "stavka2": "ТБ(231)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(231,5)",
    "stavka2": "ТБ(231,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(232)",
    "stavka2": "ТБ(231,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(232)",
    "stavka2": "ТБ(232)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(232,5)",
    "stavka2": "ТБ(228,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(232,5)",
    "stavka2": "ТБ(229,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(232,5)",
    "stavka2": "ТБ(232)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(232,5)",
    "stavka2": "ТБ(232,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(233)",
    "stavka2": "ТБ(232,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(233)",
    "stavka2": "ТБ(232,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(233)",
    "stavka2": "ТБ(233)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(233,5)",
    "stavka2": "ТБ(232,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(233,5)",
    "stavka2": "ТБ(233)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(233,5)",
    "stavka2": "ТБ(233)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(233,5)",
    "stavka2": "ТБ(233,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(230,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(230,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(231)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(231)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(232)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(232)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(233)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(233,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(234)",
    "stavka2": "ТБ(234)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(234,5)",
    "stavka2": "ТБ(231)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(234,5)",
    "stavka2": "ТБ(231)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(234,5)",
    "stavka2": "ТБ(233,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(234,5)",
    "stavka2": "ТБ(234)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(234,5)",
    "stavka2": "ТБ(234,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(235)",
    "stavka2": "ТБ(234)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(235)",
    "stavka2": "ТБ(234)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(235)",
    "stavka2": "ТБ(234,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(235)",
    "stavka2": "ТБ(235)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(235,5)",
    "stavka2": "ТБ(234)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(235,5)",
    "stavka2": "ТБ(234,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(235,5)",
    "stavka2": "ТБ(235)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(235,5)",
    "stavka2": "ТБ(235)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(235,5)",
    "stavka2": "ТБ(235,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(236)",
    "stavka2": "ТБ(235)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(236)",
    "stavka2": "ТБ(235,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(236)",
    "stavka2": "ТБ(235,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(236)",
    "stavka2": "ТБ(236)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(236,5)",
    "stavka2": "ТБ(235)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(236,5)",
    "stavka2": "ТБ(235,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(236,5)",
    "stavka2": "ТБ(236)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(236,5)",
    "stavka2": "ТБ(236,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(237)",
    "stavka2": "ТБ(235,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(237)",
    "stavka2": "ТБ(236,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(237)",
    "stavka2": "ТБ(236,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(237)",
    "stavka2": "ТБ(237)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(237,5)",
    "stavka2": "ТБ(236)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(237,5)",
    "stavka2": "ТБ(236,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(237,5)",
    "stavka2": "ТБ(237)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(237,5)",
    "stavka2": "ТБ(237)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(237,5)",
    "stavka2": "ТБ(237,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(238)",
    "stavka2": "ТБ(237)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(238)",
    "stavka2": "ТБ(237)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(238)",
    "stavka2": "ТБ(237,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(238)",
    "stavka2": "ТБ(238)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(238,5)",
    "stavka2": "ТБ(237,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(238,5)",
    "stavka2": "ТБ(238)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(238,5)",
    "stavka2": "ТБ(238)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(238,5)",
    "stavka2": "ТБ(238,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(239)",
    "stavka2": "ТБ(238,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(239)",
    "stavka2": "ТБ(239)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(239,5)",
    "stavka2": "ТБ(237)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(239,5)",
    "stavka2": "ТБ(237,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(239,5)",
    "stavka2": "ТБ(237,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(239,5)",
    "stavka2": "ТБ(238,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(239,5)",
    "stavka2": "ТБ(239)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(239,5)",
    "stavka2": "ТБ(239,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(24)",
    "stavka2": "ТБ(23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(24)",
    "stavka2": "ТБ(23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(24)",
    "stavka2": "ТБ(24)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(24,5)",
    "stavka2": "ТБ(24)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(240)",
    "stavka2": "ТБ(239)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(240)",
    "stavka2": "ТБ(239,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(240)",
    "stavka2": "ТБ(240)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(240,5)",
    "stavka2": "ТБ(239,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(240,5)",
    "stavka2": "ТБ(240)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(240,5)",
    "stavka2": "ТБ(240,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(241)",
    "stavka2": "ТБ(240,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(241)",
    "stavka2": "ТБ(241)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(241,5)",
    "stavka2": "ТБ(240,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(241,5)",
    "stavka2": "ТБ(241)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(241,5)",
    "stavka2": "ТБ(241)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(241,5)",
    "stavka2": "ТБ(241,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(242)",
    "stavka2": "ТБ(241,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(242)",
    "stavka2": "ТБ(242)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(242,5)",
    "stavka2": "ТБ(242)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(242,5)",
    "stavka2": "ТБ(242,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(243)",
    "stavka2": "ТБ(240,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(243)",
    "stavka2": "ТБ(242,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(243)",
    "stavka2": "ТБ(243)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(243,5)",
    "stavka2": "ТБ(241)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(243,5)",
    "stavka2": "ТБ(241,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(243,5)",
    "stavka2": "ТБ(243)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(243,5)",
    "stavka2": "ТБ(243,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(244)",
    "stavka2": "ТБ(244)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(244,5)",
    "stavka2": "ТБ(244,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(245)",
    "stavka2": "ТБ(244,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(245)",
    "stavka2": "ТБ(245)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(245,5)",
    "stavka2": "ТБ(242,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(245,5)",
    "stavka2": "ТБ(244,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(245,5)",
    "stavka2": "ТБ(245)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(246)",
    "stavka2": "ТБ(243,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(246)",
    "stavka2": "ТБ(244,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(246)",
    "stavka2": "ТБ(245)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(246)",
    "stavka2": "ТБ(245,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(246,5)",
    "stavka2": "ТБ(246)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(247)",
    "stavka2": "ТБ(246,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(248,5)",
    "stavka2": "ТБ(248)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(25)",
    "stavka2": "ТБ(23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(25,5)",
    "stavka2": "ТБ(25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(25,5)",
    "stavka2": "ТБ(25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(26)",
    "stavka2": "ТБ(24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(26)",
    "stavka2": "ТБ(26)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(26,5)",
    "stavka2": "ТБ(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(27,5)",
    "stavka2": "ТБ(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(28,5)",
    "stavka2": "ТБ(28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(29,5)",
    "stavka2": "ТБ(29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(299,5)",
    "stavka2": "ТБ(298)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(3)",
    "stavka2": "ТБ(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(3)",
    "stavka2": "ТБ(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(3,25)",
    "stavka2": "ТБ(3,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(3,5)",
    "stavka2": "ТБ(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(3,5)",
    "stavka2": "ТБ(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(3,5)",
    "stavka2": "ТБ(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(3,75)",
    "stavka2": "ТБ(3,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(300)",
    "stavka2": "ТБ(298)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(304)",
    "stavka2": "ТБ(303,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(306,5)",
    "stavka2": "ТБ(305,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(309,5)",
    "stavka2": "ТБ(309)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(31,5)",
    "stavka2": "ТБ(31,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(312)",
    "stavka2": "ТБ(311,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(312)",
    "stavka2": "ТБ(312)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(32,5)",
    "stavka2": "ТБ(32,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(33,5)",
    "stavka2": "ТБ(33,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(34)",
    "stavka2": "ТБ(33,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(34)",
    "stavka2": "ТБ(34)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(34,5)",
    "stavka2": "ТБ(34)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(34,5)",
    "stavka2": "ТБ(34,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(35)",
    "stavka2": "ТБ(34,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(35)",
    "stavka2": "ТБ(35)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(35,5)",
    "stavka2": "ТБ(34,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(35,5)",
    "stavka2": "ТБ(35)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(35,5)",
    "stavka2": "ТБ(35,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(36)",
    "stavka2": "ТБ(35,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(36,5)",
    "stavka2": "ТБ(36)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(36,5)",
    "stavka2": "ТБ(36,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(37)",
    "stavka2": "ТБ(36,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(37)",
    "stavka2": "ТБ(37)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(37,5)",
    "stavka2": "ТБ(36,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(37,5)",
    "stavka2": "ТБ(37)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(37,5)",
    "stavka2": "ТБ(37)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(37,5)",
    "stavka2": "ТБ(37,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(38)",
    "stavka2": "ТБ(37,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(38)",
    "stavka2": "ТБ(38)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(38,5)",
    "stavka2": "ТБ(35,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(38,5)",
    "stavka2": "ТБ(38)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(38,5)",
    "stavka2": "ТБ(38)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(38,5)",
    "stavka2": "ТБ(38,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(39)",
    "stavka2": "ТБ(38,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(39)",
    "stavka2": "ТБ(39)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(39,5)",
    "stavka2": "ТБ(38)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(39,5)",
    "stavka2": "ТБ(38,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(39,5)",
    "stavka2": "ТБ(39)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(39,5)",
    "stavka2": "ТБ(39,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(4)",
    "stavka2": "ТБ(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(4,25)",
    "stavka2": "ТБ(4,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(4,5)",
    "stavka2": "ТБ(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(4,5)",
    "stavka2": "ТБ(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(4,75)",
    "stavka2": "ТБ(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(4,75)",
    "stavka2": "ТБ(4,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(40)",
    "stavka2": "ТБ(39,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(40,5)",
    "stavka2": "ТБ(39,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(40,5)",
    "stavka2": "ТБ(40)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(40,5)",
    "stavka2": "ТБ(40)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(40,5)",
    "stavka2": "ТБ(40,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(41)",
    "stavka2": "ТБ(41)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(41,5)",
    "stavka2": "ТБ(41)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(41,5)",
    "stavka2": "ТБ(41,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(42)",
    "stavka2": "ТБ(41,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(42)",
    "stavka2": "ТБ(41,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(42,5)",
    "stavka2": "ТБ(41,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(42,5)",
    "stavka2": "ТБ(42)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(42,5)",
    "stavka2": "ТБ(42,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(43)",
    "stavka2": "ТБ(43)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(43,5)",
    "stavka2": "ТБ(41,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(43,5)",
    "stavka2": "ТБ(43)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(43,5)",
    "stavka2": "ТБ(43,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(44)",
    "stavka2": "ТБ(43,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(44,5)",
    "stavka2": "ТБ(42,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(44,5)",
    "stavka2": "ТБ(44)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(44,5)",
    "stavka2": "ТБ(44,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(45)",
    "stavka2": "ТБ(44,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(45)",
    "stavka2": "ТБ(45)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(45,5)",
    "stavka2": "ТБ(45)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(45,5)",
    "stavka2": "ТБ(45,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(46)",
    "stavka2": "ТБ(45,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(46)",
    "stavka2": "ТБ(46)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(46,5)",
    "stavka2": "ТБ(46)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(46,5)",
    "stavka2": "ТБ(46,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(47,5)",
    "stavka2": "ТБ(46)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(47,5)",
    "stavka2": "ТБ(47)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(47,5)",
    "stavka2": "ТБ(47,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(48)",
    "stavka2": "ТБ(47,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(48)",
    "stavka2": "ТБ(48)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(48,5)",
    "stavka2": "ТБ(47)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(48,5)",
    "stavka2": "ТБ(48)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(48,5)",
    "stavka2": "ТБ(48,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(49)",
    "stavka2": "ТБ(49)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(49,5)",
    "stavka2": "ТБ(49)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(49,5)",
    "stavka2": "ТБ(49,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(5)",
    "stavka2": "ТБ(4,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(5)",
    "stavka2": "ТБ(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(5,25)",
    "stavka2": "ТБ(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(5,25)",
    "stavka2": "ТБ(5,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(5,5)",
    "stavka2": "ТБ(4,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(5,5)",
    "stavka2": "ТБ(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(5,5)",
    "stavka2": "ТБ(5,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(5,5)",
    "stavka2": "ТБ(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(5,75)",
    "stavka2": "ТБ(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(5,75)",
    "stavka2": "ТБ(5,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(50)",
    "stavka2": "ТБ(50)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(50,5)",
    "stavka2": "ТБ(50)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(50,5)",
    "stavka2": "ТБ(50,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(51)",
    "stavka2": "ТБ(50,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(51)",
    "stavka2": "ТБ(51)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(51,5)",
    "stavka2": "ТБ(50,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(51,5)",
    "stavka2": "ТБ(51)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(51,5)",
    "stavka2": "ТБ(51,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(52)",
    "stavka2": "ТБ(50)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(52)",
    "stavka2": "ТБ(51,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(52)",
    "stavka2": "ТБ(52)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(52,5)",
    "stavka2": "ТБ(52)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(52,5)",
    "stavka2": "ТБ(52,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(53)",
    "stavka2": "ТБ(52)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(53)",
    "stavka2": "ТБ(52,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(53)",
    "stavka2": "ТБ(52,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(53)",
    "stavka2": "ТБ(53)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(53,5)",
    "stavka2": "ТБ(53)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(53,5)",
    "stavka2": "ТБ(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(53,5)",
    "stavka2": "ТБ(53,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(54)",
    "stavka2": "ТБ(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(54)",
    "stavka2": "ТБ(54)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(54,5)",
    "stavka2": "ТБ(52,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(54,5)",
    "stavka2": "ТБ(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(54,5)",
    "stavka2": "ТБ(54)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(54,5)",
    "stavka2": "ТБ(54)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(54,5)",
    "stavka2": "ТБ(54,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(55)",
    "stavka2": "ТБ(54)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(55)",
    "stavka2": "ТБ(54,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(55)",
    "stavka2": "ТБ(55)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(55)",
    "stavka2": "ТБ(55)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(55,5)",
    "stavka2": "ТБ(55,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(56)",
    "stavka2": "ТБ(55)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(56)",
    "stavka2": "ТБ(55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(56)",
    "stavka2": "ТБ(56)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(56,5)",
    "stavka2": "ТБ(56,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(57)",
    "stavka2": "ТБ(55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(57)",
    "stavka2": "ТБ(56,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(57)",
    "stavka2": "ТБ(57)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(57,5)",
    "stavka2": "ТБ(57)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(57,5)",
    "stavka2": "ТБ(57,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(58)",
    "stavka2": "ТБ(57,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(58)",
    "stavka2": "ТБ(58)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(58,5)",
    "stavka2": "ТБ(58,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(59)",
    "stavka2": "ТБ(58,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(59)",
    "stavka2": "ТБ(59)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(59,5)",
    "stavka2": "ТБ(59,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(6)",
    "stavka2": "ТБ(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(6)",
    "stavka2": "ТБ(5,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(6)",
    "stavka2": "ТБ(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(6,25)",
    "stavka2": "ТБ(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(6,25)",
    "stavka2": "ТБ(6,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(6,5)",
    "stavka2": "ТБ(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(6,5)",
    "stavka2": "ТБ(6,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(6,5)",
    "stavka2": "ТБ(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(6,5)",
    "stavka2": "ТБ(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(6,75)",
    "stavka2": "ТБ(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(6,75)",
    "stavka2": "ТБ(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(6,75)",
    "stavka2": "ТБ(6,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(60)",
    "stavka2": "ТБ(59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(60)",
    "stavka2": "ТБ(60)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(60,5)",
    "stavka2": "ТБ(60)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(60,5)",
    "stavka2": "ТБ(60,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(61)",
    "stavka2": "ТБ(61)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(61,5)",
    "stavka2": "ТБ(61)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(61,5)",
    "stavka2": "ТБ(61,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(62)",
    "stavka2": "ТБ(62)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(62,5)",
    "stavka2": "ТБ(62,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(63)",
    "stavka2": "ТБ(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(63)",
    "stavka2": "ТБ(63)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(63,5)",
    "stavka2": "ТБ(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(63,5)",
    "stavka2": "ТБ(63,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(64,5)",
    "stavka2": "ТБ(64,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(65,5)",
    "stavka2": "ТБ(65)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(65,5)",
    "stavka2": "ТБ(65,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(67)",
    "stavka2": "ТБ(66,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(67,5)",
    "stavka2": "ТБ(67,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(68,5)",
    "stavka2": "ТБ(67)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(68,5)",
    "stavka2": "ТБ(67,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(69)",
    "stavka2": "ТБ(67)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(69)",
    "stavka2": "ТБ(68,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(69,5)",
    "stavka2": "ТБ(69)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(7)",
    "stavka2": "ТБ(6,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(7)",
    "stavka2": "ТБ(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(7,25)",
    "stavka2": "ТБ(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(7,5)",
    "stavka2": "ТБ(7,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(7,5)",
    "stavka2": "ТБ(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(7,75)",
    "stavka2": "ТБ(7,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(71,5)",
    "stavka2": "ТБ(71,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(72)",
    "stavka2": "ТБ(71,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(73,5)",
    "stavka2": "ТБ(73,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(8)",
    "stavka2": "ТБ(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(8,25)",
    "stavka2": "ТБ(8,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(8,5)",
    "stavka2": "ТБ(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(8,75)",
    "stavka2": "ТБ(8,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(9)",
    "stavka2": "ТБ(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(9,5)",
    "stavka2": "ТБ(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "ТМ(9,5)",
    "stavka2": "ТБ(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "ТМ(99)",
    "stavka2": "ТБ(98,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "УГЛ 1",
    "stavka2": "УГЛ Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1",
    "stavka2": "УГЛ Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 1",
    "stavka2": "УГЛ 1/2 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 ТМ(3,5)",
    "stavka2": "УГЛ 1/2 ТБ(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 ТМ(4)",
    "stavka2": "УГЛ 1/2 ТБ(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 ТМ(4,5)",
    "stavka2": "УГЛ 1/2 ТБ(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 ТМ(5)",
    "stavka2": "УГЛ 1/2 ТБ(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 ТМ(5,5)",
    "stavka2": "УГЛ 1/2 ТБ(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 ТМ(6)",
    "stavka2": "УГЛ 1/2 ТБ(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-0,25)",
    "stavka2": "УГЛ 1/2 Ф2(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-0,5)",
    "stavka2": "УГЛ 1/2 Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-0,75)",
    "stavka2": "УГЛ 1/2 Ф2(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-1)",
    "stavka2": "УГЛ 1/2 Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-1)",
    "stavka2": "УГЛ 1/2 Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-1,25)",
    "stavka2": "УГЛ 1/2 Ф2(1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-1,75)",
    "stavka2": "УГЛ 1/2 Ф2(1,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-1,75)",
    "stavka2": "УГЛ 1/2 Ф2(1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-2)",
    "stavka2": "УГЛ 1/2 Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-2,25)",
    "stavka2": "УГЛ 1/2 Ф2(2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-2,5)",
    "stavka2": "УГЛ 1/2 Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-3)",
    "stavka2": "УГЛ 1/2 Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(-3,5)",
    "stavka2": "УГЛ 1/2 Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(0)",
    "stavka2": "УГЛ 1/2 Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(0,25)",
    "stavka2": "УГЛ 1/2 Ф2(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(0,5)",
    "stavka2": "УГЛ 1/2 2",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(0,5)",
    "stavka2": "УГЛ 1/2 Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(0,75)",
    "stavka2": "УГЛ 1/2 Ф2(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф1(1,5)",
    "stavka2": "УГЛ 1/2 Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф2(0,25)",
    "stavka2": "УГЛ 1/2 Ф1(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ 1/2 Ф2(0,25)",
    "stavka2": "УГЛ 1/2 Ф2(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ИТ2М(2,5)",
    "stavka2": "УГЛ ИТ2Б(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ИТ2М(3,5)",
    "stavka2": "УГЛ ИТ1Б(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ИТ2М(3,5)",
    "stavka2": "УГЛ ИТ2Б(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ТМ(10)",
    "stavka2": "УГЛ ТБ(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ТМ(10,5)",
    "stavka2": "УГЛ ТБ(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ТМ(11)",
    "stavka2": "УГЛ ТБ(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ТМ(11,5)",
    "stavka2": "УГЛ ТБ(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ТМ(12)",
    "stavka2": "УГЛ ТБ(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ТМ(8,5)",
    "stavka2": "УГЛ ТБ(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ТМ(9)",
    "stavka2": "УГЛ ТБ(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ ТМ(9,5)",
    "stavka2": "УГЛ ТБ(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-0,25)",
    "stavka2": "УГЛ Ф2(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-0,5)",
    "stavka2": "УГЛ Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-0,75)",
    "stavka2": "УГЛ Ф2(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-1)",
    "stavka2": "УГЛ Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-1,5)",
    "stavka2": "УГЛ Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-2)",
    "stavka2": "УГЛ Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-2,25)",
    "stavka2": "УГЛ Ф2(2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-2,5)",
    "stavka2": "УГЛ Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-2,75)",
    "stavka2": "УГЛ Ф2(2,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-3)",
    "stavka2": "УГЛ Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-3,25)",
    "stavka2": "УГЛ Ф2(3,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-3,5)",
    "stavka2": "УГЛ Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-3,75)",
    "stavka2": "УГЛ Ф2(3,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-4)",
    "stavka2": "УГЛ Ф2(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-4,25)",
    "stavka2": "УГЛ Ф2(4,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-4,5)",
    "stavka2": "УГЛ Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-5)",
    "stavka2": "УГЛ Ф2(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-5,25)",
    "stavka2": "УГЛ Ф1(5,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-5,25)",
    "stavka2": "УГЛ Ф2(5,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-5,5)",
    "stavka2": "УГЛ Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-5,75)",
    "stavka2": "УГЛ Ф2(5,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-6)",
    "stavka2": "УГЛ Ф2(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-6,5)",
    "stavka2": "УГЛ Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-7)",
    "stavka2": "УГЛ Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(-9)",
    "stavka2": "УГЛ Ф2(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(0)",
    "stavka2": "УГЛ Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(0)",
    "stavka2": "УГЛ Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(0,5)",
    "stavka2": "УГЛ 2",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(0,5)",
    "stavka2": "УГЛ Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(0,75)",
    "stavka2": "УГЛ Ф2(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(1)",
    "stavka2": "УГЛ Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(1,25)",
    "stavka2": "УГЛ Ф2(-1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(1,5)",
    "stavka2": "УГЛ Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(1,75)",
    "stavka2": "УГЛ Ф2(-1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(2)",
    "stavka2": "УГЛ Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(2,5)",
    "stavka2": "УГЛ Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(3)",
    "stavka2": "УГЛ Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(3,5)",
    "stavka2": "УГЛ Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(3,5)",
    "stavka2": "УГЛ Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(4)",
    "stavka2": "УГЛ Ф2(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(4,5)",
    "stavka2": "УГЛ Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(4,75)",
    "stavka2": "УГЛ Ф2(-4,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(5)",
    "stavka2": "УГЛ Ф2(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф1(5,25)",
    "stavka2": "УГЛ Ф2(-5,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(-0,5)",
    "stavka2": "УГЛ Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(-4)",
    "stavka2": "УГЛ Ф1(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(0)",
    "stavka2": "УГЛ Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(0,5)",
    "stavka2": "УГЛ 1",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(1)",
    "stavka2": "УГЛ Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(1,75)",
    "stavka2": "УГЛ Ф2(-1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(2)",
    "stavka2": "УГЛ Ф1(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(2,5)",
    "stavka2": "УГЛ Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(3,5)",
    "stavka2": "УГЛ Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(5)",
    "stavka2": "УГЛ Ф1(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "УГЛ Ф2(6)",
    "stavka2": "УГЛ Ф2(-6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,25)",
    "stavka2": "Ф1(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,25)",
    "stavka2": "Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,25)",
    "stavka2": "Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,25)",
    "stavka2": "Ф2(0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,25)",
    "stavka2": "Ф2(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,25)",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,25)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },

  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Х2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф1(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф2(0,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,5)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,75)",
    "stavka2": "Ф1(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,75)",
    "stavka2": "Ф1(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,75)",
    "stavka2": "Ф2(0,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-0,75)",
    "stavka2": "Ф2(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-0,75)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф1(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф2(1,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1)",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,25)",
    "stavka2": "Ф1(1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,25)",
    "stavka2": "Ф2(1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,25)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },

  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф1(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф1(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,5)",
    "stavka2": "Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,75)",
    "stavka2": "Ф1(1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,75)",
    "stavka2": "Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,75)",
    "stavka2": "Ф1(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,75)",
    "stavka2": "Ф2(1,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-1,75)",
    "stavka2": "Ф2(1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-1,75)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10)",
    "stavka2": "Ф1(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-10)",
    "stavka2": "Ф1(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10)",
    "stavka2": "Ф2(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-10)",
    "stavka2": "Ф2(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10)",
    "stavka2": "Ф2(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10)",
    "stavka2": "Ф2(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10)",
    "stavka2": "Ф2(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-10,5)",
    "stavka2": "Ф1(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-10,5)",
    "stavka2": "Ф1(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10,5)",
    "stavka2": "Ф2(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-10,5)",
    "stavka2": "Ф2(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10,5)",
    "stavka2": "Ф2(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10,5)",
    "stavka2": "Ф2(12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10,5)",
    "stavka2": "Ф2(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-10,5)",
    "stavka2": "Ф2(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11)",
    "stavka2": "Ф1(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-11)",
    "stavka2": "Ф1(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11)",
    "stavka2": "Ф2(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-11)",
    "stavka2": "Ф2(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11)",
    "stavka2": "Ф2(12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11)",
    "stavka2": "Ф2(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф1(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф1(12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф1(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф1(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф2(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф2(12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф2(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф2(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф2(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф2(13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-11,5)",
    "stavka2": "Ф2(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12)",
    "stavka2": "Ф1(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-12)",
    "stavka2": "Ф1(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12)",
    "stavka2": "Ф1(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12)",
    "stavka2": "Ф2(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-12)",
    "stavka2": "Ф2(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12)",
    "stavka2": "Ф2(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12)",
    "stavka2": "Ф2(13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12)",
    "stavka2": "Ф2(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12,5)",
    "stavka2": "Ф1(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-12,5)",
    "stavka2": "Ф1(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12,5)",
    "stavka2": "Ф1(13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12,5)",
    "stavka2": "Ф2(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-12,5)",
    "stavka2": "Ф2(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12,5)",
    "stavka2": "Ф2(13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-12,5)",
    "stavka2": "Ф2(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-12,5)",
    "stavka2": "Ф2(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13)",
    "stavka2": "Ф1(13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-13)",
    "stavka2": "Ф1(13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13)",
    "stavka2": "Ф1(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13)",
    "stavka2": "Ф2(13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-13)",
    "stavka2": "Ф2(13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13)",
    "stavka2": "Ф2(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф1(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф1(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф1(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф2(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф2(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф2(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф2(15)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф2(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф2(16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-13,5)",
    "stavka2": "Ф2(16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-14)",
    "stavka2": "Ф1(14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-14)",
    "stavka2": "Ф1(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14)",
    "stavka2": "Ф2(14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-14)",
    "stavka2": "Ф2(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14)",
    "stavka2": "Ф2(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-14)",
    "stavka2": "Ф2(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14)",
    "stavka2": "Ф2(16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14)",
    "stavka2": "Ф2(16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14,5)",
    "stavka2": "Ф1(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-14,5)",
    "stavka2": "Ф1(15)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14,5)",
    "stavka2": "Ф1(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14,5)",
    "stavka2": "Ф2(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-14,5)",
    "stavka2": "Ф2(15)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14,5)",
    "stavka2": "Ф2(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14,5)",
    "stavka2": "Ф2(16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-14,5)",
    "stavka2": "Ф2(17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15)",
    "stavka2": "Ф1(15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-15)",
    "stavka2": "Ф1(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15)",
    "stavka2": "Ф2(15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-15)",
    "stavka2": "Ф2(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф1(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф1(16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф2(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф2(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф2(16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф2(16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф2(17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф2(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф2(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-15,5)",
    "stavka2": "Ф2(19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16)",
    "stavka2": "Ф1(16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-16)",
    "stavka2": "Ф2(16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-16)",
    "stavka2": "Ф2(16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16)",
    "stavka2": "Ф2(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф1(16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф1(17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф1(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф2(16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф2(17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф2(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф2(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф2(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф2(19)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-16,5)",
    "stavka2": "Ф2(20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17)",
    "stavka2": "Ф1(17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-17)",
    "stavka2": "Ф1(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17)",
    "stavka2": "Ф2(17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-17)",
    "stavka2": "Ф2(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17)",
    "stavka2": "Ф2(18)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф1(17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф1(18)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф1(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф2(17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф2(18)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф2(18)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф2(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф2(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф2(19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-17,5)",
    "stavka2": "Ф2(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18)",
    "stavka2": "Ф1(18)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-18)",
    "stavka2": "Ф1(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18)",
    "stavka2": "Ф2(18)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-18)",
    "stavka2": "Ф2(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18)",
    "stavka2": "Ф2(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18)",
    "stavka2": "Ф2(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18)",
    "stavka2": "Ф2(20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-18,5)",
    "stavka2": "Ф1(18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-18,5)",
    "stavka2": "Ф1(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18,5)",
    "stavka2": "Ф2(18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-18,5)",
    "stavka2": "Ф2(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18,5)",
    "stavka2": "Ф2(19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18,5)",
    "stavka2": "Ф2(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18,5)",
    "stavka2": "Ф2(20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-18,5)",
    "stavka2": "Ф2(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-19)",
    "stavka2": "Ф1(19)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-19)",
    "stavka2": "Ф1(19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-19)",
    "stavka2": "Ф2(19)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-19)",
    "stavka2": "Ф2(19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-19,5)",
    "stavka2": "Ф1(19,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-19,5)",
    "stavka2": "Ф1(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-19,5)",
    "stavka2": "Ф2(19,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-19,5)",
    "stavka2": "Ф2(20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-19,5)",
    "stavka2": "Ф2(20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-19,5)",
    "stavka2": "Ф2(21)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-19,5)",
    "stavka2": "Ф2(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-19,5)",
    "stavka2": "Ф2(22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф1(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф1(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф1(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2)",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,25)",
    "stavka2": "Ф1(2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2,25)",
    "stavka2": "Ф2(2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2,25)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,25)",
    "stavka2": "Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф1(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф1(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф1(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф1(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф1(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,5)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,75)",
    "stavka2": "Ф2(2,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-2,75)",
    "stavka2": "Ф2(2,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-20)",
    "stavka2": "Ф1(20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-20)",
    "stavka2": "Ф1(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-20)",
    "stavka2": "Ф2(20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-20)",
    "stavka2": "Ф2(20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-20)",
    "stavka2": "Ф2(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-20,5)",
    "stavka2": "Ф1(20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-20,5)",
    "stavka2": "Ф1(21)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-20,5)",
    "stavka2": "Ф1(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-20,5)",
    "stavka2": "Ф1(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-20,5)",
    "stavka2": "Ф2(20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-20,5)",
    "stavka2": "Ф2(21)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-20,5)",
    "stavka2": "Ф2(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-21)",
    "stavka2": "Ф1(21)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-21)",
    "stavka2": "Ф1(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-21)",
    "stavka2": "Ф2(21)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-21)",
    "stavka2": "Ф2(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-21,5)",
    "stavka2": "Ф1(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-21,5)",
    "stavka2": "Ф2(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-21,5)",
    "stavka2": "Ф2(22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-21,5)",
    "stavka2": "Ф2(22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-22)",
    "stavka2": "Ф1(22)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-22)",
    "stavka2": "Ф2(22)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-22)",
    "stavka2": "Ф2(22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-22)",
    "stavka2": "Ф2(23)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-22,5)",
    "stavka2": "Ф1(22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-22,5)",
    "stavka2": "Ф2(22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-22,5)",
    "stavka2": "Ф2(23)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-22,5)",
    "stavka2": "Ф2(23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-22,5)",
    "stavka2": "Ф2(24)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-22,5)",
    "stavka2": "Ф2(25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-23)",
    "stavka2": "Ф1(23)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-23)",
    "stavka2": "Ф2(23)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-23)",
    "stavka2": "Ф2(23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-23,5)",
    "stavka2": "Ф1(23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-23,5)",
    "stavka2": "Ф2(23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-23,5)",
    "stavka2": "Ф2(24)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-23,5)",
    "stavka2": "Ф2(24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-23,5)",
    "stavka2": "Ф2(25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-23,5)",
    "stavka2": "Ф2(25,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-24)",
    "stavka2": "Ф1(25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-24)",
    "stavka2": "Ф2(24)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-24)",
    "stavka2": "Ф2(24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-24,5)",
    "stavka2": "Ф1(24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-24,5)",
    "stavka2": "Ф1(25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-24,5)",
    "stavka2": "Ф1(25,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-24,5)",
    "stavka2": "Ф2(24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-24,5)",
    "stavka2": "Ф2(25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-24,5)",
    "stavka2": "Ф2(25,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-24,5)",
    "stavka2": "Ф2(26)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-24,5)",
    "stavka2": "Ф2(26,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-25)",
    "stavka2": "Ф1(25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-25)",
    "stavka2": "Ф1(25,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-25)",
    "stavka2": "Ф2(25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-25)",
    "stavka2": "Ф2(25,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-25,5)",
    "stavka2": "Ф1(25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-25,5)",
    "stavka2": "Ф2(25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-25,5)",
    "stavka2": "Ф2(26)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-25,5)",
    "stavka2": "Ф2(27)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-26)",
    "stavka2": "Ф1(26)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-26)",
    "stavka2": "Ф2(26)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-26)",
    "stavka2": "Ф2(26,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-26,5)",
    "stavka2": "Ф1(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-26,5)",
    "stavka2": "Ф2(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-26,5)",
    "stavka2": "Ф2(27)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-26,5)",
    "stavka2": "Ф2(27,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-26,5)",
    "stavka2": "Ф2(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-27)",
    "stavka2": "Ф1(27)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-27)",
    "stavka2": "Ф1(27,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-27)",
    "stavka2": "Ф2(27)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-27,5)",
    "stavka2": "Ф1(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-27,5)",
    "stavka2": "Ф2(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-27,5)",
    "stavka2": "Ф2(28)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-27,5)",
    "stavka2": "Ф2(28,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-27,5)",
    "stavka2": "Ф2(30,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-28)",
    "stavka2": "Ф2(28)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-28)",
    "stavka2": "Ф2(28,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-28)",
    "stavka2": "Ф2(29)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-28)",
    "stavka2": "Ф2(30,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-28,5)",
    "stavka2": "Ф1(28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-28,5)",
    "stavka2": "Ф2(28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-28,5)",
    "stavka2": "Ф2(29)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-28,5)",
    "stavka2": "Ф2(29,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-28,5)",
    "stavka2": "Ф2(30,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-29)",
    "stavka2": "Ф2(29)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-29)",
    "stavka2": "Ф2(29,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-29)",
    "stavka2": "Ф2(29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-29)",
    "stavka2": "Ф2(30,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-29,5)",
    "stavka2": "Ф1(29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-29,5)",
    "stavka2": "Ф2(29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-29,5)",
    "stavka2": "Ф2(30)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-29,5)",
    "stavka2": "Ф2(30,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-29,5)",
    "stavka2": "Ф2(31)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-29,5)",
    "stavka2": "Ф2(31,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-29,5)",
    "stavka2": "Ф2(32)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-29,5)",
    "stavka2": "Ф2(33)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф1(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф1(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф1(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф1(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,25)",
    "stavka2": "Ф2(3,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,25)",
    "stavka2": "Ф2(3,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф1(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф1(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф1(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф1(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф2(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф2(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф2(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,5)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-3,75)",
    "stavka2": "Ф2(3,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-30)",
    "stavka2": "Ф1(30)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-30)",
    "stavka2": "Ф1(30,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-30)",
    "stavka2": "Ф1(31)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-30)",
    "stavka2": "Ф2(30)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-30)",
    "stavka2": "Ф2(30,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-30,5)",
    "stavka2": "Ф1(30,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-30,5)",
    "stavka2": "Ф1(31)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-30,5)",
    "stavka2": "Ф2(30,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-30,5)",
    "stavka2": "Ф2(31)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-31)",
    "stavka2": "Ф2(31)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-31,5)",
    "stavka2": "Ф1(31,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-31,5)",
    "stavka2": "Ф2(31,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-31,5)",
    "stavka2": "Ф2(32)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-32)",
    "stavka2": "Ф2(32)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-32)",
    "stavka2": "Ф2(32,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-32,5)",
    "stavka2": "Ф2(32,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-32,5)",
    "stavka2": "Ф2(33)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-32,5)",
    "stavka2": "Ф2(33,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-32,5)",
    "stavka2": "Ф2(34)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-33)",
    "stavka2": "Ф2(33)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-33)",
    "stavka2": "Ф2(33,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-33)",
    "stavka2": "Ф2(35,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-33,5)",
    "stavka2": "Ф1(33,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-33,5)",
    "stavka2": "Ф2(33,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-33,5)",
    "stavka2": "Ф2(34)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-33,5)",
    "stavka2": "Ф2(34,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-33,5)",
    "stavka2": "Ф2(35,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-34)",
    "stavka2": "Ф2(34)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-34)",
    "stavka2": "Ф2(34,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-34)",
    "stavka2": "Ф2(35)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-34)",
    "stavka2": "Ф2(35,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-34)",
    "stavka2": "Ф2(35,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-34,5)",
    "stavka2": "Ф2(34,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-34,5)",
    "stavka2": "Ф2(35)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-34,5)",
    "stavka2": "Ф2(36)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-34,5)",
    "stavka2": "Ф2(36,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-35)",
    "stavka2": "Ф2(35)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-35)",
    "stavka2": "Ф2(35,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-35,5)",
    "stavka2": "Ф2(35,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-36)",
    "stavka2": "Ф2(36)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-36)",
    "stavka2": "Ф2(36,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-36)",
    "stavka2": "Ф2(36,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-36,5)",
    "stavka2": "Ф2(36,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-36,5)",
    "stavka2": "Ф2(39)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-37)",
    "stavka2": "Ф2(37)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-37)",
    "stavka2": "Ф2(37,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-37,5)",
    "stavka2": "Ф2(37,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-38,5)",
    "stavka2": "Ф2(38,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-38,5)",
    "stavka2": "Ф2(39)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-38,5)",
    "stavka2": "Ф2(40)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-39)",
    "stavka2": "Ф2(39)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-39)",
    "stavka2": "Ф2(39,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-39,5)",
    "stavka2": "Ф2(39,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-39,5)",
    "stavka2": "Ф2(40,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф1(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф1(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф1(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф2(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф2(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф2(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4)",
    "stavka2": "Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,25)",
    "stavka2": "Ф2(4,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф1(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф1(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф1(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф1(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф1(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,5)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-4,75)",
    "stavka2": "Ф1(4,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-4,75)",
    "stavka2": "Ф2(4,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-40)",
    "stavka2": "Ф2(40)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-40,5)",
    "stavka2": "Ф2(40,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-41)",
    "stavka2": "Ф2(41,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-41,5)",
    "stavka2": "Ф2(41,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-41,5)",
    "stavka2": "Ф2(42)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-42)",
    "stavka2": "Ф2(42)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-42)",
    "stavka2": "Ф2(42,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-42,5)",
    "stavka2": "Ф2(42,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-44,5)",
    "stavka2": "Ф2(44,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-44,5)",
    "stavka2": "Ф2(46)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-45)",
    "stavka2": "Ф2(45,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-45,5)",
    "stavka2": "Ф2(46)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-46,5)",
    "stavka2": "Ф2(46,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-47,5)",
    "stavka2": "Ф2(47,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-47,5)",
    "stavka2": "Ф2(48)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-48,5)",
    "stavka2": "Ф2(48,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-49,5)",
    "stavka2": "Ф2(49,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф1(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф2(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф2(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф2(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5,25)",
    "stavka2": "Ф2(5,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф1(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф1(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф1(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф2(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-5,5)",
    "stavka2": "Ф2(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-5,75)",
    "stavka2": "Ф2(5,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-50)",
    "stavka2": "Ф2(50)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-50,5)",
    "stavka2": "Ф2(50,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-51,5)",
    "stavka2": "Ф2(51,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-51,5)",
    "stavka2": "Ф2(52,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-51,5)",
    "stavka2": "Ф2(54)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-52)",
    "stavka2": "Ф2(53,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-53,5)",
    "stavka2": "Ф2(53,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-55,5)",
    "stavka2": "Ф2(55,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6)",
    "stavka2": "Ф1(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6)",
    "stavka2": "Ф1(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6)",
    "stavka2": "Ф1(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6)",
    "stavka2": "Ф2(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6)",
    "stavka2": "Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6)",
    "stavka2": "Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6,25)",
    "stavka2": "Ф2(6,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф1(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф1(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф1(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф2(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-6,5)",
    "stavka2": "Ф2(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-6,75)",
    "stavka2": "Ф2(6,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-60,5)",
    "stavka2": "Ф2(60,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-62)",
    "stavka2": "Ф2(62,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-63,5)",
    "stavka2": "Ф2(63,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-65)",
    "stavka2": "Ф2(65)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф1(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф1(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф1(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф1(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф1(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф2(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7)",
    "stavka2": "Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7,25)",
    "stavka2": "Ф2(7,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф1(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф1(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф1(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф1(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф2(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф2(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф2(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7,5)",
    "stavka2": "Ф2(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-7,75)",
    "stavka2": "Ф2(7,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-70,5)",
    "stavka2": "Ф2(70,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-8)",
    "stavka2": "Ф1(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-8)",
    "stavka2": "Ф1(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8)",
    "stavka2": "Ф1(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-8)",
    "stavka2": "Ф1(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8)",
    "stavka2": "Ф2(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8)",
    "stavka2": "Ф2(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-8)",
    "stavka2": "Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф1(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф1(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф1(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф1(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф1(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф2(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф2(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф2(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф2(12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф2(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф2(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-8,5)",
    "stavka2": "Ф2(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9)",
    "stavka2": "Ф1(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-9)",
    "stavka2": "Ф1(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9)",
    "stavka2": "Ф2(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9)",
    "stavka2": "Ф2(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-9)",
    "stavka2": "Ф2(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9,5)",
    "stavka2": "Ф1(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9,5)",
    "stavka2": "Ф1(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-9,5)",
    "stavka2": "Ф2(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9,5)",
    "stavka2": "Ф2(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9,5)",
    "stavka2": "Ф2(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9,5)",
    "stavka2": "Ф2(12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(-9,5)",
    "stavka2": "Ф2(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(-9,5)",
    "stavka2": "Ф2(9,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф1(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0)",
    "stavka2": "Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,25)",
    "stavka2": "Ф1(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0,25)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,25)",
    "stavka2": "Ф2(-0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,25)",
    "stavka2": "Ф2(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0,25)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "Ф1(-0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "Ф2(-0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,5)",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,75)",
    "stavka2": "1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,75)",
    "stavka2": "Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(0,75)",
    "stavka2": "Ф1(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(0,75)",
    "stavka2": "Ф2(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф1(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф1(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1)",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,25)",
    "stavka2": "Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,25)",
    "stavka2": "Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,25)",
    "stavka2": "Ф1(-1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,25)",
    "stavka2": "Ф2(-1,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,25)",
    "stavka2": "Ф2(-1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф1(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф1(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(-1,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,5)",
    "stavka2": "Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,75)",
    "stavka2": "Ф1(-1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(1,75)",
    "stavka2": "Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,75)",
    "stavka2": "Ф2(-1,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(1,75)",
    "stavka2": "Ф2(-1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(10)",
    "stavka2": "Ф1(-10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(10)",
    "stavka2": "Ф1(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10)",
    "stavka2": "Ф2(-10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(10)",
    "stavka2": "Ф2(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10)",
    "stavka2": "Ф2(-9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10)",
    "stavka2": "Ф2(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10)",
    "stavka2": "Ф2(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф1(-10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф1(-10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф1(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф2(-10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф2(-10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф2(-10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф2(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф2(-9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(10,5)",
    "stavka2": "Ф2(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11)",
    "stavka2": "Ф1(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11)",
    "stavka2": "Ф1(-11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(11)",
    "stavka2": "Ф2(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11)",
    "stavka2": "Ф2(-11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф1(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф1(-11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф1(-11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф1(-11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф1(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф2(-10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф2(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф2(-11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф2(-11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф2(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(11,5)",
    "stavka2": "Ф2(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф1(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф1(-11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф1(-11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф1(-12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф2(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф2(-11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф2(-11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф2(-12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф2(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф2(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(12)",
    "stavka2": "Ф2(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12,5)",
    "stavka2": "Ф1(-12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(12,5)",
    "stavka2": "Ф2(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12,5)",
    "stavka2": "Ф2(-11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12,5)",
    "stavka2": "Ф2(-11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12,5)",
    "stavka2": "Ф2(-12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(12,5)",
    "stavka2": "Ф2(-12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(12,5)",
    "stavka2": "Ф2(-12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(13)",
    "stavka2": "Ф1(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(13)",
    "stavka2": "Ф1(-13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(13)",
    "stavka2": "Ф2(-11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(13)",
    "stavka2": "Ф2(-12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(13)",
    "stavka2": "Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(13)",
    "stavka2": "Ф2(-13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(13,5)",
    "stavka2": "Ф1(-13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(13,5)",
    "stavka2": "Ф2(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(13,5)",
    "stavka2": "Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(13,5)",
    "stavka2": "Ф2(-13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(13,5)",
    "stavka2": "Ф2(-13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(14)",
    "stavka2": "Ф1(-14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(14)",
    "stavka2": "Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(14)",
    "stavka2": "Ф2(-13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(14)",
    "stavka2": "Ф2(-14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(14,5)",
    "stavka2": "Ф1(-14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(14,5)",
    "stavka2": "Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(14,5)",
    "stavka2": "Ф2(-13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(14,5)",
    "stavka2": "Ф2(-14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(14,5)",
    "stavka2": "Ф2(-14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(15)",
    "stavka2": "Ф1(-14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(15)",
    "stavka2": "Ф1(-15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(15)",
    "stavka2": "Ф2(-14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(15)",
    "stavka2": "Ф2(-15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(15,5)",
    "stavka2": "Ф1(-15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(15,5)",
    "stavka2": "Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(15,5)",
    "stavka2": "Ф2(-14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(15,5)",
    "stavka2": "Ф2(-15)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(15,5)",
    "stavka2": "Ф2(-15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(16)",
    "stavka2": "Ф1(-15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16)",
    "stavka2": "Ф1(-16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(16)",
    "stavka2": "Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16)",
    "stavka2": "Ф2(-14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16)",
    "stavka2": "Ф2(-15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16)",
    "stavka2": "Ф2(-16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(16,5)",
    "stavka2": "Ф1(-16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16,5)",
    "stavka2": "Ф1(-16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(16,5)",
    "stavka2": "Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16,5)",
    "stavka2": "Ф2(-14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16,5)",
    "stavka2": "Ф2(-14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16,5)",
    "stavka2": "Ф2(-15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16,5)",
    "stavka2": "Ф2(-16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(16,5)",
    "stavka2": "Ф2(-16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(17)",
    "stavka2": "Ф1(-16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(17)",
    "stavka2": "Ф1(-17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(17)",
    "stavka2": "Ф2(-15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(17)",
    "stavka2": "Ф2(-16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(17)",
    "stavka2": "Ф2(-17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(17,5)",
    "stavka2": "Ф1(-17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(17,5)",
    "stavka2": "Ф1(-17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(17,5)",
    "stavka2": "Ф2(-16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(17,5)",
    "stavka2": "Ф2(-17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(17,5)",
    "stavka2": "Ф2(-17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(18)",
    "stavka2": "Ф1(-17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(18)",
    "stavka2": "Ф1(-18)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(18)",
    "stavka2": "Ф2(-16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(18)",
    "stavka2": "Ф2(-17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(18)",
    "stavka2": "Ф2(-17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(18)",
    "stavka2": "Ф2(-18)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(18,5)",
    "stavka2": "Ф1(-17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(18,5)",
    "stavka2": "Ф1(-18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(18,5)",
    "stavka2": "Ф2(-15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(18,5)",
    "stavka2": "Ф2(-16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(18,5)",
    "stavka2": "Ф2(-17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(18,5)",
    "stavka2": "Ф2(-18)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(18,5)",
    "stavka2": "Ф2(-18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(19)",
    "stavka2": "Ф1(-18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(19)",
    "stavka2": "Ф2(-17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(19)",
    "stavka2": "Ф2(-18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(19)",
    "stavka2": "Ф2(-19)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(19,5)",
    "stavka2": "Ф2(-17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(19,5)",
    "stavka2": "Ф2(-18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(19,5)",
    "stavka2": "Ф2(-19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(19,5)",
    "stavka2": "Ф2(-19,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф1(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф1(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф1(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2)",
    "stavka2": "Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2,25)",
    "stavka2": "Ф1(-2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2,25)",
    "stavka2": "Ф2(-2,25)",
    "tip": "VILKA"
  },

  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф1(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф1(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(-2,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(-2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(2,5)",
    "stavka2": "Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(2,75)",
    "stavka2": "Ф2(-2,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(20)",
    "stavka2": "Ф2(-18)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(20)",
    "stavka2": "Ф2(-19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(20)",
    "stavka2": "Ф2(-20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(20,5)",
    "stavka2": "Ф1(-19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(20,5)",
    "stavka2": "Ф1(-20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(20,5)",
    "stavka2": "Ф1(-20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(20,5)",
    "stavka2": "Ф2(-19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(20,5)",
    "stavka2": "Ф2(-20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(20,5)",
    "stavka2": "Ф2(-20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(20,5)",
    "stavka2": "Ф2(-20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(21)",
    "stavka2": "Ф1(-20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(21)",
    "stavka2": "Ф1(-21)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(21)",
    "stavka2": "Ф2(-20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(21)",
    "stavka2": "Ф2(-21)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(21,5)",
    "stavka2": "Ф1(-21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(21,5)",
    "stavka2": "Ф2(-20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(21,5)",
    "stavka2": "Ф2(-21)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(21,5)",
    "stavka2": "Ф2(-21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(22)",
    "stavka2": "Ф1(-21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(22)",
    "stavka2": "Ф1(-22)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(22)",
    "stavka2": "Ф2(-21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(22)",
    "stavka2": "Ф2(-22)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(22,5)",
    "stavka2": "Ф1(-22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(22,5)",
    "stavka2": "Ф1(-22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(22,5)",
    "stavka2": "Ф2(-21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(22,5)",
    "stavka2": "Ф2(-22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(22,5)",
    "stavka2": "Ф2(-22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(23)",
    "stavka2": "Ф1(-22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(23)",
    "stavka2": "Ф2(-20)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(23)",
    "stavka2": "Ф2(-22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(23)",
    "stavka2": "Ф2(-22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(23)",
    "stavka2": "Ф2(-23)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(23,5)",
    "stavka2": "Ф1(-22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(23,5)",
    "stavka2": "Ф1(-23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(23,5)",
    "stavka2": "Ф2(-23)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(23,5)",
    "stavka2": "Ф2(-23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(24)",
    "stavka2": "Ф2(-22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(24)",
    "stavka2": "Ф2(-23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(24)",
    "stavka2": "Ф2(-24)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(24,5)",
    "stavka2": "Ф1(-24)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(24,5)",
    "stavka2": "Ф1(-24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(24,5)",
    "stavka2": "Ф2(-23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(24,5)",
    "stavka2": "Ф2(-24)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(24,5)",
    "stavka2": "Ф2(-24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(25)",
    "stavka2": "Ф1(-24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(25)",
    "stavka2": "Ф2(-23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(25)",
    "stavka2": "Ф2(-24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(25,5)",
    "stavka2": "Ф1(-24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(25,5)",
    "stavka2": "Ф2(-23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(25,5)",
    "stavka2": "Ф2(-24)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(25,5)",
    "stavka2": "Ф2(-25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(25,5)",
    "stavka2": "Ф2(-25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(26)",
    "stavka2": "Ф1(-26)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(26)",
    "stavka2": "Ф2(-24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(26)",
    "stavka2": "Ф2(-26)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(26,5)",
    "stavka2": "Ф1(-26)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(26,5)",
    "stavka2": "Ф1(-26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(26,5)",
    "stavka2": "Ф2(-24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(26,5)",
    "stavka2": "Ф2(-25,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(26,5)",
    "stavka2": "Ф2(-26)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(26,5)",
    "stavka2": "Ф2(-26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(27)",
    "stavka2": "Ф1(-26,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(27)",
    "stavka2": "Ф1(-27)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(27)",
    "stavka2": "Ф2(-26)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(27)",
    "stavka2": "Ф2(-26,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(27)",
    "stavka2": "Ф2(-27)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(27,5)",
    "stavka2": "Ф1(-26,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(27,5)",
    "stavka2": "Ф1(-27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(27,5)",
    "stavka2": "Ф2(-24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(27,5)",
    "stavka2": "Ф2(-27)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(27,5)",
    "stavka2": "Ф2(-27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(28)",
    "stavka2": "Ф1(-26,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28)",
    "stavka2": "Ф1(-27,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28)",
    "stavka2": "Ф2(-24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28)",
    "stavka2": "Ф2(-25,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28)",
    "stavka2": "Ф2(-26)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28)",
    "stavka2": "Ф2(-28)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(28,5)",
    "stavka2": "Ф1(-28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(28,5)",
    "stavka2": "Ф2(-24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28,5)",
    "stavka2": "Ф2(-26,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28,5)",
    "stavka2": "Ф2(-27,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28,5)",
    "stavka2": "Ф2(-28)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(28,5)",
    "stavka2": "Ф2(-28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(29)",
    "stavka2": "Ф1(-29)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(29)",
    "stavka2": "Ф2(-27,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(29)",
    "stavka2": "Ф2(-28,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(29,5)",
    "stavka2": "Ф1(-29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(29,5)",
    "stavka2": "Ф2(-27,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(29,5)",
    "stavka2": "Ф2(-28,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(29,5)",
    "stavka2": "Ф2(-29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф1(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф1(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф1(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф1(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф2(-2,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,25)",
    "stavka2": "Ф2(-3,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф1(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,5)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(3,75)",
    "stavka2": "Ф2(-3,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(30)",
    "stavka2": "Ф2(-29,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(30)",
    "stavka2": "Ф2(-30)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(30,5)",
    "stavka2": "Ф1(-30,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(30,5)",
    "stavka2": "Ф2(-29,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(30,5)",
    "stavka2": "Ф2(-30)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(30,5)",
    "stavka2": "Ф2(-30,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(31)",
    "stavka2": "Ф2(-29,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(31)",
    "stavka2": "Ф2(-30,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(31)",
    "stavka2": "Ф2(-31)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(31,5)",
    "stavka2": "Ф2(-29,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(31,5)",
    "stavka2": "Ф2(-31,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(32)",
    "stavka2": "Ф2(-31,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(32,5)",
    "stavka2": "Ф2(-32)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(32,5)",
    "stavka2": "Ф2(-32,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(33)",
    "stavka2": "Ф2(-32,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(33)",
    "stavka2": "Ф2(-33)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(33,5)",
    "stavka2": "Ф2(-32)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(33,5)",
    "stavka2": "Ф2(-33,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(34,5)",
    "stavka2": "Ф1(-34,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(34,5)",
    "stavka2": "Ф2(-33,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(34,5)",
    "stavka2": "Ф2(-34,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(35)",
    "stavka2": "Ф2(-33,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(35,5)",
    "stavka2": "Ф2(-34,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(36)",
    "stavka2": "Ф2(-33)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(36,5)",
    "stavka2": "Ф2(-36)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(36,5)",
    "stavka2": "Ф2(-36,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(37)",
    "stavka2": "Ф2(-33,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(37,5)",
    "stavka2": "Ф2(-36,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(38)",
    "stavka2": "Ф2(-37,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(38,5)",
    "stavka2": "Ф2(-38,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(39)",
    "stavka2": "Ф2(-38,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(39)",
    "stavka2": "Ф2(-39)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(39,5)",
    "stavka2": "Ф2(-38,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(39,5)",
    "stavka2": "Ф2(-39,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф1(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф1(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф1(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф1(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф2(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф2(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,25)",
    "stavka2": "Ф2(-4,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф1(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф1(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф1(-4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф1(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф2(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф2(-4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф2(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4,5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4,75)",
    "stavka2": "Ф1(-4,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(4,75)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(4,75)",
    "stavka2": "Ф2(-4,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(40)",
    "stavka2": "Ф2(-39,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(40,5)",
    "stavka2": "Ф2(-39,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(40,5)",
    "stavka2": "Ф2(-40,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(41)",
    "stavka2": "Ф2(-40)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(41,5)",
    "stavka2": "Ф2(-41,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(45,5)",
    "stavka2": "Ф2(-43,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(45,5)",
    "stavka2": "Ф2(-44,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(45,5)",
    "stavka2": "Ф2(-45)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(45,5)",
    "stavka2": "Ф2(-45,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(46,5)",
    "stavka2": "Ф2(-46,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(47,5)",
    "stavka2": "Ф2(-47,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(48,5)",
    "stavka2": "Ф2(-48,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(49)",
    "stavka2": "Ф2(-49)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф1(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф1(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф1(-4,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф1(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф2(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф2(-4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(5)",
    "stavka2": "Ф2(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(5,25)",
    "stavka2": "Ф1(-5,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(5,25)",
    "stavka2": "Ф2(-5,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(5,5)",
    "stavka2": "Ф1(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5,5)",
    "stavka2": "Ф1(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(5,5)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5,5)",
    "stavka2": "Ф2(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5,5)",
    "stavka2": "Ф2(-4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5,5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5,5)",
    "stavka2": "Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(5,5)",
    "stavka2": "Ф2(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(5,75)",
    "stavka2": "Ф2(-5,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(50)",
    "stavka2": "Ф2(-50)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(50,5)",
    "stavka2": "Ф2(-50,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(51,5)",
    "stavka2": "Ф2(-50,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(54,5)",
    "stavka2": "Ф2(-54,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(56)",
    "stavka2": "Ф2(-55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(56,5)",
    "stavka2": "Ф2(-55,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(58)",
    "stavka2": "Ф2(-57,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(58,5)",
    "stavka2": "Ф2(-58,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(59,5)",
    "stavka2": "Ф2(-59,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(6)",
    "stavka2": "Ф1(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6)",
    "stavka2": "Ф1(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6)",
    "stavka2": "Ф1(-6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(6)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6)",
    "stavka2": "Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6)",
    "stavka2": "Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6)",
    "stavka2": "Ф2(-6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(6,25)",
    "stavka2": "Ф2(-6,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф1(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф1(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф1(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф1(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф1(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф1(-6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф1(-6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф2(-6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф2(-6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(6,5)",
    "stavka2": "Ф2(-6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(60)",
    "stavka2": "Ф2(-59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(60,5)",
    "stavka2": "Ф2(-59,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф1(-6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф1(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф1(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф1(-7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф2(-6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф2(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7)",
    "stavka2": "Ф2(-7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(7,25)",
    "stavka2": "Ф2(-7,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф1(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф1(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(7,5)",
    "stavka2": "Ф2(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(72)",
    "stavka2": "Ф2(-71,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(72)",
    "stavka2": "Ф2(-72)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(74,5)",
    "stavka2": "Ф2(-74,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф1(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф1(-7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф1(-8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф2(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф2(-7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф2(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(8)",
    "stavka2": "Ф2(-8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(8,5)",
    "stavka2": "Ф1(-7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8,5)",
    "stavka2": "Ф1(-8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8,5)",
    "stavka2": "Ф1(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(8,5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8,5)",
    "stavka2": "Ф2(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8,5)",
    "stavka2": "Ф2(-7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8,5)",
    "stavka2": "Ф2(-8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(8,5)",
    "stavka2": "Ф2(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф1(-8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф1(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф1(-9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф2(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф2(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф2(-7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф2(-8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф2(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9)",
    "stavka2": "Ф2(-9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(9,5)",
    "stavka2": "Ф1(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9,5)",
    "stavka2": "Ф1(-9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9,5)",
    "stavka2": "Ф1(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф1(9,5)",
    "stavka2": "Ф2(-7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9,5)",
    "stavka2": "Ф2(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9,5)",
    "stavka2": "Ф2(-9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф1(9,5)",
    "stavka2": "Ф2(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-0,25)",
    "stavka2": "Ф1(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-0,25)",
    "stavka2": "Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-0,25)",
    "stavka2": "Ф2(0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-0,5)",
    "stavka2": "Ф1(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-0,5)",
    "stavka2": "Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-0,5)",
    "stavka2": "Ф2(0,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-0,5)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-0,5)",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-0,75)",
    "stavka2": "Ф1(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-0,75)",
    "stavka2": "Ф2(0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1)",
    "stavka2": "Ф1(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1)",
    "stavka2": "Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-1)",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1)",
    "stavka2": "Ф2(1,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-1)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-1)",
    "stavka2": "Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-1,25)",
    "stavka2": "Ф1(1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1,25)",
    "stavka2": "Ф2(1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1,5)",
    "stavka2": "Ф1(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1,5)",
    "stavka2": "Ф1(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-1,5)",
    "stavka2": "Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1,5)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-1,5)",
    "stavka2": "Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-1,5)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-1,75)",
    "stavka2": "Ф1(1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-10)",
    "stavka2": "Ф1(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-10)",
    "stavka2": "Ф2(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-10)",
    "stavka2": "Ф2(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-10)",
    "stavka2": "Ф2(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-10,5)",
    "stavka2": "Ф1(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-10,5)",
    "stavka2": "Ф2(10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-10,5)",
    "stavka2": "Ф2(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-10,5)",
    "stavka2": "Ф2(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-10,5)",
    "stavka2": "Ф2(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-11)",
    "stavka2": "Ф2(11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-11)",
    "stavka2": "Ф2(11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-11,5)",
    "stavka2": "Ф1(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-11,5)",
    "stavka2": "Ф2(11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-11,5)",
    "stavka2": "Ф2(12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-11,5)",
    "stavka2": "Ф2(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-11,5)",
    "stavka2": "Ф2(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-12)",
    "stavka2": "Ф2(12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-12)",
    "stavka2": "Ф2(12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-12)",
    "stavka2": "Ф2(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-12)",
    "stavka2": "Ф2(16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-12,5)",
    "stavka2": "Ф1(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-12,5)",
    "stavka2": "Ф2(12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-12,5)",
    "stavka2": "Ф2(13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-12,5)",
    "stavka2": "Ф2(13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-13)",
    "stavka2": "Ф2(13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-13)",
    "stavka2": "Ф2(13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-13)",
    "stavka2": "Ф2(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-13,5)",
    "stavka2": "Ф2(13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-13,5)",
    "stavka2": "Ф2(14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-13,5)",
    "stavka2": "Ф2(14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-13,5)",
    "stavka2": "Ф2(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-13,5)",
    "stavka2": "Ф2(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-14)",
    "stavka2": "Ф2(14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-14)",
    "stavka2": "Ф2(14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-14)",
    "stavka2": "Ф2(16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-14,5)",
    "stavka2": "Ф1(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-14,5)",
    "stavka2": "Ф2(14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-14,5)",
    "stavka2": "Ф2(15)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-14,5)",
    "stavka2": "Ф2(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-14,5)",
    "stavka2": "Ф2(16)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-15)",
    "stavka2": "Ф2(15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-15)",
    "stavka2": "Ф2(15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-15)",
    "stavka2": "Ф2(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-15,5)",
    "stavka2": "Ф1(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-15,5)",
    "stavka2": "Ф2(15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-16)",
    "stavka2": "Ф2(16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-16)",
    "stavka2": "Ф2(16,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-16,5)",
    "stavka2": "Ф2(16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-16,5)",
    "stavka2": "Ф2(17)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-17)",
    "stavka2": "Ф2(17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-17)",
    "stavka2": "Ф2(17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-17,5)",
    "stavka2": "Ф1(17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-17,5)",
    "stavka2": "Ф2(17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-17,5)",
    "stavka2": "Ф2(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-18)",
    "stavka2": "Ф2(18)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-18)",
    "stavka2": "Ф2(18,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-18,5)",
    "stavka2": "Ф1(18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-18,5)",
    "stavka2": "Ф2(18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-18,5)",
    "stavka2": "Ф2(19)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-19)",
    "stavka2": "Ф2(19)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-19)",
    "stavka2": "Ф2(19,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-19,5)",
    "stavka2": "Ф1(19,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-19,5)",
    "stavka2": "Ф2(19,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-2)",
    "stavka2": "Ф1(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-2)",
    "stavka2": "Ф2(2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-2)",
    "stavka2": "Ф2(2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-2)",
    "stavka2": "Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-2,5)",
    "stavka2": "Ф1(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-2,5)",
    "stavka2": "Ф1(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-2,5)",
    "stavka2": "Ф2(2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-2,5)",
    "stavka2": "Ф2(3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-2,5)",
    "stavka2": "Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-2,75)",
    "stavka2": "Ф1(2,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-2,75)",
    "stavka2": "Ф2(2,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-20)",
    "stavka2": "Ф1(20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-20)",
    "stavka2": "Ф2(20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-20,5)",
    "stavka2": "Ф1(20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-20,5)",
    "stavka2": "Ф2(20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-20,5)",
    "stavka2": "Ф2(21)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-21)",
    "stavka2": "Ф2(21)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-21)",
    "stavka2": "Ф2(21,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-21,5)",
    "stavka2": "Ф2(21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-21,5)",
    "stavka2": "Ф2(22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-22)",
    "stavka2": "Ф2(22)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-22)",
    "stavka2": "Ф2(22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-22,5)",
    "stavka2": "Ф2(22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-23)",
    "stavka2": "Ф2(23)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-23)",
    "stavka2": "Ф2(23,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-23,5)",
    "stavka2": "Ф2(23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-24)",
    "stavka2": "Ф2(24)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-24)",
    "stavka2": "Ф2(24,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-24,5)",
    "stavka2": "Ф2(24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-24,5)",
    "stavka2": "Ф2(25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-25,5)",
    "stavka2": "Ф2(25,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-26)",
    "stavka2": "Ф2(26)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-26)",
    "stavka2": "Ф2(26,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-26)",
    "stavka2": "Ф2(27)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-26,5)",
    "stavka2": "Ф2(26,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-27)",
    "stavka2": "Ф2(27)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-27)",
    "stavka2": "Ф2(27,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-27)",
    "stavka2": "Ф2(28,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-27,5)",
    "stavka2": "Ф2(27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-27,5)",
    "stavka2": "Ф2(28,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-28)",
    "stavka2": "Ф2(28,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-28,5)",
    "stavka2": "Ф2(28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-29,5)",
    "stavka2": "Ф2(29,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-3)",
    "stavka2": "Ф1(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-3)",
    "stavka2": "Ф1(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-3)",
    "stavka2": "Ф2(3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-3)",
    "stavka2": "Ф2(3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-3)",
    "stavka2": "Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-3)",
    "stavka2": "Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-3)",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-3,25)",
    "stavka2": "Ф1(3,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-3,5)",
    "stavka2": "Ф1(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-3,5)",
    "stavka2": "Ф1(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-3,5)",
    "stavka2": "Ф1(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-3,5)",
    "stavka2": "Ф2(3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-3,5)",
    "stavka2": "Ф2(4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-3,5)",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-35,5)",
    "stavka2": "Ф2(35,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-4)",
    "stavka2": "Ф1(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-4)",
    "stavka2": "Ф1(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-4)",
    "stavka2": "Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-4)",
    "stavka2": "Ф2(4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-4)",
    "stavka2": "Ф2(4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-4)",
    "stavka2": "Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-4)",
    "stavka2": "Ф2(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-4,5)",
    "stavka2": "Ф1(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-4,5)",
    "stavka2": "Ф2(4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-4,5)",
    "stavka2": "Ф2(5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-4,5)",
    "stavka2": "Ф2(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-40,5)",
    "stavka2": "Ф1(40,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-5)",
    "stavka2": "Ф1(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-5)",
    "stavka2": "Ф1(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-5)",
    "stavka2": "Ф2(5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-5)",
    "stavka2": "Ф2(5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-5)",
    "stavka2": "Ф2(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-5,5)",
    "stavka2": "Ф1(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-5,5)",
    "stavka2": "Ф2(5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-5,5)",
    "stavka2": "Ф2(6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-6)",
    "stavka2": "Ф1(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-6)",
    "stavka2": "Ф2(6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-6)",
    "stavka2": "Ф2(6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-6)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-6)",
    "stavka2": "Ф2(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-6,5)",
    "stavka2": "Ф1(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-6,5)",
    "stavka2": "Ф2(6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-6,5)",
    "stavka2": "Ф2(7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-6,5)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-6,5)",
    "stavka2": "Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-7)",
    "stavka2": "Ф1(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-7)",
    "stavka2": "Ф2(7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-7)",
    "stavka2": "Ф2(7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-7)",
    "stavka2": "Ф2(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-7)",
    "stavka2": "Ф2(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-7)",
    "stavka2": "Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-7,5)",
    "stavka2": "Ф1(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-7,5)",
    "stavka2": "Ф2(7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-7,5)",
    "stavka2": "Ф2(8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-7,5)",
    "stavka2": "Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-8)",
    "stavka2": "Ф1(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-8)",
    "stavka2": "Ф2(8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-8)",
    "stavka2": "Ф2(8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-8)",
    "stavka2": "Ф2(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-8,5)",
    "stavka2": "Ф1(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-8,5)",
    "stavka2": "Ф2(11)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-8,5)",
    "stavka2": "Ф2(8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-8,5)",
    "stavka2": "Ф2(9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-8,5)",
    "stavka2": "Ф2(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-8,5)",
    "stavka2": "Ф2(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-9)",
    "stavka2": "Ф2(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-9)",
    "stavka2": "Ф2(9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-9)",
    "stavka2": "Ф2(9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-9)",
    "stavka2": "Ф2(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-9,5)",
    "stavka2": "Ф1(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-9,5)",
    "stavka2": "Ф2(10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-9,5)",
    "stavka2": "Ф2(10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(-9,5)",
    "stavka2": "Ф2(10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(-9,5)",
    "stavka2": "Ф2(9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "П1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф1(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф1(0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф1(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф1(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф2(0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф2(1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф2(1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0,25)",
    "stavka2": "Ф1(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,25)",
    "stavka2": "Ф2(-0,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "1",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "Ф1(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "Ф2(-0,25)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "Ф2(-0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "Ф2(0)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "Ф2(0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(0,5)",
    "stavka2": "Ф2(0,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,75)",
    "stavka2": "Ф1(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(0,75)",
    "stavka2": "Ф2(-0,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "П1",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "Ф1(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "Ф2(-0,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "Ф2(-0,75)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "Ф2(1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1)",
    "stavka2": "Ф2(2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1,25)",
    "stavka2": "Ф1(-1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1,25)",
    "stavka2": "Ф2(-1,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "Ф1(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "Ф1(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "Ф2(-1)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "Ф2(0)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1,5)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(1,75)",
    "stavka2": "Ф1(-1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(1,75)",
    "stavka2": "Ф2(-1,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(10)",
    "stavka2": "Ф1(-10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(10)",
    "stavka2": "Ф2(-10)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(10)",
    "stavka2": "Ф2(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(10)",
    "stavka2": "Ф2(-9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(10)",
    "stavka2": "Ф2(-9,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(10,5)",
    "stavka2": "Ф1(-10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(10,5)",
    "stavka2": "Ф2(-10)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(10,5)",
    "stavka2": "Ф2(-10,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(11)",
    "stavka2": "Ф2(-10,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(11)",
    "stavka2": "Ф2(-11)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(11,5)",
    "stavka2": "Ф1(-11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(11,5)",
    "stavka2": "Ф2(-11,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(12)",
    "stavka2": "Ф2(-11,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(12)",
    "stavka2": "Ф2(-12)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(12,5)",
    "stavka2": "Ф1(-12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(12,5)",
    "stavka2": "Ф2(-12)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(12,5)",
    "stavka2": "Ф2(-12,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(13)",
    "stavka2": "Ф2(-12,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(13)",
    "stavka2": "Ф2(-13)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(13,5)",
    "stavka2": "Ф1(-13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(13,5)",
    "stavka2": "Ф2(-13)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(13,5)",
    "stavka2": "Ф2(-13,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(14)",
    "stavka2": "Ф2(-13,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(14)",
    "stavka2": "Ф2(-14)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(14,5)",
    "stavka2": "Ф1(-14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(14,5)",
    "stavka2": "Ф2(-14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(14,5)",
    "stavka2": "Ф2(-14,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(15)",
    "stavka2": "Ф1(-14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(15)",
    "stavka2": "Ф2(-14)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(15)",
    "stavka2": "Ф2(-14,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(15)",
    "stavka2": "Ф2(-15)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(15,5)",
    "stavka2": "Ф1(-15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(15,5)",
    "stavka2": "Ф2(-15)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(15,5)",
    "stavka2": "Ф2(-15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(16)",
    "stavka2": "Ф1(-16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(16)",
    "stavka2": "Ф2(-15)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(16)",
    "stavka2": "Ф2(-15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(16)",
    "stavka2": "Ф2(-15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(16)",
    "stavka2": "Ф2(-16)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(16,5)",
    "stavka2": "Ф2(-15,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(16,5)",
    "stavka2": "Ф2(-15,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(16,5)",
    "stavka2": "Ф2(-16,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(17)",
    "stavka2": "Ф2(-17)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(17,5)",
    "stavka2": "Ф1(-17,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(18)",
    "stavka2": "Ф2(-17,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(18,5)",
    "stavka2": "Ф2(-18,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(2)",
    "stavka2": "Ф1(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2)",
    "stavka2": "Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(2)",
    "stavka2": "Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(2)",
    "stavka2": "Ф2(-1,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2)",
    "stavka2": "Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(2,25)",
    "stavka2": "Ф1(-2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2,25)",
    "stavka2": "Ф2(-2,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(2,5)",
    "stavka2": "П2",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2,5)",
    "stavka2": "Ф1(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2,5)",
    "stavka2": "Ф2(-1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(2,5)",
    "stavka2": "Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(2,5)",
    "stavka2": "Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2,5)",
    "stavka2": "Ф2(-2,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(2,5)",
    "stavka2": "Ф2(1)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(20)",
    "stavka2": "Ф2(-20)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(20,5)",
    "stavka2": "Ф1(-20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(20,5)",
    "stavka2": "Ф2(-20,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(21)",
    "stavka2": "Ф1(-21)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(21,5)",
    "stavka2": "Ф2(-20,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(21,5)",
    "stavka2": "Ф2(-21)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(21,5)",
    "stavka2": "Ф2(-21,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(22,5)",
    "stavka2": "Ф2(-22)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(22,5)",
    "stavka2": "Ф2(-22,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(23)",
    "stavka2": "Ф2(-23)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(23,5)",
    "stavka2": "Ф1(-22,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(23,5)",
    "stavka2": "Ф2(-23,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(24,5)",
    "stavka2": "Ф2(-24,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(27,5)",
    "stavka2": "Ф2(-27,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(28,5)",
    "stavka2": "Ф2(-28)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(28,5)",
    "stavka2": "Ф2(-28,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(29)",
    "stavka2": "Ф2(-29)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(3)",
    "stavka2": "Ф1(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3)",
    "stavka2": "Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(3)",
    "stavka2": "Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(3)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(3)",
    "stavka2": "Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3,25)",
    "stavka2": "Ф1(-3,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3,25)",
    "stavka2": "Ф2(-3,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(3,5)",
    "stavka2": "Ф1(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3,5)",
    "stavka2": "Ф2(-1,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(3,5)",
    "stavka2": "Ф2(-2)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3,5)",
    "stavka2": "Ф2(-2,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(3,5)",
    "stavka2": "Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(3,5)",
    "stavka2": "Ф2(-3)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3,5)",
    "stavka2": "Ф2(-3,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(3,75)",
    "stavka2": "Ф2(-3,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(30)",
    "stavka2": "Ф1(-30)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(30,5)",
    "stavka2": "Ф1(-30,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(30,5)",
    "stavka2": "Ф2(-30,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(32,5)",
    "stavka2": "Ф2(-32,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(4)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(4)",
    "stavka2": "Ф1(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(4)",
    "stavka2": "Ф1(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(4)",
    "stavka2": "Ф2(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(4)",
    "stavka2": "Ф2(-4)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(4,5)",
    "stavka2": "П2",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(4,5)",
    "stavka2": "Ф1(-4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(4,5)",
    "stavka2": "Ф1(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(4,5)",
    "stavka2": "Ф2(-2)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(4,5)",
    "stavka2": "Ф2(-4)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(4,5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(4,75)",
    "stavka2": "Ф1(-4,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(4,75)",
    "stavka2": "Ф2(-4,75)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(49,5)",
    "stavka2": "Ф2(-49,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(5)",
    "stavka2": "Ф1(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(5)",
    "stavka2": "Ф2(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(5)",
    "stavka2": "Ф2(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(5,25)",
    "stavka2": "Ф2(-5,25)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(5,5)",
    "stavka2": "Ф1(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(5,5)",
    "stavka2": "Ф1(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(5,5)",
    "stavka2": "Ф2(-3)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(5,5)",
    "stavka2": "Ф2(-3,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(5,5)",
    "stavka2": "Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(5,5)",
    "stavka2": "Ф2(-5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(5,5)",
    "stavka2": "Ф2(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(58,5)",
    "stavka2": "Ф1(-58,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(6)",
    "stavka2": "Ф1(-6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(6)",
    "stavka2": "Ф2(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(6)",
    "stavka2": "Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(6)",
    "stavka2": "Ф2(-6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(6,5)",
    "stavka2": "Ф1(-6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(6,5)",
    "stavka2": "Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(6,5)",
    "stavka2": "Ф2(-6)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(6,5)",
    "stavka2": "Ф2(-6)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(6,5)",
    "stavka2": "Ф2(-6,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(7)",
    "stavka2": "Ф1(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(7)",
    "stavka2": "Ф1(-7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(7)",
    "stavka2": "Ф2(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(7)",
    "stavka2": "Ф2(-7)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(7,5)",
    "stavka2": "Ф1(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(7,5)",
    "stavka2": "Ф2(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(7,5)",
    "stavka2": "Ф2(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(8)",
    "stavka2": "Ф1(-8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(8)",
    "stavka2": "Ф2(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8)",
    "stavka2": "Ф2(-6,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8)",
    "stavka2": "Ф2(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8)",
    "stavka2": "Ф2(-7,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8)",
    "stavka2": "Ф2(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(8)",
    "stavka2": "Ф2(-8)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф1(-4,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф1(-4,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф1(-5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф1(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф1(-5,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф1(-7,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф1(-8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф1(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф2(-7)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф2(-8)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(8,5)",
    "stavka2": "Ф2(-8,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(9)",
    "stavka2": "Ф2(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(9)",
    "stavka2": "Ф2(-9)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Ф2(9,5)",
    "stavka2": "Ф1(-5,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(9,5)",
    "stavka2": "Ф2(-8,5)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(9,5)",
    "stavka2": "Ф2(-9)",
    "tip": "KARIDOR"
  },
  {
    "stavka1": "Ф2(9,5)",
    "stavka2": "Ф2(-9,5)",
    "tip": "VILKA"
  },
  {
    "stavka1": "Чёт",
    "stavka2": "Нечёт",
    "tip": "VILKA"
  }
]
`
