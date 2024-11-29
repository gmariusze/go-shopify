package goshopify

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestShopGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/shop.json",
		httpmock.NewBytesResponder(200, loadFixture("shop.json")))

	shop, err := client.Shop.Get(nil)
	if err != nil {
		t.Errorf("Shop.Get returned error: %v", err)
	}

	// Check that dates are parsed
	d := time.Date(2007, time.December, 31, 19, 00, 00, 0, time.UTC)
	if !d.Equal(*shop.CreatedAt) {
		t.Errorf("Shop.CreatedAt returned %+v, expected %+v", shop.CreatedAt, d)
	}

	// Test a few fields
	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"ID", 690933842, shop.ID},
		{"ShopOwner", "Steve Jobs", shop.ShopOwner},
		{"Address1", "1 Infinite Loop", shop.Address1},
		{"Name", "Apple Computers", shop.Name},
		{"Email", "steve@apple.com", shop.Email},
		{"HasStorefront", true, shop.HasStorefront},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("Shop.%v returned %v, expected %v", c.field, c.actual, c.expected)
		}
	}
}
