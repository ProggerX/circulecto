package main

import "github.com/ProggerX/circulecto/pkg/draw"

func main() {
	// debugFlag := flag.Bool("debug", false, "print debug data")
	// flag.Parse()
	// if len(flag.Args()) != 1 {
	// 	log.Fatal("Usage: circulecto <string>")
	// }
	// if *debugFlag {
	// 	log.SetLevel(log.DebugLevel)
	// }
	// str := flag.Args()[0]
	// err := json.Unmarshal([]byte(str), &data)
	// if err != nil {
	// 	log.Fatal("Error while parsing json", "err", err)
	// }
	var dr draw.Drawing
	dr.Init(500, 500)
	dr.DrawCorrectLine(100, 100, 300, 200)
	dr.Save()
}
