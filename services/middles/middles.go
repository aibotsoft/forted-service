package middles

import (
	"github.com/pkg/errors"
)

func CalcMiddle(marketName1 string, marketName2 string) (diff float64, err error) {
	m1, err := Convert(marketName1)
	if err != nil {
		return diff, errors.Wrapf(err, "name: %q", marketName1)
	}
	m2, err := Convert(marketName2)
	if err != nil {
		return diff, errors.Wrapf(err, "name: %q", marketName2)
	}

	if m1.BetType == TOTAL_POINTS && m2.BetType == TOTAL_POINTS {
		if m1.Side == OVER {
			return *m2.Handicap - *m1.Handicap, nil
		} else {
			return *m1.Handicap - *m2.Handicap, nil

		}
	}
	if m1.BetType == TEAM_TOTAL_POINTS && m2.BetType == TEAM_TOTAL_POINTS {
		if m1.Side == OVER {
			return *m2.Handicap - *m1.Handicap, nil
		} else {
			return *m1.Handicap - *m2.Handicap, nil
		}
	}
	if m1.BetType == SPREAD && m2.BetType == SPREAD {
		return *m1.Handicap + *m2.Handicap, nil
		//return diff, nil
	}
	if m1.BetType == MONEYLINE && m2.BetType == MONEYLINE {
		return *m1.Handicap + *m2.Handicap, nil
		//return diff, nil
	}
	if m1.BetType == MONEYLINE && m2.BetType == SPREAD {
		return *m2.Handicap + *m1.Handicap, nil
		//return diff, nil
	}
	if m1.BetType == SPREAD && m2.BetType == MONEYLINE {
		return *m1.Handicap + *m2.Handicap, nil
		//return diff, nil
	}
	return diff, nil
}
