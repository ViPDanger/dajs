import asyncio
from aiogram import Bot, Dispatcher,Router
from options import BOT_TOKEN
from handlers import routers



async def main():
    bot = Bot(token=BOT_TOKEN)
    dp = Dispatcher()
    router = Router()
    for r in routers:
        router.include_router(r)
    dp.include_router(router)
    await dp.start_polling(bot)

if __name__ == "__main__":
    asyncio.run(main())

     