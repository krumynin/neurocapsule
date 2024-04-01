package repositories

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	lpgx "stash.lamoda.ru/gotools/pgx"
	"strings"

	"stash.lamoda.ru/neurocapsule/neurocapsule/app/models"
)

type Product interface {
	Get(ctx context.Context, query lpgx.PoolMeasured, query2 lpgx.PoolMeasured, category string, colors []string, excludedSKUList string, price []int) (*models.Product, error)
	Save(ctx context.Context, query lpgx.PoolMeasured, product *models.Product)
}

type product struct {
}

func NewProduct() Product {
	return &product{}
}

func (p *product) Save(ctx context.Context, query lpgx.PoolMeasured, product *models.Product) {
	sql1 := `INSERT INTO product (
					sku, name, price_amount, gallery, colors
				) VALUES ($1, $2, $3, $4, $5) 
				ON CONFLICT (sku) DO NOTHING;`

	_, err := query.Exec(ctx, "insert", sql1, product.SKU, product.Name, product.PriceAmount, product.Gallery, product.Color)
	if err != nil {
		fmt.Println(err)
	}

	sql2 := `INSERT INTO product_size (
					full_sku, product_sku
				) VALUES ($1, $2) 
				ON CONFLICT (full_sku) DO NOTHING;`

	_, err = query.Exec(ctx, "insert", sql2, product.FullSKU, product.SKU)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *product) Get(ctx context.Context, query lpgx.PoolMeasured, query2 lpgx.PoolMeasured, category string, colors []string, excludedSKUList string, price []int) (*models.Product, error) {
	var product models.Product

	for i, color := range colors {
		colors[i] = getColor(color)
	}

	sql := `
		SELECT
			p.sku,
			ps.full_sku,
			p.name,
			p.price_amount,
			p.gallery,
			p.colors
		FROM product p
		JOIN product_size ps ON ps.product_sku = p.sku
		WHERE p.name = $1
		  	AND p.colors = ANY($2)
			AND p.price_amount BETWEEN $3 AND $4
		ORDER BY random() LIMIT 1
	`

	err := query2.Get(ctx, &product, "query2_name", sql, category, colors, price[0], price[1])
	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil
}

func (p *product) Get2(ctx context.Context, query lpgx.PoolMeasured, query2 lpgx.PoolMeasured, category string, colors []string, excludedSKUList string, price []int) (*models.Product, error) {
	var product models.Product

	sql := `
		SELECT
			p.sku,
			ps.full_sku,
			p.name,
			p.price_amount,
			p.thumbnail,
			p.gallery,
			p.colors
		FROM product p
		JOIN product_size ps ON ps.product_sku = p.sku
		WHERE p.name = $1
		  	AND p.colors = ANY($2)
		  	AND p.is_sellable = true
		  	AND ps.qty != 0 AND ps.qty IS NOT NULL
			AND p.gender = 'men'
			AND p.price_amount BETWEEN $3 AND $4
	`

	if excludedSKUList != "" {
		SKUs := strings.Split(excludedSKUList, ",")
		str := ""
		for i, sku := range SKUs {
			if i == 0 {
				str = "'" + sku + "'"
				continue
			}
			str += ",'" + sku + "'"
		}

		sql += ` AND p.sku NOT IN(` + str + `)`
	}

	sql += ` ORDER BY random() LIMIT 1`
	err := query.Get(ctx, &product, "query_name", sql, category, colors, price[0], price[1])
	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	product.Color = getColor(product.Color)

	p.Save(ctx, query2, &product)

	return &product, nil
}

func getColor(color string) string {
	colors := strings.Split(color, "=>")
	if len(colors) == 0 {
		return ""
	}

	return strings.Trim(strings.Trim(colors[1], "\""), " ")
}
