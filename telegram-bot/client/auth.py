import requests
import json
from dto.user import User,AccessToken,RefreshToken,Tokens
from dataclasses import asdict
register = '/register'
auth = '/login'
refresh = '/refresh'

class Authorizator:
    url  :          str = ""
    user :          User = User()
    access_token  : AccessToken = AccessToken()
    refresh_token : RefreshToken = RefreshToken()
    def __init__(self,url :str,login :str,password :str):
        self.url           = url
        self.user = User(username=login,password=password)
        self.register()
    def auth(self) -> bool:
        response = requests.post(self.url+auth, json=self.user.model_dump(by_alias=True))
        if response.status_code != 200:
            if self.register() == False:
                return False
            return self.auth()
        tokens = Tokens.model_validate(response.json())
        self.access_token =  tokens.access
        self.refresh_token = tokens.refresh
        # проверка?
        print(self.access_token)
        print(self.refresh_token)
        return True
    def register(self):
        response = requests.post(self.url+register, json=self.user.model_dump(by_alias=True))
        data = response.json()
        if response.status_code == 400: 
            print(data)
            return False
        return True
    def refresh(self): 
        response = requests.post(self.url+refresh, json=asdict(self.refresh_token.token)).json()
        self.access_token =  AccessToken.model_validate(response.json())
    def GetAccessToken(self)->str:
        if not self.access_token:
            self.register()
            self.auth()
            return self.access_token.token
        if not self.access_token.token:
            if not self.auth():
                self.register()
                self.auth()
        return self.access_token.token
    