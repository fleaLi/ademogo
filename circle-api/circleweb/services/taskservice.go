package services

import (
	"circleweb/dao/dto"
	"circleweb/models"
)

func (s *Service) Insert(task *models.TaskTemplate) (int64,error)  {
	return s.db.Insert(task)
}

func (s *Service) TaskItems(pageno ,pagesize,taskstatus int64,key string,uid string)(count int64,totalpage int64,result []models.TaskTemplate) {
    dd :=dto.TaskDto{
    	PageSize:int64(pagesize),
    	PageNo:int64(pageno),
    	Key:key,
    	TaskStatus:taskstatus,
    	Uid:uid,
	}
   return s.db.TaskItems(dd)
}
func (s *Service) ChangeTask( taskId string,src,target int8 )  {
	s.db.ChangeTaskStatus(taskId,src,target)
}
func (s *Service) UpdateTask(dto *models.TaskTemplate)(int64,error)  {
	return s.db.UpdateTask(dto)
}