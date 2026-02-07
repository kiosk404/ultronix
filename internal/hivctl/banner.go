package hivctl

import (
	"fmt"

	"github.com/kiosk404/ultronix/pkg/version"
)

const bannerText = `
 _   _ _ _                   _      
| | | | | |_ _ __ ___  _ __ (_)_  __
| | | | | __| '__/ _ \| '_ \| \ \/ /
| |_| | | |_| | | (_) | | | | |>  < 
 \___/|_|\__|_|  \___/|_| |_|_/_/\_\
                                      
  Distributed AI Agent System
`

// Banner returns the CLI banner string.
func Banner() string {
	return fmt.Sprintf("%s\n  Version: %s\n", bannerText, version.Get().String())
}
