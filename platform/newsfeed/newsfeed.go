package newsfeed

type GetAll interface {
	GetAll()[]Item
}

type Adder interface {
	Add(item *Item)
}

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(item *Item) {
	r.Items = append(r.Items, *item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
