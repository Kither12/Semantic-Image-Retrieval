# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ModelService.proto
# Protobuf Python Version: 5.27.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    3,
    '',
    'ModelService.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12ModelService.proto\"*\n\x15TextEmbeddingResponse\x12\x11\n\tembedding\x18\x01 \x03(\x01\"+\n\x16ImageEmbeddingResponse\x12\x11\n\tembedding\x18\x01 \x03(\x01\"$\n\x14TextEmbeddingRequest\x12\x0c\n\x04text\x18\x01 \x01(\t\"&\n\x15ImageEmbeddingRequest\x12\r\n\x05\x63hunk\x18\x01 \x01(\x0c\x32\x93\x01\n\x0cModelService\x12>\n\rTextEmbedding\x12\x15.TextEmbeddingRequest\x1a\x16.TextEmbeddingResponse\x12\x43\n\x0eImageEmbedding\x12\x16.ImageEmbeddingRequest\x1a\x17.ImageEmbeddingResponse(\x01\x42\x07Z\x05./apib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ModelService_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\005./api'
  _globals['_TEXTEMBEDDINGRESPONSE']._serialized_start=22
  _globals['_TEXTEMBEDDINGRESPONSE']._serialized_end=64
  _globals['_IMAGEEMBEDDINGRESPONSE']._serialized_start=66
  _globals['_IMAGEEMBEDDINGRESPONSE']._serialized_end=109
  _globals['_TEXTEMBEDDINGREQUEST']._serialized_start=111
  _globals['_TEXTEMBEDDINGREQUEST']._serialized_end=147
  _globals['_IMAGEEMBEDDINGREQUEST']._serialized_start=149
  _globals['_IMAGEEMBEDDINGREQUEST']._serialized_end=187
  _globals['_MODELSERVICE']._serialized_start=190
  _globals['_MODELSERVICE']._serialized_end=337
# @@protoc_insertion_point(module_scope)
