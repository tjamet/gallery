package url

import "net/url"

var NoUpscale url.Values = url.Values{
	"fit": {"max"},
}

var Small url.Values = url.Values{
	"w": {"200"},
	"h": {"200"},
}

var Medium url.Values = url.Values{
	"w": {"1000"},
	"h": {"1000"},
}

var Large url.Values = url.Values{
	"w": {"3000"},
	"h": {"3000"},
}

var LowQuality url.Values = url.Values{
	"q": {"45"},
}

var MediumQuality url.Values = url.Values{
	"q": {"75"},
}

var HighQuality url.Values = url.Values{
	"q": {"95"},
}

var Jpeg url.Values = url.Values{
	"fm": {"jpg"},
}

var Webp url.Values = url.Values{
	"fm": {"webp"},
}
