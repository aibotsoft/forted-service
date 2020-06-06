package middles

import (
	"github.com/aibotsoft/micro/util"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

type Market struct {
	BetType  string
	Team     string
	Side     string
	Handicap *float64
}

const (
	SPREAD            = "SPREAD"
	MONEYLINE         = "MONEYLINE"
	TOTAL_POINTS      = "TOTAL_POINTS"
	TEAM_TOTAL_POINTS = "TEAM_TOTAL_POINTS"
	OVER              = "OVER"
	UNDER             = "UNDER"
	Team1             = "Team1"
	Team2             = "Team2"
	Draw              = "Draw"
)

var teamSide = map[string]string{"1": Team1, "2": Team2, "Х": Draw}
var totalSide = map[string]string{"Б": OVER, "М": UNDER}

var moneyRe = regexp.MustCompile(`П(\d)$`)
var threeRe = regexp.MustCompile(`(^|\s)([12Х]$)`)
var handicapRe = regexp.MustCompile(`Ф(\d)\((-?\d+,?\d{0,3})\)`)
var totalRe = regexp.MustCompile(`Т([МБ])\((-?\d+,?\d{0,3})\)`)
var teamTotalRe = regexp.MustCompile(`ИТ(\d)([МБ])\((-?\d+,?\d{0,3})\)`)

var DoubleChanceRe = regexp.MustCompile(`(1Х|Х2|12)$`)
var OddEvenRe = regexp.MustCompile(`(Чёт|Нечёт)$`)
var NotConvertedError = errors.New("NotConverted")

func Convert(name string) (Market, error) {
	if name == "" {
		return Market{}, errors.Errorf("cannot convert empty str")
	}
	var found []string

	found = moneyRe.FindStringSubmatch(name)
	if len(found) == 2 {
		return processMoney(name, found)
	}

	found = threeRe.FindStringSubmatch(name)
	if len(found) == 3 {
		return processThreeWay(name, found)
	}

	found = handicapRe.FindStringSubmatch(name)
	if len(found) == 3 {
		return processHandicap(name, found)
	}

	found = totalRe.FindStringSubmatch(name)
	if len(found) == 3 {
		return processTotal(name, found)
	}
	found = teamTotalRe.FindStringSubmatch(name)
	if len(found) == 4 {
		return processTeamTotal(name, found)
	}
	found = DoubleChanceRe.FindStringSubmatch(name)
	if len(found) == 2 {
		return processDoubleChance(name, found)
	}
	found = OddEvenRe.FindStringSubmatch(name)
	if len(found) == 2 {
		return Market{}, nil
	}
	return Market{}, NotConvertedError
}
func processPoint(point string) (float64, error) {
	return strconv.ParseFloat(strings.Replace(point, ",", ".", -1), 64)
}

func processMoney(market string, found []string) (Market, error) {
	m := Market{
		BetType:  MONEYLINE,
		Team:     teamSide[found[1]],
		Handicap: util.PtrFloat64(0),
	}
	return m, nil
}
func processThreeWay(market string, found []string) (Market, error) {
	m := Market{
		BetType:  MONEYLINE,
		Team:     teamSide[found[2]],
		Handicap: util.PtrFloat64(-0.5),
	}
	return m, nil
}

func processDoubleChance(market string, found []string) (Market, error) {
	m := Market{
		BetType: MONEYLINE,
		//Team:     teamSide[found[2]],
		Handicap: util.PtrFloat64(0.5),
	}
	return m, nil
}
func processHandicap(market string, found []string) (Market, error) {
	point, err := processPoint(found[2])
	if err != nil {
		return Market{}, err
	}
	m := Market{
		BetType:  SPREAD,
		Team:     teamSide[found[1]],
		Handicap: &point,
	}
	return m, nil
}

func processTotal(market string, found []string) (Market, error) {
	point, err := processPoint(found[2])
	if err != nil {
		return Market{}, err
	}
	m := Market{
		BetType:  TOTAL_POINTS,
		Side:     totalSide[found[1]],
		Handicap: &point,
	}
	return m, nil
}
func processTeamTotal(market string, found []string) (Market, error) {
	point, err := processPoint(found[3])
	if err != nil {
		return Market{}, err
	}
	m := Market{
		BetType:  TEAM_TOTAL_POINTS,
		Team:     teamSide[found[1]],
		Side:     totalSide[found[2]],
		Handicap: &point,
	}
	return m, nil
}
