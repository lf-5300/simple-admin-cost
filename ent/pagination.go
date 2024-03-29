// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/hf/simple-admin-cost-api/ent/project"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type ProjectPager struct {
	Order  project.OrderOption
	Filter func(*ProjectQuery) (*ProjectQuery, error)
}

// ProjectPaginateOption enables pagination customization.
type ProjectPaginateOption func(*ProjectPager)

// DefaultProjectOrder is the default ordering of Project.
var DefaultProjectOrder = Desc(project.FieldID)

func newProjectPager(opts []ProjectPaginateOption) (*ProjectPager, error) {
	pager := &ProjectPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultProjectOrder
	}
	return pager, nil
}

func (p *ProjectPager) ApplyFilter(query *ProjectQuery) (*ProjectQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// ProjectPageList is Project PageList result.
type ProjectPageList struct {
	List        []*Project   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (pr *ProjectQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...ProjectPaginateOption,
) (*ProjectPageList, error) {

	pager, err := newProjectPager(opts)
	if err != nil {
		return nil, err
	}

	if pr, err = pager.ApplyFilter(pr); err != nil {
		return nil, err
	}

	ret := &ProjectPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := pr.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		pr = pr.Order(pager.Order)
	} else {
		pr = pr.Order(DefaultProjectOrder)
	}

	pr = pr.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := pr.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
