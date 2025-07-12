from aiogram import  types, Router, types
from aiogram.filters import Command
from aiogram.types import ReplyKeyboardMarkup, KeyboardButton
import asyncio

start_router = Router()

main_kb = ReplyKeyboardMarkup(
    keyboard=[
        [KeyboardButton(text="👋 Поздороваться")],
        [KeyboardButton(text="❓ Что ты умеешь?")]
    ],
    resize_keyboard=True
)

@start_router.message(Command("start"))
async def start_handler(message: types.Message):
    await message.answer("Привет! Я бот!", reply_markup=main_kb)


# main_kb
@start_router.message(lambda msg: msg.text == main_kb.keyboard[0][0].text)
async def greet_handler(message: types.Message):
    await message.answer("Привет-привет!")

@start_router.message(lambda msg: msg.text == main_kb.keyboard[1][0].text)
async def help_handler(message: types.Message):
    await message.answer("Я умею здороваться и показывать кнопки!")

@start_router.message()
async def fallback_handler(message: types.Message):
    await message.answer("Я тебя не понял 🤖. Нажми кнопку или /start.")