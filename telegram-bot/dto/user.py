from pydantic import BaseModel, Field
from datetime import datetime


class User(BaseModel):
    username: str = Field("", alias="username")
    password: str = Field("", alias="password")


class AccessToken(BaseModel):
    token:    str = Field("", alias="access_token")
    expire_time: datetime = Field("", alias="access_exp")


class RefreshToken(BaseModel):
    token:    str = Field("", alias="refresh_token")
    expire_time: datetime = Field("", alias="refresh_exp")


class Tokens(BaseModel):
    access: AccessToken = Field(default_factory=AccessToken, alias="Access")
    refresh: RefreshToken = Field(default_factory=RefreshToken,alias="Refresh")
