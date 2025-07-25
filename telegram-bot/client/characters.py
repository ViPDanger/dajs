import dto
from dto.character import Character 
import aiohttp
from client.auth import Authorizator
from options import BASE_URL
import json
from dataclasses import asdict
from typing import Optional

CHARACTER_PATH = "/character"






async def New(authorizator: Authorizator ,character: Character):
    async with aiohttp.ClientSession() as session:
        session.headers.add("Authorization",Authorizator.access_token.token)
        async with session.post(f"{BASE_URL}{CHARACTER_PATH}") as resp:
            if resp.status != 200:
                raise Exception("Ошибка запроса к API")
            data = await resp.json()
            # Предположим, сервер вернёт [{"name": "Гендальф"}, {"name": "Фродо"}]
            return [char["name"] for char in data]
        
async def post_character(authorizator: Authorizator,character: Character, base_url: str = BASE_URL) -> Optional[str]:
    url = f"{base_url}{CHARACTER_PATH}"
    headers = {"Content-Type": "application/json","Authorization":authorizator.GetAccessToken()}
    async with aiohttp.ClientSession() as session:
        async with session.post(url, data=character.model_dump_json(by_alias=True), headers=headers) as resp:
            if resp.status == 200:
                data = await resp.json()
                print(data)
                return data.get("message")
            else:

                text = await resp.text()
                raise Exception(f"Ошибка отправки персонажа: {resp.status} — {text}")
            
async def get_character(authorizator: Authorizator,id: str, base_url: str = BASE_URL) -> Optional[Character]:
    url = f"{base_url}{CHARACTER_PATH}"
    headers = {"Content-Type": "application/json","Authorization":authorizator.GetAccessToken(),"id":id}
    async with aiohttp.ClientSession() as session:
        async with session.get(url, headers=headers) as resp:
            if resp.status == 200:
                char = Character.model_validate(await resp.json())
                print(char)
                return char
            else:
                text = await resp.text()
                raise Exception(f"Ошибка отправки персонажа: {resp.status} — {text}")
            
