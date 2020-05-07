package gojumanpp

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mattn/go-pipeline"
)

type Jumanpp struct {
	mrph_list []*Morpheme
}

func NewJumanpp() *Jumanpp {
	j := new(Jumanpp)
	j.mrph_list = make([]*Morpheme, 0)

	return j
}

func (j *Jumanpp) Analysis(text string) error {
	out, err := pipeline.Output(
		[]string{"echo", text},
		[]string{"jumanpp"},
	)
	if err != nil {
		return err
	}

	lines := strings.Split(string(out), "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		if line != "EOS" {
			j.parseLine(line)
		}
	}

	return nil
}
func (j *Jumanpp) Result(line string) {
	j.parseLine(line)
}
func (j *Jumanpp) ResultAll(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "EOS" {
			j.parseLine(line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
func (j *Jumanpp) parseLine(line string) {
	split := strings.Split(line, " ")

	m := NewMorpheme()
	m.Midasi = split[0]
	m.Yomi = split[1]
	m.Genkei = split[2]
	m.Hinsi = split[3]
	m.Hinsi_id, _ = strconv.Atoi(split[4])
	m.Bunrui = split[5]
	m.Bunrui_id, _ = strconv.Atoi(split[6])
	m.Katuyou1 = split[7]
	m.Katuyou1_id, _ = strconv.Atoi(split[8])
	m.Katuyou2 = split[9]
	m.Katuyou2_id, _ = strconv.Atoi(split[10])

	rep, _ := regexp.Compile(`".*"`)
	m.Imis = rep.FindString(line)
	m.Imis = strings.Replace(m.Imis, `"`, "", -1)
	if m.Imis == "" {
		m.Imis = "NIL"
	}

	split2 := strings.Split(m.Imis, " ")
	for _, str := range split2 {
		if strings.Contains(str, "代表表記") {
			m.Repname = strings.Split(str, ":")[1]
		}
	}

	j.mrph_list = append(j.mrph_list, m)
}

func (j *Jumanpp) MrphList() []*Morpheme {
	return j.mrph_list
}
