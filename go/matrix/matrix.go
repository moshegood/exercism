package matrix

type matrix struct{}
type Matrix = *matrix

func (m *matrix) Cols() [][]int {
	return nil
}

func (m *matrix) Rows() [][]int {
	return nil
}

func (m *matrix) Set(row, col, value int) bool {
	return true
}

func New(input string) (Matrix, error) {
	return &matrix{}, nil
}
