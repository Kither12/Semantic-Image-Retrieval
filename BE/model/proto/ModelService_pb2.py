# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ModelService.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12ModelService.proto\"*\n\x15TextEmbeddingResponse\x12\x11\n\tembedding\x18\x01 \x03(\x01\"+\n\x16ImageEmbeddingResponse\x12\x11\n\tembedding\x18\x01 \x03(\x01\"$\n\x14TextEmbeddingRequest\x12\x0c\n\x04text\x18\x01 \x01(\t\"&\n\x15ImageEmbeddingRequest\x12\r\n\x05\x63hunk\x18\x01 \x01(\x0c\x32\x93\x01\n\x0cModelService\x12>\n\rTextEmbedding\x12\x15.TextEmbeddingRequest\x1a\x16.TextEmbeddingResponse\x12\x43\n\x0eImageEmbedding\x12\x16.ImageEmbeddingRequest\x1a\x17.ImageEmbeddingResponse(\x01\x42\x07Z\x05./apib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ModelService_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\005./api'
  _TEXTEMBEDDINGRESPONSE._serialized_start=22
  _TEXTEMBEDDINGRESPONSE._serialized_end=64
  _IMAGEEMBEDDINGRESPONSE._serialized_start=66
  _IMAGEEMBEDDINGRESPONSE._serialized_end=109
  _TEXTEMBEDDINGREQUEST._serialized_start=111
  _TEXTEMBEDDINGREQUEST._serialized_end=147
  _IMAGEEMBEDDINGREQUEST._serialized_start=149
  _IMAGEEMBEDDINGREQUEST._serialized_end=187
  _MODELSERVICE._serialized_start=190
  _MODELSERVICE._serialized_end=337
# @@protoc_insertion_point(module_scope)