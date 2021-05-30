# go-shopee-v2

Shopee API v2 with golang

https://open.shopee.com/documents?module=87&type=2&id=58&version=2

[go-shopee](https://github.com/passwind/go-shopee) for `Shopee API v1`

Frankly, Shopee v2 is too different from v1. Have to set up this repo for v2.

## How to use

Initialize Client And request shop info

```
  app := goshopee.App{
		PartnerID:  133456,
		PartnerKey: "xxxxxxxxx",
		APIURL:     "https://xxxxx",
		RedirectURL: "https://yourdomain/usercallback",
	}

	client := goshopee.NewClient(app, goshopee.WithRetry(3), goshopee.WithLogger(NewLogger()))

  // auth url
  authUrl, err := client.Auth.GetAuthURL()

  // fetch access token
  // code from https://yourdomain/usercallback?code=xxxxx&shop_id=123456
  res, err: client.Auth.GetAccessToken(sid, 0, code)
  tok:=res.AccessToken

  // fetch shop info
  client.Shop.GetShopInfo(sid, tok)
```

## Thanks to

- [go-shopify](https://github.com/bold-commerce/go-shopify) Inspire me and provide a base structure
