package perfume

//PrintBuffer operate texts for printing
type PrintBuffer struct {
	size       Size
	colChanges []uint
	rowChanges []uint
	buffer     string //It doesn't contain \n. \n only adds in print
}

//NewPrintBuffer makes new print buffer
func NewPrintBuffer(s Size) PrintBuffer {

	pb := PrintBuffer{
		colChanges: make([]uint, 0),
		rowChanges: make([]uint, 0),
	}
	pb.Resize(s)

	return pb
}

//Resize resizes buffer size & sets
func (buffer *PrintBuffer) Resize(s Size) {
	for i := 0; i < int(s.Height*s.Width); i++ {
		buffer.buffer += " "
	}
	buffer.size = s
}

//GetLine return line using split
func (buffer PrintBuffer) GetLine(index uint) string {
	width := buffer.size.Width
	start := index * width
	return buffer.buffer[start : start+width]
}

func insert(origin string, index int, str string) string {
	return origin[:index] + str + origin[index+len(str):]
}
func startEndIdxCheck(start uint, end uint, max uint) (bool, error) {
	if start > end {
		return false, ErrStartIndexOverEndIndex
	} else if end > max {
		return false, ErrEndIndexOverMax
	}
	return true, nil
}
func getFillPattern(length int, pattern string) string {
	patternGroup := ""
	fillRadio := (length / len(pattern))
	for i := 0; i < fillRadio; i++ {
		patternGroup += pattern
	}
	patternGroup += pattern[:length-len(patternGroup)]
	return patternGroup
}

//SetRow set row 0 ~ width
func (buffer *PrintBuffer) SetRow(pattern string, row uint, start uint, end uint) error {
	size := buffer.size
	startIdx := int(row*size.Width + start)
	setLength := int(end - start)

	if ok, err := startEndIdxCheck(start, end, size.Width); !ok {
		return err
	} else if row >= size.Height {
		return ErrOverSize
	} else if len(pattern) > setLength {
		return ErrOverSize
	}

	text := getFillPattern(setLength, pattern)

	buffer.buffer = insert(buffer.buffer, startIdx, text)

	buffer.rowChanges = append(buffer.rowChanges, row)

	return nil
}

//SetColumn set column 0 ~ height
func (buffer *PrintBuffer) SetColumn(pattern string, col uint, start uint, end uint) error {
	size := buffer.size
	setLength := int(end - start)

	if ok, err := startEndIdxCheck(start, end, size.Height); !ok {
		return err
	} else if col >= size.Width {
		return ErrOverSize
	} else if len(pattern) > setLength {
		return ErrOverSize
	}

	text := getFillPattern(setLength, pattern)
	buff := []rune(buffer.buffer)
	for i := start; i < end; i++ {
		row := i * size.Width
		buff[row+col] = rune(text[i-start])
	}
	buffer.buffer = string(buff)

	buffer.colChanges = append(buffer.colChanges, col)

	return nil
}

//GetChanges return edited lines - it re-initalizes when calling Reinital
func (buffer PrintBuffer) GetChanges() (col []uint, row []uint) {
	col = buffer.colChanges
	row = buffer.rowChanges
	return
}

//ApplyChanges initalize changes array.
func (buffer *PrintBuffer) ApplyChanges() {
	//fmt.Printf("warning, reinitalize printbuffer changes")
	buffer.colChanges = make([]uint, 0)
	buffer.rowChanges = make([]uint, 0)
}
