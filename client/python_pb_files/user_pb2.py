# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x1a\x1bgoogle/protobuf/empty.proto\">\n\x07\x43ontact\x12\x11\n\thome_addr\x18\x01 \x01(\t\x12\x0f\n\x07mob_num\x18\x02 \x01(\t\x12\x0f\n\x07mail_id\x18\x03 \x01(\t\"\x92\x01\n\x04User\x12\n\n\x02id\x18\x01 \x01(\t\x12\x12\n\nfirst_name\x18\x02 \x01(\t\x12\x11\n\tlast_name\x18\x03 \x01(\t\x12\x0c\n\x04\x63ity\x18\x04 \x01(\t\x12\r\n\x05phone\x18\x05 \x01(\t\x12\x0e\n\x06height\x18\x06 \x01(\x02\x12\x0f\n\x07married\x18\x07 \x01(\x08\x12\x19\n\x07\x63ontact\x18\x08 \x01(\x0b\x32\x08.Contact\"\x14\n\x06UserID\x12\n\n\x02id\x18\x01 \x01(\t\"\x16\n\x07UserIDs\x12\x0b\n\x03ids\x18\x01 \x03(\t\" \n\x08UserList\x12\x14\n\x05users\x18\x01 \x03(\x0b\x32\x05.User\">\n\x0eSearchCriteria\x12\x0c\n\x04\x63ity\x18\x01 \x01(\t\x12\r\n\x05phone\x18\x02 \x01(\t\x12\x0f\n\x07married\x18\x03 \x01(\x08\x32\x9f\x02\n\x0bUserService\x12\x19\n\x07GetUser\x12\x07.UserID\x1a\x05.User\x12.\n\tListUsers\x12\x16.google.protobuf.Empty\x1a\t.UserList\x12\x19\n\x07SetUser\x12\x05.User\x1a\x07.UserID\x12+\n\nUpdateUser\x12\x05.User\x1a\x16.google.protobuf.Empty\x12-\n\nDeleteUser\x12\x07.UserID\x1a\x16.google.protobuf.Empty\x12)\n\x0bSearchUsers\x12\x0f.SearchCriteria\x1a\t.UserList\x12#\n\x0cGetUsersByID\x12\x08.UserIDs\x1a\t.UserListB\x07Z\x05/userb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\005/user'
  _globals['_CONTACT']._serialized_start=43
  _globals['_CONTACT']._serialized_end=105
  _globals['_USER']._serialized_start=108
  _globals['_USER']._serialized_end=254
  _globals['_USERID']._serialized_start=256
  _globals['_USERID']._serialized_end=276
  _globals['_USERIDS']._serialized_start=278
  _globals['_USERIDS']._serialized_end=300
  _globals['_USERLIST']._serialized_start=302
  _globals['_USERLIST']._serialized_end=334
  _globals['_SEARCHCRITERIA']._serialized_start=336
  _globals['_SEARCHCRITERIA']._serialized_end=398
  _globals['_USERSERVICE']._serialized_start=401
  _globals['_USERSERVICE']._serialized_end=688
# @@protoc_insertion_point(module_scope)
