# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/feature/feature_last_used_info.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n*proto/feature/feature_last_used_info.proto\x12\x11\x62ucketeer.feature\"\xa2\x01\n\x13\x46\x65\x61tureLastUsedInfo\x12\x12\n\nfeature_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\x05\x12\x14\n\x0clast_used_at\x18\x03 \x01(\x03\x12\x12\n\ncreated_at\x18\x04 \x01(\x03\x12\x1d\n\x15\x63lient_oldest_version\x18\x05 \x01(\t\x12\x1d\n\x15\x63lient_latest_version\x18\x06 \x01(\tB1Z/github.com/bucketeer-io/bucketeer/proto/featureb\x06proto3')



_FEATURELASTUSEDINFO = DESCRIPTOR.message_types_by_name['FeatureLastUsedInfo']
FeatureLastUsedInfo = _reflection.GeneratedProtocolMessageType('FeatureLastUsedInfo', (_message.Message,), {
  'DESCRIPTOR' : _FEATURELASTUSEDINFO,
  '__module__' : 'proto.feature.feature_last_used_info_pb2'
  # @@protoc_insertion_point(class_scope:bucketeer.feature.FeatureLastUsedInfo)
  })
_sym_db.RegisterMessage(FeatureLastUsedInfo)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z/github.com/bucketeer-io/bucketeer/proto/feature'
  _FEATURELASTUSEDINFO._serialized_start=66
  _FEATURELASTUSEDINFO._serialized_end=228
# @@protoc_insertion_point(module_scope)