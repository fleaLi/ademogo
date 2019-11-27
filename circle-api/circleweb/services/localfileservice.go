package services

import (
    "circleweb/kit"
    "circleweb/models"
)

func (s *Service) GetOne(md5str string) (bool,*models.CircleCsv) {
  return s.db.GetOne(md5str)
}
func (s *Service) InsertCircleCsv(ccsv *models.CircleCsv,bs *map[string][]map[string]string) int64 {
    rows:=kit.Map2CsvRow(*bs)
 return  s.db.InsertCircleCsv(ccsv,&rows)
}