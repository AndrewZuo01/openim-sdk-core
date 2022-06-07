package sdk_params_callback

import "open_im_sdk/pkg/db"

type UserInDepartment struct {
	DepartmentInfo *db.LocalDepartment       `json:"department"`
	MemberInfo     *db.LocalDepartmentMember `json:"member"`
}

type DepartmentAndUser struct {
	db.LocalDepartment
	db.LocalDepartmentMember
}

type GetSubDepartmentCallback []*db.LocalDepartment

type GetParentDepartmentListCallback []*db.LocalDepartment

type GetDepartmentMemberCallback []*db.LocalDepartmentMember

type GetUserInDepartmentCallback []*UserInDepartment

type ParentDepartmentCallback struct {
	Name         string `json:"name"`
	DepartmentID string `json:"departmentID"`
}

type GetDepartmentMemberAndSubDepartmentCallback struct {
	DepartmentList       []*db.LocalDepartment       `json:"departmentList"`
	DepartmentMemberList []*db.LocalDepartmentMember `json:"departmentMemberList"`
	ParentDepartmentList []ParentDepartmentCallback  `json:"parentDepartmentList"`
}

type GetDepartmentInfoCallback *db.LocalDepartment

type SearchOrganizationCallback struct {
	DepartmentList       []*db.LocalDepartment `json:"departmentList"`
	DepartmentMemberList []*struct {
		*db.SearchDepartmentMemberResult
		Path []*ParentDepartmentCallback `json:"path"`
	} `json:"departmentMemberList"`
}
