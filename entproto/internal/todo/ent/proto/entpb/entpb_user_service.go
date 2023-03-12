// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	ent "entgo.io/contrib/entproto/internal/todo/ent"
	attachment "entgo.io/contrib/entproto/internal/todo/ent/attachment"
	group "entgo.io/contrib/entproto/internal/todo/ent/group"
	pet "entgo.io/contrib/entproto/internal/todo/ent/pet"
	schema "entgo.io/contrib/entproto/internal/todo/ent/schema"
	user "entgo.io/contrib/entproto/internal/todo/ent/user"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	errors "errors"
	fmt "fmt"
	uuid "github.com/google/uuid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	strconv "strconv"
	strings "strings"
)

// UserService implements UserServiceServer
type UserService struct {
	client *ent.Client
	UnimplementedUserServiceServer
}

// NewUserService returns a new UserService
func NewUserService(client *ent.Client) *UserService {
	return &UserService{
		client: client,
	}
}

func toProtoUser_DeviceType(e user.DeviceType) User_DeviceType {
	if v, ok := User_DeviceType_value[strings.ToUpper("DEVICE_TYPE_"+string(e))]; ok {
		return User_DeviceType(v)
	}
	return User_DeviceType(0)
}

func toEntUser_DeviceType(e User_DeviceType) user.DeviceType {
	if v, ok := User_DeviceType_name[int32(e)]; ok {
		entVal := map[string]string{
			"DEVICE_TYPE_GLOWY9000": "GLOWY9000",
			"DEVICE_TYPE_SPEEDY300": "SPEEDY300",
		}[v]
		return user.DeviceType(entVal)
	}
	return ""
}

func toProtoUser_OmitPrefix(e user.OmitPrefix) User_OmitPrefix {
	if v, ok := User_OmitPrefix_value[strings.ToUpper(string(e))]; ok {
		return User_OmitPrefix(v)
	}
	return User_OmitPrefix(0)
}

func toEntUser_OmitPrefix(e User_OmitPrefix) user.OmitPrefix {
	if v, ok := User_OmitPrefix_name[int32(e)]; ok {
		entVal := map[string]string{
			"FOO": "foo",
			"BAR": "bar",
		}[v]
		return user.OmitPrefix(entVal)
	}
	return ""
}

func toProtoUser_Status(e user.Status) User_Status {
	if v, ok := User_Status_value[strings.ToUpper("STATUS_"+string(e))]; ok {
		return User_Status(v)
	}
	return User_Status(0)
}

func toEntUser_Status(e User_Status) user.Status {
	if v, ok := User_Status_name[int32(e)]; ok {
		entVal := map[string]string{
			"STATUS_PENDING": "pending",
			"STATUS_ACTIVE":  "active",
		}[v]
		return user.Status(entVal)
	}
	return ""
}

// toProtoUser transforms the ent type to the pb type
func toProtoUser(e *ent.User) (*User, error) {
	v := &User{}
	account_balance := e.AccountBalance
	v.AccountBalance = account_balance
	b_user_1 := wrapperspb.Int64(int64(e.BUser1))
	v.BUser_1 = b_user_1
	banned := e.Banned
	v.Banned = banned
	big_intValue, err := e.BigInt.Value()
	if err != nil {
		return nil, err
	}
	big_intTyped, ok := big_intValue.(string)
	if !ok {
		return nil, errors.New("casting value to string")
	}
	big_int := wrapperspb.String(big_intTyped)
	v.BigInt = big_int
	crm_id, err := e.CrmID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.CrmId = crm_id
	custom_pb := uint64(e.CustomPb)
	v.CustomPb = custom_pb
	device_type := toProtoUser_DeviceType(e.DeviceType)
	v.DeviceType = device_type
	exp := e.Exp
	v.Exp = exp
	external_id := int64(e.ExternalID)
	v.ExternalId = external_id
	height_in_cm := e.HeightInCm
	v.HeightInCm = height_in_cm
	id := e.ID
	v.Id = id
	joined := timestamppb.New(e.Joined)
	v.Joined = joined
	labels := e.Labels
	v.Labels = labels
	omit_prefix := toProtoUser_OmitPrefix(e.OmitPrefix)
	v.OmitPrefix = omit_prefix
	opt_bool := wrapperspb.Bool(e.OptBool)
	v.OptBool = opt_bool
	opt_num := wrapperspb.Int64(int64(e.OptNum))
	v.OptNum = opt_num
	opt_str := wrapperspb.String(e.OptStr)
	v.OptStr = opt_str
	points := uint32(e.Points)
	v.Points = points
	status := toProtoUser_Status(e.Status)
	v.Status = status
	_type := wrapperspb.String(e.Type)
	v.Type = _type
	user_name := e.UserName
	v.UserName = user_name
	if edg := e.Edges.Attachment; edg != nil {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Attachment = &Attachment{
			Id: id,
		}
	}
	if edg := e.Edges.Group; edg != nil {
		id := int64(edg.ID)
		v.Group = &Group{
			Id: id,
		}
	}
	if edg := e.Edges.Pet; edg != nil {
		id := int64(edg.ID)
		v.Pet = &Pet{
			Id: id,
		}
	}
	for _, edg := range e.Edges.Received1 {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Received_1 = append(v.Received_1, &Attachment{
			Id: id,
		})
	}
	return v, nil
}

