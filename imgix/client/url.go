package client

import "net/url"

func (c *client) URL(path string, values url.Values) string {
	if values == nil || len(values) == 0 {
		return c.client.Path(path)
	}
	return c.client.PathWithParams(path, values)
}
