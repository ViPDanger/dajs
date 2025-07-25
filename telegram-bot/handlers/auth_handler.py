from aiogram import  types, Router, types
from aiogram.filters import Command
from aiogram.types import ReplyKeyboardMarkup, KeyboardButton
from client import Authorizator
from options import BASE_URL
from .str import start_str,character_menu_str,help_str
import asyncio

router = Router()
user_data: dict[int,Authorizator] = {}

@router.message(Command("start"))
async def start_handler(message: types.Message):
    user_id = message.from_user.id
    authorizator = user_data.get(user_id)
    if authorizator:
        authorizator.GetAccessToken()
    else:
        authorizator = Authorizator(BASE_URL,str(user_id),str(user_id))
        user_data[user_id] = authorizator
        authorizator.GetAccessToken()

    start_kb = ReplyKeyboardMarkup(
    keyboard=[
        [KeyboardButton(text=character_menu_str)],
        [KeyboardButton(text=help_str)]
    ],
    resize_keyboard=True
    )
    await message.answer("Привет! Я бот!", reply_markup=start_kb)

@router.message(lambda msg: msg.text == start_str)
async def return_handler(message: types.Message):
    await start_handler(message)

# ? вопросики
@router.message(lambda msg: msg.text == help_str)
async def help_handler(message: types.Message):
    help_kb = ReplyKeyboardMarkup(
    keyboard=[
        [KeyboardButton(text=start_str)],
    ],
    resize_keyboard=True
    )
    await message.answer("Вообщем то, я создан для того чтобы хранить персонажей и инвентари, а так же иметь быстрый доступ к описаниям заклинаний и прочему. Вот такой я крутой, да", reply_markup=help_kb)

#@router.message()
#async def fallback_handler(message: types.Message):
#    await message.answer("Я тебя не понял 🤖. Нажми кнопку или /start.")