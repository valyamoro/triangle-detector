package artifact

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherchan2006/go-triangle-detector/internal/detect"
)

type Names struct {
	GroupDir                 string
	HTMLTmp                  string
	PNG                      string
	DebugTxt                 string
	CalcATRTxt               string
	SwingTxt                 string
	HorizTxt                 string
	CheckTimingTxt           string
	FindValleysTxt           string
	ValidateValleysTxt       string
	FitSupportLineTxt        string
	CheckGeometryTxt         string
	CheckVolumeTxt           string
}

func NewNames(baseDir, stem string) Names {
	groupDir := filepath.Join(baseDir, stem)
	return Names{
		GroupDir:           groupDir,
		HTMLTmp:            filepath.Join(baseDir, stem+"_render.tmp.html"),
		PNG:                filepath.Join(groupDir, fmt.Sprintf("1_%s_1.png", stem)),
		DebugTxt:           filepath.Join(groupDir, fmt.Sprintf("2_%s_2.txt", stem)),
		CalcATRTxt:         filepath.Join(groupDir, fmt.Sprintf("3_%s_calcATR_3.txt", stem)),
		SwingTxt:           filepath.Join(groupDir, fmt.Sprintf("4_%s_findSwingHighs_4.txt", stem)),
		HorizTxt:           filepath.Join(groupDir, fmt.Sprintf("5_%s_findHorizontalResistance_5.txt", stem)),
		CheckTimingTxt:     filepath.Join(groupDir, fmt.Sprintf("6_%s_checkTimingAndHighs_6.txt", stem)),
		FindValleysTxt:     filepath.Join(groupDir, fmt.Sprintf("7_%s_findValleys_7.txt", stem)),
		ValidateValleysTxt: filepath.Join(groupDir, fmt.Sprintf("8_%s_validateValleys_8.txt", stem)),
		FitSupportLineTxt:  filepath.Join(groupDir, fmt.Sprintf("9_%s_fitSupportLine_9.txt", stem)),
		CheckGeometryTxt:   filepath.Join(groupDir, fmt.Sprintf("10_%s_checkGeometry_10.txt", stem)),
		CheckVolumeTxt:     filepath.Join(groupDir, fmt.Sprintf("11_%s_checkVolume_11.txt", stem)),
	}
}

func WriteTexts(names Names, result detect.Result, writeFn func(path string, result detect.Result)) {
	writeFn(names.DebugTxt, result)
	writeLogTxt(names.CalcATRTxt, result.Debug.Logs.CalcATR)
	writeLogTxt(names.SwingTxt, result.Debug.Logs.FindSwingHighs)
	writeLogTxt(names.HorizTxt, result.Debug.Logs.FindHorizontalResistance)
	writeLogTxt(names.CheckTimingTxt, result.Debug.Logs.CheckTimingAndHighs)
	writeLogTxt(names.FindValleysTxt, result.Debug.Logs.FindValleys)
	writeLogTxt(names.ValidateValleysTxt, result.Debug.Logs.ValidateValleys)
	writeLogTxt(names.FitSupportLineTxt, result.Debug.Logs.FitSupportLine)
	writeLogTxt(names.CheckGeometryTxt, result.Debug.Logs.CheckGeometry)
	writeLogTxt(names.CheckVolumeTxt, result.Debug.Logs.CheckVolume)
}

func WriteLogTxt(path, content string) {
	writeLogTxt(path, content)
}

func writeLogTxt(path, content string) {
	if strings.TrimSpace(content) == "" {
		return
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		log.Printf("writeLogTxt %s: %v", path, err)
	}
}
