package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type matrix struct {
	storage [][]int
}
type Matrix = *matrix

func (m *matrix) Size() (int, int) {
	x := len(m.storage)
	if x == 0 {
		return 0, 0
	}
	return x, len(m.storage[0])
}

func (m *matrix) Cols() [][]int {
	y, x := m.Size()
	if x == 0 {
		return nil
	}
	clone := make([][]int, x)
	for i := range clone {
		clone[i] = make([]int, y)
	}
	for i, row := range m.storage {
		for j, value := range row {
			clone[j][i] = value
		}
	}
	return clone
}

func (m *matrix) Rows() [][]int {
	x, y := m.Size()
	if x == 0 {
		return nil
	}
	clone := make([][]int, x)
	for i := range clone {
		clone[i] = make([]int, y)
	}
	for i, row := range m.storage {
		for j, value := range row {
			clone[i][j] = value
		}
	}
	return clone
}

func (m *matrix) Set(i, j, value int) bool {
	x, y := m.Size()
	if i < 0 || j < 0 || i >= x || j >= y {
		return false
	}
	m.storage[i][j] = value

	return true
}

func New(input string) (Matrix, error) {
	m := &matrix{}

	rows := strings.Split(input, "\n")
	columnCount := 0
	for i, row := range rows {
		cols := strings.Fields(row)
		newRow := []int{}
		for _, value := range cols {
			number, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			newRow = append(newRow, number)
		}
		if i == 0 {
			columnCount = len(cols)
		} else {
			if columnCount != len(cols) {
				return nil, fmt.Errorf("Unequal number of columns per row")
			}
		}
		m.storage = append(m.storage, newRow)
	}

	return m, nil
}
