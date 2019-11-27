package services

import "circleweb/models"
/**
查找列表
 */
func (s *Service) FindItems(city,area,name,address string)[]models.Poi  {
	return s.poiDb.FindItems(name,address,city,area)
}
