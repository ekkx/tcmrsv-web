// @generated by protoc-gen-es v2.5.2 with parameter "target=ts"
// @generated from file user/v1/user.proto (package user.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file user/v1/user.proto.
 */
export const file_user_v1_user: GenFile = /*@__PURE__*/
  fileDesc("ChJ1c2VyL3YxL3VzZXIucHJvdG8SB3VzZXIudjEifQoEVXNlchIKCgJpZBgBIAEoCRIUCgxkaXNwbGF5X25hbWUYAiABKAkSIgoLbWFzdGVyX3VzZXIYAyABKAsyDS51c2VyLnYxLlVzZXISLwoLY3JlYXRlX3RpbWUYBCABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wIiEKDkdldFVzZXJSZXF1ZXN0Eg8KB3VzZXJfaWQYASABKAkiLgoPR2V0VXNlclJlc3BvbnNlEhsKBHVzZXIYASABKAsyDS51c2VyLnYxLlVzZXIiFwoVTGlzdFNsYXZlVXNlcnNSZXF1ZXN0IjYKFkxpc3RTbGF2ZVVzZXJzUmVzcG9uc2USHAoFdXNlcnMYASADKAsyDS51c2VyLnYxLlVzZXIiQAoWQ3JlYXRlU2xhdmVVc2VyUmVxdWVzdBIUCgxkaXNwbGF5X25hbWUYASABKAkSEAoIcGFzc3dvcmQYAiABKAkiNgoXQ3JlYXRlU2xhdmVVc2VyUmVzcG9uc2USGwoEdXNlchgBIAEoCzINLnVzZXIudjEuVXNlciIpChFVcGRhdGVVc2VyUmVxdWVzdBIUCgxkaXNwbGF5X25hbWUYAiABKAkiMQoSVXBkYXRlVXNlclJlc3BvbnNlEhsKBHVzZXIYASABKAsyDS51c2VyLnYxLlVzZXIiEwoRRGVsZXRlVXNlclJlcXVlc3QiFAoSRGVsZXRlVXNlclJlc3BvbnNlIikKFkRlbGV0ZVNsYXZlVXNlclJlcXVlc3QSDwoHdXNlcl9pZBgBIAEoCSIZChdEZWxldGVTbGF2ZVVzZXJSZXNwb25zZTLYAwoLVXNlclNlcnZpY2USPAoHR2V0VXNlchIXLnVzZXIudjEuR2V0VXNlclJlcXVlc3QaGC51c2VyLnYxLkdldFVzZXJSZXNwb25zZRJRCg5MaXN0U2xhdmVVc2VycxIeLnVzZXIudjEuTGlzdFNsYXZlVXNlcnNSZXF1ZXN0Gh8udXNlci52MS5MaXN0U2xhdmVVc2Vyc1Jlc3BvbnNlElQKD0NyZWF0ZVNsYXZlVXNlchIfLnVzZXIudjEuQ3JlYXRlU2xhdmVVc2VyUmVxdWVzdBogLnVzZXIudjEuQ3JlYXRlU2xhdmVVc2VyUmVzcG9uc2USRQoKVXBkYXRlVXNlchIaLnVzZXIudjEuVXBkYXRlVXNlclJlcXVlc3QaGy51c2VyLnYxLlVwZGF0ZVVzZXJSZXNwb25zZRJFCgpEZWxldGVVc2VyEhoudXNlci52MS5EZWxldGVVc2VyUmVxdWVzdBobLnVzZXIudjEuRGVsZXRlVXNlclJlc3BvbnNlElQKD0RlbGV0ZVNsYXZlVXNlchIfLnVzZXIudjEuRGVsZXRlU2xhdmVVc2VyUmVxdWVzdBogLnVzZXIudjEuRGVsZXRlU2xhdmVVc2VyUmVzcG9uc2VCmgEKC2NvbS51c2VyLnYxQglVc2VyUHJvdG9QAVpDZ2l0aHViLmNvbS9la2t4L3RjbXJzdi13ZWIvc2VydmVyL2ludGVybmFsL3NoYXJlZC9wYi91c2VyL3YxO3VzZXJ2MaICA1VYWKoCB1VzZXIuVjHKAgdVc2VyXFYx4gITVXNlclxWMVxHUEJNZXRhZGF0YeoCCFVzZXI6OlYxYgZwcm90bzM", [file_google_protobuf_timestamp]);

