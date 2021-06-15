package client

import (
	"context"
	"fmt"
)

type OrgChangeRolesInput struct {
	AuthCookie string
	Emails     []string `json:"emails"`
	OrgKey     string
	RoleName   string `json:"role_name"`
}

type OrgChangeRolesOutput struct {
	Message string `json:"message"`
}

func OrgChangeRoles(c context.Context, input OrgChangeRolesInput) (out []OrgChangeRolesOutput, err error) {
	resty := RestyFromContext(c)
	url := fmt.Sprintf("/v1/organizations/%v/change_roles", input.OrgKey)
	resp, err := resty.R().
		SetError(&ErrorResponse{}).
		SetBody(&input).
		SetResult(&out).
		SetHeader("Cookie", input.AuthCookie).
		Post(url)
	if err != nil {
		return out, err
	}
	if err := CoerceError(resp); err != nil {
		return out, err
	}
	return out, nil
}

type OrgCreateInput struct {
	AuthCookie string
	OrgName    string `json:"name"`
}

type OrgCreateOutput struct {
	Message string `json:"message"`
}

func OrgCreate(c context.Context, input OrgCreateInput) (out OrgCreateOutput, err error) {
	resty := RestyFromContext(c)
	resp, err := resty.R().
		SetError(&ErrorResponse{}).
		SetBody(&input).
		SetResult(&out).
		SetHeader("Cookie", input.AuthCookie).
		Post("/v1/organizations/create")
	if err != nil {
		return out, err
	}
	if err := CoerceError(resp); err != nil {
		return out, err
	}
	return out, nil
}

type OrgInviteInput struct {
	AuthCookie string
	Email      string `json:"email"`
	OrgKey     string
}

type OrgInviteOutput struct {
	Message string `json:"message"`
}

func OrgInvite(c context.Context, input OrgInviteInput) (out OrgInviteOutput, err error) {
	resty := RestyFromContext(c)
	url := fmt.Sprintf("/v1/organizations/%v/invite", input.OrgKey)
	resp, err := resty.R().
		SetError(&ErrorResponse{}).
		SetBody(&input).
		SetResult(&out).
		SetHeader("Cookie", input.AuthCookie).
		Post(url)
	if err != nil {
		return out, err
	}
	if err := CoerceError(resp); err != nil {
		return out, err
	}
	return out, nil
}

type OrgJoinInput struct {
	AuthCookie     string
	InvitationCode string `json:"invitation_code"`
}

type OrgJoinOutput struct {
	Message string `json:"message"`
}

func OrgJoin(c context.Context, input OrgJoinInput) (out OrgJoinOutput, err error) {
	resty := RestyFromContext(c)
	resp, err := resty.R().
		SetError(&ErrorResponse{}).
		SetBody(&input).
		SetResult(&out).
		SetHeader("Cookie", input.AuthCookie).
		Post("/v1/organizations/join")
	if err != nil {
		return out, err
	}
	if err := CoerceError(resp); err != nil {
		return out, err
	}
	return out, nil
}

type OrgListInput struct {
	AuthCookie string
}

type OrgListOutput struct {
	Data []OrganizationEntity `json:"items"`
}

func OrgList(c context.Context, input OrgListInput) (out OrgListOutput, err error) {
	resty := RestyFromContext(c)
	resp, err := resty.R().
		SetError(&ErrorResponse{}).
		SetResult(&out).
		SetHeader("Cookie", input.AuthCookie).
		Get("/v1/organizations/")
	if err != nil {
		return out, err
	}
	if err := CoerceError(resp); err != nil {
		return out, err
	}
	return out, nil
}

type OrgMembersInput struct {
	AuthCookie string
	OrgKey     string
}

type OrgMembersOutput struct {
	Data []OrganizationMembershipEntity `json:"items"`
}

func OrgMembers(c context.Context, input OrgMembersInput) (out OrgMembersOutput, err error) {
	resty := RestyFromContext(c)
	url := fmt.Sprintf("/v1/organizations/%v/members", input.OrgKey)
	resp, err := resty.R().
		SetError(&ErrorResponse{}).
		SetResult(&out).
		SetHeader("Cookie", input.AuthCookie).
		Get(url)
	if err != nil {
		return out, err
	}
	if err := CoerceError(resp); err != nil {
		return out, err
	}
	return out, nil
}

type OrgRemoveInput struct {
	AuthCookie string
	Email      string `json:"email"`
	OrgKey     string
}

type OrgRemoveOutput struct {
	Message string `json:"message"`
}

func OrgRemove(c context.Context, input OrgRemoveInput) (out OrgRemoveOutput, err error) {
	resty := RestyFromContext(c)
	url := fmt.Sprintf("/v1/organizations/%v/remove", input.OrgKey)
	resp, err := resty.R().
		SetError(&ErrorResponse{}).
		SetBody(&input).
		SetResult(&out).
		SetHeader("Cookie", input.AuthCookie).
		Post(url)
	if err != nil {
		return out, err
	}
	if err := CoerceError(resp); err != nil {
		return out, err
	}
	return out, nil
}
