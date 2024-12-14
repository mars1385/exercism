package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Score struct {
	mp int
	w  int
	l  int
	d  int
	p  int
}

type FinalData struct {
	Key   string
	Value int
}

func createTeam(scoreList map[string]Score, team string) {
	_, has := scoreList[team]

	if !has {
		scoreList[team] = Score{
			mp: 0,
			w:  0,
			l:  0,
			d:  0,
			p:  0,
		}

	}

}

func Tally(reader io.Reader, writer io.Writer) error {

	text, err := io.ReadAll(reader)

	if err != nil {
		return err
	}

	texts := strings.Split(strings.Trim(string(text), ""), "\n")

	scoreList := map[string]Score{}
	for _, val := range texts {
		slice := strings.Split(val, ";")

		if len(slice) == 3 {
			if slice[2] != "draw" && slice[2] != "win" && slice[2] != "loss" {
				return errors.New("invalid data")
			}
			createTeam(scoreList, strings.TrimSpace(slice[0]))
			createTeam(scoreList, strings.TrimSpace(slice[1]))

			teamA := scoreList[strings.TrimSpace(slice[0])]
			teamB := scoreList[strings.TrimSpace(slice[1])]

			teamA.mp += 1
			teamB.mp += 1

			if slice[2] == "draw" {
				teamA.d += 1
				teamB.d += 1

				teamA.p += 1
				teamB.p += 1
			}

			if slice[2] == "win" {
				teamA.w += 1
				teamB.l += 1

				teamA.p += 3

			}

			if slice[2] == "loss" {
				teamA.l += 1
				teamB.w += 1

				teamB.p += 3

			}
			scoreList[strings.TrimSpace(slice[0])] = teamA
			scoreList[strings.TrimSpace(slice[1])] = teamB
		}

	}

	if len(scoreList) <= 0 {
		return errors.New("invalid data")
	}

	writer.Write([]byte(fmt.Sprintf("%-*s| MP |  W |  D |  L |  P", 31, "Team")))

	var items []FinalData
	for key, val := range scoreList {

		items = append(items, FinalData{
			Key:   fmt.Sprintf("%-*s|  %v |  %v |  %v |  %v |  %v", 31, key, val.mp, val.w, val.d, val.l, val.p),
			Value: val.p,
		})

	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Key < items[j].Key
	})
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	for _, val := range items {
		writer.Write([]byte("\n"))

		writer.Write([]byte(val.Key))

	}
	writer.Write([]byte("\n"))

	return nil
}
