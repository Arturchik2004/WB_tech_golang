package main

import "fmt"

type LegacyPrinter interface {
	PrintLegacy(text string) string
}

type OldPrinter struct{}

func (p *OldPrinter) PrintLegacy(text string) string {
	return text + " - хорошо забытое старое"
}

type ModernPrinter interface {
	PrintModern(text string) string
}

type PrinterAdapter struct {
	LegacyPrinter LegacyPrinter
}

func (adapter *PrinterAdapter) PrintModern(text string) string {
	return adapter.LegacyPrinter.PrintLegacy(text)
}

func main() {

	oldPrinter := &OldPrinter{}

	adapter := &PrinterAdapter{
		LegacyPrinter: oldPrinter,
	}

	var modernPrinter ModernPrinter = adapter
	result := modernPrinter.PrintModern("Все новое")
	fmt.Println(result)
}
