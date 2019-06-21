package test

import (
	"context"
	"spider/conf"
	"spider/dao"
	"spider/model"
)

// Service service.
type Service struct {
	//ac  *paladin.Map
	dao *dao.Dao
}

// New new a service and return.
func New(c *conf.Config) (s *Service) {
	//var ac = new(paladin.TOML)
	//if err := paladin.Watch("spider.toml", ac); err != nil {
	//	panic(err)
	//}
	s = &Service{
		//ac:  ac,
		dao: dao.New(),
	}
	return s
}



// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

func (s *Service)Hello() model.Response{
	return model.Response{
		Message:"hello",
	}
}

func  (s *Service) test(){

}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
