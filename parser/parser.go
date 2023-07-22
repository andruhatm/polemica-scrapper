package parser

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"polemica_scrapper/models"
	"strings"
)

func ParseGameData(body []byte) ([]*models.RatingPlayer, error) {
	ratings := make([]*models.RatingPlayer, 10)
	reader := strings.NewReader(string(body))
	tokenizer := html.NewTokenizer(reader)
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return nil, io.EOF
			}
			return nil, tokenizer.Err()
		}
		tag, hasAttr := tokenizer.TagName()

		if string(tag) == "gamestats" && hasAttr {
			for {
				attrKey, attrValue, moreAttr := tokenizer.TagAttr()
				fmt.Printf("Tag: %v\n", string(tag))
				if string(attrKey) == ":game-data" {
					err := ParseRatings(ratings, attrValue)
					if err != nil {
						return nil, err
					}
					break
				}
				if !moreAttr {
					break
				}
			}
		}
	}
}

func ParseRatings(ratings []*models.RatingPlayer, data []byte) error {
	fmt.Println(string(data))

	var gameStats models.GameResults
	//parse resp
	err := json.Unmarshal(data, &gameStats)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", gameStats)

	//insert game stats to pg "game_history"

	//insert ratings to pg "players"

	return nil
}
