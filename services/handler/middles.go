package handler

import (
	pb "github.com/aibotsoft/gen/fortedpb"
)

func CalcMiddle(sb *pb.Surebet) {
	m1 := sb.Members[0].Market
	m2 := sb.Members[1].Market
	if m1.Side == OVER {
		sb.Calc.MiddleDiff = m2.Handicap - m1.Handicap
	} else if m1.Side == UNDER {
		sb.Calc.MiddleDiff = m1.Handicap - m2.Handicap
	} else {
		sb.Calc.MiddleDiff = m1.Handicap + m2.Handicap
	}
}

//if m1.BetType == TOTAL_POINTS && m2.BetType == TOTAL_POINTS {
//	fmt.Println(m1, m2)
//	if m1.Side == OVER {
//		fmt.Println("m1.Side == OVER")
//		sb.Calc.MiddleDiff = m2.Handicap - m1.Handicap
//	} else {
//		sb.Calc.MiddleDiff = m1.Handicap - m2.Handicap
//	}
//}
//if m1.BetType == TEAM_TOTAL_POINTS && m2.BetType == TEAM_TOTAL_POINTS {
//	if m1.Side == OVER {
//		sb.Calc.MiddleDiff = m2.Handicap - m1.Handicap
//	} else {
//		sb.Calc.MiddleDiff = m1.Handicap - m2.Handicap
//	}
//}
//if m1.BetType == SPREAD && m2.BetType == SPREAD {
//	sb.Calc.MiddleDiff = m1.Handicap + m2.Handicap
//}
//if m1.BetType == MONEYLINE && m2.BetType == MONEYLINE {
//	sb.Calc.MiddleDiff = m1.Handicap + m2.Handicap
//}
//if m1.BetType == MONEYLINE && m2.BetType == SPREAD {
//	sb.Calc.MiddleDiff = m2.Handicap + m1.Handicap
//}
//if m1.BetType == SPREAD && m2.BetType == MONEYLINE {
//	sb.Calc.MiddleDiff = m1.Handicap + m2.Handicap
//}
