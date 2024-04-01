package neurocapsule

import (
	"context"
	"fmt"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/dto"
	"strings"
	"sync"
	"time"
)

const (
	brown    = "641 => коричневый"
	black    = "645 => черный"
	darkBlue = "643 => синий"
	white    = "615 => белый"
	blue     = "859 => голубой"
	gray     = "635 => серый"
	beige    = "647 => бежевый"
	orange   = "629 => оранжевый"
	yellow   = "613 => желтый"
	khaki    = "3847 => хаки"
	green    = "637 => зеленый"
	violet   = "689 => фиолетовый"
)

const min = 300
const max = 1000000

func (p *neurocapsuleProcessor) GetNeurocapsule(ctx context.Context, params *dto.GetNeurocapsule) (*dto.Neurocapsule, error) {
	if params.ColorScheme == "classic" && params.ColorBase != "blackWhite" {
		return nil, nil
	}

	if params.ColorScheme != "classic" && params.ColorBase == "blackWhite" {
		return nil, nil
	}

	time.Sleep(5 * time.Second)

	priceMap := map[string]map[string][]int{
		"low": {
			"trousers":   []int{min, 1180},
			"jeans":      []int{min, 1574},
			"sweatpants": []int{min, 1180},

			"shirt1":  []int{min, 1165},
			"shirt2":  []int{min, 1165},
			"tShirt1": []int{min, 599},
			"tShirt2": []int{min, 599},

			"hoodie":  []int{min, 1226},
			"sweater": []int{min, 1226},
			"blazer":  []int{min, 2532},
		},
		"middle": {
			"trousers":   []int{1181, 3141},
			"jeans":      []int{1575, 3388},
			"sweatpants": []int{1181, 3141},

			"shirt1":  []int{1166, 2503},
			"shirt2":  []int{1166, 2503},
			"tShirt1": []int{600, 1468},
			"tShirt2": []int{600, 1468},

			"hoodie":  []int{1227, 3374},
			"sweater": []int{1227, 3374},
			"blazer":  []int{2533, 6373},
		},
		"high": {
			"trousers":   []int{3142, max},
			"jeans":      []int{3389, max},
			"sweatpants": []int{3142, max},

			"shirt1":  []int{2504, max},
			"shirt2":  []int{2504, max},
			"tShirt1": []int{1469, max},
			"tShirt2": []int{1469, max},

			"hoodie":  []int{3375, max},
			"sweater": []int{3375, max},
			"blazer":  []int{6374, max},
		},
	}

	colorMap := map[string]map[string]map[string][]string{
		"contrast": {
			"brown": {
				"trousers":   []string{brown, black},
				"jeans":      []string{darkBlue},
				"sweatpants": []string{brown, black},

				"shirt1":  []string{white},
				"shirt2":  []string{blue},
				"tShirt1": []string{white},
				"tShirt2": []string{gray, beige, orange},

				"hoodie":  []string{black, brown, gray, orange},
				"sweater": []string{black, brown, gray, orange},
				"blazer":  []string{black, brown, gray, orange},
			},
			"khaki": {
				"trousers":   []string{khaki, black},
				"jeans":      []string{darkBlue},
				"sweatpants": []string{khaki, black},

				"shirt1":  []string{white},
				"shirt2":  []string{gray},
				"tShirt1": []string{white},
				"tShirt2": []string{blue, yellow, orange},

				"hoodie":  []string{black, khaki, gray, orange},
				"sweater": []string{black, khaki, gray, orange},
				"blazer":  []string{black, khaki, gray, orange},
			},
			"darkBlue": {
				"trousers":   []string{darkBlue, black},
				"jeans":      []string{darkBlue},
				"sweatpants": []string{darkBlue, black},

				"shirt1":  []string{white},
				"shirt2":  []string{blue},
				"tShirt1": []string{white},
				"tShirt2": []string{gray, darkBlue, orange},

				"hoodie":  []string{black, gray, darkBlue, orange},
				"sweater": []string{black, gray, darkBlue, orange},
				"blazer":  []string{black, gray, darkBlue, orange},
			},
		},
		"monochrome": {
			"brown": {
				"trousers":   []string{brown, black},
				"jeans":      []string{brown, black},
				"sweatpants": []string{brown, black},

				"shirt1":  []string{white},
				"shirt2":  []string{beige},
				"tShirt1": []string{beige},
				"tShirt2": []string{brown},

				"hoodie":  []string{brown, beige, black},
				"sweater": []string{brown, beige, black},
				"blazer":  []string{brown, beige, black},
			},
			"khaki": {
				"trousers":   []string{khaki, black},
				"jeans":      []string{khaki, black},
				"sweatpants": []string{khaki, black},

				"shirt1":  []string{white},
				"shirt2":  []string{green},
				"tShirt1": []string{green},
				"tShirt2": []string{khaki},

				"hoodie":  []string{khaki, green, black},
				"sweater": []string{khaki, green, black},
				"blazer":  []string{khaki, green, black},
			},
			"darkBlue": {
				"trousers":   []string{darkBlue, black},
				"jeans":      []string{darkBlue, black},
				"sweatpants": []string{darkBlue, black},

				"shirt1":  []string{white},
				"shirt2":  []string{blue},
				"tShirt1": []string{blue},
				"tShirt2": []string{darkBlue},

				"hoodie":  []string{blue, darkBlue, violet},
				"sweater": []string{blue, darkBlue, violet},
				"blazer":  []string{blue, darkBlue, violet},
			},
		},
		"classic": {
			"blackWhite": {
				"trousers":   []string{black, gray},
				"jeans":      []string{black, gray},
				"sweatpants": []string{black, gray},

				"shirt1":  []string{white},
				"shirt2":  []string{gray},
				"tShirt1": []string{white},
				"tShirt2": []string{gray},

				"hoodie":  []string{black, gray},
				"sweater": []string{black, gray},
				"blazer":  []string{black, gray},
			},
		},
	}

	excludedSkuList := ""
	if params.ExcludedSkuList != nil {
		excludedSkuList = *params.ExcludedSkuList
	}

	var wg sync.WaitGroup
	var errors []string
	var resp dto.Neurocapsule
	var err error

	db := p.storage.Conn
	db2 := p.storage2.Conn

	// bottom
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Bottom.Trousers, err = p.repository.Product.Get(ctx, db, db2, "Брюки", colorMap[params.ColorScheme][params.ColorBase]["trousers"], excludedSkuList, priceMap[params.PriceSegment]["trousers"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Bottom.Jeans, err = p.repository.Product.Get(ctx, db, db2, "Джинсы", colorMap[params.ColorScheme][params.ColorBase]["jeans"], excludedSkuList, priceMap[params.PriceSegment]["jeans"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Bottom.Sweatpants, err = p.repository.Product.Get(ctx, db, db2, "Брюки спортивные", colorMap[params.ColorScheme][params.ColorBase]["sweatpants"], excludedSkuList, priceMap[params.PriceSegment]["sweatpants"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	// top1
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Top1.Shirt1, err = p.repository.Product.Get(ctx, db, db2, "Рубашка", colorMap[params.ColorScheme][params.ColorBase]["shirt1"], excludedSkuList, priceMap[params.PriceSegment]["shirt1"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Top1.Shirt2, err = p.repository.Product.Get(ctx, db, db2, "Рубашка", colorMap[params.ColorScheme][params.ColorBase]["shirt2"], excludedSkuList, priceMap[params.PriceSegment]["shirt2"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Top1.TShirt1, err = p.repository.Product.Get(ctx, db, db2, "Футболка", colorMap[params.ColorScheme][params.ColorBase]["tShirt1"], excludedSkuList, priceMap[params.PriceSegment]["tShirt1"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Top1.TShirt2, err = p.repository.Product.Get(ctx, db, db2, "Футболка", colorMap[params.ColorScheme][params.ColorBase]["tShirt2"], excludedSkuList, priceMap[params.PriceSegment]["tShirt2"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	// top2
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Top2.Hoodie, err = p.repository.Product.Get(ctx, db, db2, "Худи", colorMap[params.ColorScheme][params.ColorBase]["hoodie"], excludedSkuList, priceMap[params.PriceSegment]["hoodie"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Top2.Sweater, err = p.repository.Product.Get(ctx, db, db2, "Свитер", colorMap[params.ColorScheme][params.ColorBase]["sweater"], excludedSkuList, priceMap[params.PriceSegment]["sweater"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp.Top2.Blazer, err = p.repository.Product.Get(ctx, db, db2, "Пиджак", colorMap[params.ColorScheme][params.ColorBase]["blazer"], excludedSkuList, priceMap[params.PriceSegment]["blazer"])
		if err != nil {
			errors = append(errors, err.Error())
		}
	}()

	wg.Wait()

	if len(errors) > 0 {
		return nil, fmt.Errorf(strings.Join(errors, "; "))
	}

	return &resp, nil
}
