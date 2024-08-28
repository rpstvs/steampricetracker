package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseResults(results SearchResult) {

	for _, result := range results.Results {
		//fmt.Println(result.HashName)
		name := strings.Split(result.HashName, "|")
		price := PriceConverter(result.SalePriceText)

		if len(name) == 1 {
			caixa := parseCase(name, price)
			fmt.Println(caixa)
			continue
		}

		if name[0] == "Sticker " {
			//fmt.Printf("Estou a entrar no if dos das stickers %s", name)
			sticker := ParseSticker(name, price)
			fmt.Println(sticker)
		} else if strings.Contains(result.AssetDescription.Type, "Agente") {
			//fmt.Printf("Estou a entrar no if dos  agentes %s", name)
			agent := parseAgent(name, price)
			fmt.Println(agent)

		} else {
			//fmt.Printf("Estou a entrar no if dos das armas %s", name)
			skin := parseSkin(name, price)
			fmt.Println(skin)
		}

	}
}

func ParseSticker(sticker []string, price float64) Sticker {
	if len(sticker) == 0 {
		log.Fatal("no input")
		return Sticker{}
	}

	condition := "none"
	tournamentName := "none"
	var pcondition []string

	tmp := strings.Split(sticker[1], "(")

	if len(tmp) == 2 {
		pcondition = parseCondition(sticker[1])
	}

	name := strings.TrimSpace(tmp[0])

	if len(sticker) > 2 {
		tournamentName = sticker[2]
	}

	if pcondition != nil {
		condition = pcondition[1]
	}

	//fmt.Println(name, tournamentName, condition)

	return Sticker{
		Name:       name,
		Tournament: tournamentName,
		Condition:  condition,
		Price:      price,
	}
}

func parseSkin(skin []string, price float64) Skin {

	if len(skin) == 0 {
		log.Fatal("no input")
		return Skin{}
	}

	condition := "none"

	gunName := skin[0]

	tmp := strings.Split(skin[1], "(")
	skinName := strings.TrimSpace(tmp[0])

	pCondition := parseCondition(skin[1])

	if pCondition != nil {
		condition = pCondition[1]
	}

	return Skin{
		GunName:   gunName,
		SkinName:  skinName,
		Condition: condition,
		Price:     price,
	}
}

func parseCase(name []string, price float64) Case {
	return Case{
		CaseName: name[0],
		Price:    price,
	}
}

func parseAgent(agent []string, price float64) Agent {

	return Agent{
		Name:  agent[0],
		Group: agent[1],
		Price: price,
	}
}

func parseCondition(condition string) []string {
	re := regexp.MustCompile(`\(([^)]+)\)`)
	matches := re.FindStringSubmatch(condition)

	if len(matches) == 0 {
		return nil
	}

	return matches
}

func PriceConverter(priceStr string) float64 {
	if len(priceStr) > 7 {
		priceStr = strings.ReplaceAll(priceStr, "$", "")
		priceStr = strings.ReplaceAll(priceStr, ",", "")
		priceStr = strings.ReplaceAll(priceStr, "-", "0")

		price, err := strconv.ParseFloat(priceStr, 64)

		if err != nil {
			log.Println("error parsing price")
		}

		return price
	}
	priceStr = strings.ReplaceAll(priceStr, "$", "")
	priceStr = strings.ReplaceAll(priceStr, ",", ".")
	priceStr = strings.ReplaceAll(priceStr, "-", "0")

	price, err := strconv.ParseFloat(priceStr, 64)

	if err != nil {
		log.Println("error parsing price")
	}

	return price
}

func WriteToFile(results SearchResult) {
	file, err := os.OpenFile("tmp2.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("couldnt open file")
	}
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	writer := csv.NewWriter(file)

	for _, result := range results.Results {
		row := []string{result.HashName, result.AssetDescription.Classid}
		writer.Write(row)
	}

}
