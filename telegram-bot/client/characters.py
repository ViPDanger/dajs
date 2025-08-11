import dto
from dto.character import Character 
import aiohttp
from client.auth import Authorizator
from options import BASE_URL
import json
from dataclasses import asdict
from typing import Optional
from pydantic import TypeAdapter

CHARACTER_PATH = "/char"
async def post_character(authorizator: Authorizator,character: Character, base_url: str = BASE_URL) -> Optional[str]:
    url = f"{base_url}{CHARACTER_PATH}"
    headers = {"Content-Type": "application/json","Authorization":authorizator.GetAccessToken()}
    async with aiohttp.ClientSession() as session:
        async with session.post(url, data=character.model_dump_json(by_alias=True), headers=headers) as resp:
            if resp.status == 201:
                data = await resp.json()
                print(data)
                return data.get("message")
            else:
                text = await resp.text()
                raise Exception(f"Ошибка отправки персонажа: {resp.status} — {text}")
            
async def characters(authorizator: Authorizator,base_url: str = BASE_URL) -> Optional[list[Character]]:
    url = f"{base_url}{CHARACTER_PATH}"
    headers = {"Content-Type": "application/json","Authorization":authorizator.GetAccessToken()}
    async with aiohttp.ClientSession() as session:
        async with session.get(url, headers=headers) as resp:
            if resp.status == 200:
                adapter = TypeAdapter(list[Character])
                characters = adapter.validate_python(await resp.json())
                print(characters)
                return characters
            else:
                text = await resp.text()
                raise Exception(f"Ошибка отправки персонажа: {resp.status} — {text}")
            
