package model

import "flag"

type Config struct {
	url             string
	qtdeConcurrency int
	qtdeRequests    int
}

func (c *Config) Validar() {
	c.validarUrl()
	c.validarQtdConcurrency()
	c.validarQtdeRequests()
	c.validarQtdeRequestMenorQtdeConcurrency()
}

func (c *Config) validarQtdeRequestMenorQtdeConcurrency() {
	if c.qtdeRequests < c.qtdeConcurrency {
		panic("Informe uma quantidade de requests maior que de concurrency. ")
	}
}

func (c *Config) validarUrl() {
	if c.url == "" {
		panic("URL para o teste de stress não informado.")
	}
}

func (c *Config) validarQtdConcurrency() {
	if c.qtdeConcurrency == 0 {
		panic("Quantidade de requisições simultaneas não informado.")
	}
}

func (c *Config) validarQtdeRequests() {
	if c.qtdeRequests == 0 {
		panic("Quantidade de requisições não informado.")
	}
}

func (c *Config) GetURL() string {
	return c.url
}

func (c *Config) GetQtdeRequests() int {
	return c.qtdeRequests
}

func (c *Config) GetQtdeConcurrency() int {
	return c.qtdeConcurrency
}

func (c *Config) GetQtdeResquesLoop() int {
	return c.qtdeRequests / c.qtdeConcurrency
}

func CarregarParams() Config {
	xUrl := flag.String("url", "", "URL para realizar o teste de stress.")
	xRequests := flag.Int("requests", 0, "Quantidade total de requisições.")
	xConcurrency := flag.Int("concurrency", 1, "Quantidade de requisições simultaneas.")
	flag.Parse()

	config := Config{
		url:             *xUrl,
		qtdeConcurrency: *xConcurrency,
		qtdeRequests:    *xRequests,
	}

	config.Validar()

	return config
}

func CarregarParamsParaTest() Config {
	config := Config{
		url:             "http://google.com",
		qtdeConcurrency: 10,
		qtdeRequests:    50,
	}
	return config
}
