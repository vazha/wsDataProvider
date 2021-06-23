package kraken

type Kraken struct {
	Url string
	Channels []string
}

func New(url string) *Kraken {
	return &Kraken{
		Url: url,
	}
}

func (*Kraken) Start() {

}

func (*Kraken) Stop() {

}

func (*Kraken) Subscribe([]string) {

}