package service

import (
	"github.com/zhixian0949/gin-blog/internal/model"
	"github.com/zhixian0949/gin-blog/internal/model/request"
	"github.com/zhixian0949/gin-blog/pkg/app"
)

func (svc *Service) CountTag(param *request.CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *request.TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *request.CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *request.UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *request.DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
