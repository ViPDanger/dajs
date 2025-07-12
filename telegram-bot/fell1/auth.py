import requests

register = '/register'
auth = '/login'
refresh = '/refresh'

class Authorizator:
    def __init__(self,url :str,login :str,password :str):
        self.url = url
        self.username = login
        self.password = password

    def auth(self):
        data = requests.post(self.url+auth, json={'user': self.username, 'password': self.password}).json()
        self.access_token =  data['Access']
        self.refresh_token = data['Refresh']
        print(self.access_token)
        print(self.refresh_token)
        return True
    def register(self):
        response = requests.post(self.url+register, json={'user': self.username, 'password': self.password})
        data = response.json()
        print(data)
        return True
    def refresh(self): 
        data = requests.post(self.url+refresh, json=self.refresh_token).json()
        self.access_token =  data['Access']


    def GetAccessToken(self):
        return self.access_token
    