package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rpstvs/steamprice-api/internals/database"
	"github.com/rpstvs/steamprice-api/internals/utils"
)

func (cfg *Client) UpdateDB(index int) {

	resultados := cfg.GetSkins(index)

	start := resultados.Start
	end := resultados.TotalCount

	ctx := context.Background()
	for _, result := range resultados.Results {

		image := utils.BuildImageURL(result.AssetDescription.IconURL)

		_, err := cfg.DB.GetItemByName(ctx, result.HashName)
		if err != nil {
			cfg.WriteToDB(result.HashName, image, ctx)
		}
		fmt.Println(result.SalePriceText)
		cfg.PriceUpdate(result.HashName, result.SalePriceText, ctx)
		cfg.PriceChangeDaily(result.HashName)
		cfg.WeeklyPriceChange(result.HashName)

	}

	if start < end {
		start += 100
		fmt.Printf("Dormir 15s - Next Index %d /%d \n", start, end)
		time.Sleep(15 * time.Second)
		cfg.UpdateDB(start)
	}

}

func (cfg *Client) WriteToDB(itemName, url string, ctx context.Context) {

	_, err := cfg.DB.CreateItem(ctx, database.CreateItemParams{
		Itemname:   itemName,
		Imageurl:   url,
		Daychange:  0.00,
		Weekchange: 0.00,
	})

	if err != nil {
		log.Print(err)
	}

}

func (cfg *Client) PriceUpdate(itemName string, price string, ctx context.Context) {

	priceDb := utils.PriceConverter(price)
	if err != nil {
		fmt.Println(err)
	}

	date := utils.ConvertDate()

	cfg.DB.AddPrice(ctx, database.AddPriceParams{
		ItemID:    id,
		Pricedate: date,
		Price:     priceDb,
	})

}
