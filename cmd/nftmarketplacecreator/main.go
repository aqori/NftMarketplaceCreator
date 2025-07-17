// cmd/nftmarketplacecreator/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacecreator/internal/nftmarketplacecreator"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacecreator.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
