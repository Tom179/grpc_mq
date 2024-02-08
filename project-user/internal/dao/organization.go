package dao

import (
	"context"
	"test.com/project-user/internal/data/organization"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/database/gorm"
)

type OrganizationDao struct {
	conn *gorm.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		conn: gorm.New(),
	}
}

func (o *OrganizationDao) FindOrganizationByMemId(ctx context.Context, memId int64) ([]organization.Organization, error) {
	var orgs []organization.Organization
	err := o.conn.Session(ctx).Where("member_id=?", memId).Find(&orgs).Error
	return orgs, err
}

func (o *OrganizationDao) SaveOrganization(conn database.DbConn, ctx context.Context, org *organization.Organization) error {
	o.conn = conn.(*gorm.GormConn)
	return o.conn.Tx(ctx).Create(org).Error
}
