"""Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT."""

from __future__ import annotations
from petstore.types import BaseModel
from petstore.utils import FieldMetadata, SecurityMetadata
from typing_extensions import Annotated, TypedDict


class SecurityTypedDict(TypedDict):
    api_key: str


class Security(BaseModel):
    api_key: Annotated[
        str,
        FieldMetadata(
            security=SecurityMetadata(
                scheme=True,
                scheme_type="apiKey",
                sub_type="header",
                field_name="api_key",
            )
        ),
    ]
