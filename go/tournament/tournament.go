package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type teamScores struct {
	mp, w, d, l, p int
}

func Tally(input io.Reader, output io.Writer) error {
	results := make(map[string]teamScores)
	reader := bufio.NewScanner(input)
	for reader.Scan() {
		game := reader.Text()
		if game == "" {
			continue
		}
		if strings.HasPrefix(game, "#") {
			continue
		}
		parts := strings.Split(game, ";")
		if len(parts) != 3 {
			return fmt.Errorf("bad line %q", game)
		}

		switch parts[2] {
		case "win":
			results[parts[0]] = addWin(results[parts[0]])
			results[parts[1]] = addLoss(results[parts[1]])
		case "loss":
			results[parts[0]] = addLoss(results[parts[0]])
			results[parts[1]] = addWin(results[parts[1]])
		case "draw":
			results[parts[0]] = addDraw(results[parts[0]])
			results[parts[1]] = addDraw(results[parts[1]])
		default:
			return fmt.Errorf("bad line %q", game)
		}
	}

	allTeams := []string{}
	for k := range results {
		allTeams = append(allTeams, k)
	}

	sort.Slice(allTeams, func(i, j int) bool {
		if diff := results[allTeams[i]].p - results[allTeams[j]].p; diff > 0 {
			return true
		} else if diff < 0 {
			return false
		}
		return allTeams[i] < allTeams[j]
	})

	fmt.Fprintln(output, "Team                           | MP |  W |  D |  L |  P")
	for _, team := range allTeams {
		scores := results[team]
		fmt.Fprintf(output, "%-31s|%3d |%3d |%3d |%3d |%3d\n", team, scores.mp, scores.w, scores.d, scores.l, scores.p)
	}
	return nil
}

func addWin(score teamScores) teamScores {
	score.mp++
	score.w++
	score.p += 3
	return score
}

func addLoss(score teamScores) teamScores {
	score.mp++
	score.l++
	return score
}

func addDraw(score teamScores) teamScores {
	score.mp++
	score.d++
	score.p++
	return score
}