/**
 * @generated from message user.v1.User
 */
export type User = Message<"user.v1.User"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string display_name = 2;
   */
  displayName: string;

  /**
   * @generated from field: user.v1.User master_user = 3;
   */
  masterUser?: User;

  /**
   * @generated from field: google.protobuf.Timestamp create_time = 4;
   */
  createTime?: Timestamp;
};

/**
 * Describes the message user.v1.User.
 * Use `create(UserSchema)` to create a new message.
 */
export const UserSchema: GenMessage<User> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 0);

/**
 * @generated from message user.v1.GetUserRequest
 */
export type GetUserRequest = Message<"user.v1.GetUserRequest"> & {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;
};

/**
 * Describes the message user.v1.GetUserRequest.
 * Use `create(GetUserRequestSchema)` to create a new message.
 */
export const GetUserRequestSchema: GenMessage<GetUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 1);

/**
 * @generated from message user.v1.GetUserResponse
 */
export type GetUserResponse = Message<"user.v1.GetUserResponse"> & {
  /**
   * @generated from field: user.v1.User user = 1;
   */
  user?: User;
};

/**
 * Describes the message user.v1.GetUserResponse.
 * Use `create(GetUserResponseSchema)` to create a new message.
 */
export const GetUserResponseSchema: GenMessage<GetUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 2);

/**
 * @generated from message user.v1.ListSlaveUsersRequest
 */
export type ListSlaveUsersRequest = Message<"user.v1.ListSlaveUsersRequest"> & {
};

/**
 * Describes the message user.v1.ListSlaveUsersRequest.
 * Use `create(ListSlaveUsersRequestSchema)` to create a new message.
 */
export const ListSlaveUsersRequestSchema: GenMessage<ListSlaveUsersRequest> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 3);

/**
 * @generated from message user.v1.ListSlaveUsersResponse
 */
export type ListSlaveUsersResponse = Message<"user.v1.ListSlaveUsersResponse"> & {
  /**
   * @generated from field: repeated user.v1.User users = 1;
   */
  users: User[];
};

/**
 * Describes the message user.v1.ListSlaveUsersResponse.
 * Use `create(ListSlaveUsersResponseSchema)` to create a new message.
 */
export const ListSlaveUsersResponseSchema: GenMessage<ListSlaveUsersResponse> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 4);

/**
 * @generated from message user.v1.CreateSlaveUserRequest
 */
export type CreateSlaveUserRequest = Message<"user.v1.CreateSlaveUserRequest"> & {
  /**
   * @generated from field: string display_name = 1;
   */
  displayName: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;
};

/**
 * Describes the message user.v1.CreateSlaveUserRequest.
 * Use `create(CreateSlaveUserRequestSchema)` to create a new message.
 */
export const CreateSlaveUserRequestSchema: GenMessage<CreateSlaveUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 5);

/**
 * @generated from message user.v1.CreateSlaveUserResponse
 */
export type CreateSlaveUserResponse = Message<"user.v1.CreateSlaveUserResponse"> & {
  /**
   * @generated from field: user.v1.User user = 1;
   */
  user?: User;
};

/**
 * Describes the message user.v1.CreateSlaveUserResponse.
 * Use `create(CreateSlaveUserResponseSchema)` to create a new message.
 */
export const CreateSlaveUserResponseSchema: GenMessage<CreateSlaveUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 6);

/**
 * @generated from message user.v1.UpdateUserRequest
 */
export type UpdateUserRequest = Message<"user.v1.UpdateUserRequest"> & {
  /**
   * @generated from field: string display_name = 2;
   */
  displayName: string;
};

