package migration

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetCreateTableSQLFromFile(filePath string) ([]string, error) {
	return nil, nil
}
