from aiogram import  types, Router, types
from aiogram.filters import Command
from aiogram.types import ReplyKeyboardMarkup, KeyboardButton
from client import Authorizator
from .auth_handler import user_data
from .str import start_str,character_menu_str,get_character_str,post_character_str
from options import BASE_URL
import asyncio
from dto.character import Character 
import client.characters

router = Router()

@router.message(lambda msg: msg.text == character_menu_str)
async def character_menu(message: types.Message):
    kb = ReplyKeyboardMarkup(
    keyboard=[
        [KeyboardButton(text=get_character_str)],
        [KeyboardButton(text=post_character_str)],
        [KeyboardButton(text=start_str)],
    ],
    )
    await message.answer("Возможные действия с персонажами:",reply_markup=kb)
@router.message(lambda msg: msg.text == post_character_str)
async def post_character(message: types.Message):
    user_id = message.from_user.id
    authorizator = user_data.get(user_id)
    if authorizator:
        authorizator.GetAccessToken()
    else:
        authorizator = Authorizator(BASE_URL,str(user_id),str(user_id))
        user_data[user_id] = authorizator
        authorizator.auth()
    await client.characters.post_character(
            authorizator= authorizator,
            character=Character(id="01",name="vitold")
            )
    kb = ReplyKeyboardMarkup(
    keyboard=[
        [KeyboardButton(text=start_str)],
    ],
    resize_keyboard=True
    )
    await message.answer("гы гы пенис!", reply_markup=kb)

@router.message(lambda msg: msg.text == get_character_str)
async def get_character(message: types.Message):
    user_id = message.from_user.id
    authorizator = user_data.get(user_id)
    if authorizator:
        authorizator.GetAccessToken()
    else:
        authorizator = Authorizator(BASE_URL,str(user_id),str(user_id))
        user_data[user_id] = authorizator
        authorizator.auth()
    await client.characters.get_creator_characters(
        authorizator= authorizator,
        id="01"
        )
    kb = ReplyKeyboardMarkup(
    keyboard=[
        [KeyboardButton(text=start_str)],
    ],
    resize_keyboard=True
    )
    await message.answer("гы гы пенис!", reply_markup=kb)