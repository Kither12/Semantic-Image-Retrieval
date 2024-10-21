from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchResponse(_message.Message):
    __slots__ = ("info", "chunk")
    INFO_FIELD_NUMBER: _ClassVar[int]
    CHUNK_FIELD_NUMBER: _ClassVar[int]
    info: Info
    chunk: bytes
    def __init__(self, info: _Optional[_Union[Info, _Mapping]] = ..., chunk: _Optional[bytes] = ...) -> None: ...

class SearchRequest(_message.Message):
    __slots__ = ("prompt", "total")
    PROMPT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    prompt: str
    total: int
    def __init__(self, prompt: _Optional[str] = ..., total: _Optional[int] = ...) -> None: ...

class Info(_message.Message):
    __slots__ = ("file_name",)
    FILE_NAME_FIELD_NUMBER: _ClassVar[int]
    file_name: str
    def __init__(self, file_name: _Optional[str] = ...) -> None: ...

class UploadResponse(_message.Message):
    __slots__ = ("id", "size")
    ID_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    id: str
    size: int
    def __init__(self, id: _Optional[str] = ..., size: _Optional[int] = ...) -> None: ...

class UploadRequest(_message.Message):
    __slots__ = ("info", "chunk")
    INFO_FIELD_NUMBER: _ClassVar[int]
    CHUNK_FIELD_NUMBER: _ClassVar[int]
    info: Info
    chunk: bytes
    def __init__(self, info: _Optional[_Union[Info, _Mapping]] = ..., chunk: _Optional[bytes] = ...) -> None: ...
