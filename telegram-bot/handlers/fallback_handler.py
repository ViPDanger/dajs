from aiogram import  types, Router, types
from aiogram.filters import Command
from aiogram.types import ReplyKeyboardMarkup, KeyboardButton
from client import Authorizator
from .auth_handler import user_data
from .str import start_str
from options import BASE_URL
import asyncio
from dto.character import Character 
import client.characters

router = Router()

@router.message() 
async def fallback_handler(message: types.Message):
    kb = ReplyKeyboardMarkup(
    keyboard=[
        [KeyboardButton(text=start_str)],
    ],
    resize_keyboard=True
    )
    await message.answer("Неизвестная команда.",reply_markup=kb)