/**
 * Describes the message user.v1.UpdateUserRequest.
 * Use `create(UpdateUserRequestSchema)` to create a new message.
 */
export const UpdateUserRequestSchema: GenMessage<UpdateUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 7);

/**
 * @generated from message user.v1.UpdateUserResponse
 */
export type UpdateUserResponse = Message<"user.v1.UpdateUserResponse"> & {
  /**
   * @generated from field: user.v1.User user = 1;
   */
  user?: User;
};

/**
 * Describes the message user.v1.UpdateUserResponse.
 * Use `create(UpdateUserResponseSchema)` to create a new message.
 */
export const UpdateUserResponseSchema: GenMessage<UpdateUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 8);

/**
 * @generated from message user.v1.DeleteUserRequest
 */
export type DeleteUserRequest = Message<"user.v1.DeleteUserRequest"> & {
};

/**
 * Describes the message user.v1.DeleteUserRequest.
 * Use `create(DeleteUserRequestSchema)` to create a new message.
 */
export const DeleteUserRequestSchema: GenMessage<DeleteUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 9);

/**
 * @generated from message user.v1.DeleteUserResponse
 */
export type DeleteUserResponse = Message<"user.v1.DeleteUserResponse"> & {
};

/**
 * Describes the message user.v1.DeleteUserResponse.
 * Use `create(DeleteUserResponseSchema)` to create a new message.
 */
export const DeleteUserResponseSchema: GenMessage<DeleteUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 10);

/**
 * @generated from message user.v1.DeleteSlaveUserRequest
 */
export type DeleteSlaveUserRequest = Message<"user.v1.DeleteSlaveUserRequest"> & {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;
};

/**
 * Describes the message user.v1.DeleteSlaveUserRequest.
 * Use `create(DeleteSlaveUserRequestSchema)` to create a new message.
 */
export const DeleteSlaveUserRequestSchema: GenMessage<DeleteSlaveUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 11);

/**
 * @generated from message user.v1.DeleteSlaveUserResponse
 */
export type DeleteSlaveUserResponse = Message<"user.v1.DeleteSlaveUserResponse"> & {
};

/**
 * Describes the message user.v1.DeleteSlaveUserResponse.
 * Use `create(DeleteSlaveUserResponseSchema)` to create a new message.
 */
export const DeleteSlaveUserResponseSchema: GenMessage<DeleteSlaveUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1_user, 12);

/**
 * @generated from service user.v1.UserService
 */
export const UserService: GenService<{
  /**
   * @generated from rpc user.v1.UserService.GetUser
   */
  getUser: {
    methodKind: "unary";
    input: typeof GetUserRequestSchema;
    output: typeof GetUserResponseSchema;
  },
  /**
   * @generated from rpc user.v1.UserService.ListSlaveUsers
   */
  listSlaveUsers: {
    methodKind: "unary";
    input: typeof ListSlaveUsersRequestSchema;
    output: typeof ListSlaveUsersResponseSchema;
  },
  /**
   * @generated from rpc user.v1.UserService.CreateSlaveUser
   */
  createSlaveUser: {
    methodKind: "unary";
    input: typeof CreateSlaveUserRequestSchema;
    output: typeof CreateSlaveUserResponseSchema;
  },
  /**
   * @generated from rpc user.v1.UserService.UpdateUser
   */
  updateUser: {
    methodKind: "unary";
    input: typeof UpdateUserRequestSchema;
    output: typeof UpdateUserResponseSchema;
  },
  /**
   * @generated from rpc user.v1.UserService.DeleteUser
   */
  deleteUser: {
    methodKind: "unary";
    input: typeof DeleteUserRequestSchema;
    output: typeof DeleteUserResponseSchema;
  },
  /**
   * @generated from rpc user.v1.UserService.DeleteSlaveUser
   */
  deleteSlaveUser: {
    methodKind: "unary";
    input: typeof DeleteSlaveUserRequestSchema;
    output: typeof DeleteSlaveUserResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_user_v1_user, 0);