// toProtoUserList transforms a list of ent type to a list of pb type
func toProtoUserList(e []*ent.User) ([]*User, error) {
	var pbList []*User
	for _, entEntity := range e {
		pbEntity, err := toProtoUser(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements UserServiceServer.Create
func (svc *UserService) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
	user := req.GetUser()
	m, err := svc.createBuilder(user)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoUser(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements UserServiceServer.Get
func (svc *UserService) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	var (
		err error
		get *ent.User
	)
	id := uint32(req.GetId())
	switch req.GetView() {
	case GetUserRequest_VIEW_UNSPECIFIED, GetUserRequest_BASIC:
		get, err = svc.client.User.Get(ctx, id)
	case GetUserRequest_WITH_EDGE_IDS:
		get, err = svc.client.User.Query().
			Where(user.ID(id)).
			WithAttachment(func(query *ent.AttachmentQuery) {
				query.Select(attachment.FieldID)
			}).
			WithGroup(func(query *ent.GroupQuery) {
				query.Select(group.FieldID)
			}).
			WithPet(func(query *ent.PetQuery) {
				query.Select(pet.FieldID)
			}).
			WithReceived1(func(query *ent.AttachmentQuery) {
				query.Select(attachment.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoUser(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements UserServiceServer.Update
func (svc *UserService) Update(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	user := req.GetUser()
	userID := uint32(user.GetId())
	m := svc.client.User.UpdateOneID(userID)
	userAccountBalance := float64(user.GetAccountBalance())
	m.SetAccountBalance(userAccountBalance)
	if user.GetBUser_1() != nil {
		userBUser1 := int(user.GetBUser_1().GetValue())
		m.SetBUser1(userBUser1)
	}
	userBanned := user.GetBanned()
	m.SetBanned(userBanned)
	if user.GetBigInt() != nil {
		userBigInt := schema.BigInt{}
		if err := (&userBigInt).Scan(user.GetBigInt().GetValue()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetBigInt(userBigInt)
	}
	var userCrmID uuid.UUID
	if err := (&userCrmID).UnmarshalBinary(user.GetCrmId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetCrmID(userCrmID)
	userCustomPb := uint8(user.GetCustomPb())
	m.SetCustomPb(userCustomPb)
	userDeviceType := toEntUser_DeviceType(user.GetDeviceType())
	m.SetDeviceType(userDeviceType)
	userExp := uint64(user.GetExp())
	m.SetExp(userExp)
	userExternalID := int(user.GetExternalId())
	m.SetExternalID(userExternalID)
	userHeightInCm := float32(user.GetHeightInCm())
	m.SetHeightInCm(userHeightInCm)
	if user.GetLabels() != nil {
		userLabels := user.GetLabels()
		m.SetLabels(userLabels)
	}
	userOmitPrefix := toEntUser_OmitPrefix(user.GetOmitPrefix())
	m.SetOmitPrefix(userOmitPrefix)
	if user.GetOptBool() != nil {
		userOptBool := user.GetOptBool().GetValue()
		m.SetOptBool(userOptBool)
	}
	if user.GetOptNum() != nil {
		userOptNum := int(user.GetOptNum().GetValue())
		m.SetOptNum(userOptNum)
	}
	if user.GetOptStr() != nil {
		userOptStr := user.GetOptStr().GetValue()
		m.SetOptStr(userOptStr)
	}
	userPoints := uint(user.GetPoints())
	m.SetPoints(userPoints)
	userStatus := toEntUser_Status(user.GetStatus())
	m.SetStatus(userStatus)
	if user.GetType() != nil {
		userType := user.GetType().GetValue()
		m.SetType(userType)
	}
	userUserName := user.GetUserName()
	m.SetUserName(userUserName)
	if user.GetAttachment() != nil {
		var userAttachment uuid.UUID
		if err := (&userAttachment).UnmarshalBinary(user.GetAttachment().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetAttachmentID(userAttachment)
	}
	if user.GetGroup() != nil {
		userGroup := int(user.GetGroup().GetId())
		m.SetGroupID(userGroup)
	}
	if user.GetPet() != nil {
		userPet := int(user.GetPet().GetId())
		m.SetPetID(userPet)
	}
	for _, item := range user.GetReceived_1() {
		var received1 uuid.UUID
		if err := (&received1).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddReceived1IDs(received1)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoUser(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements UserServiceServer.Delete
func (svc *UserService) Delete(ctx context.Context, req *DeleteUserRequest) (*emptypb.Empty, error) {
	var err error
	id := uint32(req.GetId())
	err = svc.client.User.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements UserServiceServer.List
func (svc *UserService) List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	var (
		err      error
		entList  []*ent.User
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.User.Query().
		Order(ent.Desc(user.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := uint32(token)
		listQuery = listQuery.
			Where(user.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListUserRequest_VIEW_UNSPECIFIED, ListUserRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListUserRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithAttachment(func(query *ent.AttachmentQuery) {
				query.Select(attachment.FieldID)
			}).
			WithGroup(func(query *ent.GroupQuery) {
				query.Select(group.FieldID)
			}).
			WithPet(func(query *ent.PetQuery) {
				query.Select(pet.FieldID)
			}).
			WithReceived1(func(query *ent.AttachmentQuery) {
				query.Select(attachment.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoUserList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListUserResponse{
			UserList:      protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements UserServiceServer.BatchCreate
func (svc *UserService) BatchCreate(ctx context.Context, req *BatchCreateUsersRequest) (*BatchCreateUsersResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.UserCreate, len(requests))
	for i, req := range requests {
		user := req.GetUser()
		var err error
		bulk[i], err = svc.createBuilder(user)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.User.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoUserList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateUsersResponse{
			Users: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *UserService) createBuilder(user *User) (*ent.UserCreate, error) {
	m := svc.client.User.Create()
	userAccountBalance := float64(user.GetAccountBalance())
	m.SetAccountBalance(userAccountBalance)
	if user.GetBUser_1() != nil {
		userBUser1 := int(user.GetBUser_1().GetValue())
		m.SetBUser1(userBUser1)
	}
	userBanned := user.GetBanned()
	m.SetBanned(userBanned)
	if user.GetBigInt() != nil {
		userBigInt := schema.BigInt{}
		if err := (&userBigInt).Scan(user.GetBigInt().GetValue()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetBigInt(userBigInt)
	}
	var userCrmID uuid.UUID
	if err := (&userCrmID).UnmarshalBinary(user.GetCrmId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetCrmID(userCrmID)
	userCustomPb := uint8(user.GetCustomPb())
	m.SetCustomPb(userCustomPb)
	userDeviceType := toEntUser_DeviceType(user.GetDeviceType())
	m.SetDeviceType(userDeviceType)
	userExp := uint64(user.GetExp())
	m.SetExp(userExp)
	userExternalID := int(user.GetExternalId())
	m.SetExternalID(userExternalID)
	userHeightInCm := float32(user.GetHeightInCm())
	m.SetHeightInCm(userHeightInCm)
	userJoined := runtime.ExtractTime(user.GetJoined())
	m.SetJoined(userJoined)
	if user.GetLabels() != nil {
		userLabels := user.GetLabels()
		m.SetLabels(userLabels)
	}
	userOmitPrefix := toEntUser_OmitPrefix(user.GetOmitPrefix())
	m.SetOmitPrefix(userOmitPrefix)
	if user.GetOptBool() != nil {
		userOptBool := user.GetOptBool().GetValue()
		m.SetOptBool(userOptBool)
	}
	if user.GetOptNum() != nil {
		userOptNum := int(user.GetOptNum().GetValue())
		m.SetOptNum(userOptNum)
	}
	if user.GetOptStr() != nil {
		userOptStr := user.GetOptStr().GetValue()
		m.SetOptStr(userOptStr)
	}
	userPoints := uint(user.GetPoints())
	m.SetPoints(userPoints)
	userStatus := toEntUser_Status(user.GetStatus())
	m.SetStatus(userStatus)
	if user.GetType() != nil {
		userType := user.GetType().GetValue()
		m.SetType(userType)
	}
	userUserName := user.GetUserName()
	m.SetUserName(userUserName)
	if user.GetAttachment() != nil {
		var userAttachment uuid.UUID
		if err := (&userAttachment).UnmarshalBinary(user.GetAttachment().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetAttachmentID(userAttachment)
	}
	if user.GetGroup() != nil {
		userGroup := int(user.GetGroup().GetId())
		m.SetGroupID(userGroup)
	}
	if user.GetPet() != nil {
		userPet := int(user.GetPet().GetId())
		m.SetPetID(userPet)
	}
	for _, item := range user.GetReceived_1() {
		var received1 uuid.UUID
		if err := (&received1).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddReceived1IDs(received1)
	}
	return m, nil
}
