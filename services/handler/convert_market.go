package handler

import (
	"fmt"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	SPREAD            = "SPREAD"
	MONEYLINE         = "MONEYLINE"
	TOTAL_POINTS      = "TOTAL_POINTS"
	TEAM_TOTAL_POINTS = "TEAM_TOTAL_POINTS"
	OddEven           = "ODD_EVEN"
	OVER              = "OVER"
	UNDER             = "UNDER"
	Team1             = "Team1"
	Team2             = "Team2"
	Draw              = "Draw"
	Team1_or_Team2    = "12"
)

var teamSide = map[string]string{
	"1:false":     Team1,
	"1:true":      Team2,
	"2:false":     Team2,
	"2:true":      Team1,
	"Чёт:false":   "even",
	"Чёт:true":    "odd",
	"Нечёт:false": "odd",
	"Нечёт:true":  "even",
}
var threeWaySide = map[string]string{
	"1:false":  Team1,
	"1Х:false": Team1,
	"1X:false": Team1,
	"1Х:true":  Team2,
	"1X:true":  Team2,

	"1:true": Team2,

	"2:false":  Team2,
	"Х2:false": Team2,
	"X2:false": Team2,
	"Х2:true":  Team1,
	"X2:true":  Team1,

	"2:true": Team1,

	"Х:false": Draw,
	"X:false": Draw,
	"Х:true":  Team1_or_Team2,
	"X:true":  Team1_or_Team2,

	"12:false": Team1_or_Team2,
	"12:true":  Draw,
}

var totalSide = map[string]string{
	"Б:false": OVER,
	"Б:true":  UNDER,
	"М:false": UNDER,
	"М:true":  OVER,
}

var moneyRe = regexp.MustCompile(`П(\d)$`)
var threeRe = regexp.MustCompile(`(^|\s)([12ХX]$)`)
var handicapRe = regexp.MustCompile(`Ф(\d)\((-?\d+,?\d{0,3})\)`)
var totalRe = regexp.MustCompile(`Т([МБ])\((-?\d+,?\d{0,3})\)`)
var teamTotalRe = regexp.MustCompile(`ИТ(\d)([МБ])\((-?\d+,?\d{0,3})\)`)
var DoubleChanceRe = regexp.MustCompile(`(1Х|Х2|12|1X|X2)$`)
var OddEvenRe = regexp.MustCompile(`(Чёт|Нечёт)$`)
var EuroHandRe = regexp.MustCompile(`(\d)\((\d+):(\d+)\)`)
var NotConvertedError = errors.New("NotConverted")

func Convert(side *pb.SurebetSide) error {
	var found []string
	side.Market = &pb.Market{}
	if strings.Index(side.MarketName, "Против") != -1 {
		side.Market.IsLay = true
	}

	found = moneyRe.FindStringSubmatch(side.MarketName)
	if len(found) == 2 {
		processMoney(found, side.Market)
		return nil
	}

	found = threeRe.FindStringSubmatch(side.MarketName)
	if len(found) == 3 {
		processThreeWay(found, side.Market)
		return nil
	}

	found = handicapRe.FindStringSubmatch(side.MarketName)
	if len(found) == 3 {
		processHandicap(found, side.Market)
		return nil
	}

	found = totalRe.FindStringSubmatch(side.MarketName)
	if len(found) == 3 {
		processTotal(found, side.Market)
		return nil
	}
	found = teamTotalRe.FindStringSubmatch(side.MarketName)
	if len(found) == 4 {
		processTeamTotal(found, side.Market)
		return nil
	}
	found = DoubleChanceRe.FindStringSubmatch(side.MarketName)
	if len(found) == 2 {
		processDoubleChance(found, side.Market)
		return nil
	}
	found = OddEvenRe.FindStringSubmatch(side.MarketName)
	if len(found) == 2 {
		processOddEven(found, side.Market)
		return nil
	}
	found = EuroHandRe.FindStringSubmatch(side.MarketName)
	if len(found) == 4 {
		processEuroHand(found, side.Market)
		return nil
	}
	return NotConvertedError
}
func processPoint(point string) (float64, error) {
	return strconv.ParseFloat(strings.Replace(point, ",", ".", -1), 64)
}

func processOddEven(found []string, m *pb.Market) {
	m.BetType = OddEven
	teamKey := fmt.Sprintf("%v:%v", found[1], m.IsLay)
	m.Team = teamSide[teamKey]
	m.Handicap = 0
}
func processMoney(found []string, m *pb.Market) {
	m.BetType = MONEYLINE
	teamKey := fmt.Sprintf("%v:%v", found[1], m.IsLay)

	m.Team = teamSide[teamKey]
	m.Handicap = 0

}
func processThreeWay(found []string, m *pb.Market) {
	m.BetType = MONEYLINE
	teamKey := fmt.Sprintf("%v:%v", found[2], m.IsLay)
	//fmt.Println("team:", m.Team, "teamKey", teamKey)

	m.Team = threeWaySide[teamKey]
	m.Handicap = -0.5
	if m.IsLay {
		m.Handicap = 0.5
	}
}

func processDoubleChance(found []string, m *pb.Market) {
	m.BetType = MONEYLINE
	teamKey := fmt.Sprintf("%v:%v", found[1], m.IsLay)
	m.Team = threeWaySide[teamKey]
	m.Handicap = 0.5
	if m.IsLay {
		m.Handicap = -0.5
	}
}
func processHandicap(found []string, m *pb.Market) {
	m.BetType = SPREAD
	teamKey := fmt.Sprintf("%v:%v", found[1], m.IsLay)

	m.Team = teamSide[teamKey]

	point, err := processPoint(found[2])
	if err != nil {
		return
	}
	m.Handicap = point
	if m.IsLay {
		m.Handicap = m.Handicap * -1
	}
}
func processEuroHand(found []string, m *pb.Market) {
	m.BetType = SPREAD
	teamKey := fmt.Sprintf("%v:%v", found[1], m.IsLay)

	m.Team = teamSide[teamKey]

	hp, err := processPoint(found[2])
	if err != nil {
		return
	}
	ap, err := processPoint(found[3])
	if err != nil {
		return
	}
	if found[1] == "1" {
		m.Handicap = hp - ap - 0.5
	} else if found[1] == "2" {
		m.Handicap = ap - hp - 0.5
	}
	if m.IsLay {
		m.Handicap = m.Handicap * -1
	}
}

func processTotal(found []string, m *pb.Market) {
	m.BetType = TOTAL_POINTS
	totalKey := fmt.Sprintf("%v:%v", found[1], m.IsLay)
	m.Side = totalSide[totalKey]
	//fmt.Println(totalKey)
	point, err := processPoint(found[2])
	if err != nil {
		return
	}
	m.Handicap = point
}
func processTeamTotal(found []string, m *pb.Market) {
	m.BetType = TEAM_TOTAL_POINTS
	teamKey := fmt.Sprintf("%v:%v", found[1], false)
	m.Team = teamSide[teamKey]
	totalKey := fmt.Sprintf("%v:%v", found[2], m.IsLay)
	//fmt.Println(m.Team)
	m.Side = totalSide[totalKey]
	point, err := processPoint(found[3])
	if err != nil {
		return
	}
	m.Handicap = point
}
