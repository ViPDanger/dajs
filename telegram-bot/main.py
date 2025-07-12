import asyncio
from aiogram import Bot, Dispatcher
from handlers import routers


API_TOKEN = '7620989653:AAE8eiBN0MvCTnjip8lMRjbOAHTeyZRlMoo'

async def main():
    bot = Bot(token=API_TOKEN)
    dp = Dispatcher()
    for router in routers:
        dp.include_router(router)
    await dp.start_polling(bot)

if __name__ == "__main__":
    asyncio.run(main())

    