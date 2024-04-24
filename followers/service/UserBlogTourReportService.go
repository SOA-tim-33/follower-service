package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type UserBlogTourReportService struct {
	UserBlogTourReportRepo *repo.UserBlogTourReportRepository
}

func (service *UserBlogTourReportService) CreateUserBlogTourReport(userBlogTourReport *model.UserBlogTourReport) error {
	err := service.UserBlogTourReportRepo.CreateUserBlogTourReport(userBlogTourReport)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserBlogTourReportService) FindUserBlogTourReport(id string) (*model.UserBlogTourReport, error) {
	userBlogTourReport, err := service.UserBlogTourReportRepo.FindUserBlogTourReportById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("user blog tour report with id %s not found", id))
	}
	return &userBlogTourReport, nil
}

func (service *UserBlogTourReportService) FindAllUserBlogTourReports() ([]model.UserBlogTourReport, error) {
	userBlogTourReports, err := service.UserBlogTourReportRepo.FindAllUserBlogTourReports()
	if err != nil {
		return nil, err
	}
	return userBlogTourReports, nil
}

func (service *UserBlogTourReportService) UpdateUserBlogTourReport(userBlogTourReport *model.UserBlogTourReport) error {
	err := service.UserBlogTourReportRepo.UpdateUserBlogTourReport(userBlogTourReport)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserBlogTourReportService) DeleteUserBlogTourReport(id string) error {
	userBlogTourReport, err := service.UserBlogTourReportRepo.FindUserBlogTourReportById(id)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("user blog tour report with id %s not found", id))
	}
	err = service.UserBlogTourReportRepo.DeleteUserBlogTourReport(&userBlogTourReport)
	if err != nil {
		return err
	}
	return nil
}
