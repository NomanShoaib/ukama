// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users.proto

package gen

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddInternalRequest) Validate() error {
	if this.Org == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must not be an empty string`, this.Org))
	}
	if !(len(this.Org) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Org))
	}
	if nil == this.User {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf("message must exist"))
	}
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *AddInternalResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *AddRequest) Validate() error {
	if this.Org == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must not be an empty string`, this.Org))
	}
	if !(len(this.Org) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Org))
	}
	if nil == this.User {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf("message must exist"))
	}
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}

var _regex_AddResponse_Iccid = regexp.MustCompile(`^[0-9]{18,19}$`)

func (this *AddResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if !_regex_AddResponse_Iccid.MatchString(this.Iccid) {
		return github_com_mwitkow_go_proto_validators.FieldError("Iccid", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9]{18,19}$"`, this.Iccid))
	}
	return nil
}

var _regex_DeleteRequest_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *DeleteRequest) Validate() error {
	if !_regex_DeleteRequest_UserId.MatchString(this.UserId) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
func (this *DeleteResponse) Validate() error {
	return nil
}
func (this *ListRequest) Validate() error {
	if this.Org == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must not be an empty string`, this.Org))
	}
	if !(len(this.Org) > 2) {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must have a length greater than '2'`, this.Org))
	}
	return nil
}
func (this *ListResponse) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}

var _regex_GenerateSimTokenRequest_Iccid = regexp.MustCompile(`^[0-9]{18,19}$`)

func (this *GenerateSimTokenRequest) Validate() error {
	if !_regex_GenerateSimTokenRequest_Iccid.MatchString(this.Iccid) {
		return github_com_mwitkow_go_proto_validators.FieldError("Iccid", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9]{18,19}$"`, this.Iccid))
	}
	return nil
}
func (this *GenerateSimTokenResponse) Validate() error {
	return nil
}

var _regex_GetRequest_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetRequest) Validate() error {
	if !_regex_GetRequest_UserId.MatchString(this.UserId) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
func (this *GetResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if this.Sim != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sim); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sim", err)
		}
	}
	return nil
}

var _regex_UpdateRequest_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *UpdateRequest) Validate() error {
	if !_regex_UpdateRequest_UserId.MatchString(this.UserId) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}

var _regex_DeactivateUserRequest_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *DeactivateUserRequest) Validate() error {
	if !_regex_DeactivateUserRequest_UserId.MatchString(this.UserId) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
func (this *DeactivateUserResponse) Validate() error {
	return nil
}

var _regex_UserAttributes_Email = regexp.MustCompile(`^$|^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
var _regex_UserAttributes_Phone = regexp.MustCompile(`^$|^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

func (this *UserAttributes) Validate() error {
	if !_regex_UserAttributes_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`must be an email format`))
	}
	if !_regex_UserAttributes_Phone.MatchString(this.Phone) {
		return github_com_mwitkow_go_proto_validators.FieldError("Phone", fmt.Errorf(`must be a phone number format`))
	}
	return nil
}

var _regex_User_Email = regexp.MustCompile(`^$|^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
var _regex_User_Phone = regexp.MustCompile(`^$|^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
var _regex_User_Uuid = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *User) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !(len(this.Name) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Name))
	}
	if !_regex_User_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`must be an email format`))
	}
	if !_regex_User_Phone.MatchString(this.Phone) {
		return github_com_mwitkow_go_proto_validators.FieldError("Phone", fmt.Errorf(`must be a phone number format`))
	}
	if !_regex_User_Uuid.MatchString(this.Uuid) {
		return github_com_mwitkow_go_proto_validators.FieldError("Uuid", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.Uuid))
	}
	return nil
}
func (this *Sim) Validate() error {
	if this.Ukama != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ukama); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ukama", err)
		}
	}
	if this.Carrier != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Carrier); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Carrier", err)
		}
	}
	return nil
}
func (this *SimStatus) Validate() error {
	if this.Services != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Services); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Services", err)
		}
	}
	if this.Usage != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Usage); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Usage", err)
		}
	}
	return nil
}
func (this *Services) Validate() error {
	return nil
}
func (this *Usage) Validate() error {
	return nil
}

var _regex_SetSimStatusRequest_Iccid = regexp.MustCompile(`^[0-9]{18,19}$`)

func (this *SetSimStatusRequest) Validate() error {
	if !_regex_SetSimStatusRequest_Iccid.MatchString(this.Iccid) {
		return github_com_mwitkow_go_proto_validators.FieldError("Iccid", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9]{18,19}$"`, this.Iccid))
	}
	if this.Ukama != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ukama); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ukama", err)
		}
	}
	if this.Carrier != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Carrier); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Carrier", err)
		}
	}
	return nil
}
func (this *SetSimStatusRequest_SetServices) Validate() error {
	if this.Voice != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Voice); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Voice", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Sms != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sms); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sms", err)
		}
	}
	return nil
}
func (this *SetSimStatusResponse) Validate() error {
	if this.Sim != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sim); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sim", err)
		}
	}
	return nil
}