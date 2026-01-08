// main.go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/tealeg/xlsx"
)

// xlsx_helper is a mock helper library.
var xlsx_helper = &mockXLSXHelper{}

// Info holds basic game configuration.
type Info struct {
	Width int
}

// Symbol represents a single game symbol with its properties.
type Symbol struct {
	Index      int
	SymbolCode string
	Name       string
	Type       string
	ReelDist   []int
	PayoutMap  map[int]decimal.Decimal
}

// Game holds the final parsed configuration.
type Game struct {
	premiumSymbolList []*Symbol
	f2pSymbolList     []*Symbol
}

const (
	kSymbolPremiumSheetName = "Symbol_Premium"
	kSymbolF2PSheetName     = "Symbol_F2P"
	kConfigFilename         = "game_config.xlsx"
)

func LoadGameConfig(filePath string, info *Info) (*Game, error) {
	xlsxFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}

	game := &Game{}
	var parsingErr error
	// === Parse Premium Symbols ===
	sheetPremium := xlsxFile.Sheet[kSymbolPremiumSheetName]
	symbolListPremium := []*Symbol{}
	if sheetPremium != nil {
		symbolCounts := []int{}
		for rowIndex, row := range sheetPremium.Rows {
			symbol := &Symbol{
				ReelDist:  make([]int, info.Width),
				PayoutMap: make(map[int]decimal.Decimal),
			}
			index := 0
			for colIndex, cell := range row.Cells {
				if rowIndex == 0 { // Header row
					if colIndex < 3+info.Width {
						continue
					}
					val := strings.TrimPrefix(cell.String(), "Payout ")
					value := xlsx_helper.CellValToInt(val)
					symbolCounts = append(symbolCounts, value)
				} else { // Data row
					if cell.String() == "" && colIndex == 0 {
						break // Stop processing row if the first cell is empty
					}
					if colIndex == 0 {
						symbol.SymbolCode = strings.TrimSpace(cell.String())
					} else if colIndex == 1 {
						symbol.Name = strings.TrimSpace(cell.String())
					} else if colIndex == 2 {
						symbol.Type = strings.TrimSpace(cell.String())
					} else if colIndex >= 3 && colIndex < 3+info.Width {
						// reels
						symbol.ReelDist[colIndex-3], parsingErr = cell.Int()
						if parsingErr != nil {
							symbol.ReelDist[colIndex-3] = 0 // Silently ignores error
						}
					} else if colIndex >= 3+info.Width {
						// payout
						payout := xlsx_helper.ReadDecimalFromCell(cell)
						symbol.PayoutMap[symbolCounts[index]] = payout
						index++
					}
				}
			}
			if symbol.SymbolCode != "" {
				symbol.Index = len(symbolListPremium)
				symbolListPremium = append(symbolListPremium, symbol)
			}
		}
		game.premiumSymbolList = symbolListPremium
	}

	// === Parse F2P Symbols ===
	sheetF2P := xlsxFile.Sheet[kSymbolF2PSheetName]
	symbolListF2P := []*Symbol{}
	if sheetF2P != nil {
		symbolCounts := []int{}
		for rowIndex, row := range sheetF2P.Rows {
			symbol := &Symbol{
				ReelDist:  make([]int, info.Width),
				PayoutMap: make(map[int]decimal.Decimal),
			}
			index := 0
			for colIndex, cell := range row.Cells {
				if rowIndex == 0 { // Header row
					if colIndex < 3+info.Width {
						continue
					}
					val := strings.TrimPrefix(cell.String(), "Payout ")
					value := xlsx_helper.CellValToInt(val)
					symbolCounts = append(symbolCounts, value)
				} else { // Data row
					if cell.String() == "" && colIndex == 0 {
						break // Stop processing row if the first cell is empty
					}
					if colIndex == 0 {
						symbol.SymbolCode = strings.TrimSpace(cell.String())
					} else if colIndex == 1 {
						symbol.Name = strings.TrimSpace(cell.String())
					} else if colIndex == 2 {
						symbol.Type = strings.TrimSpace(cell.String())
					} else if colIndex >= 3 && colIndex < 3+info.Width {
						// reels
						symbol.ReelDist[colIndex-3], parsingErr = cell.Int()
						if parsingErr != nil {
							symbol.ReelDist[colIndex-3] = 0 // Silently ignores error
						}
					} else if colIndex >= 3+info.Width {
						// payout
						payout := xlsx_helper.ReadDecimalFromCell(cell)
						symbol.PayoutMap[symbolCounts[index]] = payout
						index++
					}
				}
			}
			if symbol.SymbolCode != "" {
				symbol.Index = len(symbolListF2P)
				symbolListF2P = append(symbolListF2P, symbol)
			}
		}
		game.f2pSymbolList = symbolListF2P
	}

	// This function currently never returns a parsing error, which is a problem.
	return game, nil
}

