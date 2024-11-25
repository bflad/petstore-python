"""Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT."""

from __future__ import annotations
from enum import Enum
from petstore.types import BaseModel
from petstore.utils import FieldMetadata, QueryParamMetadata
from typing import Optional
from typing_extensions import Annotated, NotRequired, TypedDict


class QueryParamStatus(str, Enum):
    r"""Status values that need to be considered for filter"""

    AVAILABLE = "available"
    PENDING = "pending"
    SOLD = "sold"


class FindPetsByStatusRequestTypedDict(TypedDict):
    status: NotRequired[QueryParamStatus]
    r"""Status values that need to be considered for filter"""


class FindPetsByStatusRequest(BaseModel):
    status: Annotated[
        Optional[QueryParamStatus],
        FieldMetadata(query=QueryParamMetadata(style="form", explode=True)),
    ] = QueryParamStatus.AVAILABLE
    r"""Status values that need to be considered for filter"""
