from aiogram import  types, Router, types
from aiogram.filters import Command
from aiogram.types import ReplyKeyboardMarkup, KeyboardButton,WebAppInfo
from client import Authorizator
from .auth_handler import user_data
from .str import start_str,character_menu_str,character_str,post_character_str
from options import BASE_URL,WEBAPP_URL
import asyncio
import ssl
import aiohttp
from dto.character import Character 
import client.characters


router = Router()

@router.message(lambda msg: msg.text == character_menu_str)
async def character_menu(message: types.Message):
    kb = ReplyKeyboardMarkup(
    keyboard=[
        [KeyboardButton(text="Открыть Mini App", web_app=WebAppInfo(url="https://dajs.vipdanger.keenetic.pro:8085/char"))],
        [KeyboardButton(text=start_str)],
    ],
    )
    await message.answer("Возможные действия с персонажами:",reply_markup=kb)