// --- Mock Implementations and Main Runner ---

type mockXLSXHelper struct{}

func (m *mockXLSXHelper) CellValToInt(val string) int {
	if val == "3" {
		return 3
	}
	if val == "4" {
		return 4
	}
	if val == "5" {
		return 5
	}
	return 0
}
func (m *mockXLSXHelper) ReadDecimalFromCell(cell *xlsx.Cell) decimal.Decimal {
	d, _ := decimal.NewFromString(cell.String())
	return d
}

// main function to create the file and run the parser.
func main() {
	info := &Info{Width: 5}
	game, err := LoadGameConfig(kConfigFilename, info)
	if err != nil {
		log.Fatalf("Failed to load game config: %v", err)
	}

	fmt.Println("--- Premium Symbols ---")
	for _, s := range game.premiumSymbolList {
		fmt.Printf("%+v\n", s)
	}

	fmt.Println("\n--- F2P Symbols ---")
	for _, s := range game.f2pSymbolList {
		fmt.Printf("%+v\n", s)
	}
}

// createTestExcelFile generates an xlsx file on disk for testing.
func createTestExcelFile(filePath string) error {
	file := xlsx.NewFile()

	// Create Premium Sheet
	sheetPremium, _ := file.AddSheet(kSymbolPremiumSheetName)
	row1 := sheetPremium.AddRow()
	row1.AddCell().SetString("SymbolCode")
	row1.AddCell().SetString("Name")
	row1.AddCell().SetString("Type")
	for i := 1; i <= 5; i++ {
		row1.AddCell().SetString(fmt.Sprintf("Reel %d", i))
	}
	row1.AddCell().SetString("Payout 3")
	row1.AddCell().SetString("Payout 4")
	row1.AddCell().SetString("Payout 5")
	row2 := sheetPremium.AddRow()
	row2.AddCell().SetString("P1")
	row2.AddCell().SetString("Premium Symbol 1")
	row2.AddCell().SetString("HIGH")
	for i := 0; i < 5; i++ {
		row2.AddCell().SetInt(10 + i)
	}
	row2.AddCell().SetString("1.5")
	row2.AddCell().SetString("2.5")
	row2.AddCell().SetString("5.0")

	// Create F2P Sheet
	sheetF2P, _ := file.AddSheet(kSymbolF2PSheetName)
	row3 := sheetF2P.AddRow()
	row3.AddCell().SetString("SymbolCode")
	row3.AddCell().SetString("Name")
	row3.AddCell().SetString("Type")
	for i := 1; i <= 5; i++ {
		row3.AddCell().SetString(fmt.Sprintf("Reel %d", i))
	}
	row3.AddCell().SetString("Payout 3")
	row3.AddCell().SetString("Payout 4")
	row3.AddCell().SetString("Payout 5")
	row4 := sheetF2P.AddRow()
	row4.AddCell().SetString("F1")
	row4.AddCell().SetString("F2P Symbol 1")
	row4.AddCell().SetString("LOW")
	for i := 0; i < 5; i++ {
		row4.AddCell().SetInt(1 + i)
	}
	row4.AddCell().SetString("0.5")
	row4.AddCell().SetString("1.0")
	row4.AddCell().SetString("2.0")

	return file.Save(filePath)
}
