package models

import (
	"strings"

	"github.com/grafana/grafana/pkg/services/sqlstore/searchstore"
)

type SortOption struct {
	Name        string
	DisplayName string
	Description string
	Index       int
	MetaName    string
	Filter      []SortOptionFilter
}

type SortOptionFilter interface {
	searchstore.FilterOrderBy
}

type FindPersistedDashboardsQuery struct {
	Title        string
	OrgId        int64
	SignedInUser *SignedInUser
	IsStarred    bool
	DashboardIds []int64
	Type         string
	FolderIds    []int64
	Tags         []string
	Limit        int64
	Page         int64
	Permission   PermissionType
	Sort         SortOption

	Filters []interface{}

	Result  HitList
	CountDB bool
}

type CountDashboardsResult []DashboardCounter

type DashboardCounter struct {
	FolderUID string `xorm:"folder_uid"`
	Count     int64
}

type HitType string

const (
	DashHitDB     HitType = "dash-db"
	DashHitHome   HitType = "dash-home"
	DashHitFolder HitType = "dash-folder"
)

type Hit struct {
	ID           int64    `json:"id"`
	UID          string   `json:"uid"`
	Title        string   `json:"title"`
	URI          string   `json:"uri"`
	URL          string   `json:"url"`
	Slug         string   `json:"slug"`
	Type         HitType  `json:"type"`
	Tags         []string `json:"tags"`
	IsStarred    bool     `json:"isStarred"`
	FolderID     int64    `json:"folderId,omitempty"`
	FolderUID    string   `json:"folderUid,omitempty"`
	FolderTitle  string   `json:"folderTitle,omitempty"`
	FolderURL    string   `json:"folderUrl,omitempty"`
	SortMeta     int64    `json:"sortMeta"`
	SortMetaName string   `json:"sortMetaName,omitempty"`
	DsCounter    int64    `json:"dscount,omitempty"`
}

type HitList []*Hit

func (s HitList) Len() int      { return len(s) }
func (s HitList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s HitList) Less(i, j int) bool {
	if s[i].Type == "dash-folder" && s[j].Type == "dash-db" {
		return true
	}

	if s[i].Type == "dash-db" && s[j].Type == "dash-folder" {
		return false
	}

	return strings.ToLower(s[i].Title) < strings.ToLower(s[j].Title)
}
