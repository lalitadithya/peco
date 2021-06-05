package peco

import "github.com/lalitadithya/termbox-go"

func (t *Termbox) PostInit(cfg *Config) error {
	// Windows handle Esc/Alt self
	termbox.SetInputMode(termbox.InputEsc | termbox.InputAlt)

	return nil
}
