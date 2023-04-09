package domain

type Group struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

type GroupModel interface {
	Create()
}

func (g *Group) GetTable() string {
	return "groups"
}

func (g *Group) GetFillableFields() []string {
	fields := []string{"title"}

	return fields
}

func (g *Group) GetValues() []interface{} {
	data := make([]interface{}, len(g.GetFillableFields()))

	data = append(data, g.Title)

	return data
}
